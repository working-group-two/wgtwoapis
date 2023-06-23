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
// source: wgtwo/sipbreakout/v1/sipbreakout.proto

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

const (
	SipBreakoutService_UpsertRegistration_FullMethodName = "/wgtwo.sipbreakout.v1.SipBreakoutService/UpsertRegistration"
	SipBreakoutService_DeleteRegistration_FullMethodName = "/wgtwo.sipbreakout.v1.SipBreakoutService/DeleteRegistration"
)

// SipBreakoutServiceClient is the client API for SipBreakoutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SipBreakoutServiceClient interface {
	// Add or replace a registration
	UpsertRegistration(ctx context.Context, in *UpsertRegistrationRequest, opts ...grpc.CallOption) (*UpsertRegistrationResponse, error)
	// Delete an existing registration
	DeleteRegistration(ctx context.Context, in *DeleteRegistrationRequest, opts ...grpc.CallOption) (*DeleteRegistrationResponse, error)
}

type sipBreakoutServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSipBreakoutServiceClient(cc grpc.ClientConnInterface) SipBreakoutServiceClient {
	return &sipBreakoutServiceClient{cc}
}

func (c *sipBreakoutServiceClient) UpsertRegistration(ctx context.Context, in *UpsertRegistrationRequest, opts ...grpc.CallOption) (*UpsertRegistrationResponse, error) {
	out := new(UpsertRegistrationResponse)
	err := c.cc.Invoke(ctx, SipBreakoutService_UpsertRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sipBreakoutServiceClient) DeleteRegistration(ctx context.Context, in *DeleteRegistrationRequest, opts ...grpc.CallOption) (*DeleteRegistrationResponse, error) {
	out := new(DeleteRegistrationResponse)
	err := c.cc.Invoke(ctx, SipBreakoutService_DeleteRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SipBreakoutServiceServer is the server API for SipBreakoutService service.
// All implementations should embed UnimplementedSipBreakoutServiceServer
// for forward compatibility
type SipBreakoutServiceServer interface {
	// Add or replace a registration
	UpsertRegistration(context.Context, *UpsertRegistrationRequest) (*UpsertRegistrationResponse, error)
	// Delete an existing registration
	DeleteRegistration(context.Context, *DeleteRegistrationRequest) (*DeleteRegistrationResponse, error)
}

// UnimplementedSipBreakoutServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSipBreakoutServiceServer struct {
}

func (UnimplementedSipBreakoutServiceServer) UpsertRegistration(context.Context, *UpsertRegistrationRequest) (*UpsertRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertRegistration not implemented")
}
func (UnimplementedSipBreakoutServiceServer) DeleteRegistration(context.Context, *DeleteRegistrationRequest) (*DeleteRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegistration not implemented")
}

// UnsafeSipBreakoutServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SipBreakoutServiceServer will
// result in compilation errors.
type UnsafeSipBreakoutServiceServer interface {
	mustEmbedUnimplementedSipBreakoutServiceServer()
}

func RegisterSipBreakoutServiceServer(s grpc.ServiceRegistrar, srv SipBreakoutServiceServer) {
	s.RegisterService(&SipBreakoutService_ServiceDesc, srv)
}

func _SipBreakoutService_UpsertRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SipBreakoutServiceServer).UpsertRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SipBreakoutService_UpsertRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SipBreakoutServiceServer).UpsertRegistration(ctx, req.(*UpsertRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SipBreakoutService_DeleteRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SipBreakoutServiceServer).DeleteRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SipBreakoutService_DeleteRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SipBreakoutServiceServer).DeleteRegistration(ctx, req.(*DeleteRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SipBreakoutService_ServiceDesc is the grpc.ServiceDesc for SipBreakoutService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SipBreakoutService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.sipbreakout.v1.SipBreakoutService",
	HandlerType: (*SipBreakoutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertRegistration",
			Handler:    _SipBreakoutService_UpsertRegistration_Handler,
		},
		{
			MethodName: "DeleteRegistration",
			Handler:    _SipBreakoutService_DeleteRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/sipbreakout/v1/sipbreakout.proto",
}
