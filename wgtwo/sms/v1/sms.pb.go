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
// source: wgtwo/sms/v1/sms.proto

package v1

import (
	_ "github.com/working-group-two/wgtwoapis/wgtwo/annotations/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
// The class of the message.
//
// See https://en.wikipedia.org/wiki/Data_Coding_Scheme#Message_Classes
type MessageClass int32

const (
	// No message class specified.
	MessageClass_MESSAGE_CLASS_UNSPECIFIED MessageClass = 0
	// CLASS 0
	// A flash message is a message that is not stored on the device. It is handled
	// immediately or dropped, and also cannot be multi-fragment. Data messages of
	// this type needs a destination application port to designate what app will
	// handle it.
	MessageClass_MESSAGE_CLASS_FLASH_MESSAGE MessageClass = 1
	// CLASS 1
	// Mobile Equipment specific messages are handled by by an application on the
	// handset device itself, e.g. WAP push messages. Data messages of this type
	// needs a destination application port to designate what app will handle it.
	MessageClass_MESSAGE_CLASS_ME_SPECIFIC MessageClass = 2
	// CLASS 2
	// Handled by the sim card. SIM updates have special requirements tied to the
	// SIM card itself, and must be signed with a SIM specific private key only known
	// by the operator.
	MessageClass_MESSAGE_CLASS_SIM_SPECIFIC MessageClass = 3
	// CLASS 3
	// Terminal Equipment specific message are handled by the handset itself, or by
	// a SIM application, and may require an application port to designate who will
	// handle it.
	MessageClass_MESSAGE_CLASS_TE_SPECIFIC MessageClass = 4
)

// Enum value maps for MessageClass.
var (
	MessageClass_name = map[int32]string{
		0: "MESSAGE_CLASS_UNSPECIFIED",
		1: "MESSAGE_CLASS_FLASH_MESSAGE",
		2: "MESSAGE_CLASS_ME_SPECIFIC",
		3: "MESSAGE_CLASS_SIM_SPECIFIC",
		4: "MESSAGE_CLASS_TE_SPECIFIC",
	}
	MessageClass_value = map[string]int32{
		"MESSAGE_CLASS_UNSPECIFIED":   0,
		"MESSAGE_CLASS_FLASH_MESSAGE": 1,
		"MESSAGE_CLASS_ME_SPECIFIC":   2,
		"MESSAGE_CLASS_SIM_SPECIFIC":  3,
		"MESSAGE_CLASS_TE_SPECIFIC":   4,
	}
)

func (x MessageClass) Enum() *MessageClass {
	p := new(MessageClass)
	*p = x
	return p
}

func (x MessageClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageClass) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_sms_v1_sms_proto_enumTypes[0].Descriptor()
}

func (MessageClass) Type() protoreflect.EnumType {
	return &file_wgtwo_sms_v1_sms_proto_enumTypes[0]
}

func (x MessageClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageClass.Descriptor instead.
func (MessageClass) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{0}
}

type SendMessageResponse_SendStatus int32

const (
	SendMessageResponse_SEND_STATUS_UNSPECIFIED SendMessageResponse_SendStatus = 0
	// Sending the message is accepted. Note that the message is not delivered yet.
	SendMessageResponse_SEND_STATUS_OK SendMessageResponse_SendStatus = 1
	// Sending the message is rejected (not allowed). This may be because of
	// subscriber policy limitations, rights of the product or content of the
	// message. See description for details.
	SendMessageResponse_SEND_STATUS_REJECT SendMessageResponse_SendStatus = 2
	// Sending the message failed. This error should be treated as temporary, and
	// sending the message again may work.
	SendMessageResponse_SEND_STATUS_ERROR SendMessageResponse_SendStatus = 3
)

// Enum value maps for SendMessageResponse_SendStatus.
var (
	SendMessageResponse_SendStatus_name = map[int32]string{
		0: "SEND_STATUS_UNSPECIFIED",
		1: "SEND_STATUS_OK",
		2: "SEND_STATUS_REJECT",
		3: "SEND_STATUS_ERROR",
	}
	SendMessageResponse_SendStatus_value = map[string]int32{
		"SEND_STATUS_UNSPECIFIED": 0,
		"SEND_STATUS_OK":          1,
		"SEND_STATUS_REJECT":      2,
		"SEND_STATUS_ERROR":       3,
	}
)

func (x SendMessageResponse_SendStatus) Enum() *SendMessageResponse_SendStatus {
	p := new(SendMessageResponse_SendStatus)
	*p = x
	return p
}

func (x SendMessageResponse_SendStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMessageResponse_SendStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_wgtwo_sms_v1_sms_proto_enumTypes[1].Descriptor()
}

func (SendMessageResponse_SendStatus) Type() protoreflect.EnumType {
	return &file_wgtwo_sms_v1_sms_proto_enumTypes[1]
}

func (x SendMessageResponse_SendStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMessageResponse_SendStatus.Descriptor instead.
func (SendMessageResponse_SendStatus) EnumDescriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{4, 0}
}

type SendTextFromSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message text content. Minimum 1 character, maximum 2000 characters. Supports
	// unicode, though completeness is dependent on the receiver handset.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The subscriber number to send this. The sending product must have a right
	// to send as the subscriber specifically. E.g. operators can only send as
	// their own subscribers, third party products only as subscribers who have
	// enabled the product. Must be international number starting with '+'.
	FromSubscriber string `protobuf:"bytes,2,opt,name=from_subscriber,json=fromSubscriber,proto3" json:"from_subscriber,omitempty"`
	// The destination number of the message. Can be international starting
	// with '+', short form number or network specific numbers.
	ToAddress string `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Deadline to which the message needs to be delivered. If this is passed
	// and a delivery is not succeeded, the message delivery will fail. If not set
	// will use the maximum deadline. Maximum is 7 days.
	DeliveryDeadline *durationpb.Duration `protobuf:"bytes,4,opt,name=delivery_deadline,json=deliveryDeadline,proto3" json:"delivery_deadline,omitempty"`
}

func (x *SendTextFromSubscriberRequest) Reset() {
	*x = SendTextFromSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTextFromSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTextFromSubscriberRequest) ProtoMessage() {}

func (x *SendTextFromSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTextFromSubscriberRequest.ProtoReflect.Descriptor instead.
func (*SendTextFromSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SendTextFromSubscriberRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendTextFromSubscriberRequest) GetFromSubscriber() string {
	if x != nil {
		return x.FromSubscriber
	}
	return ""
}

func (x *SendTextFromSubscriberRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SendTextFromSubscriberRequest) GetDeliveryDeadline() *durationpb.Duration {
	if x != nil {
		return x.DeliveryDeadline
	}
	return nil
}

type SendTextToSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message text content. Minimum 1 character, maximum 2000 characters. Supports
	// unicode, though completeness is dependent on the receiver handset.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The destination phone number of the subscriber to receive the
	// message. Must be international number starting with '+'.
	ToSubscriber string `protobuf:"bytes,2,opt,name=to_subscriber,json=toSubscriber,proto3" json:"to_subscriber,omitempty"`
	// Origin address of the message.
	//
	// It can either be a
	// - a phone number formatted as E.164 starting with '+'.
	// - a alphanumeric sender ID.
	// - short form number.
	// - network specific number.
	//
	// Typical values here would be to send from your product's name.
	//
	// Important: Address must be pre-approved by Working Group Two.
	// See docs on origin addresses for what is allowed.
	FromAddress string `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// Deadline to which the message needs to be delivered. If this is passed
	// and a delivery is not succeeded, the message delivery will fail. If not set
	// will use the maximum deadline. Maximum is 7 days.
	DeliveryDeadline *durationpb.Duration `protobuf:"bytes,4,opt,name=delivery_deadline,json=deliveryDeadline,proto3" json:"delivery_deadline,omitempty"`
}

func (x *SendTextToSubscriberRequest) Reset() {
	*x = SendTextToSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTextToSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTextToSubscriberRequest) ProtoMessage() {}

func (x *SendTextToSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTextToSubscriberRequest.ProtoReflect.Descriptor instead.
func (*SendTextToSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{1}
}

func (x *SendTextToSubscriberRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendTextToSubscriberRequest) GetToSubscriber() string {
	if x != nil {
		return x.ToSubscriber
	}
	return ""
}

func (x *SendTextToSubscriberRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *SendTextToSubscriberRequest) GetDeliveryDeadline() *durationpb.Duration {
	if x != nil {
		return x.DeliveryDeadline
	}
	return nil
}

type SendDataToSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The binary content of the data SMS. Must be at least 1 byte, and maximum
	// 2000 bytes.
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The destination phone number of the subscriber to receive the
	// message. Must be international number starting with '+'.
	ToSubscriber string `protobuf:"bytes,2,opt,name=to_subscriber,json=toSubscriber,proto3" json:"to_subscriber,omitempty"`
	// Origin address of the message.
	//
	// It can either be a
	// - a phone number formatted as E.164 starting with '+'.
	// - a alphanumeric sender ID.
	// - short form number.
	// - network specific number.
	//
	// Typical values here would be to send from your product's name.
	//
	// Important: Address must be pre-approved by Working Group Two.
	// See docs on origin addresses for what is allowed.
	FromAddress string `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// Deadline to which the message needs to be delivered. If this is passed
	// and a delivery is not succeeded, the message delivery will fail. If not set
	// will use the maximum deadline. Maximum is 7 days.
	DeliveryDeadline *durationpb.Duration `protobuf:"bytes,4,opt,name=delivery_deadline,json=deliveryDeadline,proto3" json:"delivery_deadline,omitempty"`
	// The message class to use for the data SMS message. Must be set.
	MessageClass MessageClass `protobuf:"varint,5,opt,name=message_class,json=messageClass,proto3,enum=wgtwo.sms.v1.MessageClass" json:"message_class,omitempty"`
	// The application port for the message.
	ApplicationPort *ApplicationPort `protobuf:"bytes,6,opt,name=application_port,json=applicationPort,proto3" json:"application_port,omitempty"`
}

func (x *SendDataToSubscriberRequest) Reset() {
	*x = SendDataToSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDataToSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDataToSubscriberRequest) ProtoMessage() {}

func (x *SendDataToSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDataToSubscriberRequest.ProtoReflect.Descriptor instead.
func (*SendDataToSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{2}
}

func (x *SendDataToSubscriberRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SendDataToSubscriberRequest) GetToSubscriber() string {
	if x != nil {
		return x.ToSubscriber
	}
	return ""
}

func (x *SendDataToSubscriberRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *SendDataToSubscriberRequest) GetDeliveryDeadline() *durationpb.Duration {
	if x != nil {
		return x.DeliveryDeadline
	}
	return nil
}

func (x *SendDataToSubscriberRequest) GetMessageClass() MessageClass {
	if x != nil {
		return x.MessageClass
	}
	return MessageClass_MESSAGE_CLASS_UNSPECIFIED
}

func (x *SendDataToSubscriberRequest) GetApplicationPort() *ApplicationPort {
	if x != nil {
		return x.ApplicationPort
	}
	return nil
}

// Application ports are used to send data SMS messages to specific applications
// on the handset. If a reply to the message is sent, it should use the same
// ports but swap originator and destination port numbers.
type ApplicationPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The originator application port.
	OriginatorPort uint32 `protobuf:"varint,1,opt,name=originator_port,json=originatorPort,proto3" json:"originator_port,omitempty"`
	// The destination application port.
	DestinationPort uint32 `protobuf:"varint,2,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
}

func (x *ApplicationPort) Reset() {
	*x = ApplicationPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationPort) ProtoMessage() {}

func (x *ApplicationPort) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationPort.ProtoReflect.Descriptor instead.
func (*ApplicationPort) Descriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationPort) GetOriginatorPort() uint32 {
	if x != nil {
		return x.OriginatorPort
	}
	return 0
}

func (x *ApplicationPort) GetDestinationPort() uint32 {
	if x != nil {
		return x.DestinationPort
	}
	return 0
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An ID representing the message. For events etc related to the message, this ID
	// will be used in the event as identifier.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The response status for attempting to send the message.
	Status SendMessageResponse_SendStatus `protobuf:"varint,2,opt,name=status,proto3,enum=wgtwo.sms.v1.SendMessageResponse_SendStatus" json:"status,omitempty"`
	// Human readable description for what failed or rejected the message.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Number of fragments sent. This is because of the underlying SMS protocols only
	// supports sending 140 bytes per message after encoding and packing. This is including
	// extra encoding info and correlation and part handling for multi-fragmented messages.
	NumFragments uint32 `protobuf:"varint,4,opt,name=num_fragments,json=numFragments,proto3" json:"num_fragments,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_sms_v1_sms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_sms_v1_sms_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *SendMessageResponse) GetStatus() SendMessageResponse_SendStatus {
	if x != nil {
		return x.Status
	}
	return SendMessageResponse_SEND_STATUS_UNSPECIFIED
}

func (x *SendMessageResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SendMessageResponse) GetNumFragments() uint32 {
	if x != nil {
		return x.NumFragments
	}
	return 0
}

var File_wgtwo_sms_v1_sms_proto protoreflect.FileDescriptor

var file_wgtwo_sms_v1_sms_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9,
	0x01, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x1b, 0x53,
	0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x11,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x22, 0xd2, 0x02, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3f,
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x48, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x65, 0x0a, 0x0f, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0xaf, 0x02, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x46, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x6c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x4b, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x45, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x03, 0x2a, 0xac, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x5f, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43,
	0x4c, 0x41, 0x53, 0x53, 0x5f, 0x4d, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43,
	0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x5f, 0x53, 0x49, 0x4d, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43,
	0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c,
	0x41, 0x53, 0x53, 0x5f, 0x54, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x10,
	0x04, 0x32, 0xaa, 0x03, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x8b, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x77, 0x67,
	0x74, 0x77, 0x6f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54,
	0x65, 0x78, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xea, 0xb5, 0x18,
	0x1d, 0x73, 0x6d, 0x73, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x3a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x85,
	0x01, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x78, 0x74, 0x54,
	0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xea, 0xb5, 0x18, 0x1b, 0x73, 0x6d, 0x73, 0x2e, 0x74,
	0x65, 0x78, 0x74, 0x3a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12,
	0x29, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x6f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xea,
	0xb5, 0x18, 0x1b, 0x73, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x42, 0x55,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x6d, 0x73, 0x42, 0x08, 0x53, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x73,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wgtwo_sms_v1_sms_proto_rawDescOnce sync.Once
	file_wgtwo_sms_v1_sms_proto_rawDescData = file_wgtwo_sms_v1_sms_proto_rawDesc
)

func file_wgtwo_sms_v1_sms_proto_rawDescGZIP() []byte {
	file_wgtwo_sms_v1_sms_proto_rawDescOnce.Do(func() {
		file_wgtwo_sms_v1_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_sms_v1_sms_proto_rawDescData)
	})
	return file_wgtwo_sms_v1_sms_proto_rawDescData
}

var file_wgtwo_sms_v1_sms_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_wgtwo_sms_v1_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wgtwo_sms_v1_sms_proto_goTypes = []interface{}{
	(MessageClass)(0),                     // 0: wgtwo.sms.v1.MessageClass
	(SendMessageResponse_SendStatus)(0),   // 1: wgtwo.sms.v1.SendMessageResponse.SendStatus
	(*SendTextFromSubscriberRequest)(nil), // 2: wgtwo.sms.v1.SendTextFromSubscriberRequest
	(*SendTextToSubscriberRequest)(nil),   // 3: wgtwo.sms.v1.SendTextToSubscriberRequest
	(*SendDataToSubscriberRequest)(nil),   // 4: wgtwo.sms.v1.SendDataToSubscriberRequest
	(*ApplicationPort)(nil),               // 5: wgtwo.sms.v1.ApplicationPort
	(*SendMessageResponse)(nil),           // 6: wgtwo.sms.v1.SendMessageResponse
	(*durationpb.Duration)(nil),           // 7: google.protobuf.Duration
}
var file_wgtwo_sms_v1_sms_proto_depIdxs = []int32{
	7, // 0: wgtwo.sms.v1.SendTextFromSubscriberRequest.delivery_deadline:type_name -> google.protobuf.Duration
	7, // 1: wgtwo.sms.v1.SendTextToSubscriberRequest.delivery_deadline:type_name -> google.protobuf.Duration
	7, // 2: wgtwo.sms.v1.SendDataToSubscriberRequest.delivery_deadline:type_name -> google.protobuf.Duration
	0, // 3: wgtwo.sms.v1.SendDataToSubscriberRequest.message_class:type_name -> wgtwo.sms.v1.MessageClass
	5, // 4: wgtwo.sms.v1.SendDataToSubscriberRequest.application_port:type_name -> wgtwo.sms.v1.ApplicationPort
	1, // 5: wgtwo.sms.v1.SendMessageResponse.status:type_name -> wgtwo.sms.v1.SendMessageResponse.SendStatus
	2, // 6: wgtwo.sms.v1.SmsService.SendTextFromSubscriber:input_type -> wgtwo.sms.v1.SendTextFromSubscriberRequest
	3, // 7: wgtwo.sms.v1.SmsService.SendTextToSubscriber:input_type -> wgtwo.sms.v1.SendTextToSubscriberRequest
	4, // 8: wgtwo.sms.v1.SmsService.SendDataToSubscriber:input_type -> wgtwo.sms.v1.SendDataToSubscriberRequest
	6, // 9: wgtwo.sms.v1.SmsService.SendTextFromSubscriber:output_type -> wgtwo.sms.v1.SendMessageResponse
	6, // 10: wgtwo.sms.v1.SmsService.SendTextToSubscriber:output_type -> wgtwo.sms.v1.SendMessageResponse
	6, // 11: wgtwo.sms.v1.SmsService.SendDataToSubscriber:output_type -> wgtwo.sms.v1.SendMessageResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_wgtwo_sms_v1_sms_proto_init() }
func file_wgtwo_sms_v1_sms_proto_init() {
	if File_wgtwo_sms_v1_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_sms_v1_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTextFromSubscriberRequest); i {
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
		file_wgtwo_sms_v1_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTextToSubscriberRequest); i {
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
		file_wgtwo_sms_v1_sms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDataToSubscriberRequest); i {
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
		file_wgtwo_sms_v1_sms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationPort); i {
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
		file_wgtwo_sms_v1_sms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
			RawDescriptor: file_wgtwo_sms_v1_sms_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_sms_v1_sms_proto_goTypes,
		DependencyIndexes: file_wgtwo_sms_v1_sms_proto_depIdxs,
		EnumInfos:         file_wgtwo_sms_v1_sms_proto_enumTypes,
		MessageInfos:      file_wgtwo_sms_v1_sms_proto_msgTypes,
	}.Build()
	File_wgtwo_sms_v1_sms_proto = out.File
	file_wgtwo_sms_v1_sms_proto_rawDesc = nil
	file_wgtwo_sms_v1_sms_proto_goTypes = nil
	file_wgtwo_sms_v1_sms_proto_depIdxs = nil
}
