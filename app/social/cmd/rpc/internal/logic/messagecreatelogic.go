package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageCreateLogic {
	return &MessageCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageCreateLogic) MessageCreate(in *pb.MessageCreateReq) (*pb.MessageCreateResp, error) {
	var messages []model.Message
	for _, id := range in.UserIds {
		var message model.Message
		_ = copier.Copy(&message, in)
		message.UserId = id
		messages = append(messages, message)
	}

	l.svcCtx.DB.Create(&messages)

	return &pb.MessageCreateResp{}, nil
}
