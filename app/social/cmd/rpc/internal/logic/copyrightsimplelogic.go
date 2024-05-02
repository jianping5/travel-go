package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightSimpleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCopyrightSimpleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightSimpleLogic {
	return &CopyrightSimpleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CopyrightSimpleLogic) CopyrightSimple(in *pb.CopyrightSimpleReq) (*pb.CopyrightSimpleResp, error) {
	var accountAddress string
	l.svcCtx.DB.Model(model.Copyright{}).Select("account_address").Where("id = ?", in.CopyrightId).Scan(&accountAddress)

	return &pb.CopyrightSimpleResp{
		AccountAddress: accountAddress,
	}, nil
}
