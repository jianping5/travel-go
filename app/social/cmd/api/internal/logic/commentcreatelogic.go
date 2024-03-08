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

type CommentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentCreateLogic {
	return &CommentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentCreateLogic) CommentCreate(req *types.CommentCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	comment := &model.Comment{
		UserId:        loginUserId,
		CommentItemId: req.CommentItemId,
		ParentUserId:  req.ParentUserId,
		TopId:         req.TopId,
		Content:       req.Content,
	}
	if err := l.svcCtx.DB.Create(comment).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建失败")
	}

	// 更新对应评论量
	l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", comment.CommentItemId).Update("commentCount", gorm.Expr("commentCount + ?", 1))
	return nil
}
