// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto_defs/common/common.proto

package common

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AssetService_RegisterAsset_FullMethodName = "/AssetService/RegisterAsset"
	AssetService_SendResponse_FullMethodName  = "/AssetService/SendResponse"
	AssetService_CheckIn_FullMethodName       = "/AssetService/CheckIn"
)

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetServiceClient interface {
	RegisterAsset(ctx context.Context, in *AssetRegistration, opts ...grpc.CallOption) (*RegistrationConfirmation, error)
	SendResponse(ctx context.Context, in *AssetResponse, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CheckIn(ctx context.Context, in *AssetResponse, opts ...grpc.CallOption) (*ServerOrder, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) RegisterAsset(ctx context.Context, in *AssetRegistration, opts ...grpc.CallOption) (*RegistrationConfirmation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegistrationConfirmation)
	err := c.cc.Invoke(ctx, AssetService_RegisterAsset_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) SendResponse(ctx context.Context, in *AssetResponse, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AssetService_SendResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) CheckIn(ctx context.Context, in *AssetResponse, opts ...grpc.CallOption) (*ServerOrder, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerOrder)
	err := c.cc.Invoke(ctx, AssetService_CheckIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
// All implementations must embed UnimplementedAssetServiceServer
// for forward compatibility.
type AssetServiceServer interface {
	RegisterAsset(context.Context, *AssetRegistration) (*RegistrationConfirmation, error)
	SendResponse(context.Context, *AssetResponse) (*emptypb.Empty, error)
	CheckIn(context.Context, *AssetResponse) (*ServerOrder, error)
	mustEmbedUnimplementedAssetServiceServer()
}

// UnimplementedAssetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAssetServiceServer struct{}

func (UnimplementedAssetServiceServer) RegisterAsset(context.Context, *AssetRegistration) (*RegistrationConfirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAsset not implemented")
}
func (UnimplementedAssetServiceServer) SendResponse(context.Context, *AssetResponse) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResponse not implemented")
}
func (UnimplementedAssetServiceServer) CheckIn(context.Context, *AssetResponse) (*ServerOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIn not implemented")
}
func (UnimplementedAssetServiceServer) mustEmbedUnimplementedAssetServiceServer() {}
func (UnimplementedAssetServiceServer) testEmbeddedByValue()                      {}

// UnsafeAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServiceServer will
// result in compilation errors.
type UnsafeAssetServiceServer interface {
	mustEmbedUnimplementedAssetServiceServer()
}

func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	// If the following call pancis, it indicates UnimplementedAssetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AssetService_ServiceDesc, srv)
}

func _AssetService_RegisterAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).RegisterAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_RegisterAsset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).RegisterAsset(ctx, req.(*AssetRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_SendResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).SendResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_SendResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).SendResponse(ctx, req.(*AssetResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_CheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssetResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).CheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetService_CheckIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).CheckIn(ctx, req.(*AssetResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetService_ServiceDesc is the grpc.ServiceDesc for AssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAsset",
			Handler:    _AssetService_RegisterAsset_Handler,
		},
		{
			MethodName: "SendResponse",
			Handler:    _AssetService_SendResponse_Handler,
		},
		{
			MethodName: "CheckIn",
			Handler:    _AssetService_CheckIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_defs/common/common.proto",
}
