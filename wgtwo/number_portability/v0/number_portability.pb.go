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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: wgtwo/number_portability/v0/number_portability.proto

package v0

import (
	_ "github.com/working-group-two/wgtwoapis/wgtwo/annotations"
	v1 "github.com/working-group-two/wgtwoapis/wgtwo/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message to create porting records.
type CreatePortingRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Porting records to be created.
	Records []*PortingRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *CreatePortingRecordsRequest) Reset() {
	*x = CreatePortingRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortingRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortingRecordsRequest) ProtoMessage() {}

func (x *CreatePortingRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortingRecordsRequest.ProtoReflect.Descriptor instead.
func (*CreatePortingRecordsRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_number_portability_v0_number_portability_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePortingRecordsRequest) GetRecords() []*PortingRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// Response message for create porting records.
type CreatePortingRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePortingRecordsResponse) Reset() {
	*x = CreatePortingRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortingRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortingRecordsResponse) ProtoMessage() {}

func (x *CreatePortingRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortingRecordsResponse.ProtoReflect.Descriptor instead.
func (*CreatePortingRecordsResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_number_portability_v0_number_portability_proto_rawDescGZIP(), []int{1}
}

// Request message to list porting records.
type ListPortingRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional subscriber number prefix to filter porting records.
	OperatorCode *string `protobuf:"bytes,1,opt,name=operator_code,json=operatorCode,proto3,oneof" json:"operator_code,omitempty"`
	// Optional operator code to filter porting records.
	SubscriberNumberPrefix *string `protobuf:"bytes,2,opt,name=subscriber_number_prefix,json=subscriberNumberPrefix,proto3,oneof" json:"subscriber_number_prefix,omitempty"`
	// Optional porting date to filter porting records.
	ValidFrom *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=valid_from,json=validFrom,proto3,oneof" json:"valid_from,omitempty"`
}

func (x *ListPortingRecordsRequest) Reset() {
	*x = ListPortingRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortingRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortingRecordsRequest) ProtoMessage() {}

func (x *ListPortingRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortingRecordsRequest.ProtoReflect.Descriptor instead.
func (*ListPortingRecordsRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_number_portability_v0_number_portability_proto_rawDescGZIP(), []int{2}
}

func (x *ListPortingRecordsRequest) GetOperatorCode() string {
	if x != nil && x.OperatorCode != nil {
		return *x.OperatorCode
	}
	return ""
}

func (x *ListPortingRecordsRequest) GetSubscriberNumberPrefix() string {
	if x != nil && x.SubscriberNumberPrefix != nil {
		return *x.SubscriberNumberPrefix
	}
	return ""
}

func (x *ListPortingRecordsRequest) GetValidFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

// Response message for list porting records.
type ListPortingRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Porting records.
	Records []*PortingRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListPortingRecordsResponse) Reset() {
	*x = ListPortingRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortingRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortingRecordsResponse) ProtoMessage() {}

func (x *ListPortingRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortingRecordsResponse.ProtoReflect.Descriptor instead.
func (*ListPortingRecordsResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_number_portability_v0_number_portability_proto_rawDescGZIP(), []int{3}
}

func (x *ListPortingRecordsResponse) GetRecords() []*PortingRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// Message representing a porting record.
type PortingRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subscriber number with a country code.
	SubscriberNumber *v1.E164 `protobuf:"bytes,1,opt,name=subscriber_number,json=subscriberNumber,proto3" json:"subscriber_number,omitempty"`
	// A new operator's code (in Sweden this is SPID allocated by SNPAC).
	OperatorCode string `protobuf:"bytes,2,opt,name=operator_code,json=operatorCode,proto3" json:"operator_code,omitempty"`
	// Routing code of the new operator, used at least in Sweden.
	RoutingCode string `protobuf:"bytes,3,opt,name=routing_code,json=routingCode,proto3" json:"routing_code,omitempty"`
	// Porting date and time.
	ValidFrom *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	// May be used to store tenant-specific data.
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PortingRecord) Reset() {
	*x = PortingRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortingRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortingRecord) ProtoMessage() {}

func (x *PortingRecord) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortingRecord.ProtoReflect.Descriptor instead.
func (*PortingRecord) Descriptor() ([]byte, []int) {
	return file_wgtwo_number_portability_v0_number_portability_proto_rawDescGZIP(), []int{4}
}

func (x *PortingRecord) GetSubscriberNumber() *v1.E164 {
	if x != nil {
		return x.SubscriberNumber
	}
	return nil
}

func (x *PortingRecord) GetOperatorCode() string {
	if x != nil {
		return x.OperatorCode
	}
	return ""
}

func (x *PortingRecord) GetRoutingCode() string {
	if x != nil {
		return x.RoutingCode
	}
	return ""
}

func (x *PortingRecord) GetValidFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidFrom
	}
	return nil
}

func (x *PortingRecord) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_wgtwo_number_portability_v0_number_portability_proto protoreflect.FileDescriptor

var file_wgtwo_number_portability_v0_number_portability_proto_rawDesc = []byte{
	0x0a, 0x34, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x30, 0x2f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x30, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x1e, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x82, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x18, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x16, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46,
	0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x22, 0x62, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x30, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xe9, 0x02, 0x0a, 0x0d, 0x50, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x42, 0x0a, 0x11, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x31, 0x36, 0x34, 0x52, 0x10, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x54, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x30, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xe8, 0x03, 0x0a, 0x18, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xe9, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x38, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x30, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x5c, 0xea, 0xb5, 0x18, 0x27, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x3a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0xdf, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x36, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0xea, 0xb5, 0x18, 0x26, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x3a, 0x72, 0x65, 0x61,
	0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x30, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42,
	0x83, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x30, 0x2e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x16, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_number_portability_v0_number_portability_proto_rawDescOnce sync.Once
	file_wgtwo_number_portability_v0_number_portability_proto_rawDescData = file_wgtwo_number_portability_v0_number_portability_proto_rawDesc
)

func file_wgtwo_number_portability_v0_number_portability_proto_rawDescGZIP() []byte {
	file_wgtwo_number_portability_v0_number_portability_proto_rawDescOnce.Do(func() {
		file_wgtwo_number_portability_v0_number_portability_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_number_portability_v0_number_portability_proto_rawDescData)
	})
	return file_wgtwo_number_portability_v0_number_portability_proto_rawDescData
}

var file_wgtwo_number_portability_v0_number_portability_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wgtwo_number_portability_v0_number_portability_proto_goTypes = []any{
	(*CreatePortingRecordsRequest)(nil),  // 0: wgtwo.number_portability.v0.CreatePortingRecordsRequest
	(*CreatePortingRecordsResponse)(nil), // 1: wgtwo.number_portability.v0.CreatePortingRecordsResponse
	(*ListPortingRecordsRequest)(nil),    // 2: wgtwo.number_portability.v0.ListPortingRecordsRequest
	(*ListPortingRecordsResponse)(nil),   // 3: wgtwo.number_portability.v0.ListPortingRecordsResponse
	(*PortingRecord)(nil),                // 4: wgtwo.number_portability.v0.PortingRecord
	nil,                                  // 5: wgtwo.number_portability.v0.PortingRecord.MetadataEntry
	(*timestamppb.Timestamp)(nil),        // 6: google.protobuf.Timestamp
	(*v1.E164)(nil),                      // 7: wgtwo.common.v1.E164
}
var file_wgtwo_number_portability_v0_number_portability_proto_depIdxs = []int32{
	4, // 0: wgtwo.number_portability.v0.CreatePortingRecordsRequest.records:type_name -> wgtwo.number_portability.v0.PortingRecord
	6, // 1: wgtwo.number_portability.v0.ListPortingRecordsRequest.valid_from:type_name -> google.protobuf.Timestamp
	4, // 2: wgtwo.number_portability.v0.ListPortingRecordsResponse.records:type_name -> wgtwo.number_portability.v0.PortingRecord
	7, // 3: wgtwo.number_portability.v0.PortingRecord.subscriber_number:type_name -> wgtwo.common.v1.E164
	6, // 4: wgtwo.number_portability.v0.PortingRecord.valid_from:type_name -> google.protobuf.Timestamp
	5, // 5: wgtwo.number_portability.v0.PortingRecord.metadata:type_name -> wgtwo.number_portability.v0.PortingRecord.MetadataEntry
	0, // 6: wgtwo.number_portability.v0.NumberPortabilityService.CreatePortingRecords:input_type -> wgtwo.number_portability.v0.CreatePortingRecordsRequest
	2, // 7: wgtwo.number_portability.v0.NumberPortabilityService.ListPortingRecords:input_type -> wgtwo.number_portability.v0.ListPortingRecordsRequest
	1, // 8: wgtwo.number_portability.v0.NumberPortabilityService.CreatePortingRecords:output_type -> wgtwo.number_portability.v0.CreatePortingRecordsResponse
	3, // 9: wgtwo.number_portability.v0.NumberPortabilityService.ListPortingRecords:output_type -> wgtwo.number_portability.v0.ListPortingRecordsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_wgtwo_number_portability_v0_number_portability_proto_init() }
func file_wgtwo_number_portability_v0_number_portability_proto_init() {
	if File_wgtwo_number_portability_v0_number_portability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePortingRecordsRequest); i {
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
		file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePortingRecordsResponse); i {
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
		file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListPortingRecordsRequest); i {
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
		file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListPortingRecordsResponse); i {
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
		file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PortingRecord); i {
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
	file_wgtwo_number_portability_v0_number_portability_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_number_portability_v0_number_portability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_number_portability_v0_number_portability_proto_goTypes,
		DependencyIndexes: file_wgtwo_number_portability_v0_number_portability_proto_depIdxs,
		MessageInfos:      file_wgtwo_number_portability_v0_number_portability_proto_msgTypes,
	}.Build()
	File_wgtwo_number_portability_v0_number_portability_proto = out.File
	file_wgtwo_number_portability_v0_number_portability_proto_rawDesc = nil
	file_wgtwo_number_portability_v0_number_portability_proto_goTypes = nil
	file_wgtwo_number_portability_v0_number_portability_proto_depIdxs = nil
}
