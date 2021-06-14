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

// EventsServiceClient is the client API for EventsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsServiceClient interface {
	Subscribe(ctx context.Context, in *SubscribeEventsRequest, opts ...grpc.CallOption) (EventsService_SubscribeClient, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResponse, error)
}

type eventsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsServiceClient(cc grpc.ClientConnInterface) EventsServiceClient {
	return &eventsServiceClient{cc}
}

func (c *eventsServiceClient) Subscribe(ctx context.Context, in *SubscribeEventsRequest, opts ...grpc.CallOption) (EventsService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventsService_serviceDesc.Streams[0], "/wgtwo.events.v0.EventsService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventsService_SubscribeClient interface {
	Recv() (*SubscribeEventsResponse, error)
	grpc.ClientStream
}

type eventsServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventsServiceSubscribeClient) Recv() (*SubscribeEventsResponse, error) {
	m := new(SubscribeEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventsServiceClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResponse, error) {
	out := new(AckResponse)
	err := c.cc.Invoke(ctx, "/wgtwo.events.v0.EventsService/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServiceServer is the server API for EventsService service.
// All implementations should embed UnimplementedEventsServiceServer
// for forward compatibility
type EventsServiceServer interface {
	Subscribe(*SubscribeEventsRequest, EventsService_SubscribeServer) error
	Ack(context.Context, *AckRequest) (*AckResponse, error)
}

// UnimplementedEventsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEventsServiceServer struct {
}

func (UnimplementedEventsServiceServer) Subscribe(*SubscribeEventsRequest, EventsService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventsServiceServer) Ack(context.Context, *AckRequest) (*AckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}

// UnsafeEventsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServiceServer will
// result in compilation errors.
type UnsafeEventsServiceServer interface {
	mustEmbedUnimplementedEventsServiceServer()
}

func RegisterEventsServiceServer(s grpc.ServiceRegistrar, srv EventsServiceServer) {
	s.RegisterService(&_EventsService_serviceDesc, srv)
}

func _EventsService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServiceServer).Subscribe(m, &eventsServiceSubscribeServer{stream})
}

type EventsService_SubscribeServer interface {
	Send(*SubscribeEventsResponse) error
	grpc.ServerStream
}

type eventsServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventsServiceSubscribeServer) Send(m *SubscribeEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventsService_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServiceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wgtwo.events.v0.EventsService/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServiceServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.events.v0.EventsService",
	HandlerType: (*EventsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ack",
			Handler:    _EventsService_Ack_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _EventsService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wgtwo/events/v0/events.proto",
}
