// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: bot.proto

package raketapb

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
	RaketaService_SignUp_FullMethodName       = "/raketa.RaketaService/SignUp"
	RaketaService_CreateTask_FullMethodName   = "/raketa.RaketaService/CreateTask"
	RaketaService_DeleteTask_FullMethodName   = "/raketa.RaketaService/DeleteTask"
	RaketaService_AssignWorker_FullMethodName = "/raketa.RaketaService/AssignWorker"
	RaketaService_CloseTask_FullMethodName    = "/raketa.RaketaService/CloseTask"
	RaketaService_GetOpenTasks_FullMethodName = "/raketa.RaketaService/GetOpenTasks"
)

// RaketaServiceClient is the client API for RaketaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaketaServiceClient interface {
	// SignUp user
	SignUp(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	CreateTask(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	DeleteTask(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	AssignWorker(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*AssignResponse, error)
	CloseTask(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
	GetOpenTasks(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type raketaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaketaServiceClient(cc grpc.ClientConnInterface) RaketaServiceClient {
	return &raketaServiceClient{cc}
}

func (c *raketaServiceClient) SignUp(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, RaketaService_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raketaServiceClient) CreateTask(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, RaketaService_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raketaServiceClient) DeleteTask(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, RaketaService_DeleteTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raketaServiceClient) AssignWorker(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*AssignResponse, error) {
	out := new(AssignResponse)
	err := c.cc.Invoke(ctx, RaketaService_AssignWorker_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raketaServiceClient) CloseTask(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, RaketaService_CloseTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raketaServiceClient) GetOpenTasks(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, RaketaService_GetOpenTasks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaketaServiceServer is the server API for RaketaService service.
// All implementations must embed UnimplementedRaketaServiceServer
// for forward compatibility
type RaketaServiceServer interface {
	// SignUp user
	SignUp(context.Context, *RegisterRequest) (*RegisterResponse, error)
	CreateTask(context.Context, *CreateRequest) (*CreateResponse, error)
	DeleteTask(context.Context, *DeleteRequest) (*DeleteResponse, error)
	AssignWorker(context.Context, *AssignRequest) (*AssignResponse, error)
	CloseTask(context.Context, *CloseRequest) (*CloseResponse, error)
	GetOpenTasks(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedRaketaServiceServer()
}

// UnimplementedRaketaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRaketaServiceServer struct {
}

func (UnimplementedRaketaServiceServer) SignUp(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedRaketaServiceServer) CreateTask(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedRaketaServiceServer) DeleteTask(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedRaketaServiceServer) AssignWorker(context.Context, *AssignRequest) (*AssignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignWorker not implemented")
}
func (UnimplementedRaketaServiceServer) CloseTask(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTask not implemented")
}
func (UnimplementedRaketaServiceServer) GetOpenTasks(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenTasks not implemented")
}
func (UnimplementedRaketaServiceServer) mustEmbedUnimplementedRaketaServiceServer() {}

// UnsafeRaketaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaketaServiceServer will
// result in compilation errors.
type UnsafeRaketaServiceServer interface {
	mustEmbedUnimplementedRaketaServiceServer()
}

func RegisterRaketaServiceServer(s grpc.ServiceRegistrar, srv RaketaServiceServer) {
	s.RegisterService(&RaketaService_ServiceDesc, srv)
}

func _RaketaService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaketaServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaketaService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaketaServiceServer).SignUp(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaketaService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaketaServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaketaService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaketaServiceServer).CreateTask(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaketaService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaketaServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaketaService_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaketaServiceServer).DeleteTask(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaketaService_AssignWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaketaServiceServer).AssignWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaketaService_AssignWorker_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaketaServiceServer).AssignWorker(ctx, req.(*AssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaketaService_CloseTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaketaServiceServer).CloseTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaketaService_CloseTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaketaServiceServer).CloseTask(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaketaService_GetOpenTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaketaServiceServer).GetOpenTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaketaService_GetOpenTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaketaServiceServer).GetOpenTasks(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RaketaService_ServiceDesc is the grpc.ServiceDesc for RaketaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaketaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raketa.RaketaService",
	HandlerType: (*RaketaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _RaketaService_SignUp_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _RaketaService_CreateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _RaketaService_DeleteTask_Handler,
		},
		{
			MethodName: "AssignWorker",
			Handler:    _RaketaService_AssignWorker_Handler,
		},
		{
			MethodName: "CloseTask",
			Handler:    _RaketaService_CloseTask_Handler,
		},
		{
			MethodName: "GetOpenTasks",
			Handler:    _RaketaService_GetOpenTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot.proto",
}
