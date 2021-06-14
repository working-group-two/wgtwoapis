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

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsServiceClient interface {
	SendTextToSubscriber(ctx context.Context, in *SendTextToSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error)
	SendTextFromSubscriber(ctx context.Context, in *SendTextFromSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error)
	SendBinaryToSubscriber(ctx context.Context, in *SendBinaryToSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) SendTextToSubscriber(ctx context.Context, in *SendTextToSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.sms.v0.SmsService/SendTextToSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsServiceClient) SendTextFromSubscriber(ctx context.Context, in *SendTextFromSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.sms.v0.SmsService/SendTextFromSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsServiceClient) SendBinaryToSubscriber(ctx context.Context, in *SendBinaryToSubscriberRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.sms.v0.SmsService/SendBinaryToSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
// All implementations should embed UnimplementedSmsServiceServer
// for forward compatibility
type SmsServiceServer interface {
	SendTextToSubscriber(context.Context, *SendTextToSubscriberRequest) (*SendResponse, error)
	SendTextFromSubscriber(context.Context, *SendTextFromSubscriberRequest) (*SendResponse, error)
	SendBinaryToSubscriber(context.Context, *SendBinaryToSubscriberRequest) (*SendResponse, error)
}

// UnimplementedSmsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSmsServiceServer struct {
}

func (UnimplementedSmsServiceServer) SendTextToSubscriber(context.Context, *SendTextToSubscriberRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTextToSubscriber not implemented")
}
func (UnimplementedSmsServiceServer) SendTextFromSubscriber(context.Context, *SendTextFromSubscriberRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTextFromSubscriber not implemented")
}
func (UnimplementedSmsServiceServer) SendBinaryToSubscriber(context.Context, *SendBinaryToSubscriberRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBinaryToSubscriber not implemented")
}

// UnsafeSmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServiceServer will
// result in compilation errors.
type UnsafeSmsServiceServer interface {
	mustEmbedUnimplementedSmsServiceServer()
}

func RegisterSmsServiceServer(s grpc.ServiceRegistrar, srv SmsServiceServer) {
	s.RegisterService(&_SmsService_serviceDesc, srv)
}

func _SmsService_SendTextToSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTextToSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SendTextToSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.sms.v0.SmsService/SendTextToSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SendTextToSubscriber(ctx, req.(*SendTextToSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsService_SendTextFromSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTextFromSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SendTextFromSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.sms.v0.SmsService/SendTextFromSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SendTextFromSubscriber(ctx, req.(*SendTextFromSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsService_SendBinaryToSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendBinaryToSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SendBinaryToSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.sms.v0.SmsService/SendBinaryToSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SendBinaryToSubscriber(ctx, req.(*SendBinaryToSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.sms.v0.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTextToSubscriber",
			Handler:    _SmsService_SendTextToSubscriber_Handler,
		},
		{
			MethodName: "SendTextFromSubscriber",
			Handler:    _SmsService_SendTextFromSubscriber_Handler,
		},
		{
			MethodName: "SendBinaryToSubscriber",
			Handler:    _SmsService_SendBinaryToSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/sms/v0/sms.proto",
}
