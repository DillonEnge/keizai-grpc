// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: keizai_grpc.proto

package keizai_grpc

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

// KeizaiGrpcClient is the client API for KeizaiGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeizaiGrpcClient interface {
	GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (KeizaiGrpc_GetPositionClient, error)
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (KeizaiGrpc_GetMessageClient, error)
	UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*Empty, error)
	GetEntityIds(ctx context.Context, in *Empty, opts ...grpc.CallOption) (KeizaiGrpc_GetEntityIdsClient, error)
	AddEntity(ctx context.Context, in *AddEntityRequest, opts ...grpc.CallOption) (*Empty, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error)
}

type keizaiGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewKeizaiGrpcClient(cc grpc.ClientConnInterface) KeizaiGrpcClient {
	return &keizaiGrpcClient{cc}
}

func (c *keizaiGrpcClient) GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (KeizaiGrpc_GetPositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &KeizaiGrpc_ServiceDesc.Streams[0], "/KeizaiGrpc/GetPosition", opts...)
	if err != nil {
		return nil, err
	}
	x := &keizaiGrpcGetPositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KeizaiGrpc_GetPositionClient interface {
	Recv() (*GetPositionResponse, error)
	grpc.ClientStream
}

type keizaiGrpcGetPositionClient struct {
	grpc.ClientStream
}

func (x *keizaiGrpcGetPositionClient) Recv() (*GetPositionResponse, error) {
	m := new(GetPositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keizaiGrpcClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (KeizaiGrpc_GetMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &KeizaiGrpc_ServiceDesc.Streams[1], "/KeizaiGrpc/GetMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &keizaiGrpcGetMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KeizaiGrpc_GetMessageClient interface {
	Recv() (*GetMessageResponse, error)
	grpc.ClientStream
}

type keizaiGrpcGetMessageClient struct {
	grpc.ClientStream
}

func (x *keizaiGrpcGetMessageClient) Recv() (*GetMessageResponse, error) {
	m := new(GetMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keizaiGrpcClient) UpdatePosition(ctx context.Context, in *UpdatePositionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/KeizaiGrpc/UpdatePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keizaiGrpcClient) GetEntityIds(ctx context.Context, in *Empty, opts ...grpc.CallOption) (KeizaiGrpc_GetEntityIdsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KeizaiGrpc_ServiceDesc.Streams[2], "/KeizaiGrpc/GetEntityIds", opts...)
	if err != nil {
		return nil, err
	}
	x := &keizaiGrpcGetEntityIdsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KeizaiGrpc_GetEntityIdsClient interface {
	Recv() (*GetEntityIdsResponse, error)
	grpc.ClientStream
}

type keizaiGrpcGetEntityIdsClient struct {
	grpc.ClientStream
}

func (x *keizaiGrpcGetEntityIdsClient) Recv() (*GetEntityIdsResponse, error) {
	m := new(GetEntityIdsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keizaiGrpcClient) AddEntity(ctx context.Context, in *AddEntityRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/KeizaiGrpc/AddEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keizaiGrpcClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/KeizaiGrpc/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeizaiGrpcServer is the server API for KeizaiGrpc service.
// All implementations must embed UnimplementedKeizaiGrpcServer
// for forward compatibility
type KeizaiGrpcServer interface {
	GetPosition(*GetPositionRequest, KeizaiGrpc_GetPositionServer) error
	GetMessage(*GetMessageRequest, KeizaiGrpc_GetMessageServer) error
	UpdatePosition(context.Context, *UpdatePositionRequest) (*Empty, error)
	GetEntityIds(*Empty, KeizaiGrpc_GetEntityIdsServer) error
	AddEntity(context.Context, *AddEntityRequest) (*Empty, error)
	SendMessage(context.Context, *SendMessageRequest) (*Empty, error)
	mustEmbedUnimplementedKeizaiGrpcServer()
}

// UnimplementedKeizaiGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedKeizaiGrpcServer struct {
}

func (UnimplementedKeizaiGrpcServer) GetPosition(*GetPositionRequest, KeizaiGrpc_GetPositionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPosition not implemented")
}
func (UnimplementedKeizaiGrpcServer) GetMessage(*GetMessageRequest, KeizaiGrpc_GetMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedKeizaiGrpcServer) UpdatePosition(context.Context, *UpdatePositionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosition not implemented")
}
func (UnimplementedKeizaiGrpcServer) GetEntityIds(*Empty, KeizaiGrpc_GetEntityIdsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEntityIds not implemented")
}
func (UnimplementedKeizaiGrpcServer) AddEntity(context.Context, *AddEntityRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEntity not implemented")
}
func (UnimplementedKeizaiGrpcServer) SendMessage(context.Context, *SendMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedKeizaiGrpcServer) mustEmbedUnimplementedKeizaiGrpcServer() {}

// UnsafeKeizaiGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeizaiGrpcServer will
// result in compilation errors.
type UnsafeKeizaiGrpcServer interface {
	mustEmbedUnimplementedKeizaiGrpcServer()
}

func RegisterKeizaiGrpcServer(s grpc.ServiceRegistrar, srv KeizaiGrpcServer) {
	s.RegisterService(&KeizaiGrpc_ServiceDesc, srv)
}

func _KeizaiGrpc_GetPosition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeizaiGrpcServer).GetPosition(m, &keizaiGrpcGetPositionServer{stream})
}

type KeizaiGrpc_GetPositionServer interface {
	Send(*GetPositionResponse) error
	grpc.ServerStream
}

type keizaiGrpcGetPositionServer struct {
	grpc.ServerStream
}

func (x *keizaiGrpcGetPositionServer) Send(m *GetPositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KeizaiGrpc_GetMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeizaiGrpcServer).GetMessage(m, &keizaiGrpcGetMessageServer{stream})
}

type KeizaiGrpc_GetMessageServer interface {
	Send(*GetMessageResponse) error
	grpc.ServerStream
}

type keizaiGrpcGetMessageServer struct {
	grpc.ServerStream
}

func (x *keizaiGrpcGetMessageServer) Send(m *GetMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KeizaiGrpc_UpdatePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeizaiGrpcServer).UpdatePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeizaiGrpc/UpdatePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeizaiGrpcServer).UpdatePosition(ctx, req.(*UpdatePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeizaiGrpc_GetEntityIds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KeizaiGrpcServer).GetEntityIds(m, &keizaiGrpcGetEntityIdsServer{stream})
}

type KeizaiGrpc_GetEntityIdsServer interface {
	Send(*GetEntityIdsResponse) error
	grpc.ServerStream
}

type keizaiGrpcGetEntityIdsServer struct {
	grpc.ServerStream
}

func (x *keizaiGrpcGetEntityIdsServer) Send(m *GetEntityIdsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KeizaiGrpc_AddEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeizaiGrpcServer).AddEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeizaiGrpc/AddEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeizaiGrpcServer).AddEntity(ctx, req.(*AddEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeizaiGrpc_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeizaiGrpcServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeizaiGrpc/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeizaiGrpcServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeizaiGrpc_ServiceDesc is the grpc.ServiceDesc for KeizaiGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeizaiGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KeizaiGrpc",
	HandlerType: (*KeizaiGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePosition",
			Handler:    _KeizaiGrpc_UpdatePosition_Handler,
		},
		{
			MethodName: "AddEntity",
			Handler:    _KeizaiGrpc_AddEntity_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _KeizaiGrpc_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPosition",
			Handler:       _KeizaiGrpc_GetPosition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMessage",
			Handler:       _KeizaiGrpc_GetMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetEntityIds",
			Handler:       _KeizaiGrpc_GetEntityIds_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "keizai_grpc.proto",
}
