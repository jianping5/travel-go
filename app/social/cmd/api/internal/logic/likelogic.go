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
	l.svcCtx.DB.Create(like)
	// 更新对应点赞量
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Create(like)

		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Article{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Article{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount + ?", 1))
		}
		break
	case enum.VIDEO:
		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Video{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Video{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount + ?", 1))
		}
		break
	case enum.DYNAMIC:
		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Dynamic{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Dynamic{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount + ?", 1))
		}
		break
	case enum.COMMENT:
		if req.LikedStatus {
			l.svcCtx.DB.Model(&model.Comment{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount - ?", 1))
		} else {
			l.svcCtx.DB.Model(&model.Comment{}).Where("id = ?", req.ItemId).Update("likeCount", gorm.Expr("likeCount + ?", 1))
		}
		break
	default:
		break
	}
	return nil
}
