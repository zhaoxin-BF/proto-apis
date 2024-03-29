// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: disks/v1/service.proto

package disks

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
	DisksService_CreateDisks_FullMethodName = "/disks.v1.DisksService/CreateDisks"
	DisksService_ListDisks_FullMethodName   = "/disks.v1.DisksService/ListDisks"
	DisksService_GetDisk_FullMethodName     = "/disks.v1.DisksService/GetDisk"
	DisksService_UpdateDisk_FullMethodName  = "/disks.v1.DisksService/UpdateDisk"
	DisksService_DeleteDisk_FullMethodName  = "/disks.v1.DisksService/DeleteDisk"
	DisksService_AttachDisk_FullMethodName  = "/disks.v1.DisksService/AttachDisk"
	DisksService_DetachDisk_FullMethodName  = "/disks.v1.DisksService/DetachDisk"
)

// DisksServiceClient is the client API for DisksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisksServiceClient interface {
	CreateDisks(ctx context.Context, in *CreateDisksRequest, opts ...grpc.CallOption) (*CreateDisksResponse, error)
	ListDisks(ctx context.Context, in *ListDisksRequest, opts ...grpc.CallOption) (*ListDisksResponse, error)
	GetDisk(ctx context.Context, in *GetDiskRequest, opts ...grpc.CallOption) (*GetDiskResponse, error)
	UpdateDisk(ctx context.Context, in *UpdateDiskRequest, opts ...grpc.CallOption) (*UpdateDiskResponse, error)
	DeleteDisk(ctx context.Context, in *DeleteDisksRequest, opts ...grpc.CallOption) (*DeleteDisksResponse, error)
	AttachDisk(ctx context.Context, in *AttachDiskRequest, opts ...grpc.CallOption) (*AttachDiskResponse, error)
	DetachDisk(ctx context.Context, in *DetachDiskRequest, opts ...grpc.CallOption) (*DetachDiskResponse, error)
}

type disksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDisksServiceClient(cc grpc.ClientConnInterface) DisksServiceClient {
	return &disksServiceClient{cc}
}

func (c *disksServiceClient) CreateDisks(ctx context.Context, in *CreateDisksRequest, opts ...grpc.CallOption) (*CreateDisksResponse, error) {
	out := new(CreateDisksResponse)
	err := c.cc.Invoke(ctx, DisksService_CreateDisks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksServiceClient) ListDisks(ctx context.Context, in *ListDisksRequest, opts ...grpc.CallOption) (*ListDisksResponse, error) {
	out := new(ListDisksResponse)
	err := c.cc.Invoke(ctx, DisksService_ListDisks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksServiceClient) GetDisk(ctx context.Context, in *GetDiskRequest, opts ...grpc.CallOption) (*GetDiskResponse, error) {
	out := new(GetDiskResponse)
	err := c.cc.Invoke(ctx, DisksService_GetDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksServiceClient) UpdateDisk(ctx context.Context, in *UpdateDiskRequest, opts ...grpc.CallOption) (*UpdateDiskResponse, error) {
	out := new(UpdateDiskResponse)
	err := c.cc.Invoke(ctx, DisksService_UpdateDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksServiceClient) DeleteDisk(ctx context.Context, in *DeleteDisksRequest, opts ...grpc.CallOption) (*DeleteDisksResponse, error) {
	out := new(DeleteDisksResponse)
	err := c.cc.Invoke(ctx, DisksService_DeleteDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksServiceClient) AttachDisk(ctx context.Context, in *AttachDiskRequest, opts ...grpc.CallOption) (*AttachDiskResponse, error) {
	out := new(AttachDiskResponse)
	err := c.cc.Invoke(ctx, DisksService_AttachDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksServiceClient) DetachDisk(ctx context.Context, in *DetachDiskRequest, opts ...grpc.CallOption) (*DetachDiskResponse, error) {
	out := new(DetachDiskResponse)
	err := c.cc.Invoke(ctx, DisksService_DetachDisk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisksServiceServer is the server API for DisksService service.
// All implementations must embed UnimplementedDisksServiceServer
// for forward compatibility
type DisksServiceServer interface {
	CreateDisks(context.Context, *CreateDisksRequest) (*CreateDisksResponse, error)
	ListDisks(context.Context, *ListDisksRequest) (*ListDisksResponse, error)
	GetDisk(context.Context, *GetDiskRequest) (*GetDiskResponse, error)
	UpdateDisk(context.Context, *UpdateDiskRequest) (*UpdateDiskResponse, error)
	DeleteDisk(context.Context, *DeleteDisksRequest) (*DeleteDisksResponse, error)
	AttachDisk(context.Context, *AttachDiskRequest) (*AttachDiskResponse, error)
	DetachDisk(context.Context, *DetachDiskRequest) (*DetachDiskResponse, error)
	mustEmbedUnimplementedDisksServiceServer()
}

// UnimplementedDisksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDisksServiceServer struct {
}

func (UnimplementedDisksServiceServer) CreateDisks(context.Context, *CreateDisksRequest) (*CreateDisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDisks not implemented")
}
func (UnimplementedDisksServiceServer) ListDisks(context.Context, *ListDisksRequest) (*ListDisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDisks not implemented")
}
func (UnimplementedDisksServiceServer) GetDisk(context.Context, *GetDiskRequest) (*GetDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisk not implemented")
}
func (UnimplementedDisksServiceServer) UpdateDisk(context.Context, *UpdateDiskRequest) (*UpdateDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDisk not implemented")
}
func (UnimplementedDisksServiceServer) DeleteDisk(context.Context, *DeleteDisksRequest) (*DeleteDisksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDisk not implemented")
}
func (UnimplementedDisksServiceServer) AttachDisk(context.Context, *AttachDiskRequest) (*AttachDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachDisk not implemented")
}
func (UnimplementedDisksServiceServer) DetachDisk(context.Context, *DetachDiskRequest) (*DetachDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachDisk not implemented")
}
func (UnimplementedDisksServiceServer) mustEmbedUnimplementedDisksServiceServer() {}

// UnsafeDisksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisksServiceServer will
// result in compilation errors.
type UnsafeDisksServiceServer interface {
	mustEmbedUnimplementedDisksServiceServer()
}

func RegisterDisksServiceServer(s grpc.ServiceRegistrar, srv DisksServiceServer) {
	s.RegisterService(&DisksService_ServiceDesc, srv)
}

func _DisksService_CreateDisks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).CreateDisks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_CreateDisks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).CreateDisks(ctx, req.(*CreateDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisksService_ListDisks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).ListDisks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_ListDisks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).ListDisks(ctx, req.(*ListDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisksService_GetDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).GetDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_GetDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).GetDisk(ctx, req.(*GetDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisksService_UpdateDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).UpdateDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_UpdateDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).UpdateDisk(ctx, req.(*UpdateDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisksService_DeleteDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDisksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).DeleteDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_DeleteDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).DeleteDisk(ctx, req.(*DeleteDisksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisksService_AttachDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).AttachDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_AttachDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).AttachDisk(ctx, req.(*AttachDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisksService_DetachDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisksServiceServer).DetachDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisksService_DetachDisk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisksServiceServer).DetachDisk(ctx, req.(*DetachDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DisksService_ServiceDesc is the grpc.ServiceDesc for DisksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DisksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "disks.v1.DisksService",
	HandlerType: (*DisksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDisks",
			Handler:    _DisksService_CreateDisks_Handler,
		},
		{
			MethodName: "ListDisks",
			Handler:    _DisksService_ListDisks_Handler,
		},
		{
			MethodName: "GetDisk",
			Handler:    _DisksService_GetDisk_Handler,
		},
		{
			MethodName: "UpdateDisk",
			Handler:    _DisksService_UpdateDisk_Handler,
		},
		{
			MethodName: "DeleteDisk",
			Handler:    _DisksService_DeleteDisk_Handler,
		},
		{
			MethodName: "AttachDisk",
			Handler:    _DisksService_AttachDisk_Handler,
		},
		{
			MethodName: "DetachDisk",
			Handler:    _DisksService_DetachDisk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "disks/v1/service.proto",
}
