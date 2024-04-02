package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageDeleteLogic {
	return &MessageDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageDeleteLogic) MessageDelete(req *types.MessageDeleteReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var userId int64
	l.svcCtx.DB.Model(&model.Message{}).Select("user_id").Where("id = ?", req.Id).Scan(&userId)
	if loginUserId != userId {
		return errors.Wrap(xerr.NewErrMsg("没有权限删除"), "没有权限删除")
	}

	// 删除
	if err := l.svcCtx.DB.Delete(&model.Message{}, "id = ?", req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	return nil
}
