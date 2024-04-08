package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/xerr"

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
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// 先查找收藏的内容id
	var favor model.Favor
	l.svcCtx.DB.Model(&model.Favor{}).Select("item_type", "item_id").Where("favorite_id = ? and item_id = ?", req.FavoriteId, req.ItemId).Scan(&favor)

	// 删除
	if err := l.svcCtx.DB.Where("favorite_id = ? and item_id = ?", req.FavoriteId, req.ItemId).Delete(&model.Favor{}).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	// 更新对应收藏量
	// 查看删除之后是否还存在该 item 其他的收藏（即在其他收藏夹中）
	var id int64
	// 若不存在，则减少收藏量
	if l.svcCtx.DB.Model(&model.Favor{}).Select("id").Where("user_id = ? and item_id = ?", loginUserId, favor.ItemId).First(&id); id == 0 {
		l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", favor.ItemId).Update("favor_count", gorm.Expr("favor_count - ?", 1))
	}

	return nil
}
