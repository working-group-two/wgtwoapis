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

// MmsServiceClient is the client API for MmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MmsServiceClient interface {
	SendMessageToSubscriber(ctx context.Context, in *SendMessageToSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error)
	SendMessageFromSubscriber(ctx context.Context, in *SendMessageFromSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type mmsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMmsServiceClient(cc grpc.ClientConnInterface) MmsServiceClient {
	return &mmsServiceClient{cc}
}

func (c *mmsServiceClient) SendMessageToSubscriber(ctx context.Context, in *SendMessageToSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.mms.v0.MmsService/SendMessageToSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mmsServiceClient) SendMessageFromSubscriber(ctx context.Context, in *SendMessageFromSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.mms.v0.MmsService/SendMessageFromSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MmsServiceServer is the server API for MmsService service.
// All implementations should embed UnimplementedMmsServiceServer
// for forward compatibility
type MmsServiceServer interface {
	SendMessageToSubscriber(context.Context, *SendMessageToSubscriberRequest) (*SendResponse, error)
	SendMessageFromSubscriber(context.Context, *SendMessageFromSubscriberRequest) (*SendResponse, error)
}

// UnimplementedMmsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMmsServiceServer struct {
}

func (UnimplementedMmsServiceServer) SendMessageToSubscriber(context.Context, *SendMessageToSubscriberRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToSubscriber not implemented")
}
func (UnimplementedMmsServiceServer) SendMessageFromSubscriber(context.Context, *SendMessageFromSubscriberRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageFromSubscriber not implemented")
}

// UnsafeMmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MmsServiceServer will
// result in compilation errors.
type UnsafeMmsServiceServer interface {
	mustEmbedUnimplementedMmsServiceServer()
}

func RegisterMmsServiceServer(s grpc.ServiceRegistrar, srv MmsServiceServer) {
	s.RegisterService(&_MmsService_serviceDesc, srv)
}

func _MmsService_SendMessageToSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageToSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MmsServiceServer).SendMessageToSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.mms.v0.MmsService/SendMessageToSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MmsServiceServer).SendMessageToSubscriber(ctx, req.(*SendMessageToSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MmsService_SendMessageFromSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageFromSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MmsServiceServer).SendMessageFromSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.mms.v0.MmsService/SendMessageFromSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MmsServiceServer).SendMessageFromSubscriber(ctx, req.(*SendMessageFromSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MmsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.mms.v0.MmsService",
	HandlerType: (*MmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageToSubscriber",
			Handler:    _MmsService_SendMessageToSubscriber_Handler,
		},
		{
			MethodName: "SendMessageFromSubscriber",
			Handler:    _MmsService_SendMessageFromSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/mms/v0/mms.proto",
}
