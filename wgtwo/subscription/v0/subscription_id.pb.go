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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: wgtwo/subscription/v0/subscription_id.proto

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

type GetSubscriptionIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber *v0.PhoneNumber `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *GetSubscriptionIdRequest) Reset() {
	*x = GetSubscriptionIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionIdRequest) ProtoMessage() {}

func (x *GetSubscriptionIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionIdRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriptionIdRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_subscription_v0_subscription_id_proto_rawDescGZIP(), []int{0}
}

func (x *GetSubscriptionIdRequest) GetPhoneNumber() *v0.PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

type GetSubscriptionIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId *v0.SubscriptionIdentifier `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	StatusCode     v0.StatusCode              `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3,enum=wgtwo.common.v0.StatusCode" json:"status_code,omitempty"`
	ErrorMessage   string                     `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *GetSubscriptionIdResponse) Reset() {
	*x = GetSubscriptionIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriptionIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriptionIdResponse) ProtoMessage() {}

func (x *GetSubscriptionIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriptionIdResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriptionIdResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_subscription_v0_subscription_id_proto_rawDescGZIP(), []int{1}
}

func (x *GetSubscriptionIdResponse) GetSubscriptionId() *v0.SubscriptionIdentifier {
	if x != nil {
		return x.SubscriptionId
	}
	return nil
}

func (x *GetSubscriptionIdResponse) GetStatusCode() v0.StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return v0.StatusCode(0)
}

func (x *GetSubscriptionIdResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetMsisdnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId *v0.SubscriptionIdentifier `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *GetMsisdnRequest) Reset() {
	*x = GetMsisdnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMsisdnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMsisdnRequest) ProtoMessage() {}

func (x *GetMsisdnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMsisdnRequest.ProtoReflect.Descriptor instead.
func (*GetMsisdnRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_subscription_v0_subscription_id_proto_rawDescGZIP(), []int{2}
}

func (x *GetMsisdnRequest) GetSubscriptionId() *v0.SubscriptionIdentifier {
	if x != nil {
		return x.SubscriptionId
	}
	return nil
}

type GetMsisdnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber  *v0.PhoneNumber `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	StatusCode   v0.StatusCode   `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3,enum=wgtwo.common.v0.StatusCode" json:"status_code,omitempty"`
	ErrorMessage string          `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *GetMsisdnResponse) Reset() {
	*x = GetMsisdnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMsisdnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMsisdnResponse) ProtoMessage() {}

func (x *GetMsisdnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMsisdnResponse.ProtoReflect.Descriptor instead.
func (*GetMsisdnResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_subscription_v0_subscription_id_proto_rawDescGZIP(), []int{3}
}

func (x *GetMsisdnResponse) GetPhoneNumber() *v0.PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *GetMsisdnResponse) GetStatusCode() v0.StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return v0.StatusCode(0)
}

func (x *GetMsisdnResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_wgtwo_subscription_v0_subscription_id_proto protoreflect.FileDescriptor

var file_wgtwo_subscription_v0_subscription_id_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x30, 0x1a, 0x23, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0xd0, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x67,
	0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x73,
	0x69, 0x73, 0x64, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb7, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa4, 0x02, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x90, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2f, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xea, 0xb5, 0x18, 0x14,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x64, 0x3a,
	0x72, 0x65, 0x61, 0x64, 0x12, 0x78, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x69, 0x73, 0x64,
	0x6e, 0x12, 0x27, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x69,
	0x73, 0x64, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xea, 0xb5, 0x18, 0x14, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x64, 0x3a, 0x72, 0x65, 0x61, 0x64, 0x42, 0x73,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x30, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74,
	0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_subscription_v0_subscription_id_proto_rawDescOnce sync.Once
	file_wgtwo_subscription_v0_subscription_id_proto_rawDescData = file_wgtwo_subscription_v0_subscription_id_proto_rawDesc
)

func file_wgtwo_subscription_v0_subscription_id_proto_rawDescGZIP() []byte {
	file_wgtwo_subscription_v0_subscription_id_proto_rawDescOnce.Do(func() {
		file_wgtwo_subscription_v0_subscription_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_subscription_v0_subscription_id_proto_rawDescData)
	})
	return file_wgtwo_subscription_v0_subscription_id_proto_rawDescData
}

var file_wgtwo_subscription_v0_subscription_id_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wgtwo_subscription_v0_subscription_id_proto_goTypes = []interface{}{
	(*GetSubscriptionIdRequest)(nil),  // 0: wgtwo.subscription.v0.GetSubscriptionIdRequest
	(*GetSubscriptionIdResponse)(nil), // 1: wgtwo.subscription.v0.GetSubscriptionIdResponse
	(*GetMsisdnRequest)(nil),          // 2: wgtwo.subscription.v0.GetMsisdnRequest
	(*GetMsisdnResponse)(nil),         // 3: wgtwo.subscription.v0.GetMsisdnResponse
	(*v0.PhoneNumber)(nil),            // 4: wgtwo.common.v0.PhoneNumber
	(*v0.SubscriptionIdentifier)(nil), // 5: wgtwo.common.v0.SubscriptionIdentifier
	(v0.StatusCode)(0),                // 6: wgtwo.common.v0.StatusCode
}
var file_wgtwo_subscription_v0_subscription_id_proto_depIdxs = []int32{
	4, // 0: wgtwo.subscription.v0.GetSubscriptionIdRequest.phone_number:type_name -> wgtwo.common.v0.PhoneNumber
	5, // 1: wgtwo.subscription.v0.GetSubscriptionIdResponse.subscription_id:type_name -> wgtwo.common.v0.SubscriptionIdentifier
	6, // 2: wgtwo.subscription.v0.GetSubscriptionIdResponse.status_code:type_name -> wgtwo.common.v0.StatusCode
	5, // 3: wgtwo.subscription.v0.GetMsisdnRequest.subscription_id:type_name -> wgtwo.common.v0.SubscriptionIdentifier
	4, // 4: wgtwo.subscription.v0.GetMsisdnResponse.phone_number:type_name -> wgtwo.common.v0.PhoneNumber
	6, // 5: wgtwo.subscription.v0.GetMsisdnResponse.status_code:type_name -> wgtwo.common.v0.StatusCode
	0, // 6: wgtwo.subscription.v0.SubscriptionIdService.GetSubscriptionId:input_type -> wgtwo.subscription.v0.GetSubscriptionIdRequest
	2, // 7: wgtwo.subscription.v0.SubscriptionIdService.GetMsisdn:input_type -> wgtwo.subscription.v0.GetMsisdnRequest
	1, // 8: wgtwo.subscription.v0.SubscriptionIdService.GetSubscriptionId:output_type -> wgtwo.subscription.v0.GetSubscriptionIdResponse
	3, // 9: wgtwo.subscription.v0.SubscriptionIdService.GetMsisdn:output_type -> wgtwo.subscription.v0.GetMsisdnResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_wgtwo_subscription_v0_subscription_id_proto_init() }
func file_wgtwo_subscription_v0_subscription_id_proto_init() {
	if File_wgtwo_subscription_v0_subscription_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionIdRequest); i {
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
		file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriptionIdResponse); i {
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
		file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMsisdnRequest); i {
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
		file_wgtwo_subscription_v0_subscription_id_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMsisdnResponse); i {
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
			RawDescriptor: file_wgtwo_subscription_v0_subscription_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_subscription_v0_subscription_id_proto_goTypes,
		DependencyIndexes: file_wgtwo_subscription_v0_subscription_id_proto_depIdxs,
		MessageInfos:      file_wgtwo_subscription_v0_subscription_id_proto_msgTypes,
	}.Build()
	File_wgtwo_subscription_v0_subscription_id_proto = out.File
	file_wgtwo_subscription_v0_subscription_id_proto_rawDesc = nil
	file_wgtwo_subscription_v0_subscription_id_proto_goTypes = nil
	file_wgtwo_subscription_v0_subscription_id_proto_depIdxs = nil
}
