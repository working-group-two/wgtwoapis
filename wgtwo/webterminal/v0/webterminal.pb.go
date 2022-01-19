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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: wgtwo/webterminal/v0/webterminal.proto

package v0

import (
	_ "github.com/working-group-two/wgtwoapis/wgtwo/annotations/v0"
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

type Error_Code int32

const (
	Error_UNKNOWN         Error_Code = 0
	Error_NOT_INCALL      Error_Code = 1
	Error_TRY_AGAIN       Error_Code = 2
	Error_INVALID_CALL_ID Error_Code = 3
	Error_RATE_LIMIT_HIT  Error_Code = 4
)

// Enum value maps for Error_Code.
var (
	Error_Code_name = map[int32]string{
		0: "UNKNOWN",
		1: "NOT_INCALL",
		2: "TRY_AGAIN",
		3: "INVALID_CALL_ID",
		4: "RATE_LIMIT_HIT",
	}
	Error_Code_value = map[string]int32{
		"UNKNOWN":         0,
		"NOT_INCALL":      1,
		"TRY_AGAIN":       2,
		"INVALID_CALL_ID": 3,
		"RATE_LIMIT_HIT":  4,
	}
)

func (x Error_Code) Enum() *Error_Code {
	p := new(Error_Code)
	*p = x
	return p
}

func (x Error_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_webterminal_v0_webterminal_proto_enumTypes[0].Descriptor()
}

func (Error_Code) Type() protoreflect.EnumType {
	return &file_wgtwo_webterminal_v0_webterminal_proto_enumTypes[0]
}

func (x Error_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Code.Descriptor instead.
func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{7, 0}
}

type WebTerminalMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*WebTerminalMessage_Offer
	//	*WebTerminalMessage_Answer
	//	*WebTerminalMessage_Ringing
	//	*WebTerminalMessage_Bye
	//	*WebTerminalMessage_Idle
	//	*WebTerminalMessage_InCall
	//	*WebTerminalMessage_Error
	//	*WebTerminalMessage_Action
	Message isWebTerminalMessage_Message `protobuf_oneof:"message"`
	CallId  string                       `protobuf:"bytes,10,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
}

func (x *WebTerminalMessage) Reset() {
	*x = WebTerminalMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebTerminalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebTerminalMessage) ProtoMessage() {}

func (x *WebTerminalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebTerminalMessage.ProtoReflect.Descriptor instead.
func (*WebTerminalMessage) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{0}
}

func (m *WebTerminalMessage) GetMessage() isWebTerminalMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *WebTerminalMessage) GetOffer() *Offer {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Offer); ok {
		return x.Offer
	}
	return nil
}

func (x *WebTerminalMessage) GetAnswer() *Answer {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Answer); ok {
		return x.Answer
	}
	return nil
}

func (x *WebTerminalMessage) GetRinging() *Ringing {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Ringing); ok {
		return x.Ringing
	}
	return nil
}

func (x *WebTerminalMessage) GetBye() *Bye {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Bye); ok {
		return x.Bye
	}
	return nil
}

func (x *WebTerminalMessage) GetIdle() *Idle {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Idle); ok {
		return x.Idle
	}
	return nil
}

func (x *WebTerminalMessage) GetInCall() *InCall {
	if x, ok := x.GetMessage().(*WebTerminalMessage_InCall); ok {
		return x.InCall
	}
	return nil
}

func (x *WebTerminalMessage) GetError() *Error {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Error); ok {
		return x.Error
	}
	return nil
}

func (x *WebTerminalMessage) GetAction() *Action {
	if x, ok := x.GetMessage().(*WebTerminalMessage_Action); ok {
		return x.Action
	}
	return nil
}

func (x *WebTerminalMessage) GetCallId() string {
	if x != nil {
		return x.CallId
	}
	return ""
}

type isWebTerminalMessage_Message interface {
	isWebTerminalMessage_Message()
}

type WebTerminalMessage_Offer struct {
	Offer *Offer `protobuf:"bytes,1,opt,name=offer,proto3,oneof"`
}

type WebTerminalMessage_Answer struct {
	Answer *Answer `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

type WebTerminalMessage_Ringing struct {
	Ringing *Ringing `protobuf:"bytes,3,opt,name=ringing,proto3,oneof"`
}

type WebTerminalMessage_Bye struct {
	Bye *Bye `protobuf:"bytes,4,opt,name=bye,proto3,oneof"`
}

type WebTerminalMessage_Idle struct {
	Idle *Idle `protobuf:"bytes,5,opt,name=idle,proto3,oneof"`
}

type WebTerminalMessage_InCall struct {
	InCall *InCall `protobuf:"bytes,6,opt,name=in_call,json=inCall,proto3,oneof"`
}

type WebTerminalMessage_Error struct {
	Error *Error `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

type WebTerminalMessage_Action struct {
	Action *Action `protobuf:"bytes,9,opt,name=action,proto3,oneof"`
}

func (*WebTerminalMessage_Offer) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_Answer) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_Ringing) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_Bye) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_Idle) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_InCall) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_Error) isWebTerminalMessage_Message() {}

func (*WebTerminalMessage_Action) isWebTerminalMessage_Message() {}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdp    string          `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
	Msisdn *v0.PhoneNumber `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{1}
}

func (x *Offer) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

func (x *Offer) GetMsisdn() *v0.PhoneNumber {
	if x != nil {
		return x.Msisdn
	}
	return nil
}

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdp string `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{2}
}

func (x *Answer) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type Ringing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ringing) Reset() {
	*x = Ringing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ringing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ringing) ProtoMessage() {}

func (x *Ringing) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ringing.ProtoReflect.Descriptor instead.
func (*Ringing) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{3}
}

type Bye struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Bye) Reset() {
	*x = Bye{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bye) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bye) ProtoMessage() {}

func (x *Bye) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bye.ProtoReflect.Descriptor instead.
func (*Bye) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{4}
}

type Idle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Idle) Reset() {
	*x = Idle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Idle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Idle) ProtoMessage() {}

func (x *Idle) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Idle.ProtoReflect.Descriptor instead.
func (*Idle) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{5}
}

type InCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msisdn *v0.PhoneNumber `protobuf:"bytes,1,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
}

func (x *InCall) Reset() {
	*x = InCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InCall) ProtoMessage() {}

func (x *InCall) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InCall.ProtoReflect.Descriptor instead.
func (*InCall) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{6}
}

func (x *InCall) GetMsisdn() *v0.PhoneNumber {
	if x != nil {
		return x.Msisdn
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err Error_Code `protobuf:"varint,1,opt,name=err,proto3,enum=wgtwo.webterminal.v0.Error_Code" json:"err,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{7}
}

func (x *Error) GetErr() Error_Code {
	if x != nil {
		return x.Err
	}
	return Error_UNKNOWN
}

type ToPhone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ToPhone) Reset() {
	*x = ToPhone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToPhone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToPhone) ProtoMessage() {}

func (x *ToPhone) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToPhone.ProtoReflect.Descriptor instead.
func (*ToPhone) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{8}
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*Action_ToPhone
	Message isAction_Message `protobuf_oneof:"message"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP(), []int{9}
}

func (m *Action) GetMessage() isAction_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *Action) GetToPhone() *ToPhone {
	if x, ok := x.GetMessage().(*Action_ToPhone); ok {
		return x.ToPhone
	}
	return nil
}

type isAction_Message interface {
	isAction_Message()
}

type Action_ToPhone struct {
	ToPhone *ToPhone `protobuf:"bytes,1,opt,name=toPhone,proto3,oneof"`
}

func (*Action_ToPhone) isAction_Message() {}

var File_wgtwo_webterminal_v0_webterminal_proto protoreflect.FileDescriptor

var file_wgtwo_webterminal_v0_webterminal_proto_rawDesc = []byte{
	0x0a, 0x26, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x30, 0x2f, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x1a, 0x23,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x03, 0x0a, 0x12, 0x57, 0x65, 0x62, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a,
	0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x2e, 0x76, 0x30, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x69,
	0x6e, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x67,
	0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e,
	0x76, 0x30, 0x2e, 0x52, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x07, 0x72, 0x69,
	0x6e, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x03, 0x62, 0x79, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x42, 0x79, 0x65, 0x48, 0x00, 0x52,
	0x03, 0x62, 0x79, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x64, 0x6c, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x69, 0x64, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x49,
	0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12,
	0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x61, 0x6c, 0x6c, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x4f, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x64, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x12, 0x34, 0x0a, 0x06, 0x6d,
	0x73, 0x69, 0x73, 0x64, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67,
	0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x73, 0x69, 0x73, 0x64,
	0x6e, 0x22, 0x1a, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x64, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x22, 0x09, 0x0a,
	0x07, 0x52, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x05, 0x0a, 0x03, 0x42, 0x79, 0x65, 0x22,
	0x06, 0x0a, 0x04, 0x49, 0x64, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x49, 0x6e, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x34, 0x0a, 0x06, 0x6d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x32, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x5b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f,
	0x54, 0x5f, 0x49, 0x4e, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x52,
	0x59, 0x5f, 0x41, 0x47, 0x41, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x48, 0x49, 0x54,
	0x10, 0x04, 0x22, 0x09, 0x0a, 0x07, 0x54, 0x6f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x4e, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x07, 0x74, 0x6f, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e,
	0x54, 0x6f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x48, 0x00, 0x52, 0x07, 0x74, 0x6f, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa7, 0x02,
	0x0a, 0x12, 0x57, 0x65, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x04, 0x50, 0x69, 0x70, 0x65, 0x12, 0x28, 0x2e,
	0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x65, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x28, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x57,
	0x65, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x24, 0xea, 0xb5, 0x18, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x89, 0x01, 0x0a, 0x09,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x69, 0x70, 0x65, 0x12, 0x28, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30,
	0x2e, 0x57, 0x65, 0x62, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x28, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x77, 0x65, 0x62, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x30, 0x2e, 0x57, 0x65, 0x62, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0xea,
	0xb5, 0x18, 0x20, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x6f, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x77, 0x65, 0x62, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x42, 0x10, 0x57, 0x65, 0x62, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2d,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_webterminal_v0_webterminal_proto_rawDescOnce sync.Once
	file_wgtwo_webterminal_v0_webterminal_proto_rawDescData = file_wgtwo_webterminal_v0_webterminal_proto_rawDesc
)

func file_wgtwo_webterminal_v0_webterminal_proto_rawDescGZIP() []byte {
	file_wgtwo_webterminal_v0_webterminal_proto_rawDescOnce.Do(func() {
		file_wgtwo_webterminal_v0_webterminal_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_webterminal_v0_webterminal_proto_rawDescData)
	})
	return file_wgtwo_webterminal_v0_webterminal_proto_rawDescData
}

var file_wgtwo_webterminal_v0_webterminal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wgtwo_webterminal_v0_webterminal_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_wgtwo_webterminal_v0_webterminal_proto_goTypes = []interface{}{
	(Error_Code)(0),            // 0: wgtwo.webterminal.v0.Error.Code
	(*WebTerminalMessage)(nil), // 1: wgtwo.webterminal.v0.WebTerminalMessage
	(*Offer)(nil),              // 2: wgtwo.webterminal.v0.Offer
	(*Answer)(nil),             // 3: wgtwo.webterminal.v0.Answer
	(*Ringing)(nil),            // 4: wgtwo.webterminal.v0.Ringing
	(*Bye)(nil),                // 5: wgtwo.webterminal.v0.Bye
	(*Idle)(nil),               // 6: wgtwo.webterminal.v0.Idle
	(*InCall)(nil),             // 7: wgtwo.webterminal.v0.InCall
	(*Error)(nil),              // 8: wgtwo.webterminal.v0.Error
	(*ToPhone)(nil),            // 9: wgtwo.webterminal.v0.ToPhone
	(*Action)(nil),             // 10: wgtwo.webterminal.v0.Action
	(*v0.PhoneNumber)(nil),     // 11: wgtwo.common.v0.PhoneNumber
}
var file_wgtwo_webterminal_v0_webterminal_proto_depIdxs = []int32{
	2,  // 0: wgtwo.webterminal.v0.WebTerminalMessage.offer:type_name -> wgtwo.webterminal.v0.Offer
	3,  // 1: wgtwo.webterminal.v0.WebTerminalMessage.answer:type_name -> wgtwo.webterminal.v0.Answer
	4,  // 2: wgtwo.webterminal.v0.WebTerminalMessage.ringing:type_name -> wgtwo.webterminal.v0.Ringing
	5,  // 3: wgtwo.webterminal.v0.WebTerminalMessage.bye:type_name -> wgtwo.webterminal.v0.Bye
	6,  // 4: wgtwo.webterminal.v0.WebTerminalMessage.idle:type_name -> wgtwo.webterminal.v0.Idle
	7,  // 5: wgtwo.webterminal.v0.WebTerminalMessage.in_call:type_name -> wgtwo.webterminal.v0.InCall
	8,  // 6: wgtwo.webterminal.v0.WebTerminalMessage.error:type_name -> wgtwo.webterminal.v0.Error
	10, // 7: wgtwo.webterminal.v0.WebTerminalMessage.action:type_name -> wgtwo.webterminal.v0.Action
	11, // 8: wgtwo.webterminal.v0.Offer.msisdn:type_name -> wgtwo.common.v0.PhoneNumber
	11, // 9: wgtwo.webterminal.v0.InCall.msisdn:type_name -> wgtwo.common.v0.PhoneNumber
	0,  // 10: wgtwo.webterminal.v0.Error.err:type_name -> wgtwo.webterminal.v0.Error.Code
	9,  // 11: wgtwo.webterminal.v0.Action.toPhone:type_name -> wgtwo.webterminal.v0.ToPhone
	1,  // 12: wgtwo.webterminal.v0.WebTerminalService.Pipe:input_type -> wgtwo.webterminal.v0.WebTerminalMessage
	1,  // 13: wgtwo.webterminal.v0.WebTerminalService.MultiPipe:input_type -> wgtwo.webterminal.v0.WebTerminalMessage
	1,  // 14: wgtwo.webterminal.v0.WebTerminalService.Pipe:output_type -> wgtwo.webterminal.v0.WebTerminalMessage
	1,  // 15: wgtwo.webterminal.v0.WebTerminalService.MultiPipe:output_type -> wgtwo.webterminal.v0.WebTerminalMessage
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_wgtwo_webterminal_v0_webterminal_proto_init() }
func file_wgtwo_webterminal_v0_webterminal_proto_init() {
	if File_wgtwo_webterminal_v0_webterminal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebTerminalMessage); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ringing); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bye); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Idle); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InCall); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ToPhone); i {
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
		file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
	file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WebTerminalMessage_Offer)(nil),
		(*WebTerminalMessage_Answer)(nil),
		(*WebTerminalMessage_Ringing)(nil),
		(*WebTerminalMessage_Bye)(nil),
		(*WebTerminalMessage_Idle)(nil),
		(*WebTerminalMessage_InCall)(nil),
		(*WebTerminalMessage_Error)(nil),
		(*WebTerminalMessage_Action)(nil),
	}
	file_wgtwo_webterminal_v0_webterminal_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Action_ToPhone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_webterminal_v0_webterminal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_webterminal_v0_webterminal_proto_goTypes,
		DependencyIndexes: file_wgtwo_webterminal_v0_webterminal_proto_depIdxs,
		EnumInfos:         file_wgtwo_webterminal_v0_webterminal_proto_enumTypes,
		MessageInfos:      file_wgtwo_webterminal_v0_webterminal_proto_msgTypes,
	}.Build()
	File_wgtwo_webterminal_v0_webterminal_proto = out.File
	file_wgtwo_webterminal_v0_webterminal_proto_rawDesc = nil
	file_wgtwo_webterminal_v0_webterminal_proto_goTypes = nil
	file_wgtwo_webterminal_v0_webterminal_proto_depIdxs = nil
}
