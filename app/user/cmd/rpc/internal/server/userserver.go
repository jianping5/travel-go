// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"travel/app/user/cmd/rpc/internal/logic"
	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UserServer) SearchUser(ctx context.Context, in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	l := logic.NewSearchUserLogic(ctx, s.svcCtx)
	return l.SearchUser(in)
}

func (s *UserServer) GetUserIds(ctx context.Context, in *pb.GetUserIdsReq) (*pb.GetUserIdsResp, error) {
	l := logic.NewGetUserIdsLogic(ctx, s.svcCtx)
	return l.GetUserIds(in)
}

func (s *UserServer) GetFans(ctx context.Context, in *pb.GetFansReq) (*pb.GetFansResp, error) {
	l := logic.NewGetFansLogic(ctx, s.svcCtx)
	return l.GetFans(in)
}
