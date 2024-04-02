// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: data.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Data_ContentSimilar_FullMethodName   = "/pb.Data/ContentSimilar"
	Data_UserLikeContent_FullMethodName  = "/pb.Data/UserLikeContent"
	Data_ContentTagCreate_FullMethodName = "/pb.Data/ContentTagCreate"
)

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	ContentSimilar(ctx context.Context, in *ContentSimilarReq, opts ...grpc.CallOption) (*ContentSimilarResp, error)
	UserLikeContent(ctx context.Context, in *UserLikeContentReq, opts ...grpc.CallOption) (*UserLikeContentResp, error)
	ContentTagCreate(ctx context.Context, in *ContentTagCreateReq, opts ...grpc.CallOption) (*ContentTagCreateResp, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) ContentSimilar(ctx context.Context, in *ContentSimilarReq, opts ...grpc.CallOption) (*ContentSimilarResp, error) {
	out := new(ContentSimilarResp)
	err := c.cc.Invoke(ctx, Data_ContentSimilar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) UserLikeContent(ctx context.Context, in *UserLikeContentReq, opts ...grpc.CallOption) (*UserLikeContentResp, error) {
	out := new(UserLikeContentResp)
	err := c.cc.Invoke(ctx, Data_UserLikeContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) ContentTagCreate(ctx context.Context, in *ContentTagCreateReq, opts ...grpc.CallOption) (*ContentTagCreateResp, error) {
	out := new(ContentTagCreateResp)
	err := c.cc.Invoke(ctx, Data_ContentTagCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
// All implementations must embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	ContentSimilar(context.Context, *ContentSimilarReq) (*ContentSimilarResp, error)
	UserLikeContent(context.Context, *UserLikeContentReq) (*UserLikeContentResp, error)
	ContentTagCreate(context.Context, *ContentTagCreateReq) (*ContentTagCreateResp, error)
	mustEmbedUnimplementedDataServer()
}

// UnimplementedDataServer must be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (UnimplementedDataServer) ContentSimilar(context.Context, *ContentSimilarReq) (*ContentSimilarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentSimilar not implemented")
}
func (UnimplementedDataServer) UserLikeContent(context.Context, *UserLikeContentReq) (*UserLikeContentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLikeContent not implemented")
}
func (UnimplementedDataServer) ContentTagCreate(context.Context, *ContentTagCreateReq) (*ContentTagCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentTagCreate not implemented")
}
func (UnimplementedDataServer) mustEmbedUnimplementedDataServer() {}

// UnsafeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServer will
// result in compilation errors.
type UnsafeDataServer interface {
	mustEmbedUnimplementedDataServer()
}

func RegisterDataServer(s grpc.ServiceRegistrar, srv DataServer) {
	s.RegisterService(&Data_ServiceDesc, srv)
}

func _Data_ContentSimilar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentSimilarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).ContentSimilar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_ContentSimilar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).ContentSimilar(ctx, req.(*ContentSimilarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_UserLikeContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLikeContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).UserLikeContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_UserLikeContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).UserLikeContent(ctx, req.(*UserLikeContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_ContentTagCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentTagCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).ContentTagCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_ContentTagCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).ContentTagCreate(ctx, req.(*ContentTagCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Data_ServiceDesc is the grpc.ServiceDesc for Data service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Data_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContentSimilar",
			Handler:    _Data_ContentSimilar_Handler,
		},
		{
			MethodName: "UserLikeContent",
			Handler:    _Data_UserLikeContent_Handler,
		},
		{
			MethodName: "ContentTagCreate",
			Handler:    _Data_ContentTagCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
