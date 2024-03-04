package logic

import (
	"context"

	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConversationGenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConversationGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConversationGenerateLogic {
	return &ConversationGenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConversationGenerateLogic) ConversationGenerate(req *types.ConversationGenerateReq) (resp *types.ConversationListResp, err error) {
	// ask := req.Content

	// TODO：调用 AI 接口

	// 处理响应值，并返回

	return
}
