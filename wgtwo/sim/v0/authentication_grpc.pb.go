// Copyright 2024 [Working Group Two]/[Cisco Systems]
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
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: wgtwo/sim/v0/authentication.proto

package v0

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SimAuthenticationService_GenerateEapAkaAuthenticationVector_FullMethodName = "/wgtwo.sim.v0.SimAuthenticationService/GenerateEapAkaAuthenticationVector"
)

// SimAuthenticationServiceClient is the client API for SimAuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimAuthenticationServiceClient interface {
	GenerateEapAkaAuthenticationVector(ctx context.Context, in *GenerateEapAkaAuthenticationVectorRequest, opts ...grpc.CallOption) (*GenerateEapAkaAuthenticationVectorResponse, error)
}

type simAuthenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimAuthenticationServiceClient(cc grpc.ClientConnInterface) SimAuthenticationServiceClient {
	return &simAuthenticationServiceClient{cc}
}

func (c *simAuthenticationServiceClient) GenerateEapAkaAuthenticationVector(ctx context.Context, in *GenerateEapAkaAuthenticationVectorRequest, opts ...grpc.CallOption) (*GenerateEapAkaAuthenticationVectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateEapAkaAuthenticationVectorResponse)
	err := c.cc.Invoke(ctx, SimAuthenticationService_GenerateEapAkaAuthenticationVector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimAuthenticationServiceServer is the server API for SimAuthenticationService service.
// All implementations should embed UnimplementedSimAuthenticationServiceServer
// for forward compatibility.
type SimAuthenticationServiceServer interface {
	GenerateEapAkaAuthenticationVector(context.Context, *GenerateEapAkaAuthenticationVectorRequest) (*GenerateEapAkaAuthenticationVectorResponse, error)
}

// UnimplementedSimAuthenticationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSimAuthenticationServiceServer struct{}

func (UnimplementedSimAuthenticationServiceServer) GenerateEapAkaAuthenticationVector(context.Context, *GenerateEapAkaAuthenticationVectorRequest) (*GenerateEapAkaAuthenticationVectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateEapAkaAuthenticationVector not implemented")
}
func (UnimplementedSimAuthenticationServiceServer) testEmbeddedByValue() {}

// UnsafeSimAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimAuthenticationServiceServer will
// result in compilation errors.
type UnsafeSimAuthenticationServiceServer interface {
	mustEmbedUnimplementedSimAuthenticationServiceServer()
}

func RegisterSimAuthenticationServiceServer(s grpc.ServiceRegistrar, srv SimAuthenticationServiceServer) {
	// If the following call pancis, it indicates UnimplementedSimAuthenticationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SimAuthenticationService_ServiceDesc, srv)
}

func _SimAuthenticationService_GenerateEapAkaAuthenticationVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateEapAkaAuthenticationVectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimAuthenticationServiceServer).GenerateEapAkaAuthenticationVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimAuthenticationService_GenerateEapAkaAuthenticationVector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimAuthenticationServiceServer).GenerateEapAkaAuthenticationVector(ctx, req.(*GenerateEapAkaAuthenticationVectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimAuthenticationService_ServiceDesc is the grpc.ServiceDesc for SimAuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimAuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.sim.v0.SimAuthenticationService",
	HandlerType: (*SimAuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateEapAkaAuthenticationVector",
			Handler:    _SimAuthenticationService_GenerateEapAkaAuthenticationVector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/sim/v0/authentication.proto",
}
