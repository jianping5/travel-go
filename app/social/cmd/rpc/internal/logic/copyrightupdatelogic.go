package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCopyrightUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightUpdateLogic {
	return &CopyrightUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CopyrightUpdateLogic) CopyrightUpdate(in *pb.CopyrightUpdateReq) (*pb.CopyrightUpdateResp, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.DB.Model(&model.Copyright{}).Where("id = ?", in.CopyrightId).Updates(map[string]interface{}{
		"user_id":         in.UserId,
		"account_address": in.AccountAddress,
	})

	return &pb.CopyrightUpdateResp{}, nil
}
