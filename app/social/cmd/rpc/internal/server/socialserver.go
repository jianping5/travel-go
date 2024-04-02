// Code generated by goctl. DO NOT EDIT.
// Source: social.proto

package server

import (
	"context"

	"travel/app/social/cmd/rpc/internal/logic"
	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"
)

type SocialServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSocialServer
}

func NewSocialServer(svcCtx *svc.ServiceContext) *SocialServer {
	return &SocialServer{
		svcCtx: svcCtx,
	}
}

func (s *SocialServer) MessageCreate(ctx context.Context, in *pb.MessageCreateReq) (*pb.MessageCreateResp, error) {
	l := logic.NewMessageCreateLogic(ctx, s.svcCtx)
	return l.MessageCreate(in)
}

func (s *SocialServer) CopyrightDetail(ctx context.Context, in *pb.CopyrightDetailReq) (*pb.CopyrightDetailResp, error) {
	l := logic.NewCopyrightDetailLogic(ctx, s.svcCtx)
	return l.CopyrightDetail(in)
}

func (s *SocialServer) ContentSimple(ctx context.Context, in *pb.ContentSimpleReq) (*pb.ContentSimpleResp, error) {
	l := logic.NewContentSimpleLogic(ctx, s.svcCtx)
	return l.ContentSimple(in)
}

func (s *SocialServer) ContentDelete(ctx context.Context, in *pb.ContentDeleteReq) (*pb.ContentDeleteResp, error) {
	l := logic.NewContentDeleteLogic(ctx, s.svcCtx)
	return l.ContentDelete(in)
}
