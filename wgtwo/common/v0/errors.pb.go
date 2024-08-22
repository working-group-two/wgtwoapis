// Copyright 2021 [Working Group Two]/[Cisco Systems]
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
// source: wgtwo/common/v0/errors.proto

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

// Shared status codes
type StatusCode int32

const (
	// Unspecified/default status code; do not use.
	StatusCode_STATUS_CODE_UNSPECIFIED StatusCode = 0
	// Success.
	StatusCode_STATUS_CODE_OK StatusCode = 1
	// The request is not acceptable.
	StatusCode_STATUS_CODE_NOT_ACCEPTABLE StatusCode = 2
	// Access denied. The request is not authorized; the access token may not have the necessary scopes or targeting the
	// wrong resource.
	StatusCode_STATUS_CODE_ACCESS_DENIED StatusCode = 3
	// Internal error. The request could not be processed due to an error on the server side.
	StatusCode_STATUS_CODE_INTERNAL_ERROR StatusCode = 4
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "STATUS_CODE_UNSPECIFIED",
		1: "STATUS_CODE_OK",
		2: "STATUS_CODE_NOT_ACCEPTABLE",
		3: "STATUS_CODE_ACCESS_DENIED",
		4: "STATUS_CODE_INTERNAL_ERROR",
	}
	StatusCode_value = map[string]int32{
		"STATUS_CODE_UNSPECIFIED":    0,
		"STATUS_CODE_OK":             1,
		"STATUS_CODE_NOT_ACCEPTABLE": 2,
		"STATUS_CODE_ACCESS_DENIED":  3,
		"STATUS_CODE_INTERNAL_ERROR": 4,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_common_v0_errors_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_wgtwo_common_v0_errors_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_common_v0_errors_proto_rawDescGZIP(), []int{0}
}

var File_wgtwo_common_v0_errors_proto protoreflect.FileDescriptor

var file_wgtwo_common_v0_errors_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2a,
	0x9c, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12,
	0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12,
	0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1e,
	0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x42, 0x5e,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_common_v0_errors_proto_rawDescOnce sync.Once
	file_wgtwo_common_v0_errors_proto_rawDescData = file_wgtwo_common_v0_errors_proto_rawDesc
)

func file_wgtwo_common_v0_errors_proto_rawDescGZIP() []byte {
	file_wgtwo_common_v0_errors_proto_rawDescOnce.Do(func() {
		file_wgtwo_common_v0_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_common_v0_errors_proto_rawDescData)
	})
	return file_wgtwo_common_v0_errors_proto_rawDescData
}

var file_wgtwo_common_v0_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wgtwo_common_v0_errors_proto_goTypes = []any{
	(StatusCode)(0), // 0: wgtwo.common.v0.StatusCode
}
var file_wgtwo_common_v0_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wgtwo_common_v0_errors_proto_init() }
func file_wgtwo_common_v0_errors_proto_init() {
	if File_wgtwo_common_v0_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_common_v0_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wgtwo_common_v0_errors_proto_goTypes,
		DependencyIndexes: file_wgtwo_common_v0_errors_proto_depIdxs,
		EnumInfos:         file_wgtwo_common_v0_errors_proto_enumTypes,
	}.Build()
	File_wgtwo_common_v0_errors_proto = out.File
	file_wgtwo_common_v0_errors_proto_rawDesc = nil
	file_wgtwo_common_v0_errors_proto_goTypes = nil
	file_wgtwo_common_v0_errors_proto_depIdxs = nil
}
