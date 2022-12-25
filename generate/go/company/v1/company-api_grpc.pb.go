// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/company/v1/company-api.proto

package companyv1

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

// CrudApiClient is the client API for CrudApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrudApiClient interface {
	ReadItem(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error)
	ReadItem3(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error)
}

type crudApiClient struct {
	cc grpc.ClientConnInterface
}

func NewCrudApiClient(cc grpc.ClientConnInterface) CrudApiClient {
	return &crudApiClient{cc}
}

func (c *crudApiClient) ReadItem(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/proto.company.v1.CrudApi/ReadItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crudApiClient) ReadItem3(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/proto.company.v1.CrudApi/ReadItem3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrudApiServer is the server API for CrudApi service.
// All implementations should embed UnimplementedCrudApiServer
// for forward compatibility
type CrudApiServer interface {
	ReadItem(context.Context, *ID) (*Employee, error)
	ReadItem3(context.Context, *ID) (*Employee, error)
}

// UnimplementedCrudApiServer should be embedded to have forward compatible implementations.
type UnimplementedCrudApiServer struct {
}

func (UnimplementedCrudApiServer) ReadItem(context.Context, *ID) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadItem not implemented")
}
func (UnimplementedCrudApiServer) ReadItem3(context.Context, *ID) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadItem3 not implemented")
}

// UnsafeCrudApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrudApiServer will
// result in compilation errors.
type UnsafeCrudApiServer interface {
	mustEmbedUnimplementedCrudApiServer()
}

func RegisterCrudApiServer(s grpc.ServiceRegistrar, srv CrudApiServer) {
	s.RegisterService(&CrudApi_ServiceDesc, srv)
}

func _CrudApi_ReadItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudApiServer).ReadItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company.v1.CrudApi/ReadItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudApiServer).ReadItem(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrudApi_ReadItem3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrudApiServer).ReadItem3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.company.v1.CrudApi/ReadItem3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrudApiServer).ReadItem3(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// CrudApi_ServiceDesc is the grpc.ServiceDesc for CrudApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrudApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.company.v1.CrudApi",
	HandlerType: (*CrudApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadItem",
			Handler:    _CrudApi_ReadItem_Handler,
		},
		{
			MethodName: "ReadItem3",
			Handler:    _CrudApi_ReadItem3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/company/v1/company-api.proto",
}
