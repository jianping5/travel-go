package logic

import (
	"context"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/rpc/social"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageCreateLogic {
	return &MessageCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageCreateLogic) MessageCreate(req *types.MessageCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	_, _ = l.svcCtx.SocialRpc.MessageCreate(l.ctx, &social.MessageCreateReq{
		UserIds:       req.UserIds,
		ItemType:      int32(req.ItemType),
		ItemId:        req.ItemId,
		MessageType:   int32(req.MessageType),
		MessageUserId: loginUserId,
		Content:       req.Content,
	})

	return nil
}
