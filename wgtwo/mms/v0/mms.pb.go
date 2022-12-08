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
// source: wgtwo/mms/v0/mms.proto

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

type SendResponse_SendStatus int32

const (
	SendResponse_UNKNOWN     SendResponse_SendStatus = 0
	SendResponse_SEND_OK     SendResponse_SendStatus = 1
	SendResponse_SEND_REJECT SendResponse_SendStatus = 2
	SendResponse_SEND_ERROR  SendResponse_SendStatus = 3
)

// Enum value maps for SendResponse_SendStatus.
var (
	SendResponse_SendStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SEND_OK",
		2: "SEND_REJECT",
		3: "SEND_ERROR",
	}
	SendResponse_SendStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"SEND_OK":     1,
		"SEND_REJECT": 2,
		"SEND_ERROR":  3,
	}
)

func (x SendResponse_SendStatus) Enum() *SendResponse_SendStatus {
	p := new(SendResponse_SendStatus)
	*p = x
	return p
}

func (x SendResponse_SendStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendResponse_SendStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_mms_v0_mms_proto_enumTypes[0].Descriptor()
}

func (SendResponse_SendStatus) Type() protoreflect.EnumType {
	return &file_wgtwo_mms_v0_mms_proto_enumTypes[0]
}

func (x SendResponse_SendStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendResponse_SendStatus.Descriptor instead.
func (SendResponse_SendStatus) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{6, 0}
}

type MessageContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*MessageContent_Audio
	//	*MessageContent_Text
	//	*MessageContent_Image
	Content isMessageContent_Content `protobuf_oneof:"content"`
}

func (x *MessageContent) Reset() {
	*x = MessageContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageContent) ProtoMessage() {}

func (x *MessageContent) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageContent.ProtoReflect.Descriptor instead.
func (*MessageContent) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{0}
}

func (m *MessageContent) GetContent() isMessageContent_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *MessageContent) GetAudio() *AudioContent {
	if x, ok := x.GetContent().(*MessageContent_Audio); ok {
		return x.Audio
	}
	return nil
}

func (x *MessageContent) GetText() *TextContent {
	if x, ok := x.GetContent().(*MessageContent_Text); ok {
		return x.Text
	}
	return nil
}

func (x *MessageContent) GetImage() *ImageContent {
	if x, ok := x.GetContent().(*MessageContent_Image); ok {
		return x.Image
	}
	return nil
}

type isMessageContent_Content interface {
	isMessageContent_Content()
}

type MessageContent_Audio struct {
	Audio *AudioContent `protobuf:"bytes,1,opt,name=audio,proto3,oneof"`
}

type MessageContent_Text struct {
	Text *TextContent `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type MessageContent_Image struct {
	Image *ImageContent `protobuf:"bytes,3,opt,name=image,proto3,oneof"`
}

func (*MessageContent_Audio) isMessageContent_Content() {}

func (*MessageContent_Text) isMessageContent_Content() {}

func (*MessageContent_Image) isMessageContent_Content() {}

type AudioContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*AudioContent_Mp3
	//	*AudioContent_Opus
	//	*AudioContent_Amr
	//	*AudioContent_Wav
	Content isAudioContent_Content `protobuf_oneof:"content"`
}

func (x *AudioContent) Reset() {
	*x = AudioContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioContent) ProtoMessage() {}

func (x *AudioContent) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioContent.ProtoReflect.Descriptor instead.
func (*AudioContent) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{1}
}

func (m *AudioContent) GetContent() isAudioContent_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *AudioContent) GetMp3() []byte {
	if x, ok := x.GetContent().(*AudioContent_Mp3); ok {
		return x.Mp3
	}
	return nil
}

func (x *AudioContent) GetOpus() []byte {
	if x, ok := x.GetContent().(*AudioContent_Opus); ok {
		return x.Opus
	}
	return nil
}

func (x *AudioContent) GetAmr() []byte {
	if x, ok := x.GetContent().(*AudioContent_Amr); ok {
		return x.Amr
	}
	return nil
}

func (x *AudioContent) GetWav() []byte {
	if x, ok := x.GetContent().(*AudioContent_Wav); ok {
		return x.Wav
	}
	return nil
}

type isAudioContent_Content interface {
	isAudioContent_Content()
}

type AudioContent_Mp3 struct {
	Mp3 []byte `protobuf:"bytes,1,opt,name=mp3,proto3,oneof"`
}

type AudioContent_Opus struct {
	Opus []byte `protobuf:"bytes,2,opt,name=opus,proto3,oneof"`
}

type AudioContent_Amr struct {
	Amr []byte `protobuf:"bytes,3,opt,name=amr,proto3,oneof"`
}

type AudioContent_Wav struct {
	Wav []byte `protobuf:"bytes,4,opt,name=wav,proto3,oneof"`
}

func (*AudioContent_Mp3) isAudioContent_Content() {}

func (*AudioContent_Opus) isAudioContent_Content() {}

func (*AudioContent_Amr) isAudioContent_Content() {}

func (*AudioContent_Wav) isAudioContent_Content() {}

type TextContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextContent) Reset() {
	*x = TextContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextContent) ProtoMessage() {}

func (x *TextContent) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextContent.ProtoReflect.Descriptor instead.
func (*TextContent) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{2}
}

func (x *TextContent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ImageContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*ImageContent_Png
	//	*ImageContent_Jpg
	Content isImageContent_Content `protobuf_oneof:"content"`
}

func (x *ImageContent) Reset() {
	*x = ImageContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageContent) ProtoMessage() {}

func (x *ImageContent) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageContent.ProtoReflect.Descriptor instead.
func (*ImageContent) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{3}
}

func (m *ImageContent) GetContent() isImageContent_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ImageContent) GetPng() []byte {
	if x, ok := x.GetContent().(*ImageContent_Png); ok {
		return x.Png
	}
	return nil
}

func (x *ImageContent) GetJpg() []byte {
	if x, ok := x.GetContent().(*ImageContent_Jpg); ok {
		return x.Jpg
	}
	return nil
}

type isImageContent_Content interface {
	isImageContent_Content()
}

type ImageContent_Png struct {
	Png []byte `protobuf:"bytes,1,opt,name=png,proto3,oneof"`
}

type ImageContent_Jpg struct {
	Jpg []byte `protobuf:"bytes,2,opt,name=jpg,proto3,oneof"`
}

func (*ImageContent_Png) isImageContent_Content() {}

func (*ImageContent_Jpg) isImageContent_Content() {}

type SendMessageToSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageContent []*MessageContent `protobuf:"bytes,1,rep,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"`
	// Types that are assignable to FromAddress:
	//	*SendMessageToSubscriberRequest_FromE164
	//	*SendMessageToSubscriberRequest_FromTextAddress
	FromAddress  isSendMessageToSubscriberRequest_FromAddress `protobuf_oneof:"from_address"`
	ToSubscriber *v0.PhoneNumber                              `protobuf:"bytes,5,opt,name=to_subscriber,json=toSubscriber,proto3" json:"to_subscriber,omitempty"`
}

func (x *SendMessageToSubscriberRequest) Reset() {
	*x = SendMessageToSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageToSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageToSubscriberRequest) ProtoMessage() {}

func (x *SendMessageToSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageToSubscriberRequest.ProtoReflect.Descriptor instead.
func (*SendMessageToSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageToSubscriberRequest) GetMessageContent() []*MessageContent {
	if x != nil {
		return x.MessageContent
	}
	return nil
}

func (m *SendMessageToSubscriberRequest) GetFromAddress() isSendMessageToSubscriberRequest_FromAddress {
	if m != nil {
		return m.FromAddress
	}
	return nil
}

func (x *SendMessageToSubscriberRequest) GetFromE164() *v0.PhoneNumber {
	if x, ok := x.GetFromAddress().(*SendMessageToSubscriberRequest_FromE164); ok {
		return x.FromE164
	}
	return nil
}

func (x *SendMessageToSubscriberRequest) GetFromTextAddress() *v0.TextAddress {
	if x, ok := x.GetFromAddress().(*SendMessageToSubscriberRequest_FromTextAddress); ok {
		return x.FromTextAddress
	}
	return nil
}

func (x *SendMessageToSubscriberRequest) GetToSubscriber() *v0.PhoneNumber {
	if x != nil {
		return x.ToSubscriber
	}
	return nil
}

type isSendMessageToSubscriberRequest_FromAddress interface {
	isSendMessageToSubscriberRequest_FromAddress()
}

type SendMessageToSubscriberRequest_FromE164 struct {
	FromE164 *v0.PhoneNumber `protobuf:"bytes,2,opt,name=from_e164,json=fromE164,proto3,oneof"`
}

type SendMessageToSubscriberRequest_FromTextAddress struct {
	FromTextAddress *v0.TextAddress `protobuf:"bytes,3,opt,name=from_text_address,json=fromTextAddress,proto3,oneof"`
}

func (*SendMessageToSubscriberRequest_FromE164) isSendMessageToSubscriberRequest_FromAddress() {}

func (*SendMessageToSubscriberRequest_FromTextAddress) isSendMessageToSubscriberRequest_FromAddress() {
}

type SendMessageFromSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageContent []*MessageContent `protobuf:"bytes,1,rep,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"`
	FromSubscriber *v0.PhoneNumber   `protobuf:"bytes,2,opt,name=from_subscriber,json=fromSubscriber,proto3" json:"from_subscriber,omitempty"`
	ToE164         *v0.PhoneNumber   `protobuf:"bytes,3,opt,name=to_e164,json=toE164,proto3" json:"to_e164,omitempty"`
}

func (x *SendMessageFromSubscriberRequest) Reset() {
	*x = SendMessageFromSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageFromSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageFromSubscriberRequest) ProtoMessage() {}

func (x *SendMessageFromSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageFromSubscriberRequest.ProtoReflect.Descriptor instead.
func (*SendMessageFromSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{5}
}

func (x *SendMessageFromSubscriberRequest) GetMessageContent() []*MessageContent {
	if x != nil {
		return x.MessageContent
	}
	return nil
}

func (x *SendMessageFromSubscriberRequest) GetFromSubscriber() *v0.PhoneNumber {
	if x != nil {
		return x.FromSubscriber
	}
	return nil
}

func (x *SendMessageFromSubscriberRequest) GetToE164() *v0.PhoneNumber {
	if x != nil {
		return x.ToE164
	}
	return nil
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string                  `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Status      SendResponse_SendStatus `protobuf:"varint,2,opt,name=status,proto3,enum=wgtwo.mms.v0.SendResponse_SendStatus" json:"status,omitempty"`
	Description string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_mms_v0_mms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_mms_v0_mms_proto_rawDescGZIP(), []int{6}
}

func (x *SendResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SendResponse) GetStatus() SendResponse_SendStatus {
	if x != nil {
		return x.Status
	}
	return SendResponse_UNKNOWN
}

func (x *SendResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_wgtwo_mms_v0_mms_proto protoreflect.FileDescriptor

var file_wgtwo_mms_v0_mms_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x6d, 0x6d, 0x73, 0x2f, 0x76, 0x30, 0x2f, 0x6d,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x1a, 0x23, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4,
	0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e,
	0x76, 0x30, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d,
	0x73, 0x2e, 0x76, 0x30, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x6d, 0x70, 0x33, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x70, 0x33, 0x12, 0x14, 0x0a, 0x04, 0x6f, 0x70, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x70, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x03, 0x61, 0x6d, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03,
	0x61, 0x6d, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x77, 0x61, 0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x03, 0x77, 0x61, 0x76, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x70, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x70, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x03, 0x6a, 0x70, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x6a, 0x70, 0x67, 0x42, 0x09, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xc3, 0x02, 0x0a, 0x1e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73,
	0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x31, 0x36, 0x34, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x31, 0x36, 0x34, 0x12,
	0x4a, 0x0a, 0x11, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6d,
	0x54, 0x65, 0x78, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x0d, 0x74,
	0x6f, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x0c, 0x74, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x42, 0x0e,
	0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xe7,
	0x01, 0x0a, 0x20, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x12, 0x35, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x65, 0x31, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x06, 0x74, 0x6f, 0x45, 0x31, 0x36, 0x34, 0x22, 0xd7, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x0a, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x4b, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x03, 0x32, 0x95, 0x02, 0x0a, 0x0a, 0x4d, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7f, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0xea, 0xb5, 0x18, 0x16, 0x6d, 0x6d, 0x73, 0x2e,
	0x73, 0x65, 0x6e, 0x64, 0x2e, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x12, 0x85, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x12, 0x2e, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x6d, 0x6d, 0x73, 0x2e, 0x76, 0x30, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0xea, 0xb5,
	0x18, 0x18, 0x6d, 0x6d, 0x73, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x42, 0x55, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x6d,
	0x6d, 0x73, 0x42, 0x08, 0x4d, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x6d, 0x6d, 0x73, 0x2f, 0x76,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_mms_v0_mms_proto_rawDescOnce sync.Once
	file_wgtwo_mms_v0_mms_proto_rawDescData = file_wgtwo_mms_v0_mms_proto_rawDesc
)

func file_wgtwo_mms_v0_mms_proto_rawDescGZIP() []byte {
	file_wgtwo_mms_v0_mms_proto_rawDescOnce.Do(func() {
		file_wgtwo_mms_v0_mms_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_mms_v0_mms_proto_rawDescData)
	})
	return file_wgtwo_mms_v0_mms_proto_rawDescData
}

var file_wgtwo_mms_v0_mms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wgtwo_mms_v0_mms_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wgtwo_mms_v0_mms_proto_goTypes = []interface{}{
	(SendResponse_SendStatus)(0),             // 0: wgtwo.mms.v0.SendResponse.SendStatus
	(*MessageContent)(nil),                   // 1: wgtwo.mms.v0.MessageContent
	(*AudioContent)(nil),                     // 2: wgtwo.mms.v0.AudioContent
	(*TextContent)(nil),                      // 3: wgtwo.mms.v0.TextContent
	(*ImageContent)(nil),                     // 4: wgtwo.mms.v0.ImageContent
	(*SendMessageToSubscriberRequest)(nil),   // 5: wgtwo.mms.v0.SendMessageToSubscriberRequest
	(*SendMessageFromSubscriberRequest)(nil), // 6: wgtwo.mms.v0.SendMessageFromSubscriberRequest
	(*SendResponse)(nil),                     // 7: wgtwo.mms.v0.SendResponse
	(*v0.PhoneNumber)(nil),                   // 8: wgtwo.common.v0.PhoneNumber
	(*v0.TextAddress)(nil),                   // 9: wgtwo.common.v0.TextAddress
}
var file_wgtwo_mms_v0_mms_proto_depIdxs = []int32{
	2,  // 0: wgtwo.mms.v0.MessageContent.audio:type_name -> wgtwo.mms.v0.AudioContent
	3,  // 1: wgtwo.mms.v0.MessageContent.text:type_name -> wgtwo.mms.v0.TextContent
	4,  // 2: wgtwo.mms.v0.MessageContent.image:type_name -> wgtwo.mms.v0.ImageContent
	1,  // 3: wgtwo.mms.v0.SendMessageToSubscriberRequest.message_content:type_name -> wgtwo.mms.v0.MessageContent
	8,  // 4: wgtwo.mms.v0.SendMessageToSubscriberRequest.from_e164:type_name -> wgtwo.common.v0.PhoneNumber
	9,  // 5: wgtwo.mms.v0.SendMessageToSubscriberRequest.from_text_address:type_name -> wgtwo.common.v0.TextAddress
	8,  // 6: wgtwo.mms.v0.SendMessageToSubscriberRequest.to_subscriber:type_name -> wgtwo.common.v0.PhoneNumber
	1,  // 7: wgtwo.mms.v0.SendMessageFromSubscriberRequest.message_content:type_name -> wgtwo.mms.v0.MessageContent
	8,  // 8: wgtwo.mms.v0.SendMessageFromSubscriberRequest.from_subscriber:type_name -> wgtwo.common.v0.PhoneNumber
	8,  // 9: wgtwo.mms.v0.SendMessageFromSubscriberRequest.to_e164:type_name -> wgtwo.common.v0.PhoneNumber
	0,  // 10: wgtwo.mms.v0.SendResponse.status:type_name -> wgtwo.mms.v0.SendResponse.SendStatus
	5,  // 11: wgtwo.mms.v0.MmsService.SendMessageToSubscriber:input_type -> wgtwo.mms.v0.SendMessageToSubscriberRequest
	6,  // 12: wgtwo.mms.v0.MmsService.SendMessageFromSubscriber:input_type -> wgtwo.mms.v0.SendMessageFromSubscriberRequest
	7,  // 13: wgtwo.mms.v0.MmsService.SendMessageToSubscriber:output_type -> wgtwo.mms.v0.SendResponse
	7,  // 14: wgtwo.mms.v0.MmsService.SendMessageFromSubscriber:output_type -> wgtwo.mms.v0.SendResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_wgtwo_mms_v0_mms_proto_init() }
func file_wgtwo_mms_v0_mms_proto_init() {
	if File_wgtwo_mms_v0_mms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_mms_v0_mms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageContent); i {
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
		file_wgtwo_mms_v0_mms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioContent); i {
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
		file_wgtwo_mms_v0_mms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextContent); i {
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
		file_wgtwo_mms_v0_mms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageContent); i {
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
		file_wgtwo_mms_v0_mms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageToSubscriberRequest); i {
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
		file_wgtwo_mms_v0_mms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageFromSubscriberRequest); i {
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
		file_wgtwo_mms_v0_mms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
	file_wgtwo_mms_v0_mms_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MessageContent_Audio)(nil),
		(*MessageContent_Text)(nil),
		(*MessageContent_Image)(nil),
	}
	file_wgtwo_mms_v0_mms_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AudioContent_Mp3)(nil),
		(*AudioContent_Opus)(nil),
		(*AudioContent_Amr)(nil),
		(*AudioContent_Wav)(nil),
	}
	file_wgtwo_mms_v0_mms_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ImageContent_Png)(nil),
		(*ImageContent_Jpg)(nil),
	}
	file_wgtwo_mms_v0_mms_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*SendMessageToSubscriberRequest_FromE164)(nil),
		(*SendMessageToSubscriberRequest_FromTextAddress)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_mms_v0_mms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_mms_v0_mms_proto_goTypes,
		DependencyIndexes: file_wgtwo_mms_v0_mms_proto_depIdxs,
		EnumInfos:         file_wgtwo_mms_v0_mms_proto_enumTypes,
		MessageInfos:      file_wgtwo_mms_v0_mms_proto_msgTypes,
	}.Build()
	File_wgtwo_mms_v0_mms_proto = out.File
	file_wgtwo_mms_v0_mms_proto_rawDesc = nil
	file_wgtwo_mms_v0_mms_proto_goTypes = nil
	file_wgtwo_mms_v0_mms_proto_depIdxs = nil
}
