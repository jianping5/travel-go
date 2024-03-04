package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageUpdateLogic {
	return &MessageUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageUpdateLogic) MessageUpdate(req *types.MessageUpdateReq) error {
	l.svcCtx.DB.Model(&model.Message{}).Updates(req)

	return nil
}
