package logic

import (
	"context"
	"travel/app/intelligence/cmd/model"
	"travel/common/ctxdata"

	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConversationCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConversationCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConversationCreateLogic {
	return &ConversationCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConversationCreateLogic) ConversationCreate(req *types.ConversationCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	conversation := &model.Conversation{
		UserId:      loginUserId,
		Content:     req.Content,
		IsGenerated: req.IsGenerated,
	}
	l.svcCtx.DB.Create(conversation)

	return nil
}
