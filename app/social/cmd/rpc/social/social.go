// Code generated by goctl. DO NOT EDIT.
// Source: social.proto

package social

import (
	"context"

	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ContentDeleteReq    = pb.ContentDeleteReq
	ContentDeleteResp   = pb.ContentDeleteResp
	ContentSimpleReq    = pb.ContentSimpleReq
	ContentSimpleResp   = pb.ContentSimpleResp
	CopyrightDetailReq  = pb.CopyrightDetailReq
	CopyrightDetailResp = pb.CopyrightDetailResp
	MessageCreateReq    = pb.MessageCreateReq
	MessageCreateResp   = pb.MessageCreateResp

	Social interface {
		MessageCreate(ctx context.Context, in *MessageCreateReq, opts ...grpc.CallOption) (*MessageCreateResp, error)
		CopyrightDetail(ctx context.Context, in *CopyrightDetailReq, opts ...grpc.CallOption) (*CopyrightDetailResp, error)
		ContentSimple(ctx context.Context, in *ContentSimpleReq, opts ...grpc.CallOption) (*ContentSimpleResp, error)
		ContentDelete(ctx context.Context, in *ContentDeleteReq, opts ...grpc.CallOption) (*ContentDeleteResp, error)
	}

	defaultSocial struct {
		cli zrpc.Client
	}
)

func NewSocial(cli zrpc.Client) Social {
	return &defaultSocial{
		cli: cli,
	}
}

func (m *defaultSocial) MessageCreate(ctx context.Context, in *MessageCreateReq, opts ...grpc.CallOption) (*MessageCreateResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.MessageCreate(ctx, in, opts...)
}

func (m *defaultSocial) CopyrightDetail(ctx context.Context, in *CopyrightDetailReq, opts ...grpc.CallOption) (*CopyrightDetailResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.CopyrightDetail(ctx, in, opts...)
}

func (m *defaultSocial) ContentSimple(ctx context.Context, in *ContentSimpleReq, opts ...grpc.CallOption) (*ContentSimpleResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.ContentSimple(ctx, in, opts...)
}

func (m *defaultSocial) ContentDelete(ctx context.Context, in *ContentDeleteReq, opts ...grpc.CallOption) (*ContentDeleteResp, error) {
	client := pb.NewSocialClient(m.cli.Conn())
	return client.ContentDelete(ctx, in, opts...)
}
