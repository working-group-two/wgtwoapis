// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v0

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RightServiceClient is the client API for RightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RightServiceClient interface {
	List(ctx context.Context, in *ListRightsRequest, opts ...grpc.CallOption) (*ListRightsResponse, error)
	CheckExpansion(ctx context.Context, in *CheckExpansionRequest, opts ...grpc.CallOption) (*CheckExpansionResponse, error)
}

type rightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRightServiceClient(cc grpc.ClientConnInterface) RightServiceClient {
	return &rightServiceClient{cc}
}

func (c *rightServiceClient) List(ctx context.Context, in *ListRightsRequest, opts ...grpc.CallOption) (*ListRightsResponse, error) {
	out := new(ListRightsResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.auth.v0.RightService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightServiceClient) CheckExpansion(ctx context.Context, in *CheckExpansionRequest, opts ...grpc.CallOption) (*CheckExpansionResponse, error) {
	out := new(CheckExpansionResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.auth.v0.RightService/CheckExpansion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RightServiceServer is the server API for RightService service.
// All implementations should embed UnimplementedRightServiceServer
// for forward compatibility
type RightServiceServer interface {
	List(context.Context, *ListRightsRequest) (*ListRightsResponse, error)
	CheckExpansion(context.Context, *CheckExpansionRequest) (*CheckExpansionResponse, error)
}

// UnimplementedRightServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRightServiceServer struct {
}

func (UnimplementedRightServiceServer) List(context.Context, *ListRightsRequest) (*ListRightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRightServiceServer) CheckExpansion(context.Context, *CheckExpansionRequest) (*CheckExpansionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExpansion not implemented")
}

// UnsafeRightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RightServiceServer will
// result in compilation errors.
type UnsafeRightServiceServer interface {
	mustEmbedUnimplementedRightServiceServer()
}

func RegisterRightServiceServer(s grpc.ServiceRegistrar, srv RightServiceServer) {
	s.RegisterService(&_RightService_serviceDesc, srv)
}

func _RightService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.auth.v0.RightService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightServiceServer).List(ctx, req.(*ListRightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RightService_CheckExpansion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExpansionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightServiceServer).CheckExpansion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.auth.v0.RightService/CheckExpansion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightServiceServer).CheckExpansion(ctx, req.(*CheckExpansionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RightService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.auth.v0.RightService",
	HandlerType: (*RightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RightService_List_Handler,
		},
		{
			MethodName: "CheckExpansion",
			Handler:    _RightService_CheckExpansion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/auth/v0/rights.proto",
}
