// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto.proto

package __

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

// CenterPlatformClient is the client API for CenterPlatform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CenterPlatformClient interface {
	// 传递事件信息
	SendEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	// 服务端流式 RPC
	ServerStream(ctx context.Context, in *ControllRequest, opts ...grpc.CallOption) (CenterPlatform_ServerStreamClient, error)
}

type centerPlatformClient struct {
	cc grpc.ClientConnInterface
}

func NewCenterPlatformClient(cc grpc.ClientConnInterface) CenterPlatformClient {
	return &centerPlatformClient{cc}
}

func (c *centerPlatformClient) SendEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/deliver.CenterPlatform/SendEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerPlatformClient) ServerStream(ctx context.Context, in *ControllRequest, opts ...grpc.CallOption) (CenterPlatform_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CenterPlatform_ServiceDesc.Streams[0], "/deliver.CenterPlatform/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &centerPlatformServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CenterPlatform_ServerStreamClient interface {
	Recv() (*ControllResponse, error)
	grpc.ClientStream
}

type centerPlatformServerStreamClient struct {
	grpc.ClientStream
}

func (x *centerPlatformServerStreamClient) Recv() (*ControllResponse, error) {
	m := new(ControllResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CenterPlatformServer is the server API for CenterPlatform service.
// All implementations must embed UnimplementedCenterPlatformServer
// for forward compatibility
type CenterPlatformServer interface {
	// 传递事件信息
	SendEvent(context.Context, *EventRequest) (*EventResponse, error)
	// 服务端流式 RPC
	ServerStream(*ControllRequest, CenterPlatform_ServerStreamServer) error
	mustEmbedUnimplementedCenterPlatformServer()
}

// UnimplementedCenterPlatformServer must be embedded to have forward compatible implementations.
type UnimplementedCenterPlatformServer struct {
}

func (UnimplementedCenterPlatformServer) SendEvent(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEvent not implemented")
}
func (UnimplementedCenterPlatformServer) ServerStream(*ControllRequest, CenterPlatform_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedCenterPlatformServer) mustEmbedUnimplementedCenterPlatformServer() {}

// UnsafeCenterPlatformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CenterPlatformServer will
// result in compilation errors.
type UnsafeCenterPlatformServer interface {
	mustEmbedUnimplementedCenterPlatformServer()
}

func RegisterCenterPlatformServer(s grpc.ServiceRegistrar, srv CenterPlatformServer) {
	s.RegisterService(&CenterPlatform_ServiceDesc, srv)
}

func _CenterPlatform_SendEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterPlatformServer).SendEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliver.CenterPlatform/SendEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterPlatformServer).SendEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterPlatform_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ControllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CenterPlatformServer).ServerStream(m, &centerPlatformServerStreamServer{stream})
}

type CenterPlatform_ServerStreamServer interface {
	Send(*ControllResponse) error
	grpc.ServerStream
}

type centerPlatformServerStreamServer struct {
	grpc.ServerStream
}

func (x *centerPlatformServerStreamServer) Send(m *ControllResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CenterPlatform_ServiceDesc is the grpc.ServiceDesc for CenterPlatform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CenterPlatform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deliver.CenterPlatform",
	HandlerType: (*CenterPlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEvent",
			Handler:    _CenterPlatform_SendEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _CenterPlatform_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto.proto",
}
