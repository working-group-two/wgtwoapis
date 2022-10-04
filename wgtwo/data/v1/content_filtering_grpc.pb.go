// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ContentFilteringServiceClient is the client API for ContentFilteringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentFilteringServiceClient interface {
	SetGlobalFilter(ctx context.Context, in *SetGlobalFilterRequest, opts ...grpc.CallOption) (*SetGlobalFilterResponse, error)
	GetGlobalFilter(ctx context.Context, in *GetGlobalFilterRequest, opts ...grpc.CallOption) (*GetGlobalFilterResponse, error)
}

type contentFilteringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentFilteringServiceClient(cc grpc.ClientConnInterface) ContentFilteringServiceClient {
	return &contentFilteringServiceClient{cc}
}

func (c *contentFilteringServiceClient) SetGlobalFilter(ctx context.Context, in *SetGlobalFilterRequest, opts ...grpc.CallOption) (*SetGlobalFilterResponse, error) {
	out := new(SetGlobalFilterResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.data.v1.ContentFilteringService/SetGlobalFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentFilteringServiceClient) GetGlobalFilter(ctx context.Context, in *GetGlobalFilterRequest, opts ...grpc.CallOption) (*GetGlobalFilterResponse, error) {
	out := new(GetGlobalFilterResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.data.v1.ContentFilteringService/GetGlobalFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentFilteringServiceServer is the server API for ContentFilteringService service.
// All implementations should embed UnimplementedContentFilteringServiceServer
// for forward compatibility
type ContentFilteringServiceServer interface {
	SetGlobalFilter(context.Context, *SetGlobalFilterRequest) (*SetGlobalFilterResponse, error)
	GetGlobalFilter(context.Context, *GetGlobalFilterRequest) (*GetGlobalFilterResponse, error)
}

// UnimplementedContentFilteringServiceServer should be embedded to have forward compatible implementations.
type UnimplementedContentFilteringServiceServer struct {
}

func (UnimplementedContentFilteringServiceServer) SetGlobalFilter(context.Context, *SetGlobalFilterRequest) (*SetGlobalFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGlobalFilter not implemented")
}
func (UnimplementedContentFilteringServiceServer) GetGlobalFilter(context.Context, *GetGlobalFilterRequest) (*GetGlobalFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalFilter not implemented")
}

// UnsafeContentFilteringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentFilteringServiceServer will
// result in compilation errors.
type UnsafeContentFilteringServiceServer interface {
	mustEmbedUnimplementedContentFilteringServiceServer()
}

func RegisterContentFilteringServiceServer(s grpc.ServiceRegistrar, srv ContentFilteringServiceServer) {
	s.RegisterService(&ContentFilteringService_ServiceDesc, srv)
}

func _ContentFilteringService_SetGlobalFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGlobalFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentFilteringServiceServer).SetGlobalFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.data.v1.ContentFilteringService/SetGlobalFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentFilteringServiceServer).SetGlobalFilter(ctx, req.(*SetGlobalFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentFilteringService_GetGlobalFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentFilteringServiceServer).GetGlobalFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.data.v1.ContentFilteringService/GetGlobalFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentFilteringServiceServer).GetGlobalFilter(ctx, req.(*GetGlobalFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentFilteringService_ServiceDesc is the grpc.ServiceDesc for ContentFilteringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentFilteringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.data.v1.ContentFilteringService",
	HandlerType: (*ContentFilteringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGlobalFilter",
			Handler:    _ContentFilteringService_SetGlobalFilter_Handler,
		},
		{
			MethodName: "GetGlobalFilter",
			Handler:    _ContentFilteringService_GetGlobalFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/data/v1/content_filtering.proto",
}
