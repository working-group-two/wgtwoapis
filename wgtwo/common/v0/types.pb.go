// Copyright 2021 Working Group Two AS
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
// source: wgtwo/common/v0/types.proto

package v0

import (
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

type NetworkGeneration int32

const (
	NetworkGeneration_NETWORK_GENERATION_UNSPECIFIED NetworkGeneration = 0
	NetworkGeneration_NETWORK_GENERATION_2G          NetworkGeneration = 1
	NetworkGeneration_NETWORK_GENERATION_3G          NetworkGeneration = 2
	NetworkGeneration_NETWORK_GENERATION_4G          NetworkGeneration = 3
	NetworkGeneration_NETWORK_GENERATION_5G          NetworkGeneration = 4
)

// Enum value maps for NetworkGeneration.
var (
	NetworkGeneration_name = map[int32]string{
		0: "NETWORK_GENERATION_UNSPECIFIED",
		1: "NETWORK_GENERATION_2G",
		2: "NETWORK_GENERATION_3G",
		3: "NETWORK_GENERATION_4G",
		4: "NETWORK_GENERATION_5G",
	}
	NetworkGeneration_value = map[string]int32{
		"NETWORK_GENERATION_UNSPECIFIED": 0,
		"NETWORK_GENERATION_2G":          1,
		"NETWORK_GENERATION_3G":          2,
		"NETWORK_GENERATION_4G":          3,
		"NETWORK_GENERATION_5G":          4,
	}
)

func (x NetworkGeneration) Enum() *NetworkGeneration {
	p := new(NetworkGeneration)
	*p = x
	return p
}

func (x NetworkGeneration) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkGeneration) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_common_v0_types_proto_enumTypes[0].Descriptor()
}

func (NetworkGeneration) Type() protoreflect.EnumType {
	return &file_wgtwo_common_v0_types_proto_enumTypes[0]
}

func (x NetworkGeneration) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkGeneration.Descriptor instead.
func (NetworkGeneration) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{0}
}

// International Mobile Subscription Identity
// An IMSI is composed of three parts:
//  1) Mobile Country Code (MCC) consisting of three digits. The MCC is 3 digits long and identifies
//    uniquely the country of domicile of the mobile subscription;
//  2) Mobile Network Code (MNC), 2 or 3 digits for 3GPP network applications. The MNC identifies
//    the home PLMN of the mobile subscription. The length of the MNC depends on the value of the
//    MCC. A mixture of two and three digit MNC codes within a single MCC area is not recommended.
//  3) Mobile Subscriber Identification Number (MSIN) identifying the mobile subscription within a
//    PLMN. Normally there are 10 digits, but can be fewer in the case of a 3-digit MNC or if
//    national regulations indicate that the total length of the IMSI should be less than 15 digits.
type Imsi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Imsi) Reset() {
	*x = Imsi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v0_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Imsi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Imsi) ProtoMessage() {}

func (x *Imsi) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v0_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Imsi.ProtoReflect.Descriptor instead.
func (*Imsi) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{0}
}

func (x *Imsi) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Tadig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tadig) Reset() {
	*x = Tadig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v0_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tadig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tadig) ProtoMessage() {}

func (x *Tadig) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v0_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tadig.ProtoReflect.Descriptor instead.
func (*Tadig) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{1}
}

func (x *Tadig) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Iccid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Iccid) Reset() {
	*x = Iccid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v0_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Iccid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Iccid) ProtoMessage() {}

func (x *Iccid) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v0_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Iccid.ProtoReflect.Descriptor instead.
func (*Iccid) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{2}
}

func (x *Iccid) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NetworkIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mcc string `protobuf:"bytes,1,opt,name=mcc,proto3" json:"mcc,omitempty"`
	Mnc string `protobuf:"bytes,2,opt,name=mnc,proto3" json:"mnc,omitempty"`
}

func (x *NetworkIdentity) Reset() {
	*x = NetworkIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v0_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkIdentity) ProtoMessage() {}

func (x *NetworkIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v0_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkIdentity.ProtoReflect.Descriptor instead.
func (*NetworkIdentity) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkIdentity) GetMcc() string {
	if x != nil {
		return x.Mcc
	}
	return ""
}

func (x *NetworkIdentity) GetMnc() string {
	if x != nil {
		return x.Mnc
	}
	return ""
}

type GlobalTitle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GlobalTitle) Reset() {
	*x = GlobalTitle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v0_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalTitle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalTitle) ProtoMessage() {}

func (x *GlobalTitle) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v0_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalTitle.ProtoReflect.Descriptor instead.
func (*GlobalTitle) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{4}
}

func (x *GlobalTitle) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Fully Qualified Domain Name
type Fqdn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Fqdn) Reset() {
	*x = Fqdn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v0_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fqdn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fqdn) ProtoMessage() {}

func (x *Fqdn) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v0_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fqdn.ProtoReflect.Descriptor instead.
func (*Fqdn) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_types_proto_rawDescGZIP(), []int{5}
}

func (x *Fqdn) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_wgtwo_common_v0_types_proto protoreflect.FileDescriptor

var file_wgtwo_common_v0_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x22, 0x1c,
	0x0a, 0x04, 0x49, 0x6d, 0x73, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x05,
	0x54, 0x61, 0x64, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x49,
	0x63, 0x63, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x63, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x6e,
	0x63, 0x22, 0x23, 0x0a, 0x0b, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x46, 0x71, 0x64, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2a, 0xa3, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x4e, 0x45,
	0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x32, 0x47, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x45, 0x54,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x33, 0x47, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x34, 0x47, 0x10, 0x03, 0x12,
	0x19, 0x0a, 0x15, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x35, 0x47, 0x10, 0x04, 0x42, 0x5d, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_wgtwo_common_v0_types_proto_rawDescOnce sync.Once
	file_wgtwo_common_v0_types_proto_rawDescData = file_wgtwo_common_v0_types_proto_rawDesc
)

func file_wgtwo_common_v0_types_proto_rawDescGZIP() []byte {
	file_wgtwo_common_v0_types_proto_rawDescOnce.Do(func() {
		file_wgtwo_common_v0_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_common_v0_types_proto_rawDescData)
	})
	return file_wgtwo_common_v0_types_proto_rawDescData
}

var file_wgtwo_common_v0_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wgtwo_common_v0_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wgtwo_common_v0_types_proto_goTypes = []interface{}{
	(NetworkGeneration)(0),  // 0: wgtwo.common.v0.NetworkGeneration
	(*Imsi)(nil),            // 1: wgtwo.common.v0.Imsi
	(*Tadig)(nil),           // 2: wgtwo.common.v0.Tadig
	(*Iccid)(nil),           // 3: wgtwo.common.v0.Iccid
	(*NetworkIdentity)(nil), // 4: wgtwo.common.v0.NetworkIdentity
	(*GlobalTitle)(nil),     // 5: wgtwo.common.v0.GlobalTitle
	(*Fqdn)(nil),            // 6: wgtwo.common.v0.Fqdn
}
var file_wgtwo_common_v0_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wgtwo_common_v0_types_proto_init() }
func file_wgtwo_common_v0_types_proto_init() {
	if File_wgtwo_common_v0_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_common_v0_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Imsi); i {
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
		file_wgtwo_common_v0_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tadig); i {
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
		file_wgtwo_common_v0_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Iccid); i {
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
		file_wgtwo_common_v0_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkIdentity); i {
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
		file_wgtwo_common_v0_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalTitle); i {
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
		file_wgtwo_common_v0_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fqdn); i {
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
			RawDescriptor: file_wgtwo_common_v0_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wgtwo_common_v0_types_proto_goTypes,
		DependencyIndexes: file_wgtwo_common_v0_types_proto_depIdxs,
		EnumInfos:         file_wgtwo_common_v0_types_proto_enumTypes,
		MessageInfos:      file_wgtwo_common_v0_types_proto_msgTypes,
	}.Build()
	File_wgtwo_common_v0_types_proto = out.File
	file_wgtwo_common_v0_types_proto_rawDesc = nil
	file_wgtwo_common_v0_types_proto_goTypes = nil
	file_wgtwo_common_v0_types_proto_depIdxs = nil
}
