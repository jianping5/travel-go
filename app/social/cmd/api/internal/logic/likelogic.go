package logic

import (
	"context"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// TODO：是否考虑使用消息队列及 Redis
	like := &model.Like{
		UserId:      loginUserId,
		ItemType:    req.ItemType,
		ItemId:      req.ItemId,
		LikedStatus: !req.LikedStatus,
	}
	// 简陋写法（如果点赞记录之前存在，则将其删除）
	var id int64
	if l.svcCtx.DB.Model(&model.Like{}).Select("id").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, req.ItemType, req.ItemId).Scan(&id); id != 0 {
		l.svcCtx.DB.Delete(&model.Like{}, "id = ?", id)
	}
	l.svcCtx.DB.Create(like)
	// 更新对应点赞量
	// TODO: 点赞前应该查询数据库判断之前是否点赞还需要校验点赞对象是否存在，（前端会传输点赞状态，但安全性较低）
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
	case enum.VIDEO:
		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", req.ItemId).Update("like_count", gorm.Expr("like_count - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", req.ItemId).Update("like_count", gorm.Expr("like_count + ?", 1))
		}
		break
	case enum.DYNAMIC:
		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Dynamic{}).Where("id = ?", req.ItemId).Update("like_count", gorm.Expr("like_count - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Dynamic{}).Where("id = ?", req.ItemId).Update("like_count", gorm.Expr("like_count + ?", 1))
		}
		break
	case enum.COMMENT:
		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Comment{}).Where("id = ?", req.ItemId).Update("like_count", gorm.Expr("like_count - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Comment{}).Where("id = ?", req.ItemId).Update("like_count", gorm.Expr("like_count + ?", 1))
		}
		break
	default:
		break
	}
	return nil
}
