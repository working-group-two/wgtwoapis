// Copyright 2023 Working Group Two AS
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
// source: wgtwo/images/v0/images.proto

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image URL
	// The URL might change and should not be stored by the client.
	// Example: "https://example.com/image.png"
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_images_v0_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_images_v0_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_wgtwo_images_v0_images_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_wgtwo_images_v0_images_proto protoreflect.FileDescriptor

var file_wgtwo_images_v0_images_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76,
	0x30, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x76, 0x30, 0x22,
	0x19, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x5e, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f,
	0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wgtwo_images_v0_images_proto_rawDescOnce sync.Once
	file_wgtwo_images_v0_images_proto_rawDescData = file_wgtwo_images_v0_images_proto_rawDesc
)

func file_wgtwo_images_v0_images_proto_rawDescGZIP() []byte {
	file_wgtwo_images_v0_images_proto_rawDescOnce.Do(func() {
		file_wgtwo_images_v0_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_images_v0_images_proto_rawDescData)
	})
	return file_wgtwo_images_v0_images_proto_rawDescData
}

var file_wgtwo_images_v0_images_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wgtwo_images_v0_images_proto_goTypes = []interface{}{
	(*Image)(nil), // 0: wgtwo.images.v0.Image
}
var file_wgtwo_images_v0_images_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wgtwo_images_v0_images_proto_init() }
func file_wgtwo_images_v0_images_proto_init() {
	if File_wgtwo_images_v0_images_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_images_v0_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
			RawDescriptor: file_wgtwo_images_v0_images_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wgtwo_images_v0_images_proto_goTypes,
		DependencyIndexes: file_wgtwo_images_v0_images_proto_depIdxs,
		MessageInfos:      file_wgtwo_images_v0_images_proto_msgTypes,
	}.Build()
	File_wgtwo_images_v0_images_proto = out.File
	file_wgtwo_images_v0_images_proto_rawDesc = nil
	file_wgtwo_images_v0_images_proto_goTypes = nil
	file_wgtwo_images_v0_images_proto_depIdxs = nil
}
