// Copyright 2020 Working Group Two AS
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: wgtwo/data/v0/data.proto

package v0

import (
	_ "github.com/working-group-two/wgtwoapis/wgtwo/annotations"
	v0 "github.com/working-group-two/wgtwoapis/wgtwo/common/v0"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CustomDnsResponse_Status int32

const (
	CustomDnsResponse_UNKNOWN  CustomDnsResponse_Status = 0
	CustomDnsResponse_ACCEPTED CustomDnsResponse_Status = 1
	CustomDnsResponse_REJECTED CustomDnsResponse_Status = 2
)

// Enum value maps for CustomDnsResponse_Status.
var (
	CustomDnsResponse_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "ACCEPTED",
		2: "REJECTED",
	}
	CustomDnsResponse_Status_value = map[string]int32{
		"UNKNOWN":  0,
		"ACCEPTED": 1,
		"REJECTED": 2,
	}
)

func (x CustomDnsResponse_Status) Enum() *CustomDnsResponse_Status {
	p := new(CustomDnsResponse_Status)
	*p = x
	return p
}

func (x CustomDnsResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomDnsResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_data_v0_data_proto_enumTypes[0].Descriptor()
}

func (CustomDnsResponse_Status) Type() protoreflect.EnumType {
	return &file_wgtwo_data_v0_data_proto_enumTypes[0]
}

func (x CustomDnsResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomDnsResponse_Status.Descriptor instead.
func (CustomDnsResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{7, 0}
}

type SetCustomDnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriber *v0.PhoneNumber `protobuf:"bytes,1,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	DnsV4      *IpV4Pair       `protobuf:"bytes,2,opt,name=dnsV4,proto3" json:"dnsV4,omitempty"`
	DnsV6      *IpV6Pair       `protobuf:"bytes,3,opt,name=dnsV6,proto3" json:"dnsV6,omitempty"`
}

func (x *SetCustomDnsRequest) Reset() {
	*x = SetCustomDnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetCustomDnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCustomDnsRequest) ProtoMessage() {}

func (x *SetCustomDnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetCustomDnsRequest.ProtoReflect.Descriptor instead.
func (*SetCustomDnsRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{0}
}

func (x *SetCustomDnsRequest) GetSubscriber() *v0.PhoneNumber {
	if x != nil {
		return x.Subscriber
	}
	return nil
}

func (x *SetCustomDnsRequest) GetDnsV4() *IpV4Pair {
	if x != nil {
		return x.DnsV4
	}
	return nil
}

func (x *SetCustomDnsRequest) GetDnsV6() *IpV6Pair {
	if x != nil {
		return x.DnsV6
	}
	return nil
}

type GetCustomDnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriber *v0.PhoneNumber `protobuf:"bytes,1,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
}

func (x *GetCustomDnsRequest) Reset() {
	*x = GetCustomDnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomDnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomDnsRequest) ProtoMessage() {}

func (x *GetCustomDnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomDnsRequest.ProtoReflect.Descriptor instead.
func (*GetCustomDnsRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{1}
}

func (x *GetCustomDnsRequest) GetSubscriber() *v0.PhoneNumber {
	if x != nil {
		return x.Subscriber
	}
	return nil
}

type ClearCustomDnsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriber *v0.PhoneNumber `protobuf:"bytes,1,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
}

func (x *ClearCustomDnsRequest) Reset() {
	*x = ClearCustomDnsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearCustomDnsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearCustomDnsRequest) ProtoMessage() {}

func (x *ClearCustomDnsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearCustomDnsRequest.ProtoReflect.Descriptor instead.
func (*ClearCustomDnsRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{2}
}

func (x *ClearCustomDnsRequest) GetSubscriber() *v0.PhoneNumber {
	if x != nil {
		return x.Subscriber
	}
	return nil
}

type IpV4Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Primary   *IpAddressV4 `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	Secondary *IpAddressV4 `protobuf:"bytes,2,opt,name=secondary,proto3" json:"secondary,omitempty"`
}

func (x *IpV4Pair) Reset() {
	*x = IpV4Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpV4Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpV4Pair) ProtoMessage() {}

func (x *IpV4Pair) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpV4Pair.ProtoReflect.Descriptor instead.
func (*IpV4Pair) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{3}
}

func (x *IpV4Pair) GetPrimary() *IpAddressV4 {
	if x != nil {
		return x.Primary
	}
	return nil
}

func (x *IpV4Pair) GetSecondary() *IpAddressV4 {
	if x != nil {
		return x.Secondary
	}
	return nil
}

type IpAddressV4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipv4 string `protobuf:"bytes,1,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
}

func (x *IpAddressV4) Reset() {
	*x = IpAddressV4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpAddressV4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpAddressV4) ProtoMessage() {}

func (x *IpAddressV4) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpAddressV4.ProtoReflect.Descriptor instead.
func (*IpAddressV4) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{4}
}

func (x *IpAddressV4) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

type IpV6Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Primary   *IpAddressV6 `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	Secondary *IpAddressV6 `protobuf:"bytes,2,opt,name=secondary,proto3" json:"secondary,omitempty"`
}

func (x *IpV6Pair) Reset() {
	*x = IpV6Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpV6Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpV6Pair) ProtoMessage() {}

func (x *IpV6Pair) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpV6Pair.ProtoReflect.Descriptor instead.
func (*IpV6Pair) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{5}
}

func (x *IpV6Pair) GetPrimary() *IpAddressV6 {
	if x != nil {
		return x.Primary
	}
	return nil
}

func (x *IpV6Pair) GetSecondary() *IpAddressV6 {
	if x != nil {
		return x.Secondary
	}
	return nil
}

type IpAddressV6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipv6 string `protobuf:"bytes,1,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
}

func (x *IpAddressV6) Reset() {
	*x = IpAddressV6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpAddressV6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpAddressV6) ProtoMessage() {}

func (x *IpAddressV6) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpAddressV6.ProtoReflect.Descriptor instead.
func (*IpAddressV6) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{6}
}

func (x *IpAddressV6) GetIpv6() string {
	if x != nil {
		return x.Ipv6
	}
	return ""
}

type CustomDnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       CustomDnsResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=wgtwo.data.v0.CustomDnsResponse_Status" json:"status,omitempty"`
	ErrorMessage string                   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CustomDnsResponse) Reset() {
	*x = CustomDnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomDnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomDnsResponse) ProtoMessage() {}

func (x *CustomDnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomDnsResponse.ProtoReflect.Descriptor instead.
func (*CustomDnsResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{7}
}

func (x *CustomDnsResponse) GetStatus() CustomDnsResponse_Status {
	if x != nil {
		return x.Status
	}
	return CustomDnsResponse_UNKNOWN
}

func (x *CustomDnsResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetCustomDnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsV4 *IpV4Pair `protobuf:"bytes,1,opt,name=dnsV4,proto3" json:"dnsV4,omitempty"`
	DnsV6 *IpV6Pair `protobuf:"bytes,2,opt,name=dnsV6,proto3" json:"dnsV6,omitempty"`
}

func (x *GetCustomDnsResponse) Reset() {
	*x = GetCustomDnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_data_v0_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomDnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomDnsResponse) ProtoMessage() {}

func (x *GetCustomDnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_data_v0_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomDnsResponse.ProtoReflect.Descriptor instead.
func (*GetCustomDnsResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_data_v0_data_proto_rawDescGZIP(), []int{8}
}

func (x *GetCustomDnsResponse) GetDnsV4() *IpV4Pair {
	if x != nil {
		return x.DnsV4
	}
	return nil
}

func (x *GetCustomDnsResponse) GetDnsV6() *IpV6Pair {
	if x != nil {
		return x.DnsV6
	}
	return nil
}

var File_wgtwo_data_v0_data_proto protoreflect.FileDescriptor

var file_wgtwo_data_v0_data_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x30, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x1a, 0x23, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x64, 0x6e, 0x73, 0x56, 0x34,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70, 0x56, 0x34, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x05, 0x64, 0x6e, 0x73, 0x56, 0x34, 0x12, 0x2d, 0x0a, 0x05, 0x64, 0x6e, 0x73, 0x56, 0x36, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70, 0x56, 0x36, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x64, 0x6e, 0x73, 0x56, 0x36, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0a,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x22, 0x7a, 0x0a, 0x08, 0x49, 0x70, 0x56, 0x34, 0x50, 0x61, 0x69, 0x72, 0x12, 0x34, 0x0a,
	0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x34, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x56, 0x34, 0x52, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x22, 0x21, 0x0a,
	0x0b, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x34, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34,
	0x22, 0x7a, 0x0a, 0x08, 0x49, 0x70, 0x56, 0x36, 0x50, 0x61, 0x69, 0x72, 0x12, 0x34, 0x0a, 0x07,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x36, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56,
	0x36, 0x52, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x22, 0x21, 0x0a, 0x0b,
	0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x56, 0x36, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x70, 0x76, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x22,
	0xac, 0x01, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x74,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x64, 0x6e, 0x73, 0x56, 0x34, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70, 0x56, 0x34, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x64, 0x6e, 0x73, 0x56, 0x34, 0x12, 0x2d, 0x0a, 0x05, 0x64, 0x6e, 0x73, 0x56, 0x36, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x70, 0x56, 0x36, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x64,
	0x6e, 0x73, 0x56, 0x36, 0x32, 0xcc, 0x02, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xea, 0xb5, 0x18, 0x0c,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x2e, 0x64, 0x6e, 0x73, 0x12, 0x6a, 0x0a, 0x0e,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x12, 0x24,
	0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xea, 0xb5, 0x18, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x73, 0x65, 0x74, 0x2e, 0x64, 0x6e, 0x73, 0x12, 0x69, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x12, 0x22, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0xea, 0xb5, 0x18, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x65, 0x74, 0x2e,
	0x64, 0x6e, 0x73, 0x42, 0x58, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_data_v0_data_proto_rawDescOnce sync.Once
	file_wgtwo_data_v0_data_proto_rawDescData = file_wgtwo_data_v0_data_proto_rawDesc
)

func file_wgtwo_data_v0_data_proto_rawDescGZIP() []byte {
	file_wgtwo_data_v0_data_proto_rawDescOnce.Do(func() {
		file_wgtwo_data_v0_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_data_v0_data_proto_rawDescData)
	})
	return file_wgtwo_data_v0_data_proto_rawDescData
}

var file_wgtwo_data_v0_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wgtwo_data_v0_data_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_wgtwo_data_v0_data_proto_goTypes = []interface{}{
	(CustomDnsResponse_Status)(0), // 0: wgtwo.data.v0.CustomDnsResponse.Status
	(*SetCustomDnsRequest)(nil),   // 1: wgtwo.data.v0.SetCustomDnsRequest
	(*GetCustomDnsRequest)(nil),   // 2: wgtwo.data.v0.GetCustomDnsRequest
	(*ClearCustomDnsRequest)(nil), // 3: wgtwo.data.v0.ClearCustomDnsRequest
	(*IpV4Pair)(nil),              // 4: wgtwo.data.v0.IpV4Pair
	(*IpAddressV4)(nil),           // 5: wgtwo.data.v0.IpAddressV4
	(*IpV6Pair)(nil),              // 6: wgtwo.data.v0.IpV6Pair
	(*IpAddressV6)(nil),           // 7: wgtwo.data.v0.IpAddressV6
	(*CustomDnsResponse)(nil),     // 8: wgtwo.data.v0.CustomDnsResponse
	(*GetCustomDnsResponse)(nil),  // 9: wgtwo.data.v0.GetCustomDnsResponse
	(*v0.PhoneNumber)(nil),        // 10: wgtwo.common.v0.PhoneNumber
}
var file_wgtwo_data_v0_data_proto_depIdxs = []int32{
	10, // 0: wgtwo.data.v0.SetCustomDnsRequest.subscriber:type_name -> wgtwo.common.v0.PhoneNumber
	4,  // 1: wgtwo.data.v0.SetCustomDnsRequest.dnsV4:type_name -> wgtwo.data.v0.IpV4Pair
	6,  // 2: wgtwo.data.v0.SetCustomDnsRequest.dnsV6:type_name -> wgtwo.data.v0.IpV6Pair
	10, // 3: wgtwo.data.v0.GetCustomDnsRequest.subscriber:type_name -> wgtwo.common.v0.PhoneNumber
	10, // 4: wgtwo.data.v0.ClearCustomDnsRequest.subscriber:type_name -> wgtwo.common.v0.PhoneNumber
	5,  // 5: wgtwo.data.v0.IpV4Pair.primary:type_name -> wgtwo.data.v0.IpAddressV4
	5,  // 6: wgtwo.data.v0.IpV4Pair.secondary:type_name -> wgtwo.data.v0.IpAddressV4
	7,  // 7: wgtwo.data.v0.IpV6Pair.primary:type_name -> wgtwo.data.v0.IpAddressV6
	7,  // 8: wgtwo.data.v0.IpV6Pair.secondary:type_name -> wgtwo.data.v0.IpAddressV6
	0,  // 9: wgtwo.data.v0.CustomDnsResponse.status:type_name -> wgtwo.data.v0.CustomDnsResponse.Status
	4,  // 10: wgtwo.data.v0.GetCustomDnsResponse.dnsV4:type_name -> wgtwo.data.v0.IpV4Pair
	6,  // 11: wgtwo.data.v0.GetCustomDnsResponse.dnsV6:type_name -> wgtwo.data.v0.IpV6Pair
	1,  // 12: wgtwo.data.v0.DataService.SetCustomDns:input_type -> wgtwo.data.v0.SetCustomDnsRequest
	3,  // 13: wgtwo.data.v0.DataService.ClearCustomDns:input_type -> wgtwo.data.v0.ClearCustomDnsRequest
	2,  // 14: wgtwo.data.v0.DataService.GetCustomDns:input_type -> wgtwo.data.v0.GetCustomDnsRequest
	8,  // 15: wgtwo.data.v0.DataService.SetCustomDns:output_type -> wgtwo.data.v0.CustomDnsResponse
	8,  // 16: wgtwo.data.v0.DataService.ClearCustomDns:output_type -> wgtwo.data.v0.CustomDnsResponse
	9,  // 17: wgtwo.data.v0.DataService.GetCustomDns:output_type -> wgtwo.data.v0.GetCustomDnsResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_wgtwo_data_v0_data_proto_init() }
func file_wgtwo_data_v0_data_proto_init() {
	if File_wgtwo_data_v0_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_data_v0_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetCustomDnsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomDnsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearCustomDnsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpV4Pair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpAddressV4); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpV6Pair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpAddressV6); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomDnsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wgtwo_data_v0_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomDnsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_data_v0_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_data_v0_data_proto_goTypes,
		DependencyIndexes: file_wgtwo_data_v0_data_proto_depIdxs,
		EnumInfos:         file_wgtwo_data_v0_data_proto_enumTypes,
		MessageInfos:      file_wgtwo_data_v0_data_proto_msgTypes,
	}.Build()
	File_wgtwo_data_v0_data_proto = out.File
	file_wgtwo_data_v0_data_proto_rawDesc = nil
	file_wgtwo_data_v0_data_proto_goTypes = nil
	file_wgtwo_data_v0_data_proto_depIdxs = nil
}
