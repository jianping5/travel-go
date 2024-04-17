package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCopyrightCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightCheckLogic {
	return &CopyrightCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CopyrightCheckLogic) CopyrightCheck(in *pb.CopyrightCheckReq) (*pb.CopyrightCheckResp, error) {
	// todo: add your logic here and delete this line
	var id int64
	if l.svcCtx.DB.Model(&model.Copyright{}).Select("id").Where("id = ? and user_id = ?", in.CopyrightId, in.UserId).Scan(&id); id == 0 {
		return &pb.CopyrightCheckResp{
			IsBelonged: false,
		}, nil
	}

	return &pb.CopyrightCheckResp{
		IsBelonged: true,
	}, nil
}
