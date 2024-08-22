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
// source: wgtwo/sim/v0/authentication.proto

package v0

import (
	_ "github.com/working-group-two/wgtwoapis/wgtwo/annotations"
	v1 "github.com/working-group-two/wgtwoapis/wgtwo/common/v1"
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

type GenerateEapAkaAuthenticationVectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imsi *v1.Imsi `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
}

func (x *GenerateEapAkaAuthenticationVectorRequest) Reset() {
	*x = GenerateEapAkaAuthenticationVectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sim_v0_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateEapAkaAuthenticationVectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateEapAkaAuthenticationVectorRequest) ProtoMessage() {}

func (x *GenerateEapAkaAuthenticationVectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sim_v0_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateEapAkaAuthenticationVectorRequest.ProtoReflect.Descriptor instead.
func (*GenerateEapAkaAuthenticationVectorRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_sim_v0_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateEapAkaAuthenticationVectorRequest) GetImsi() *v1.Imsi {
	if x != nil {
		return x.Imsi
	}
	return nil
}

type GenerateEapAkaAuthenticationVectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthenticationVector *EapAkaAuthenticationVector `protobuf:"bytes,1,opt,name=authentication_vector,json=authenticationVector,proto3" json:"authentication_vector,omitempty"`
}

func (x *GenerateEapAkaAuthenticationVectorResponse) Reset() {
	*x = GenerateEapAkaAuthenticationVectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sim_v0_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateEapAkaAuthenticationVectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateEapAkaAuthenticationVectorResponse) ProtoMessage() {}

func (x *GenerateEapAkaAuthenticationVectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sim_v0_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateEapAkaAuthenticationVectorResponse.ProtoReflect.Descriptor instead.
func (*GenerateEapAkaAuthenticationVectorResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_sim_v0_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateEapAkaAuthenticationVectorResponse) GetAuthenticationVector() *EapAkaAuthenticationVector {
	if x != nil {
		return x.AuthenticationVector
	}
	return nil
}

type EapAkaAuthenticationVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rand []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Xres []byte `protobuf:"bytes,2,opt,name=xres,proto3" json:"xres,omitempty"`
	Autn []byte `protobuf:"bytes,3,opt,name=autn,proto3" json:"autn,omitempty"`
	Ck   []byte `protobuf:"bytes,4,opt,name=ck,proto3" json:"ck,omitempty"`
	Ik   []byte `protobuf:"bytes,5,opt,name=ik,proto3" json:"ik,omitempty"`
}

func (x *EapAkaAuthenticationVector) Reset() {
	*x = EapAkaAuthenticationVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sim_v0_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EapAkaAuthenticationVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EapAkaAuthenticationVector) ProtoMessage() {}

func (x *EapAkaAuthenticationVector) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sim_v0_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EapAkaAuthenticationVector.ProtoReflect.Descriptor instead.
func (*EapAkaAuthenticationVector) Descriptor() ([]byte, []int) {
	return file_wgtwo_sim_v0_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *EapAkaAuthenticationVector) GetRand() []byte {
	if x != nil {
		return x.Rand
	}
	return nil
}

func (x *EapAkaAuthenticationVector) GetXres() []byte {
	if x != nil {
		return x.Xres
	}
	return nil
}

func (x *EapAkaAuthenticationVector) GetAutn() []byte {
	if x != nil {
		return x.Autn
	}
	return nil
}

func (x *EapAkaAuthenticationVector) GetCk() []byte {
	if x != nil {
		return x.Ck
	}
	return nil
}

func (x *EapAkaAuthenticationVector) GetIk() []byte {
	if x != nil {
		return x.Ik
	}
	return nil
}

var File_wgtwo_sim_v0_authentication_proto protoreflect.FileDescriptor

var file_wgtwo_sim_v0_authentication_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x2f, 0x76, 0x30, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76,
	0x30, 0x1a, 0x23, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x29, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45,
	0x61, 0x70, 0x41, 0x6b, 0x61, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x73, 0x69, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x22, 0x8b, 0x01, 0x0a, 0x2a,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x61, 0x70, 0x41, 0x6b, 0x61, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x15, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x61, 0x70, 0x41, 0x6b, 0x61, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x78, 0x0a, 0x1a, 0x45, 0x61, 0x70,
	0x41, 0x6b, 0x61, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x78,
	0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x78, 0x72, 0x65, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61,
	0x75, 0x74, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x6b, 0x32, 0xdc, 0x01, 0x0a, 0x18, 0x53, 0x69, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xbf, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x61, 0x70,
	0x41, 0x6b, 0x61, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x37, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x73, 0x69, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45,
	0x61, 0x70, 0x41, 0x6b, 0x61, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x30, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45, 0x61, 0x70, 0x41, 0x6b, 0x61, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0xea, 0xb5, 0x18, 0x22,
	0x73, 0x69, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x3a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x42, 0xae, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x73, 0x69, 0x6d, 0x2e, 0x76, 0x30, 0x42, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x73, 0x69, 0x6d,
	0x2f, 0x76, 0x30, 0xa2, 0x02, 0x03, 0x57, 0x53, 0x56, 0xaa, 0x02, 0x0c, 0x57, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x53, 0x69, 0x6d, 0x2e, 0x56, 0x30, 0xca, 0x02, 0x0c, 0x57, 0x67, 0x74, 0x77, 0x6f,
	0x5c, 0x53, 0x69, 0x6d, 0x5c, 0x56, 0x30, 0xe2, 0x02, 0x18, 0x57, 0x67, 0x74, 0x77, 0x6f, 0x5c,
	0x53, 0x69, 0x6d, 0x5c, 0x56, 0x30, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0e, 0x57, 0x67, 0x74, 0x77, 0x6f, 0x3a, 0x3a, 0x53, 0x69, 0x6d, 0x3a,
	0x3a, 0x56, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_sim_v0_authentication_proto_rawDescOnce sync.Once
	file_wgtwo_sim_v0_authentication_proto_rawDescData = file_wgtwo_sim_v0_authentication_proto_rawDesc
)

func file_wgtwo_sim_v0_authentication_proto_rawDescGZIP() []byte {
	file_wgtwo_sim_v0_authentication_proto_rawDescOnce.Do(func() {
		file_wgtwo_sim_v0_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_sim_v0_authentication_proto_rawDescData)
	})
	return file_wgtwo_sim_v0_authentication_proto_rawDescData
}

var file_wgtwo_sim_v0_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wgtwo_sim_v0_authentication_proto_goTypes = []any{
	(*GenerateEapAkaAuthenticationVectorRequest)(nil),  // 0: wgtwo.sim.v0.GenerateEapAkaAuthenticationVectorRequest
	(*GenerateEapAkaAuthenticationVectorResponse)(nil), // 1: wgtwo.sim.v0.GenerateEapAkaAuthenticationVectorResponse
	(*EapAkaAuthenticationVector)(nil),                 // 2: wgtwo.sim.v0.EapAkaAuthenticationVector
	(*v1.Imsi)(nil),                                    // 3: wgtwo.common.v1.Imsi
}
var file_wgtwo_sim_v0_authentication_proto_depIdxs = []int32{
	3, // 0: wgtwo.sim.v0.GenerateEapAkaAuthenticationVectorRequest.imsi:type_name -> wgtwo.common.v1.Imsi
	2, // 1: wgtwo.sim.v0.GenerateEapAkaAuthenticationVectorResponse.authentication_vector:type_name -> wgtwo.sim.v0.EapAkaAuthenticationVector
	0, // 2: wgtwo.sim.v0.SimAuthenticationService.GenerateEapAkaAuthenticationVector:input_type -> wgtwo.sim.v0.GenerateEapAkaAuthenticationVectorRequest
	1, // 3: wgtwo.sim.v0.SimAuthenticationService.GenerateEapAkaAuthenticationVector:output_type -> wgtwo.sim.v0.GenerateEapAkaAuthenticationVectorResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wgtwo_sim_v0_authentication_proto_init() }
func file_wgtwo_sim_v0_authentication_proto_init() {
	if File_wgtwo_sim_v0_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_sim_v0_authentication_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateEapAkaAuthenticationVectorRequest); i {
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
		file_wgtwo_sim_v0_authentication_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateEapAkaAuthenticationVectorResponse); i {
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
		file_wgtwo_sim_v0_authentication_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EapAkaAuthenticationVector); i {
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
			RawDescriptor: file_wgtwo_sim_v0_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_sim_v0_authentication_proto_goTypes,
		DependencyIndexes: file_wgtwo_sim_v0_authentication_proto_depIdxs,
		MessageInfos:      file_wgtwo_sim_v0_authentication_proto_msgTypes,
	}.Build()
	File_wgtwo_sim_v0_authentication_proto = out.File
	file_wgtwo_sim_v0_authentication_proto_rawDesc = nil
	file_wgtwo_sim_v0_authentication_proto_goTypes = nil
	file_wgtwo_sim_v0_authentication_proto_depIdxs = nil
}
