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
// source: wgtwo/subscription/v0/subscription_id.proto

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
	SubscriptionIdService_GetSubscriptionId_FullMethodName = "/wgtwo.subscription.v0.SubscriptionIdService/GetSubscriptionId"
	SubscriptionIdService_GetMsisdn_FullMethodName         = "/wgtwo.subscription.v0.SubscriptionIdService/GetMsisdn"
)

// SubscriptionIdServiceClient is the client API for SubscriptionIdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionIdServiceClient interface {
	// Get globally unique subscription id for a msisdn
	GetSubscriptionId(ctx context.Context, in *GetSubscriptionIdRequest, opts ...grpc.CallOption) (*GetSubscriptionIdResponse, error)
	// Get msisdn for a globally unique subscription id
	GetMsisdn(ctx context.Context, in *GetMsisdnRequest, opts ...grpc.CallOption) (*GetMsisdnResponse, error)
}

type subscriptionIdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionIdServiceClient(cc grpc.ClientConnInterface) SubscriptionIdServiceClient {
	return &subscriptionIdServiceClient{cc}
}

func (c *subscriptionIdServiceClient) GetSubscriptionId(ctx context.Context, in *GetSubscriptionIdRequest, opts ...grpc.CallOption) (*GetSubscriptionIdResponse, error) {
	out := new(GetSubscriptionIdResponse)
	err := c.cc.Invoke(ctx, SubscriptionIdService_GetSubscriptionId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionIdServiceClient) GetMsisdn(ctx context.Context, in *GetMsisdnRequest, opts ...grpc.CallOption) (*GetMsisdnResponse, error) {
	out := new(GetMsisdnResponse)
	err := c.cc.Invoke(ctx, SubscriptionIdService_GetMsisdn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionIdServiceServer is the server API for SubscriptionIdService service.
// All implementations should embed UnimplementedSubscriptionIdServiceServer
// for forward compatibility
type SubscriptionIdServiceServer interface {
	// Get globally unique subscription id for a msisdn
	GetSubscriptionId(context.Context, *GetSubscriptionIdRequest) (*GetSubscriptionIdResponse, error)
	// Get msisdn for a globally unique subscription id
	GetMsisdn(context.Context, *GetMsisdnRequest) (*GetMsisdnResponse, error)
}

// UnimplementedSubscriptionIdServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubscriptionIdServiceServer struct {
}

func (UnimplementedSubscriptionIdServiceServer) GetSubscriptionId(context.Context, *GetSubscriptionIdRequest) (*GetSubscriptionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionId not implemented")
}
func (UnimplementedSubscriptionIdServiceServer) GetMsisdn(context.Context, *GetMsisdnRequest) (*GetMsisdnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsisdn not implemented")
}

// UnsafeSubscriptionIdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionIdServiceServer will
// result in compilation errors.
type UnsafeSubscriptionIdServiceServer interface {
	mustEmbedUnimplementedSubscriptionIdServiceServer()
}

func RegisterSubscriptionIdServiceServer(s grpc.ServiceRegistrar, srv SubscriptionIdServiceServer) {
	s.RegisterService(&SubscriptionIdService_ServiceDesc, srv)
}

func _SubscriptionIdService_GetSubscriptionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionIdServiceServer).GetSubscriptionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionIdService_GetSubscriptionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionIdServiceServer).GetSubscriptionId(ctx, req.(*GetSubscriptionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionIdService_GetMsisdn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsisdnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionIdServiceServer).GetMsisdn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionIdService_GetMsisdn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionIdServiceServer).GetMsisdn(ctx, req.(*GetMsisdnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionIdService_ServiceDesc is the grpc.ServiceDesc for SubscriptionIdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionIdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.subscription.v0.SubscriptionIdService",
	HandlerType: (*SubscriptionIdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscriptionId",
			Handler:    _SubscriptionIdService_GetSubscriptionId_Handler,
		},
		{
			MethodName: "GetMsisdn",
			Handler:    _SubscriptionIdService_GetMsisdn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/subscription/v0/subscription_id.proto",
}