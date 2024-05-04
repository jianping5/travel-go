package logic

import (
	"context"
	"travel/app/user/cmd/model"

	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFansLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansLogic {
	return &GetFansLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFansLogic) GetFans(in *pb.GetFansReq) (*pb.GetFansResp, error) {
	// todo: add your logic here and delete this line
	var userIds []int64
	l.svcCtx.DB.Model(model.Follow{}).Select("user_id").Where("follow_user_id = ?", in.UserId).Scan(&userIds)

	return &pb.GetFansResp{
		UserIds: userIds,
	}, nil
}
