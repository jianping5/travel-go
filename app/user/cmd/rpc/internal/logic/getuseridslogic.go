package logic

import (
	"context"
	"travel/app/user/cmd/model"

	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIdsLogic {
	return &GetUserIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserIdsLogic) GetUserIds(in *pb.GetUserIdsReq) (*pb.GetUserIdsResp, error) {
	// todo: add your logic here and delete this line
	var userIds []int64
	l.svcCtx.DB.Model(&model.User{}).Select("id").Scan(&userIds)

	return &pb.GetUserIdsResp{
		UserIds: userIds,
	}, nil
}
