package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorLogic {
	return &FavorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorLogic) Favor(req *types.FavorReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	favor := &model.Favor{
		UserId:     loginUserId,
		FavoriteId: req.FavoriteId,
		ItemType:   req.ItemType,
		ItemId:     req.ItemId,
	}
	if err := l.svcCtx.DB.Create(favor).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "收藏失败")
	}

	// 更新对应收藏量
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Model(&model.Article{}).Where("id = ?", req.ItemId).Update("favorCount", gorm.Expr("favorCount + ?", 1))
		break
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Video{}).Where("id = ?", req.ItemId).Update("favorCount", gorm.Expr("favorCount + ?", 1))
		break
	}
	return nil
}
