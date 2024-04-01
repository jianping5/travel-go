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

type CommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentDeleteLogic {
	return &CommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentDeleteLogic) CommentDelete(req *types.CommentDeleteReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// 判断是否有权限
	var comment model.Comment
	l.svcCtx.DB.Model(&model.Comment{}).Select("user_id", "comment_item_type", "comment_item_id").Scan(&comment)
	if loginUserId != comment.UserId {
		return errors.Wrap(xerr.NewErrMsg("没有权限删除"), "没有权限删除")
	}
	// 删除
	if err := l.svcCtx.DB.Delete(&model.Comment{}, "id = ?", req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	// 更新对应评论量
	switch enum.ItemType(comment.CommentItemType) {
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", comment.CommentItemId).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		break
	case enum.DYNAMIC:
		l.svcCtx.DB.Model(&model.Dynamic{}).Where("id = ?", comment.CommentItemId).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		break
	default:
		break
	}

	return nil
}
