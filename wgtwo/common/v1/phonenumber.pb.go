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
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: wgtwo/common/v1/phonenumber.proto

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

// *
// An international phone number formatted as E.164 with leading plus sign
//
// This contains three components:
// - The plus prefix
// - Country code, 1 to 3 digits
// - Subscriber number
//
// Example "+4787654321" as '+' '47' '87654321'.
//
// The number may contain sequences that do not strictly conform to the E.164
// number standard (e.g. too long), but shall always follow the three components
// as described above.
type E164 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	E164 string `protobuf:"bytes,1,opt,name=e164,proto3" json:"e164,omitempty"`
}

func (x *E164) Reset() {
	*x = E164{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v1_phonenumber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *E164) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*E164) ProtoMessage() {}

func (x *E164) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_phonenumber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use E164.ProtoReflect.Descriptor instead.
func (*E164) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v1_phonenumber_proto_rawDescGZIP(), []int{0}
}

func (x *E164) GetE164() string {
	if x != nil {
		return x.E164
	}
	return ""
}

// *
// A national number can be anything that is usually typed into a number field
// for phone numbers, SMS etc. It shall only contains digits (0-9), and the meaning
// of the number sequence is dependent on the country of the relevant operator.
type National struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *National) Reset() {
	*x = National{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v1_phonenumber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *National) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*National) ProtoMessage() {}

func (x *National) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_phonenumber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use National.ProtoReflect.Descriptor instead.
func (*National) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v1_phonenumber_proto_rawDescGZIP(), []int{1}
}

func (x *National) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

// *
// An alphanumeric address. This is usually just be the name of a product or service.
// Alphanumeric addresses are usually not routable, and can thus only be used in
// origin addresses.
type Alphanumeric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Alphanumeric) Reset() {
	*x = Alphanumeric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_common_v1_phonenumber_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alphanumeric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alphanumeric) ProtoMessage() {}

func (x *Alphanumeric) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_common_v1_phonenumber_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alphanumeric.ProtoReflect.Descriptor instead.
func (*Alphanumeric) Descriptor() ([]byte, []int) {
	return file_wgtwo_common_v1_phonenumber_proto_rawDescGZIP(), []int{2}
}

func (x *Alphanumeric) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_wgtwo_common_v1_phonenumber_proto protoreflect.FileDescriptor

var file_wgtwo_common_v1_phonenumber_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x1a, 0x0a, 0x04, 0x45, 0x31, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04,
	0x65, 0x31, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x31, 0x36, 0x34,
	0x22, 0x22, 0x0a, 0x08, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x63, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x42, 0x10, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d,
	0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67,
	0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_common_v1_phonenumber_proto_rawDescOnce sync.Once
	file_wgtwo_common_v1_phonenumber_proto_rawDescData = file_wgtwo_common_v1_phonenumber_proto_rawDesc
)

func file_wgtwo_common_v1_phonenumber_proto_rawDescGZIP() []byte {
	file_wgtwo_common_v1_phonenumber_proto_rawDescOnce.Do(func() {
		file_wgtwo_common_v1_phonenumber_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_common_v1_phonenumber_proto_rawDescData)
	})
	return file_wgtwo_common_v1_phonenumber_proto_rawDescData
}

var file_wgtwo_common_v1_phonenumber_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wgtwo_common_v1_phonenumber_proto_goTypes = []interface{}{
	(*E164)(nil),         // 0: wgtwo.common.v1.E164
	(*National)(nil),     // 1: wgtwo.common.v1.National
	(*Alphanumeric)(nil), // 2: wgtwo.common.v1.Alphanumeric
}
var file_wgtwo_common_v1_phonenumber_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wgtwo_common_v1_phonenumber_proto_init() }
func file_wgtwo_common_v1_phonenumber_proto_init() {
	if File_wgtwo_common_v1_phonenumber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_common_v1_phonenumber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*E164); i {
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
		file_wgtwo_common_v1_phonenumber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*National); i {
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
		file_wgtwo_common_v1_phonenumber_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alphanumeric); i {
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
			RawDescriptor: file_wgtwo_common_v1_phonenumber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wgtwo_common_v1_phonenumber_proto_goTypes,
		DependencyIndexes: file_wgtwo_common_v1_phonenumber_proto_depIdxs,
		MessageInfos:      file_wgtwo_common_v1_phonenumber_proto_msgTypes,
	}.Build()
	File_wgtwo_common_v1_phonenumber_proto = out.File
	file_wgtwo_common_v1_phonenumber_proto_rawDesc = nil
	file_wgtwo_common_v1_phonenumber_proto_goTypes = nil
	file_wgtwo_common_v1_phonenumber_proto_depIdxs = nil
}
