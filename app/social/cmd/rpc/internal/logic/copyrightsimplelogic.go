package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/app/social/cmd/rpc/social"

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
	var copyrightSimple social.CopyrightSimpleResp
	l.svcCtx.DB.Model(model.Copyright{}).Select("account_address, token_id").Where("id = ?", in.CopyrightId).Scan(&copyrightSimple)

	return &pb.CopyrightSimpleResp{
		AccountAddress: copyrightSimple.AccountAddress,
		TokenId:        copyrightSimple.TokenId,
	}, nil
}
