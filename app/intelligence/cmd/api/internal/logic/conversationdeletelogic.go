package logic

import (
	"context"
	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/model"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConversationDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConversationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConversationDeleteLogic {
	return &ConversationDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConversationDeleteLogic) ConversationDelete() error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)

	l.svcCtx.DB.Delete(&model.Conversation{}, "userId = ?", loginUserId)

	return nil
}
