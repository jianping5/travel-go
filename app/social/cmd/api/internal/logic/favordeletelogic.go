package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/enum"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorDeleteLogic {
	return &FavorDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorDeleteLogic) FavorDelete(req *types.FavorDeleteReq) error {
	if err := l.svcCtx.DB.Delete(&model.Favor{}, req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	var favor model.Favor
	l.svcCtx.DB.Model(&model.Favor{}).Select("itemType", "itemId").Where("id = ?", req.Id).Scan(&favor)

	// 更新对应收藏量
	switch enum.ItemType(favor.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Model(&model.Article{}).Where("id = ?", favor.ItemId).Update("favorCount", gorm.Expr("favorCount - ?", 1))
		break
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Video{}).Where("id = ?", favor.ItemId).Update("favorCount", gorm.Expr("favorCount - ?", 1))
		break
	}

	return nil
}
