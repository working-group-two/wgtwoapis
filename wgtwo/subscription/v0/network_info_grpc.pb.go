// Copyright 2022 Working Group Two AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: wgtwo/subscription/v0/network_info.proto

package v0

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
	NetworkInfoService_GetNetworkInfoForSubscriber_FullMethodName        = "/wgtwo.subscription.v0.NetworkInfoService/GetNetworkInfoForSubscriber"
	NetworkInfoService_GetAttachmentAttemptsForSubscriber_FullMethodName = "/wgtwo.subscription.v0.NetworkInfoService/GetAttachmentAttemptsForSubscriber"
)

// NetworkInfoServiceClient is the client API for NetworkInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkInfoServiceClient interface {
	// Get current network information for subscriber
	GetNetworkInfoForSubscriber(ctx context.Context, in *GetNetworkInfoForSubscriberRequest, opts ...grpc.CallOption) (*GetNetworkInfoForSubscriberResponse, error)
	// Get attachment attempts, both failed and successful, for subscriber
	GetAttachmentAttemptsForSubscriber(ctx context.Context, in *GetAttachmentAttemptsForSubscriberRequest, opts ...grpc.CallOption) (*GetAttachmentAttemptsForSubscriberResponse, error)
}

type networkInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkInfoServiceClient(cc grpc.ClientConnInterface) NetworkInfoServiceClient {
	return &networkInfoServiceClient{cc}
}

func (c *networkInfoServiceClient) GetNetworkInfoForSubscriber(ctx context.Context, in *GetNetworkInfoForSubscriberRequest, opts ...grpc.CallOption) (*GetNetworkInfoForSubscriberResponse, error) {
	out := new(GetNetworkInfoForSubscriberResponse)
	err := c.cc.Invoke(ctx, NetworkInfoService_GetNetworkInfoForSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkInfoServiceClient) GetAttachmentAttemptsForSubscriber(ctx context.Context, in *GetAttachmentAttemptsForSubscriberRequest, opts ...grpc.CallOption) (*GetAttachmentAttemptsForSubscriberResponse, error) {
	out := new(GetAttachmentAttemptsForSubscriberResponse)
	err := c.cc.Invoke(ctx, NetworkInfoService_GetAttachmentAttemptsForSubscriber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkInfoServiceServer is the server API for NetworkInfoService service.
// All implementations should embed UnimplementedNetworkInfoServiceServer
// for forward compatibility
type NetworkInfoServiceServer interface {
	// Get current network information for subscriber
	GetNetworkInfoForSubscriber(context.Context, *GetNetworkInfoForSubscriberRequest) (*GetNetworkInfoForSubscriberResponse, error)
	// Get attachment attempts, both failed and successful, for subscriber
	GetAttachmentAttemptsForSubscriber(context.Context, *GetAttachmentAttemptsForSubscriberRequest) (*GetAttachmentAttemptsForSubscriberResponse, error)
}

// UnimplementedNetworkInfoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNetworkInfoServiceServer struct {
}

func (UnimplementedNetworkInfoServiceServer) GetNetworkInfoForSubscriber(context.Context, *GetNetworkInfoForSubscriberRequest) (*GetNetworkInfoForSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetworkInfoForSubscriber not implemented")
}
func (UnimplementedNetworkInfoServiceServer) GetAttachmentAttemptsForSubscriber(context.Context, *GetAttachmentAttemptsForSubscriberRequest) (*GetAttachmentAttemptsForSubscriberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttachmentAttemptsForSubscriber not implemented")
}

// UnsafeNetworkInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkInfoServiceServer will
// result in compilation errors.
type UnsafeNetworkInfoServiceServer interface {
	mustEmbedUnimplementedNetworkInfoServiceServer()
}

func RegisterNetworkInfoServiceServer(s grpc.ServiceRegistrar, srv NetworkInfoServiceServer) {
	s.RegisterService(&NetworkInfoService_ServiceDesc, srv)
}

func _NetworkInfoService_GetNetworkInfoForSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkInfoForSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkInfoServiceServer).GetNetworkInfoForSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkInfoService_GetNetworkInfoForSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkInfoServiceServer).GetNetworkInfoForSubscriber(ctx, req.(*GetNetworkInfoForSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkInfoService_GetAttachmentAttemptsForSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttachmentAttemptsForSubscriberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkInfoServiceServer).GetAttachmentAttemptsForSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkInfoService_GetAttachmentAttemptsForSubscriber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkInfoServiceServer).GetAttachmentAttemptsForSubscriber(ctx, req.(*GetAttachmentAttemptsForSubscriberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkInfoService_ServiceDesc is the grpc.ServiceDesc for NetworkInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.subscription.v0.NetworkInfoService",
	HandlerType: (*NetworkInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkInfoForSubscriber",
			Handler:    _NetworkInfoService_GetNetworkInfoForSubscriber_Handler,
		},
		{
			MethodName: "GetAttachmentAttemptsForSubscriber",
			Handler:    _NetworkInfoService_GetAttachmentAttemptsForSubscriber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/subscription/v0/network_info.proto",
}
