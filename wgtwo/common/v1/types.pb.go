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
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: wgtwo/common/v1/types.proto

package v1

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

// Country information
//
// All programmatic use should depend on the alpha-2 code, and NOT the human readable name
type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ISO 3166-1 alpha-2 code. Examples: US, NO, SE
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Human readable name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Country) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// This is identification of the subscription
//
// Our OAuth 2.0 server is configured to use Pairwise Pseudonymous Identifiers for its sub field, where all client of
// a product will share the same identifier for a single subscription.
//
// This message contains this subject, given by the product your OAuth 2.0 client belongs to.
type SubscriptionIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SubscriptionIdentifier) Reset() {
	*x = SubscriptionIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionIdentifier) ProtoMessage() {}

func (x *SubscriptionIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionIdentifier.ProtoReflect.Descriptor instead.
func (*SubscriptionIdentifier) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionIdentifier) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
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
		mi := &file_wgtwo_common_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Imsi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Imsi) ProtoMessage() {}

func (x *Imsi) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_types_proto_msgTypes[2]
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
	return file_wgtwo_common_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Imsi) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// International Mobile station Equipment Identity and Software Version number
// An IMEI is composed of four parts:
//   1) an 8 digit Type Allocation Code (TAC);
//   2) a 6 digit Serial Number (SNR);
//   3) an optional Check Digit (CD); and
//   4) an optional 2 digit Software Version Number (handled separately).
// For more information see ETSI 123.003 Chapter 6 and Appendix B.
type ImeiSv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imei            string `protobuf:"bytes,1,opt,name=imei,proto3" json:"imei,omitempty"`
	SoftwareVersion string `protobuf:"bytes,2,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty"`
}

func (x *ImeiSv) Reset() {
	*x = ImeiSv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImeiSv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImeiSv) ProtoMessage() {}

func (x *ImeiSv) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImeiSv.ProtoReflect.Descriptor instead.
func (*ImeiSv) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *ImeiSv) GetImei() string {
	if x != nil {
		return x.Imei
	}
	return ""
}

func (x *ImeiSv) GetSoftwareVersion() string {
	if x != nil {
		return x.SoftwareVersion
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
		mi := &file_wgtwo_common_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkIdentity) ProtoMessage() {}

func (x *NetworkIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_types_proto_msgTypes[4]
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
	return file_wgtwo_common_v1_types_proto_rawDescGZIP(), []int{4}
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

var File_wgtwo_common_v1_types_proto protoreflect.FileDescriptor

var file_wgtwo_common_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x31,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2e, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x1c, 0x0a, 0x04, 0x49, 0x6d, 0x73, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x47, 0x0a, 0x06, 0x49, 0x6d, 0x65, 0x69, 0x53, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x63, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x6e, 0x63, 0x42,
	0x5d, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_common_v1_types_proto_rawDescOnce sync.Once
	file_wgtwo_common_v1_types_proto_rawDescData = file_wgtwo_common_v1_types_proto_rawDesc
)

func file_wgtwo_common_v1_types_proto_rawDescGZIP() []byte {
	file_wgtwo_common_v1_types_proto_rawDescOnce.Do(func() {
		file_wgtwo_common_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_common_v1_types_proto_rawDescData)
	})
	return file_wgtwo_common_v1_types_proto_rawDescData
}

var file_wgtwo_common_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wgtwo_common_v1_types_proto_goTypes = []interface{}{
	(*Country)(nil),                // 0: wgtwo.common.v1.Country
	(*SubscriptionIdentifier)(nil), // 1: wgtwo.common.v1.SubscriptionIdentifier
	(*Imsi)(nil),                   // 2: wgtwo.common.v1.Imsi
	(*ImeiSv)(nil),                 // 3: wgtwo.common.v1.ImeiSv
	(*NetworkIdentity)(nil),        // 4: wgtwo.common.v1.NetworkIdentity
}
var file_wgtwo_common_v1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wgtwo_common_v1_types_proto_init() }
func file_wgtwo_common_v1_types_proto_init() {
	if File_wgtwo_common_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_common_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_wgtwo_common_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionIdentifier); i {
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
		file_wgtwo_common_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_wgtwo_common_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImeiSv); i {
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
		file_wgtwo_common_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_common_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wgtwo_common_v1_types_proto_goTypes,
		DependencyIndexes: file_wgtwo_common_v1_types_proto_depIdxs,
		MessageInfos:      file_wgtwo_common_v1_types_proto_msgTypes,
	}.Build()
	File_wgtwo_common_v1_types_proto = out.File
	file_wgtwo_common_v1_types_proto_rawDesc = nil
	file_wgtwo_common_v1_types_proto_goTypes = nil
	file_wgtwo_common_v1_types_proto_depIdxs = nil
}
