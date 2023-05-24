// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/cel_service.proto

package v1

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
	CelService_Parse_FullMethodName   = "/bytebase.v1.CelService/Parse"
	CelService_Deparse_FullMethodName = "/bytebase.v1.CelService/Deparse"
)

// CelServiceClient is the client API for CelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CelServiceClient interface {
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	Deparse(ctx context.Context, in *DeparseRequest, opts ...grpc.CallOption) (*DeparseResponse, error)
}

type celServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCelServiceClient(cc grpc.ClientConnInterface) CelServiceClient {
	return &celServiceClient{cc}
}

func (c *celServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, CelService_Parse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *celServiceClient) Deparse(ctx context.Context, in *DeparseRequest, opts ...grpc.CallOption) (*DeparseResponse, error) {
	out := new(DeparseResponse)
	err := c.cc.Invoke(ctx, CelService_Deparse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CelServiceServer is the server API for CelService service.
// All implementations must embed UnimplementedCelServiceServer
// for forward compatibility
type CelServiceServer interface {
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	Deparse(context.Context, *DeparseRequest) (*DeparseResponse, error)
	mustEmbedUnimplementedCelServiceServer()
}

// UnimplementedCelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCelServiceServer struct {
}

func (UnimplementedCelServiceServer) Parse(context.Context, *ParseRequest) (*ParseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedCelServiceServer) Deparse(context.Context, *DeparseRequest) (*DeparseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deparse not implemented")
}
func (UnimplementedCelServiceServer) mustEmbedUnimplementedCelServiceServer() {}

// UnsafeCelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CelServiceServer will
// result in compilation errors.
type UnsafeCelServiceServer interface {
	mustEmbedUnimplementedCelServiceServer()
}

func RegisterCelServiceServer(s grpc.ServiceRegistrar, srv CelServiceServer) {
	s.RegisterService(&CelService_ServiceDesc, srv)
}

func _CelService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CelService_Parse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CelService_Deparse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeparseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CelServiceServer).Deparse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CelService_Deparse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CelServiceServer).Deparse(ctx, req.(*DeparseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CelService_ServiceDesc is the grpc.ServiceDesc for CelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.CelService",
	HandlerType: (*CelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _CelService_Parse_Handler,
		},
		{
			MethodName: "Deparse",
			Handler:    _CelService_Deparse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/cel_service.proto",
}
