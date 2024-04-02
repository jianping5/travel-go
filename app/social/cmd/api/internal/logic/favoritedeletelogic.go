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

var FavoriteNoExistError = xerr.NewErrMsg("收藏夹不存在")

type FavoriteDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteDeleteLogic {
	return &FavoriteDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteDeleteLogic) FavoriteDelete(req *types.FavoriteDeleteReq) error {
	// 删除收藏夹
	affected := l.svcCtx.DB.Delete(&model.Favorite{}, req.Id).RowsAffected
	if affected == 0 {
		return errors.Wrap(FavoriteNoExistError, "收藏夹不存在")
	}

	// 更新收藏量
	var itemIds []int64
	l.svcCtx.DB.Model(&model.Favor{}).Select("item_id").Where("favorite_id = ?", req.Id).Scan(&itemIds)
	if itemIds == nil || len(itemIds) == 0 {
		return nil
	}
	l.svcCtx.DB.Model(&model.Content{}).Where("id IN (?)", itemIds).Update("favor_count", gorm.Expr("favor_count - ?", 1))

	// 删除收藏夹下的收藏内容
	if err := l.svcCtx.DB.Delete(&model.Favor{}, "favorite_id = ?", req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	return nil
}
