// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: v1/database_group_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DatabaseGroupService_ListDatabaseGroups_FullMethodName  = "/bytebase.v1.DatabaseGroupService/ListDatabaseGroups"
	DatabaseGroupService_GetDatabaseGroup_FullMethodName    = "/bytebase.v1.DatabaseGroupService/GetDatabaseGroup"
	DatabaseGroupService_CreateDatabaseGroup_FullMethodName = "/bytebase.v1.DatabaseGroupService/CreateDatabaseGroup"
	DatabaseGroupService_UpdateDatabaseGroup_FullMethodName = "/bytebase.v1.DatabaseGroupService/UpdateDatabaseGroup"
	DatabaseGroupService_DeleteDatabaseGroup_FullMethodName = "/bytebase.v1.DatabaseGroupService/DeleteDatabaseGroup"
)

// DatabaseGroupServiceClient is the client API for DatabaseGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseGroupServiceClient interface {
	ListDatabaseGroups(ctx context.Context, in *ListDatabaseGroupsRequest, opts ...grpc.CallOption) (*ListDatabaseGroupsResponse, error)
	GetDatabaseGroup(ctx context.Context, in *GetDatabaseGroupRequest, opts ...grpc.CallOption) (*DatabaseGroup, error)
	CreateDatabaseGroup(ctx context.Context, in *CreateDatabaseGroupRequest, opts ...grpc.CallOption) (*DatabaseGroup, error)
	UpdateDatabaseGroup(ctx context.Context, in *UpdateDatabaseGroupRequest, opts ...grpc.CallOption) (*DatabaseGroup, error)
	DeleteDatabaseGroup(ctx context.Context, in *DeleteDatabaseGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type databaseGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseGroupServiceClient(cc grpc.ClientConnInterface) DatabaseGroupServiceClient {
	return &databaseGroupServiceClient{cc}
}

func (c *databaseGroupServiceClient) ListDatabaseGroups(ctx context.Context, in *ListDatabaseGroupsRequest, opts ...grpc.CallOption) (*ListDatabaseGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDatabaseGroupsResponse)
	err := c.cc.Invoke(ctx, DatabaseGroupService_ListDatabaseGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseGroupServiceClient) GetDatabaseGroup(ctx context.Context, in *GetDatabaseGroupRequest, opts ...grpc.CallOption) (*DatabaseGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DatabaseGroup)
	err := c.cc.Invoke(ctx, DatabaseGroupService_GetDatabaseGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseGroupServiceClient) CreateDatabaseGroup(ctx context.Context, in *CreateDatabaseGroupRequest, opts ...grpc.CallOption) (*DatabaseGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DatabaseGroup)
	err := c.cc.Invoke(ctx, DatabaseGroupService_CreateDatabaseGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseGroupServiceClient) UpdateDatabaseGroup(ctx context.Context, in *UpdateDatabaseGroupRequest, opts ...grpc.CallOption) (*DatabaseGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DatabaseGroup)
	err := c.cc.Invoke(ctx, DatabaseGroupService_UpdateDatabaseGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseGroupServiceClient) DeleteDatabaseGroup(ctx context.Context, in *DeleteDatabaseGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DatabaseGroupService_DeleteDatabaseGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseGroupServiceServer is the server API for DatabaseGroupService service.
// All implementations must embed UnimplementedDatabaseGroupServiceServer
// for forward compatibility
type DatabaseGroupServiceServer interface {
	ListDatabaseGroups(context.Context, *ListDatabaseGroupsRequest) (*ListDatabaseGroupsResponse, error)
	GetDatabaseGroup(context.Context, *GetDatabaseGroupRequest) (*DatabaseGroup, error)
	CreateDatabaseGroup(context.Context, *CreateDatabaseGroupRequest) (*DatabaseGroup, error)
	UpdateDatabaseGroup(context.Context, *UpdateDatabaseGroupRequest) (*DatabaseGroup, error)
	DeleteDatabaseGroup(context.Context, *DeleteDatabaseGroupRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDatabaseGroupServiceServer()
}

// UnimplementedDatabaseGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseGroupServiceServer struct {
}

func (UnimplementedDatabaseGroupServiceServer) ListDatabaseGroups(context.Context, *ListDatabaseGroupsRequest) (*ListDatabaseGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatabaseGroups not implemented")
}
func (UnimplementedDatabaseGroupServiceServer) GetDatabaseGroup(context.Context, *GetDatabaseGroupRequest) (*DatabaseGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatabaseGroup not implemented")
}
func (UnimplementedDatabaseGroupServiceServer) CreateDatabaseGroup(context.Context, *CreateDatabaseGroupRequest) (*DatabaseGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabaseGroup not implemented")
}
func (UnimplementedDatabaseGroupServiceServer) UpdateDatabaseGroup(context.Context, *UpdateDatabaseGroupRequest) (*DatabaseGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatabaseGroup not implemented")
}
func (UnimplementedDatabaseGroupServiceServer) DeleteDatabaseGroup(context.Context, *DeleteDatabaseGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatabaseGroup not implemented")
}
func (UnimplementedDatabaseGroupServiceServer) mustEmbedUnimplementedDatabaseGroupServiceServer() {}

// UnsafeDatabaseGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseGroupServiceServer will
// result in compilation errors.
type UnsafeDatabaseGroupServiceServer interface {
	mustEmbedUnimplementedDatabaseGroupServiceServer()
}

func RegisterDatabaseGroupServiceServer(s grpc.ServiceRegistrar, srv DatabaseGroupServiceServer) {
	s.RegisterService(&DatabaseGroupService_ServiceDesc, srv)
}

func _DatabaseGroupService_ListDatabaseGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabaseGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseGroupServiceServer).ListDatabaseGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseGroupService_ListDatabaseGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseGroupServiceServer).ListDatabaseGroups(ctx, req.(*ListDatabaseGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseGroupService_GetDatabaseGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseGroupServiceServer).GetDatabaseGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseGroupService_GetDatabaseGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseGroupServiceServer).GetDatabaseGroup(ctx, req.(*GetDatabaseGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseGroupService_CreateDatabaseGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseGroupServiceServer).CreateDatabaseGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseGroupService_CreateDatabaseGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseGroupServiceServer).CreateDatabaseGroup(ctx, req.(*CreateDatabaseGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseGroupService_UpdateDatabaseGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatabaseGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseGroupServiceServer).UpdateDatabaseGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseGroupService_UpdateDatabaseGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseGroupServiceServer).UpdateDatabaseGroup(ctx, req.(*UpdateDatabaseGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseGroupService_DeleteDatabaseGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseGroupServiceServer).DeleteDatabaseGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseGroupService_DeleteDatabaseGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseGroupServiceServer).DeleteDatabaseGroup(ctx, req.(*DeleteDatabaseGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseGroupService_ServiceDesc is the grpc.ServiceDesc for DatabaseGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.DatabaseGroupService",
	HandlerType: (*DatabaseGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDatabaseGroups",
			Handler:    _DatabaseGroupService_ListDatabaseGroups_Handler,
		},
		{
			MethodName: "GetDatabaseGroup",
			Handler:    _DatabaseGroupService_GetDatabaseGroup_Handler,
		},
		{
			MethodName: "CreateDatabaseGroup",
			Handler:    _DatabaseGroupService_CreateDatabaseGroup_Handler,
		},
		{
			MethodName: "UpdateDatabaseGroup",
			Handler:    _DatabaseGroupService_UpdateDatabaseGroup_Handler,
		},
		{
			MethodName: "DeleteDatabaseGroup",
			Handler:    _DatabaseGroupService_DeleteDatabaseGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/database_group_service.proto",
}
