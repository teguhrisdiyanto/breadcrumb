// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/profile/v1/profile_api.proto

package profilev1

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

// ProfileAPIClient is the client API for ProfileAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileAPIClient interface {
	// Create profile rpc service.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Update profile rpc service.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// ByID get profile by id rpc service.
	ByID(ctx context.Context, in *ByIDRequest, opts ...grpc.CallOption) (*ByIDResponse, error)
	// ByIDUser get profile by id rpc service.
	ByIDUser(ctx context.Context, in *ByIDUserRequest, opts ...grpc.CallOption) (*ByIDUserResponse, error)
}

type profileAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileAPIClient(cc grpc.ClientConnInterface) ProfileAPIClient {
	return &profileAPIClient{cc}
}

func (c *profileAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/proto.profile.v1.ProfileAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.profile.v1.ProfileAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileAPIClient) ByID(ctx context.Context, in *ByIDRequest, opts ...grpc.CallOption) (*ByIDResponse, error) {
	out := new(ByIDResponse)
	err := c.cc.Invoke(ctx, "/proto.profile.v1.ProfileAPI/ByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileAPIClient) ByIDUser(ctx context.Context, in *ByIDUserRequest, opts ...grpc.CallOption) (*ByIDUserResponse, error) {
	out := new(ByIDUserResponse)
	err := c.cc.Invoke(ctx, "/proto.profile.v1.ProfileAPI/ByIDUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileAPIServer is the server API for ProfileAPI service.
// All implementations should embed UnimplementedProfileAPIServer
// for forward compatibility
type ProfileAPIServer interface {
	// Create profile rpc service.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Update profile rpc service.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// ByID get profile by id rpc service.
	ByID(context.Context, *ByIDRequest) (*ByIDResponse, error)
	// ByIDUser get profile by id rpc service.
	ByIDUser(context.Context, *ByIDUserRequest) (*ByIDUserResponse, error)
}

// UnimplementedProfileAPIServer should be embedded to have forward compatible implementations.
type UnimplementedProfileAPIServer struct {
}

func (UnimplementedProfileAPIServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProfileAPIServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProfileAPIServer) ByID(context.Context, *ByIDRequest) (*ByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByID not implemented")
}
func (UnimplementedProfileAPIServer) ByIDUser(context.Context, *ByIDUserRequest) (*ByIDUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByIDUser not implemented")
}

// UnsafeProfileAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileAPIServer will
// result in compilation errors.
type UnsafeProfileAPIServer interface {
	mustEmbedUnimplementedProfileAPIServer()
}

func RegisterProfileAPIServer(s grpc.ServiceRegistrar, srv ProfileAPIServer) {
	s.RegisterService(&ProfileAPI_ServiceDesc, srv)
}

func _ProfileAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.profile.v1.ProfileAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.profile.v1.ProfileAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileAPI_ByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAPIServer).ByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.profile.v1.ProfileAPI/ByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAPIServer).ByID(ctx, req.(*ByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileAPI_ByIDUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIDUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileAPIServer).ByIDUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.profile.v1.ProfileAPI/ByIDUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileAPIServer).ByIDUser(ctx, req.(*ByIDUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileAPI_ServiceDesc is the grpc.ServiceDesc for ProfileAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.profile.v1.ProfileAPI",
	HandlerType: (*ProfileAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProfileAPI_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProfileAPI_Update_Handler,
		},
		{
			MethodName: "ByID",
			Handler:    _ProfileAPI_ByID_Handler,
		},
		{
			MethodName: "ByIDUser",
			Handler:    _ProfileAPI_ByIDUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/profile/v1/profile_api.proto",
}
