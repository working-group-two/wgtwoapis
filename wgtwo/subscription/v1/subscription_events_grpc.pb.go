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
	// First Attachment events are triggered whenever a SIM is first attached to the
	// network. It contains the IMSI to distinguish which SIM of the subscriber has
	// been attached.
	StreamFirstAttachmentEvents(ctx context.Context, in *StreamFirstAttachmentEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamFirstAttachmentEventsClient, error)
	// Manually ack a first attachment event.
	AckFirstAttachmentEvent(ctx context.Context, in *AckFirstAttachmentEventRequest, opts ...grpc.CallOption) (*AckFirstAttachmentEventResponse, error)
	// Country change events are triggered whenever a SIM changes current country
	// location. It has both the current (new) country and the previous (old) country.
	//
	// This event is triggered when the previously seen country and the currently seen country
	// are different. Note that subscribers being close to borders, or during travels may generate
	// a lot of CountryChange events. See 'PeriodicCountry' events for an alternative.
	//
	// For subscribers with multiple SIM cards you will see an event for each SIM
	// (IMSI), as they can move between countries individually.
	StreamCountryChangeEvents(ctx context.Context, in *StreamCountryChangeEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamCountryChangeEventsClient, error)
	// Manually ack a country change event.
	AckCountryChangeEvent(ctx context.Context, in *AckCountryChangeEventRequest, opts ...grpc.CallOption) (*AckCountryChangeEventResponse, error)
	// Periodic country events are triggered on a regular basis for each user for each
	// country where they are seen. It is triggered by knowingly seeing the subscriber
	// & handset in a specific country, and for each tenant will be triggered on a
	// regular interval. E.g. if 'Operator X' is configured for a 2 week interval,
	// there will be an event every 14 days (or 336 hours or 1209600 seconds) as long
	// as the subscriber is still seen in that country.
	//
	// As this event is not always triggered based on the subscriber moving between
	// countries, it does not contain the previously seen country. For getting the real-time
	// movement of the subscriber between countries, use 'CountryChange' event.
	//
	// This event is triggered:
	//
	// - When the subscriber first turns on the device and it connects to a network, it
	//   will be triggered for the country of the connected network at the same time as the
	//   corresponding 'FirstAttachment' event.
	// - When the subscriber enters a new country (not visited before). This is triggered
	//   at the same time as the corresponding 'CountryChange' event.
	// - When the subscriber is seen in a country, and the 'PeriodicCountry' event for that
	//   subscriber and country has not been triggered for the configured time delay.
	//
	// For subscribers with multiple SIM cards you will see an event for each SIM
	// (IMSI), as they can move between countries individually.
	StreamPeriodicCountryEvents(ctx context.Context, in *StreamPeriodicCountryEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamPeriodicCountryEventsClient, error)
	// Manually ack a periodic country event.
	AckPeriodicCountryEvent(ctx context.Context, in *AckPeriodicCountryEventRequest, opts ...grpc.CallOption) (*AckPeriodicCountryEventResponse, error)
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

func (c *subscriptionEventServiceClient) StreamFirstAttachmentEvents(ctx context.Context, in *StreamFirstAttachmentEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamFirstAttachmentEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionEventService_ServiceDesc.Streams[1], "/wgtwo.subscription.v1.SubscriptionEventService/StreamFirstAttachmentEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionEventServiceStreamFirstAttachmentEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionEventService_StreamFirstAttachmentEventsClient interface {
	Recv() (*StreamFirstAttachmentEventsResponse, error)
	grpc.ClientStream
}

type subscriptionEventServiceStreamFirstAttachmentEventsClient struct {
	grpc.ClientStream
}

func (x *subscriptionEventServiceStreamFirstAttachmentEventsClient) Recv() (*StreamFirstAttachmentEventsResponse, error) {
	m := new(StreamFirstAttachmentEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionEventServiceClient) AckFirstAttachmentEvent(ctx context.Context, in *AckFirstAttachmentEventRequest, opts ...grpc.CallOption) (*AckFirstAttachmentEventResponse, error) {
	out := new(AckFirstAttachmentEventResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.subscription.v1.SubscriptionEventService/AckFirstAttachmentEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionEventServiceClient) StreamCountryChangeEvents(ctx context.Context, in *StreamCountryChangeEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamCountryChangeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionEventService_ServiceDesc.Streams[2], "/wgtwo.subscription.v1.SubscriptionEventService/StreamCountryChangeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionEventServiceStreamCountryChangeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionEventService_StreamCountryChangeEventsClient interface {
	Recv() (*StreamCountryChangeEventsResponse, error)
	grpc.ClientStream
}

type subscriptionEventServiceStreamCountryChangeEventsClient struct {
	grpc.ClientStream
}

func (x *subscriptionEventServiceStreamCountryChangeEventsClient) Recv() (*StreamCountryChangeEventsResponse, error) {
	m := new(StreamCountryChangeEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionEventServiceClient) AckCountryChangeEvent(ctx context.Context, in *AckCountryChangeEventRequest, opts ...grpc.CallOption) (*AckCountryChangeEventResponse, error) {
	out := new(AckCountryChangeEventResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.subscription.v1.SubscriptionEventService/AckCountryChangeEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionEventServiceClient) StreamPeriodicCountryEvents(ctx context.Context, in *StreamPeriodicCountryEventsRequest, opts ...grpc.CallOption) (SubscriptionEventService_StreamPeriodicCountryEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionEventService_ServiceDesc.Streams[3], "/wgtwo.subscription.v1.SubscriptionEventService/StreamPeriodicCountryEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionEventServiceStreamPeriodicCountryEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionEventService_StreamPeriodicCountryEventsClient interface {
	Recv() (*StreamPeriodicCountryEventsResponse, error)
	grpc.ClientStream
}

type subscriptionEventServiceStreamPeriodicCountryEventsClient struct {
	grpc.ClientStream
}

func (x *subscriptionEventServiceStreamPeriodicCountryEventsClient) Recv() (*StreamPeriodicCountryEventsResponse, error) {
	m := new(StreamPeriodicCountryEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionEventServiceClient) AckPeriodicCountryEvent(ctx context.Context, in *AckPeriodicCountryEventRequest, opts ...grpc.CallOption) (*AckPeriodicCountryEventResponse, error) {
	out := new(AckPeriodicCountryEventResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.subscription.v1.SubscriptionEventService/AckPeriodicCountryEvent", in, out, opts...)
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
	// First Attachment events are triggered whenever a SIM is first attached to the
	// network. It contains the IMSI to distinguish which SIM of the subscriber has
	// been attached.
	StreamFirstAttachmentEvents(*StreamFirstAttachmentEventsRequest, SubscriptionEventService_StreamFirstAttachmentEventsServer) error
	// Manually ack a first attachment event.
	AckFirstAttachmentEvent(context.Context, *AckFirstAttachmentEventRequest) (*AckFirstAttachmentEventResponse, error)
	// Country change events are triggered whenever a SIM changes current country
	// location. It has both the current (new) country and the previous (old) country.
	//
	// This event is triggered when the previously seen country and the currently seen country
	// are different. Note that subscribers being close to borders, or during travels may generate
	// a lot of CountryChange events. See 'PeriodicCountry' events for an alternative.
	//
	// For subscribers with multiple SIM cards you will see an event for each SIM
	// (IMSI), as they can move between countries individually.
	StreamCountryChangeEvents(*StreamCountryChangeEventsRequest, SubscriptionEventService_StreamCountryChangeEventsServer) error
	// Manually ack a country change event.
	AckCountryChangeEvent(context.Context, *AckCountryChangeEventRequest) (*AckCountryChangeEventResponse, error)
	// Periodic country events are triggered on a regular basis for each user for each
	// country where they are seen. It is triggered by knowingly seeing the subscriber
	// & handset in a specific country, and for each tenant will be triggered on a
	// regular interval. E.g. if 'Operator X' is configured for a 2 week interval,
	// there will be an event every 14 days (or 336 hours or 1209600 seconds) as long
	// as the subscriber is still seen in that country.
	//
	// As this event is not always triggered based on the subscriber moving between
	// countries, it does not contain the previously seen country. For getting the real-time
	// movement of the subscriber between countries, use 'CountryChange' event.
	//
	// This event is triggered:
	//
	// - When the subscriber first turns on the device and it connects to a network, it
	//   will be triggered for the country of the connected network at the same time as the
	//   corresponding 'FirstAttachment' event.
	// - When the subscriber enters a new country (not visited before). This is triggered
	//   at the same time as the corresponding 'CountryChange' event.
	// - When the subscriber is seen in a country, and the 'PeriodicCountry' event for that
	//   subscriber and country has not been triggered for the configured time delay.
	//
	// For subscribers with multiple SIM cards you will see an event for each SIM
	// (IMSI), as they can move between countries individually.
	StreamPeriodicCountryEvents(*StreamPeriodicCountryEventsRequest, SubscriptionEventService_StreamPeriodicCountryEventsServer) error
	// Manually ack a periodic country event.
	AckPeriodicCountryEvent(context.Context, *AckPeriodicCountryEventRequest) (*AckPeriodicCountryEventResponse, error)
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
func (UnimplementedSubscriptionEventServiceServer) StreamFirstAttachmentEvents(*StreamFirstAttachmentEventsRequest, SubscriptionEventService_StreamFirstAttachmentEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamFirstAttachmentEvents not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) AckFirstAttachmentEvent(context.Context, *AckFirstAttachmentEventRequest) (*AckFirstAttachmentEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckFirstAttachmentEvent not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) StreamCountryChangeEvents(*StreamCountryChangeEventsRequest, SubscriptionEventService_StreamCountryChangeEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCountryChangeEvents not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) AckCountryChangeEvent(context.Context, *AckCountryChangeEventRequest) (*AckCountryChangeEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckCountryChangeEvent not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) StreamPeriodicCountryEvents(*StreamPeriodicCountryEventsRequest, SubscriptionEventService_StreamPeriodicCountryEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPeriodicCountryEvents not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) AckPeriodicCountryEvent(context.Context, *AckPeriodicCountryEventRequest) (*AckPeriodicCountryEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckPeriodicCountryEvent not implemented")
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

func _SubscriptionEventService_StreamFirstAttachmentEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamFirstAttachmentEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionEventServiceServer).StreamFirstAttachmentEvents(m, &subscriptionEventServiceStreamFirstAttachmentEventsServer{stream})
}

type SubscriptionEventService_StreamFirstAttachmentEventsServer interface {
	Send(*StreamFirstAttachmentEventsResponse) error
	grpc.ServerStream
}

type subscriptionEventServiceStreamFirstAttachmentEventsServer struct {
	grpc.ServerStream
}

func (x *subscriptionEventServiceStreamFirstAttachmentEventsServer) Send(m *StreamFirstAttachmentEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionEventService_AckFirstAttachmentEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckFirstAttachmentEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionEventServiceServer).AckFirstAttachmentEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.subscription.v1.SubscriptionEventService/AckFirstAttachmentEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionEventServiceServer).AckFirstAttachmentEvent(ctx, req.(*AckFirstAttachmentEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionEventService_StreamCountryChangeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamCountryChangeEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionEventServiceServer).StreamCountryChangeEvents(m, &subscriptionEventServiceStreamCountryChangeEventsServer{stream})
}

type SubscriptionEventService_StreamCountryChangeEventsServer interface {
	Send(*StreamCountryChangeEventsResponse) error
	grpc.ServerStream
}

type subscriptionEventServiceStreamCountryChangeEventsServer struct {
	grpc.ServerStream
}

func (x *subscriptionEventServiceStreamCountryChangeEventsServer) Send(m *StreamCountryChangeEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionEventService_AckCountryChangeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckCountryChangeEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionEventServiceServer).AckCountryChangeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.subscription.v1.SubscriptionEventService/AckCountryChangeEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionEventServiceServer).AckCountryChangeEvent(ctx, req.(*AckCountryChangeEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionEventService_StreamPeriodicCountryEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamPeriodicCountryEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionEventServiceServer).StreamPeriodicCountryEvents(m, &subscriptionEventServiceStreamPeriodicCountryEventsServer{stream})
}

type SubscriptionEventService_StreamPeriodicCountryEventsServer interface {
	Send(*StreamPeriodicCountryEventsResponse) error
	grpc.ServerStream
}

type subscriptionEventServiceStreamPeriodicCountryEventsServer struct {
	grpc.ServerStream
}

func (x *subscriptionEventServiceStreamPeriodicCountryEventsServer) Send(m *StreamPeriodicCountryEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionEventService_AckPeriodicCountryEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckPeriodicCountryEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionEventServiceServer).AckPeriodicCountryEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.subscription.v1.SubscriptionEventService/AckPeriodicCountryEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionEventServiceServer).AckPeriodicCountryEvent(ctx, req.(*AckPeriodicCountryEventRequest))
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
			MethodName: "AckFirstAttachmentEvent",
			Handler:    _SubscriptionEventService_AckFirstAttachmentEvent_Handler,
		},
		{
			MethodName: "AckCountryChangeEvent",
			Handler:    _SubscriptionEventService_AckCountryChangeEvent_Handler,
		},
		{
			MethodName: "AckPeriodicCountryEvent",
			Handler:    _SubscriptionEventService_AckPeriodicCountryEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHandsetChangeEvents",
			Handler:       _SubscriptionEventService_StreamHandsetChangeEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamFirstAttachmentEvents",
			Handler:       _SubscriptionEventService_StreamFirstAttachmentEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamCountryChangeEvents",
			Handler:       _SubscriptionEventService_StreamCountryChangeEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamPeriodicCountryEvents",
			Handler:       _SubscriptionEventService_StreamPeriodicCountryEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wgtwo/subscription/v1/subscription_events.proto",
}
