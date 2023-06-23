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
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: wgtwo/consents/v1/consent_events.proto

package v1

import (
	_ "github.com/working-group-two/wgtwoapis/wgtwo/annotations"
	v11 "github.com/working-group-two/wgtwoapis/wgtwo/common/v1"
	v1 "github.com/working-group-two/wgtwoapis/wgtwo/events/v1"
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

type StreamConsentChangeEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamConfiguration *v1.StreamConfiguration `protobuf:"bytes,1,opt,name=stream_configuration,json=streamConfiguration,proto3" json:"stream_configuration,omitempty"`
}

func (x *StreamConsentChangeEventsRequest) Reset() {
	*x = StreamConsentChangeEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamConsentChangeEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamConsentChangeEventsRequest) ProtoMessage() {}

func (x *StreamConsentChangeEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamConsentChangeEventsRequest.ProtoReflect.Descriptor instead.
func (*StreamConsentChangeEventsRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{0}
}

func (x *StreamConsentChangeEventsRequest) GetStreamConfiguration() *v1.StreamConfiguration {
	if x != nil {
		return x.StreamConfiguration
	}
	return nil
}

type StreamConsentChangeEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata           *v1.Metadata        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ConsentChangeEvent *ConsentChangeEvent `protobuf:"bytes,2,opt,name=consent_change_event,json=consentChangeEvent,proto3" json:"consent_change_event,omitempty"`
}

func (x *StreamConsentChangeEventsResponse) Reset() {
	*x = StreamConsentChangeEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamConsentChangeEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamConsentChangeEventsResponse) ProtoMessage() {}

func (x *StreamConsentChangeEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamConsentChangeEventsResponse.ProtoReflect.Descriptor instead.
func (*StreamConsentChangeEventsResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{1}
}

func (x *StreamConsentChangeEventsResponse) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StreamConsentChangeEventsResponse) GetConsentChangeEvent() *ConsentChangeEvent {
	if x != nil {
		return x.ConsentChangeEvent
	}
	return nil
}

type AckConsentChangeEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckInfo *v1.AckInfo `protobuf:"bytes,1,opt,name=ack_info,json=ackInfo,proto3" json:"ack_info,omitempty"`
}

func (x *AckConsentChangeEventRequest) Reset() {
	*x = AckConsentChangeEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckConsentChangeEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckConsentChangeEventRequest) ProtoMessage() {}

func (x *AckConsentChangeEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckConsentChangeEventRequest.ProtoReflect.Descriptor instead.
func (*AckConsentChangeEventRequest) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{2}
}

func (x *AckConsentChangeEventRequest) GetAckInfo() *v1.AckInfo {
	if x != nil {
		return x.AckInfo
	}
	return nil
}

type AckConsentChangeEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckStatus *v1.AckStatus `protobuf:"bytes,1,opt,name=ack_status,json=ackStatus,proto3" json:"ack_status,omitempty"`
}

func (x *AckConsentChangeEventResponse) Reset() {
	*x = AckConsentChangeEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckConsentChangeEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckConsentChangeEventResponse) ProtoMessage() {}

func (x *AckConsentChangeEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckConsentChangeEventResponse.ProtoReflect.Descriptor instead.
func (*AckConsentChangeEventResponse) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{3}
}

func (x *AckConsentChangeEventResponse) GetAckStatus() *v1.AckStatus {
	if x != nil {
		return x.AckStatus
	}
	return nil
}

type ConsentChangeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*ConsentChangeEvent_Added
	//	*ConsentChangeEvent_Updated
	//	*ConsentChangeEvent_Revoked
	Type isConsentChangeEvent_Type `protobuf_oneof:"type"`
	// Types that are assignable to Target:
	//
	//	*ConsentChangeEvent_Number
	//	*ConsentChangeEvent_Tenant
	Target isConsentChangeEvent_Target `protobuf_oneof:"target"`
}

func (x *ConsentChangeEvent) Reset() {
	*x = ConsentChangeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsentChangeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsentChangeEvent) ProtoMessage() {}

func (x *ConsentChangeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsentChangeEvent.ProtoReflect.Descriptor instead.
func (*ConsentChangeEvent) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{4}
}

func (m *ConsentChangeEvent) GetType() isConsentChangeEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ConsentChangeEvent) GetAdded() *ConsentAdded {
	if x, ok := x.GetType().(*ConsentChangeEvent_Added); ok {
		return x.Added
	}
	return nil
}

func (x *ConsentChangeEvent) GetUpdated() *ConsentUpdated {
	if x, ok := x.GetType().(*ConsentChangeEvent_Updated); ok {
		return x.Updated
	}
	return nil
}

func (x *ConsentChangeEvent) GetRevoked() *ConsentRevoked {
	if x, ok := x.GetType().(*ConsentChangeEvent_Revoked); ok {
		return x.Revoked
	}
	return nil
}

func (m *ConsentChangeEvent) GetTarget() isConsentChangeEvent_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *ConsentChangeEvent) GetNumber() *v11.E164 {
	if x, ok := x.GetTarget().(*ConsentChangeEvent_Number); ok {
		return x.Number
	}
	return nil
}

func (x *ConsentChangeEvent) GetTenant() string {
	if x, ok := x.GetTarget().(*ConsentChangeEvent_Tenant); ok {
		return x.Tenant
	}
	return ""
}

type isConsentChangeEvent_Type interface {
	isConsentChangeEvent_Type()
}

type ConsentChangeEvent_Added struct {
	Added *ConsentAdded `protobuf:"bytes,1,opt,name=added,proto3,oneof"`
}

type ConsentChangeEvent_Updated struct {
	Updated *ConsentUpdated `protobuf:"bytes,2,opt,name=updated,proto3,oneof"`
}

type ConsentChangeEvent_Revoked struct {
	Revoked *ConsentRevoked `protobuf:"bytes,3,opt,name=revoked,proto3,oneof"`
}

func (*ConsentChangeEvent_Added) isConsentChangeEvent_Type() {}

func (*ConsentChangeEvent_Updated) isConsentChangeEvent_Type() {}

func (*ConsentChangeEvent_Revoked) isConsentChangeEvent_Type() {}

type isConsentChangeEvent_Target interface {
	isConsentChangeEvent_Target()
}

type ConsentChangeEvent_Number struct {
	// The international number of the subscriber.
	Number *v11.E164 `protobuf:"bytes,4,opt,name=number,proto3,oneof"`
}

type ConsentChangeEvent_Tenant struct {
	Tenant string `protobuf:"bytes,5,opt,name=tenant,proto3,oneof"`
}

func (*ConsentChangeEvent_Number) isConsentChangeEvent_Target() {}

func (*ConsentChangeEvent_Tenant) isConsentChangeEvent_Target() {}

type ConsentAdded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *ConsentAdded) Reset() {
	*x = ConsentAdded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsentAdded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsentAdded) ProtoMessage() {}

func (x *ConsentAdded) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsentAdded.ProtoReflect.Descriptor instead.
func (*ConsentAdded) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{5}
}

func (x *ConsentAdded) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type ConsentUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *ConsentUpdated) Reset() {
	*x = ConsentUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsentUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsentUpdated) ProtoMessage() {}

func (x *ConsentUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsentUpdated.ProtoReflect.Descriptor instead.
func (*ConsentUpdated) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{6}
}

func (x *ConsentUpdated) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type ConsentRevoked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConsentRevoked) Reset() {
	*x = ConsentRevoked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsentRevoked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsentRevoked) ProtoMessage() {}

func (x *ConsentRevoked) ProtoReflect() protoreflect.Message {
	mi := &file_wgtwo_consents_v1_consent_events_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsentRevoked.ProtoReflect.Descriptor instead.
func (*ConsentRevoked) Descriptor() ([]byte, []int) {
	return file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP(), []int{7}
}

var File_wgtwo_consents_v1_consent_events_proto protoreflect.FileDescriptor

var file_wgtwo_consents_v1_consent_events_proto_rawDesc = []byte{
	0x0a, 0x26, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x20, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x14, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb3,
	0x01, 0x0a, 0x21, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x14, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77, 0x67, 0x74, 0x77,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x1c, 0x41, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5a, 0x0a, 0x1d, 0x41, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x63,
	0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x61, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x05,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x67,
	0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x31, 0x36, 0x34, 0x48, 0x01, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42, 0x06,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x22, 0x26, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x64, 0x32, 0xa9, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a,
	0x19, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x77, 0x67, 0x74,
	0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x34, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0xea, 0xb5, 0x18, 0x00, 0x30, 0x01, 0x12, 0x80, 0x01,
	0x0a, 0x15, 0x41, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x6b, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x6b,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0xea, 0xb5, 0x18, 0x00,
	0x42, 0x68, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x42, 0x12, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x74, 0x77, 0x6f, 0x2f, 0x77,
	0x67, 0x74, 0x77, 0x6f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x77, 0x67, 0x74, 0x77, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wgtwo_consents_v1_consent_events_proto_rawDescOnce sync.Once
	file_wgtwo_consents_v1_consent_events_proto_rawDescData = file_wgtwo_consents_v1_consent_events_proto_rawDesc
)

func file_wgtwo_consents_v1_consent_events_proto_rawDescGZIP() []byte {
	file_wgtwo_consents_v1_consent_events_proto_rawDescOnce.Do(func() {
		file_wgtwo_consents_v1_consent_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_wgtwo_consents_v1_consent_events_proto_rawDescData)
	})
	return file_wgtwo_consents_v1_consent_events_proto_rawDescData
}

var file_wgtwo_consents_v1_consent_events_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wgtwo_consents_v1_consent_events_proto_goTypes = []interface{}{
	(*StreamConsentChangeEventsRequest)(nil),  // 0: wgtwo.consents.v1.StreamConsentChangeEventsRequest
	(*StreamConsentChangeEventsResponse)(nil), // 1: wgtwo.consents.v1.StreamConsentChangeEventsResponse
	(*AckConsentChangeEventRequest)(nil),      // 2: wgtwo.consents.v1.AckConsentChangeEventRequest
	(*AckConsentChangeEventResponse)(nil),     // 3: wgtwo.consents.v1.AckConsentChangeEventResponse
	(*ConsentChangeEvent)(nil),                // 4: wgtwo.consents.v1.ConsentChangeEvent
	(*ConsentAdded)(nil),                      // 5: wgtwo.consents.v1.ConsentAdded
	(*ConsentUpdated)(nil),                    // 6: wgtwo.consents.v1.ConsentUpdated
	(*ConsentRevoked)(nil),                    // 7: wgtwo.consents.v1.ConsentRevoked
	(*v1.StreamConfiguration)(nil),            // 8: wgtwo.events.v1.StreamConfiguration
	(*v1.Metadata)(nil),                       // 9: wgtwo.events.v1.Metadata
	(*v1.AckInfo)(nil),                        // 10: wgtwo.events.v1.AckInfo
	(*v1.AckStatus)(nil),                      // 11: wgtwo.events.v1.AckStatus
	(*v11.E164)(nil),                          // 12: wgtwo.common.v1.E164
}
var file_wgtwo_consents_v1_consent_events_proto_depIdxs = []int32{
	8,  // 0: wgtwo.consents.v1.StreamConsentChangeEventsRequest.stream_configuration:type_name -> wgtwo.events.v1.StreamConfiguration
	9,  // 1: wgtwo.consents.v1.StreamConsentChangeEventsResponse.metadata:type_name -> wgtwo.events.v1.Metadata
	4,  // 2: wgtwo.consents.v1.StreamConsentChangeEventsResponse.consent_change_event:type_name -> wgtwo.consents.v1.ConsentChangeEvent
	10, // 3: wgtwo.consents.v1.AckConsentChangeEventRequest.ack_info:type_name -> wgtwo.events.v1.AckInfo
	11, // 4: wgtwo.consents.v1.AckConsentChangeEventResponse.ack_status:type_name -> wgtwo.events.v1.AckStatus
	5,  // 5: wgtwo.consents.v1.ConsentChangeEvent.added:type_name -> wgtwo.consents.v1.ConsentAdded
	6,  // 6: wgtwo.consents.v1.ConsentChangeEvent.updated:type_name -> wgtwo.consents.v1.ConsentUpdated
	7,  // 7: wgtwo.consents.v1.ConsentChangeEvent.revoked:type_name -> wgtwo.consents.v1.ConsentRevoked
	12, // 8: wgtwo.consents.v1.ConsentChangeEvent.number:type_name -> wgtwo.common.v1.E164
	0,  // 9: wgtwo.consents.v1.ConsentEventService.StreamConsentChangeEvents:input_type -> wgtwo.consents.v1.StreamConsentChangeEventsRequest
	2,  // 10: wgtwo.consents.v1.ConsentEventService.AckConsentChangeEvent:input_type -> wgtwo.consents.v1.AckConsentChangeEventRequest
	1,  // 11: wgtwo.consents.v1.ConsentEventService.StreamConsentChangeEvents:output_type -> wgtwo.consents.v1.StreamConsentChangeEventsResponse
	3,  // 12: wgtwo.consents.v1.ConsentEventService.AckConsentChangeEvent:output_type -> wgtwo.consents.v1.AckConsentChangeEventResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_wgtwo_consents_v1_consent_events_proto_init() }
func file_wgtwo_consents_v1_consent_events_proto_init() {
	if File_wgtwo_consents_v1_consent_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamConsentChangeEventsRequest); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamConsentChangeEventsResponse); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckConsentChangeEventRequest); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckConsentChangeEventResponse); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsentChangeEvent); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsentAdded); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsentUpdated); i {
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
		file_wgtwo_consents_v1_consent_events_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsentRevoked); i {
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
	file_wgtwo_consents_v1_consent_events_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ConsentChangeEvent_Added)(nil),
		(*ConsentChangeEvent_Updated)(nil),
		(*ConsentChangeEvent_Revoked)(nil),
		(*ConsentChangeEvent_Number)(nil),
		(*ConsentChangeEvent_Tenant)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wgtwo_consents_v1_consent_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wgtwo_consents_v1_consent_events_proto_goTypes,
		DependencyIndexes: file_wgtwo_consents_v1_consent_events_proto_depIdxs,
		MessageInfos:      file_wgtwo_consents_v1_consent_events_proto_msgTypes,
	}.Build()
	File_wgtwo_consents_v1_consent_events_proto = out.File
	file_wgtwo_consents_v1_consent_events_proto_rawDesc = nil
	file_wgtwo_consents_v1_consent_events_proto_goTypes = nil
	file_wgtwo_consents_v1_consent_events_proto_depIdxs = nil
}
