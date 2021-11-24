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

// SubscriptionEventServiceClient is the client API for SubscriptionEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionEventServiceClient interface {
	StreamHandsetChangeEvents(ctx context.Context, in *StreamHandsetChangeEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamHandsetChangeEventsClient, error)
	AckHandsetChangeEvent(ctx context.Context, in *AckHandsetChangeEventRequest, opts ...grpc.CallOption) (*AckHandsetChangeEventResponse, error)
	StreamRoamingEvents(ctx context.Context, in *StreamRoamingEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamRoamingEventsClient, error)
	AckRoamingEvent(ctx context.Context, in *AckRoamingEventRequest, opts ...grpc.CallOption) (*AckRoamingEventResponse, error)
}

type subscriptionEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionEventServiceClient(cc grpc.ClientConnInterface) SubscriptionEventServiceClient {
	return &subscriptionEventServiceClient{cc}
}

func (c *subscriptionEventServiceClient) StreamHandsetChangeEvents(ctx context.Context, in *StreamHandsetChangeEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamHandsetChangeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionEventService_ServiceDesc.Streams[0], "/wgtwo.subscription.v1.SubscriptionEventService/StreamHandsetChangeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionEventServiceStreamHandsetChangeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionEventService_StreamHandsetChangeEventsClient interface {
	Recv() (*StreamHandsetChangeEventsResponse, error)
	grpc.ClientStream
}

type subscriptionEventServiceStreamHandsetChangeEventsClient struct {
	grpc.ClientStream
}

func (x *subscriptionEventServiceStreamHandsetChangeEventsClient) Recv() (*StreamHandsetChangeEventsResponse, error) {
	m := new(StreamHandsetChangeEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionEventServiceClient) AckHandsetChangeEvent(ctx context.Context, in *AckHandsetChangeEventRequest, opts ...grpc.CallOption) (*AckHandsetChangeEventResponse, error) {
	out := new(AckHandsetChangeEventResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.subscription.v1.SubscriptionEventService/AckHandsetChangeEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionEventServiceClient) StreamRoamingEvents(ctx context.Context, in *StreamRoamingEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamRoamingEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionEventService_ServiceDesc.Streams[1], "/wgtwo.subscription.v1.SubscriptionEventService/StreamRoamingEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionEventServiceStreamRoamingEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionEventService_StreamRoamingEventsClient interface {
	Recv() (*StreamRoamingEventsResponse, error)
	grpc.ClientStream
}

type subscriptionEventServiceStreamRoamingEventsClient struct {
	grpc.ClientStream
}

func (x *subscriptionEventServiceStreamRoamingEventsClient) Recv() (*StreamRoamingEventsResponse, error) {
	m := new(StreamRoamingEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionEventServiceClient) AckRoamingEvent(ctx context.Context, in *AckRoamingEventRequest, opts ...grpc.CallOption) (*AckRoamingEventResponse, error) {
	out := new(AckRoamingEventResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.subscription.v1.SubscriptionEventService/AckRoamingEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionEventServiceServer is the server API for SubscriptionEventService service.
// All implementations should embed UnimplementedSubscriptionEventServiceServer
// for forward compatibility
type SubscriptionEventServiceServer interface {
	StreamHandsetChangeEvents(*StreamHandsetChangeEventsRequest, SubscriptionEventService_StreamHandsetChangeEventsServer) error
	AckHandsetChangeEvent(context.Context, *AckHandsetChangeEventRequest) (*AckHandsetChangeEventResponse, error)
	StreamRoamingEvents(*StreamRoamingEventsRequest, SubscriptionEventService_StreamRoamingEventsServer) error
	AckRoamingEvent(context.Context, *AckRoamingEventRequest) (*AckRoamingEventResponse, error)
}

// UnimplementedSubscriptionEventServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubscriptionEventServiceServer struct {
}

func (UnimplementedSubscriptionEventServiceServer) StreamHandsetChangeEvents(*StreamHandsetChangeEventsRequest, SubscriptionEventService_StreamHandsetChangeEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHandsetChangeEvents not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) AckHandsetChangeEvent(context.Context, *AckHandsetChangeEventRequest) (*AckHandsetChangeEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckHandsetChangeEvent not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) StreamRoamingEvents(*StreamRoamingEventsRequest, SubscriptionEventService_StreamRoamingEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRoamingEvents not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) AckRoamingEvent(context.Context, *AckRoamingEventRequest) (*AckRoamingEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckRoamingEvent not implemented")
}

// UnsafeSubscriptionEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionEventServiceServer will
// result in compilation errors.
type UnsafeSubscriptionEventServiceServer interface {
	mustEmbedUnimplementedSubscriptionEventServiceServer()
}

func RegisterSubscriptionEventServiceServer(s grpc.ServiceRegistrar, srv SubscriptionEventServiceServer) {
	s.RegisterService(&SubscriptionEventService_ServiceDesc, srv)
}

func _SubscriptionEventService_StreamHandsetChangeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamHandsetChangeEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionEventServiceServer).StreamHandsetChangeEvents(m, &subscriptionEventServiceStreamHandsetChangeEventsServer{stream})
}

type SubscriptionEventService_StreamHandsetChangeEventsServer interface {
	Send(*StreamHandsetChangeEventsResponse) error
	grpc.ServerStream
}

type subscriptionEventServiceStreamHandsetChangeEventsServer struct {
	grpc.ServerStream
}

func (x *subscriptionEventServiceStreamHandsetChangeEventsServer) Send(m *StreamHandsetChangeEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionEventService_AckHandsetChangeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckHandsetChangeEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionEventServiceServer).AckHandsetChangeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.subscription.v1.SubscriptionEventService/AckHandsetChangeEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionEventServiceServer).AckHandsetChangeEvent(ctx, req.(*AckHandsetChangeEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionEventService_StreamRoamingEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRoamingEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionEventServiceServer).StreamRoamingEvents(m, &subscriptionEventServiceStreamRoamingEventsServer{stream})
}

type SubscriptionEventService_StreamRoamingEventsServer interface {
	Send(*StreamRoamingEventsResponse) error
	grpc.ServerStream
}

type subscriptionEventServiceStreamRoamingEventsServer struct {
	grpc.ServerStream
}

func (x *subscriptionEventServiceStreamRoamingEventsServer) Send(m *StreamRoamingEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionEventService_AckRoamingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRoamingEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionEventServiceServer).AckRoamingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.subscription.v1.SubscriptionEventService/AckRoamingEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionEventServiceServer).AckRoamingEvent(ctx, req.(*AckRoamingEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionEventService_ServiceDesc is the grpc.ServiceDesc for SubscriptionEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.subscription.v1.SubscriptionEventService",
	HandlerType: (*SubscriptionEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AckHandsetChangeEvent",
			Handler:    _SubscriptionEventService_AckHandsetChangeEvent_Handler,
		},
		{
			MethodName: "AckRoamingEvent",
			Handler:    _SubscriptionEventService_AckRoamingEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHandsetChangeEvents",
			Handler:       _SubscriptionEventService_StreamHandsetChangeEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamRoamingEvents",
			Handler:       _SubscriptionEventService_StreamRoamingEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wgtwo/subscription/v1/subscription_events.proto",
}