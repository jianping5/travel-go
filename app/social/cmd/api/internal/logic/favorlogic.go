package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
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
		ItemId:     req.ItemId,
	}
	if err := l.svcCtx.DB.Create(favor).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "收藏失败")
	}

	// 更新对应收藏量
	// 查看除去当前收藏夹之外该 item 是否存在其他收藏夹中
	var id int64
	// 若没有，则增加收藏量
	if l.svcCtx.DB.Model(&model.Favor{}).Select("id").Where("user_id = ? and item_id = ? and favorite_id != ?", loginUserId, req.ItemId, req.FavoriteId).First(&id); id == 0 {
		l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", req.ItemId).Update("favor_count", gorm.Expr("favor_count + ?", 1))
	}

	return nil
}
