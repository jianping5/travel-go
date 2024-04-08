package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorCancelLogic {
	return &FavorCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorCancelLogic) FavorCancel(req *types.FavorCancelReq) error {
	if err := l.svcCtx.DB.Delete(&model.Favor{}).Where("favorite_id = ? and item_id = ?", req.FavoriteId, req.ItemId).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	var favor model.Favor
	l.svcCtx.DB.Model(&model.Favor{}).Select("item_type", "item_id").Where("favorite_id = ? and item_id = ?", req.FavoriteId, req.ItemId).Scan(&favor)

	// 更新对应收藏量
	l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", favor.ItemId).Update("favor_count", gorm.Expr("favor_count - ?", 1))

	return nil
}
