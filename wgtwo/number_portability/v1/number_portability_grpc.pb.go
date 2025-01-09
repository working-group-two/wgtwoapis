// Copyright 2024 Cisco Systems
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
// source: wgtwo/number_portability/v1/number_portability.proto

package v1

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
	NumberPortabilityService_CreatePortingRecords_FullMethodName = "/wgtwo.number_portability.v1.NumberPortabilityService/CreatePortingRecords"
	NumberPortabilityService_ListPortingRecords_FullMethodName   = "/wgtwo.number_portability.v1.NumberPortabilityService/ListPortingRecords"
)

// NumberPortabilityServiceClient is the client API for NumberPortabilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// NumberPortabilityService is supposed to be used by tenants or third parties to import country-specific number
// porting records into Working Group Two's number portability database. A porting record consists of a subscriber
// number, routing destination, porting date and optional routing code as well as tenant-specific metadata. Porting date
// can both be a past or a future date.
type NumberPortabilityServiceClient interface {
	// CreatePortingRecords is used to import porting records into the number portability database.
	// It may also be used to clear the porting record to return it to it's original number block -
	// for that, destination_id, operator_code and routing_code in the PortingRecords should be set to empty values.
	CreatePortingRecords(ctx context.Context, in *CreatePortingRecordsRequest, opts ...grpc.CallOption) (*CreatePortingRecordsResponse, error)
	// ListPortingRecords is used to list porting records from the number portability database.
	ListPortingRecords(ctx context.Context, in *ListPortingRecordsRequest, opts ...grpc.CallOption) (*ListPortingRecordsResponse, error)
}

type numberPortabilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberPortabilityServiceClient(cc grpc.ClientConnInterface) NumberPortabilityServiceClient {
	return &numberPortabilityServiceClient{cc}
}

func (c *numberPortabilityServiceClient) CreatePortingRecords(ctx context.Context, in *CreatePortingRecordsRequest, opts ...grpc.CallOption) (*CreatePortingRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePortingRecordsResponse)
	err := c.cc.Invoke(ctx, NumberPortabilityService_CreatePortingRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *numberPortabilityServiceClient) ListPortingRecords(ctx context.Context, in *ListPortingRecordsRequest, opts ...grpc.CallOption) (*ListPortingRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPortingRecordsResponse)
	err := c.cc.Invoke(ctx, NumberPortabilityService_ListPortingRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumberPortabilityServiceServer is the server API for NumberPortabilityService service.
// All implementations should embed UnimplementedNumberPortabilityServiceServer
// for forward compatibility.
//
// NumberPortabilityService is supposed to be used by tenants or third parties to import country-specific number
// porting records into Working Group Two's number portability database. A porting record consists of a subscriber
// number, routing destination, porting date and optional routing code as well as tenant-specific metadata. Porting date
// can both be a past or a future date.
type NumberPortabilityServiceServer interface {
	// CreatePortingRecords is used to import porting records into the number portability database.
	// It may also be used to clear the porting record to return it to it's original number block -
	// for that, destination_id, operator_code and routing_code in the PortingRecords should be set to empty values.
	CreatePortingRecords(context.Context, *CreatePortingRecordsRequest) (*CreatePortingRecordsResponse, error)
	// ListPortingRecords is used to list porting records from the number portability database.
	ListPortingRecords(context.Context, *ListPortingRecordsRequest) (*ListPortingRecordsResponse, error)
}

// UnimplementedNumberPortabilityServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNumberPortabilityServiceServer struct{}

func (UnimplementedNumberPortabilityServiceServer) CreatePortingRecords(context.Context, *CreatePortingRecordsRequest) (*CreatePortingRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePortingRecords not implemented")
}
func (UnimplementedNumberPortabilityServiceServer) ListPortingRecords(context.Context, *ListPortingRecordsRequest) (*ListPortingRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPortingRecords not implemented")
}
func (UnimplementedNumberPortabilityServiceServer) testEmbeddedByValue() {}

// UnsafeNumberPortabilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberPortabilityServiceServer will
// result in compilation errors.
type UnsafeNumberPortabilityServiceServer interface {
	mustEmbedUnimplementedNumberPortabilityServiceServer()
}

func RegisterNumberPortabilityServiceServer(s grpc.ServiceRegistrar, srv NumberPortabilityServiceServer) {
	// If the following call pancis, it indicates UnimplementedNumberPortabilityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NumberPortabilityService_ServiceDesc, srv)
}

func _NumberPortabilityService_CreatePortingRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortingRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberPortabilityServiceServer).CreatePortingRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NumberPortabilityService_CreatePortingRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberPortabilityServiceServer).CreatePortingRecords(ctx, req.(*CreatePortingRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NumberPortabilityService_ListPortingRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPortingRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberPortabilityServiceServer).ListPortingRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NumberPortabilityService_ListPortingRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberPortabilityServiceServer).ListPortingRecords(ctx, req.(*ListPortingRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NumberPortabilityService_ServiceDesc is the grpc.ServiceDesc for NumberPortabilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumberPortabilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.number_portability.v1.NumberPortabilityService",
	HandlerType: (*NumberPortabilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePortingRecords",
			Handler:    _NumberPortabilityService_CreatePortingRecords_Handler,
		},
		{
			MethodName: "ListPortingRecords",
			Handler:    _NumberPortabilityService_ListPortingRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/number_portability/v1/number_portability.proto",
}

const (
	RoutingDestinationService_CreateOrUpdateDestination_FullMethodName = "/wgtwo.number_portability.v1.RoutingDestinationService/CreateOrUpdateDestination"
	RoutingDestinationService_DeleteDestination_FullMethodName         = "/wgtwo.number_portability.v1.RoutingDestinationService/DeleteDestination"
	RoutingDestinationService_ListDestinations_FullMethodName          = "/wgtwo.number_portability.v1.RoutingDestinationService/ListDestinations"
)

// RoutingDestinationServiceClient is the client API for RoutingDestinationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RoutingDestinationService is supposed to be used by tenants or third parties to create, update, delete and list
// routing destinations, which define how SMS messages or called should be routed. A porting record may be associated
// either directly to the routing code or indirectly via routing destinations (destination_id field of the PortingRecord
// message). Indirect association facilitates routing code changes without updating all porting records.
type RoutingDestinationServiceClient interface {
	// Creates a new or updates an existing routing destination.
	CreateOrUpdateDestination(ctx context.Context, in *CreateOrUpdateDestinationRequest, opts ...grpc.CallOption) (*CreateOrUpdateDestinationResponse, error)
	// Deletes an existing routing destination.
	DeleteDestination(ctx context.Context, in *DeleteDestinationRequest, opts ...grpc.CallOption) (*DeleteDestinationResponse, error)
	// List existing mobile routing destinations.
	ListDestinations(ctx context.Context, in *ListDestinationsRequest, opts ...grpc.CallOption) (*ListDestinationsResponse, error)
}

type routingDestinationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingDestinationServiceClient(cc grpc.ClientConnInterface) RoutingDestinationServiceClient {
	return &routingDestinationServiceClient{cc}
}

func (c *routingDestinationServiceClient) CreateOrUpdateDestination(ctx context.Context, in *CreateOrUpdateDestinationRequest, opts ...grpc.CallOption) (*CreateOrUpdateDestinationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdateDestinationResponse)
	err := c.cc.Invoke(ctx, RoutingDestinationService_CreateOrUpdateDestination_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingDestinationServiceClient) DeleteDestination(ctx context.Context, in *DeleteDestinationRequest, opts ...grpc.CallOption) (*DeleteDestinationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDestinationResponse)
	err := c.cc.Invoke(ctx, RoutingDestinationService_DeleteDestination_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingDestinationServiceClient) ListDestinations(ctx context.Context, in *ListDestinationsRequest, opts ...grpc.CallOption) (*ListDestinationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDestinationsResponse)
	err := c.cc.Invoke(ctx, RoutingDestinationService_ListDestinations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingDestinationServiceServer is the server API for RoutingDestinationService service.
// All implementations should embed UnimplementedRoutingDestinationServiceServer
// for forward compatibility.
//
// RoutingDestinationService is supposed to be used by tenants or third parties to create, update, delete and list
// routing destinations, which define how SMS messages or called should be routed. A porting record may be associated
// either directly to the routing code or indirectly via routing destinations (destination_id field of the PortingRecord
// message). Indirect association facilitates routing code changes without updating all porting records.
type RoutingDestinationServiceServer interface {
	// Creates a new or updates an existing routing destination.
	CreateOrUpdateDestination(context.Context, *CreateOrUpdateDestinationRequest) (*CreateOrUpdateDestinationResponse, error)
	// Deletes an existing routing destination.
	DeleteDestination(context.Context, *DeleteDestinationRequest) (*DeleteDestinationResponse, error)
	// List existing mobile routing destinations.
	ListDestinations(context.Context, *ListDestinationsRequest) (*ListDestinationsResponse, error)
}

// UnimplementedRoutingDestinationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRoutingDestinationServiceServer struct{}

func (UnimplementedRoutingDestinationServiceServer) CreateOrUpdateDestination(context.Context, *CreateOrUpdateDestinationRequest) (*CreateOrUpdateDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateDestination not implemented")
}
func (UnimplementedRoutingDestinationServiceServer) DeleteDestination(context.Context, *DeleteDestinationRequest) (*DeleteDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDestination not implemented")
}
func (UnimplementedRoutingDestinationServiceServer) ListDestinations(context.Context, *ListDestinationsRequest) (*ListDestinationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDestinations not implemented")
}
func (UnimplementedRoutingDestinationServiceServer) testEmbeddedByValue() {}

// UnsafeRoutingDestinationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingDestinationServiceServer will
// result in compilation errors.
type UnsafeRoutingDestinationServiceServer interface {
	mustEmbedUnimplementedRoutingDestinationServiceServer()
}

func RegisterRoutingDestinationServiceServer(s grpc.ServiceRegistrar, srv RoutingDestinationServiceServer) {
	// If the following call pancis, it indicates UnimplementedRoutingDestinationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RoutingDestinationService_ServiceDesc, srv)
}

func _RoutingDestinationService_CreateOrUpdateDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingDestinationServiceServer).CreateOrUpdateDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutingDestinationService_CreateOrUpdateDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingDestinationServiceServer).CreateOrUpdateDestination(ctx, req.(*CreateOrUpdateDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingDestinationService_DeleteDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingDestinationServiceServer).DeleteDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutingDestinationService_DeleteDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingDestinationServiceServer).DeleteDestination(ctx, req.(*DeleteDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingDestinationService_ListDestinations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDestinationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingDestinationServiceServer).ListDestinations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RoutingDestinationService_ListDestinations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingDestinationServiceServer).ListDestinations(ctx, req.(*ListDestinationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoutingDestinationService_ServiceDesc is the grpc.ServiceDesc for RoutingDestinationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingDestinationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.number_portability.v1.RoutingDestinationService",
	HandlerType: (*RoutingDestinationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateDestination",
			Handler:    _RoutingDestinationService_CreateOrUpdateDestination_Handler,
		},
		{
			MethodName: "DeleteDestination",
			Handler:    _RoutingDestinationService_DeleteDestination_Handler,
		},
		{
			MethodName: "ListDestinations",
			Handler:    _RoutingDestinationService_ListDestinations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/number_portability/v1/number_portability.proto",
}

const (
	NumberBlockService_CreateOrUpdateNumberBlock_FullMethodName = "/wgtwo.number_portability.v1.NumberBlockService/CreateOrUpdateNumberBlock"
	NumberBlockService_DeleteNumberBlock_FullMethodName         = "/wgtwo.number_portability.v1.NumberBlockService/DeleteNumberBlock"
	NumberBlockService_ListNumberBlocks_FullMethodName          = "/wgtwo.number_portability.v1.NumberBlockService/ListNumberBlocks"
)

// NumberBlockServiceClient is the client API for NumberBlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// In some countries, a subscriber number itself may not be sufficient to determine the routing code for non-ported
// numbers.
// For example, in Belgium, an earliest number prefix allocated to an operator must be used, separately for fixed-line
// and mobile numbers.
// NumberBlockService allows to create, update, delete and list number blocks allocated to operators
// by a regulator.
type NumberBlockServiceClient interface {
	// Creates a new or updates an existing number block.
	CreateOrUpdateNumberBlock(ctx context.Context, in *CreateOrUpdateNumberBlockRequest, opts ...grpc.CallOption) (*CreateOrUpdateNumberBlockResponse, error)
	// Deletes an existing number block.
	DeleteNumberBlock(ctx context.Context, in *DeleteNumberBlockRequest, opts ...grpc.CallOption) (*DeleteNumberBlockResponse, error)
	// List existing number blocks.
	ListNumberBlocks(ctx context.Context, in *ListNumberBlocksRequest, opts ...grpc.CallOption) (*ListNumberBlocksResponse, error)
}

type numberBlockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberBlockServiceClient(cc grpc.ClientConnInterface) NumberBlockServiceClient {
	return &numberBlockServiceClient{cc}
}

func (c *numberBlockServiceClient) CreateOrUpdateNumberBlock(ctx context.Context, in *CreateOrUpdateNumberBlockRequest, opts ...grpc.CallOption) (*CreateOrUpdateNumberBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdateNumberBlockResponse)
	err := c.cc.Invoke(ctx, NumberBlockService_CreateOrUpdateNumberBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *numberBlockServiceClient) DeleteNumberBlock(ctx context.Context, in *DeleteNumberBlockRequest, opts ...grpc.CallOption) (*DeleteNumberBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteNumberBlockResponse)
	err := c.cc.Invoke(ctx, NumberBlockService_DeleteNumberBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *numberBlockServiceClient) ListNumberBlocks(ctx context.Context, in *ListNumberBlocksRequest, opts ...grpc.CallOption) (*ListNumberBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNumberBlocksResponse)
	err := c.cc.Invoke(ctx, NumberBlockService_ListNumberBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumberBlockServiceServer is the server API for NumberBlockService service.
// All implementations should embed UnimplementedNumberBlockServiceServer
// for forward compatibility.
//
// In some countries, a subscriber number itself may not be sufficient to determine the routing code for non-ported
// numbers.
// For example, in Belgium, an earliest number prefix allocated to an operator must be used, separately for fixed-line
// and mobile numbers.
// NumberBlockService allows to create, update, delete and list number blocks allocated to operators
// by a regulator.
type NumberBlockServiceServer interface {
	// Creates a new or updates an existing number block.
	CreateOrUpdateNumberBlock(context.Context, *CreateOrUpdateNumberBlockRequest) (*CreateOrUpdateNumberBlockResponse, error)
	// Deletes an existing number block.
	DeleteNumberBlock(context.Context, *DeleteNumberBlockRequest) (*DeleteNumberBlockResponse, error)
	// List existing number blocks.
	ListNumberBlocks(context.Context, *ListNumberBlocksRequest) (*ListNumberBlocksResponse, error)
}

// UnimplementedNumberBlockServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNumberBlockServiceServer struct{}

func (UnimplementedNumberBlockServiceServer) CreateOrUpdateNumberBlock(context.Context, *CreateOrUpdateNumberBlockRequest) (*CreateOrUpdateNumberBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateNumberBlock not implemented")
}
func (UnimplementedNumberBlockServiceServer) DeleteNumberBlock(context.Context, *DeleteNumberBlockRequest) (*DeleteNumberBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNumberBlock not implemented")
}
func (UnimplementedNumberBlockServiceServer) ListNumberBlocks(context.Context, *ListNumberBlocksRequest) (*ListNumberBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNumberBlocks not implemented")
}
func (UnimplementedNumberBlockServiceServer) testEmbeddedByValue() {}

// UnsafeNumberBlockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberBlockServiceServer will
// result in compilation errors.
type UnsafeNumberBlockServiceServer interface {
	mustEmbedUnimplementedNumberBlockServiceServer()
}

func RegisterNumberBlockServiceServer(s grpc.ServiceRegistrar, srv NumberBlockServiceServer) {
	// If the following call pancis, it indicates UnimplementedNumberBlockServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NumberBlockService_ServiceDesc, srv)
}

func _NumberBlockService_CreateOrUpdateNumberBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateNumberBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberBlockServiceServer).CreateOrUpdateNumberBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NumberBlockService_CreateOrUpdateNumberBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberBlockServiceServer).CreateOrUpdateNumberBlock(ctx, req.(*CreateOrUpdateNumberBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NumberBlockService_DeleteNumberBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNumberBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberBlockServiceServer).DeleteNumberBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NumberBlockService_DeleteNumberBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberBlockServiceServer).DeleteNumberBlock(ctx, req.(*DeleteNumberBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NumberBlockService_ListNumberBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNumberBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberBlockServiceServer).ListNumberBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NumberBlockService_ListNumberBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberBlockServiceServer).ListNumberBlocks(ctx, req.(*ListNumberBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NumberBlockService_ServiceDesc is the grpc.ServiceDesc for NumberBlockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumberBlockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wgtwo.number_portability.v1.NumberBlockService",
	HandlerType: (*NumberBlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateNumberBlock",
			Handler:    _NumberBlockService_CreateOrUpdateNumberBlock_Handler,
		},
		{
			MethodName: "DeleteNumberBlock",
			Handler:    _NumberBlockService_DeleteNumberBlock_Handler,
		},
		{
			MethodName: "ListNumberBlocks",
			Handler:    _NumberBlockService_ListNumberBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wgtwo/number_portability/v1/number_portability.proto",
}