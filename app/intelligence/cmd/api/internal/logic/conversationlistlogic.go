package logic

import (
	"context"
	"travel/app/intelligence/cmd/model"
	"travel/common/ctxdata"

	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConversationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConversationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConversationListLogic {
	return &ConversationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConversationListLogic) ConversationList() (resp *types.ConversationListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)

	var conversations []types.ConversationView
	l.svcCtx.DB.Model(&model.Conversation{}).Where("user_id = ?", loginUserId).Scan(&conversations)

	return &types.ConversationListResp{
		List: conversations,
	}, nil
}
