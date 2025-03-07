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

func (s *SocialServer) CopyrightCheck(ctx context.Context, in *pb.CopyrightCheckReq) (*pb.CopyrightCheckResp, error) {
	l := logic.NewCopyrightCheckLogic(ctx, s.svcCtx)
	return l.CopyrightCheck(in)
}

func (s *SocialServer) ContentUpdate(ctx context.Context, in *pb.ContentUpdateReq) (*pb.ContentUpdateResp, error) {
	l := logic.NewContentUpdateLogic(ctx, s.svcCtx)
	return l.ContentUpdate(in)
}

func (s *SocialServer) CopyrightUpdate(ctx context.Context, in *pb.CopyrightUpdateReq) (*pb.CopyrightUpdateResp, error) {
	l := logic.NewCopyrightUpdateLogic(ctx, s.svcCtx)
	return l.CopyrightUpdate(in)
}

func (s *SocialServer) CopyrightSimple(ctx context.Context, in *pb.CopyrightSimpleReq) (*pb.CopyrightSimpleResp, error) {
	l := logic.NewCopyrightSimpleLogic(ctx, s.svcCtx)
	return l.CopyrightSimple(in)
}
