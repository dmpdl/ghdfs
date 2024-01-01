//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
// These .proto interfaces are private and stable.
// Please see http://wiki.apache.org/hadoop/Compatibility
// for what changes are allowed for a *stable* .proto interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: HAServiceProtocol.proto

package hadoop_common

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

type HAServiceStateProto int32

const (
	HAServiceStateProto_INITIALIZING HAServiceStateProto = 0
	HAServiceStateProto_ACTIVE       HAServiceStateProto = 1
	HAServiceStateProto_STANDBY      HAServiceStateProto = 2
	HAServiceStateProto_OBSERVER     HAServiceStateProto = 3
)

// Enum value maps for HAServiceStateProto.
var (
	HAServiceStateProto_name = map[int32]string{
		0: "INITIALIZING",
		1: "ACTIVE",
		2: "STANDBY",
		3: "OBSERVER",
	}
	HAServiceStateProto_value = map[string]int32{
		"INITIALIZING": 0,
		"ACTIVE":       1,
		"STANDBY":      2,
		"OBSERVER":     3,
	}
)

func (x HAServiceStateProto) Enum() *HAServiceStateProto {
	p := new(HAServiceStateProto)
	*p = x
	return p
}

func (x HAServiceStateProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HAServiceStateProto) Descriptor() protoreflect.EnumDescriptor {
	return file_HAServiceProtocol_proto_enumTypes[0].Descriptor()
}

func (HAServiceStateProto) Type() protoreflect.EnumType {
	return &file_HAServiceProtocol_proto_enumTypes[0]
}

func (x HAServiceStateProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *HAServiceStateProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = HAServiceStateProto(num)
	return nil
}

// Deprecated: Use HAServiceStateProto.Descriptor instead.
func (HAServiceStateProto) EnumDescriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{0}
}

type HARequestSource int32

const (
	HARequestSource_REQUEST_BY_USER        HARequestSource = 0
	HARequestSource_REQUEST_BY_USER_FORCED HARequestSource = 1
	HARequestSource_REQUEST_BY_ZKFC        HARequestSource = 2
)

// Enum value maps for HARequestSource.
var (
	HARequestSource_name = map[int32]string{
		0: "REQUEST_BY_USER",
		1: "REQUEST_BY_USER_FORCED",
		2: "REQUEST_BY_ZKFC",
	}
	HARequestSource_value = map[string]int32{
		"REQUEST_BY_USER":        0,
		"REQUEST_BY_USER_FORCED": 1,
		"REQUEST_BY_ZKFC":        2,
	}
)

func (x HARequestSource) Enum() *HARequestSource {
	p := new(HARequestSource)
	*p = x
	return p
}

func (x HARequestSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HARequestSource) Descriptor() protoreflect.EnumDescriptor {
	return file_HAServiceProtocol_proto_enumTypes[1].Descriptor()
}

func (HARequestSource) Type() protoreflect.EnumType {
	return &file_HAServiceProtocol_proto_enumTypes[1]
}

func (x HARequestSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *HARequestSource) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = HARequestSource(num)
	return nil
}

// Deprecated: Use HARequestSource.Descriptor instead.
func (HARequestSource) EnumDescriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{1}
}

type HAStateChangeRequestInfoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqSource *HARequestSource `protobuf:"varint,1,req,name=reqSource,enum=hadoop.common.HARequestSource" json:"reqSource,omitempty"`
}

func (x *HAStateChangeRequestInfoProto) Reset() {
	*x = HAStateChangeRequestInfoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HAStateChangeRequestInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HAStateChangeRequestInfoProto) ProtoMessage() {}

func (x *HAStateChangeRequestInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HAStateChangeRequestInfoProto.ProtoReflect.Descriptor instead.
func (*HAStateChangeRequestInfoProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{0}
}

func (x *HAStateChangeRequestInfoProto) GetReqSource() HARequestSource {
	if x != nil && x.ReqSource != nil {
		return *x.ReqSource
	}
	return HARequestSource_REQUEST_BY_USER
}

//*
// void request
type MonitorHealthRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MonitorHealthRequestProto) Reset() {
	*x = MonitorHealthRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorHealthRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorHealthRequestProto) ProtoMessage() {}

func (x *MonitorHealthRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorHealthRequestProto.ProtoReflect.Descriptor instead.
func (*MonitorHealthRequestProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{1}
}

//*
// void response
type MonitorHealthResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MonitorHealthResponseProto) Reset() {
	*x = MonitorHealthResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorHealthResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorHealthResponseProto) ProtoMessage() {}

func (x *MonitorHealthResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorHealthResponseProto.ProtoReflect.Descriptor instead.
func (*MonitorHealthResponseProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{2}
}

//*
// void request
type TransitionToActiveRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqInfo *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
}

func (x *TransitionToActiveRequestProto) Reset() {
	*x = TransitionToActiveRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionToActiveRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionToActiveRequestProto) ProtoMessage() {}

func (x *TransitionToActiveRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionToActiveRequestProto.ProtoReflect.Descriptor instead.
func (*TransitionToActiveRequestProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{3}
}

func (x *TransitionToActiveRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if x != nil {
		return x.ReqInfo
	}
	return nil
}

//*
// void response
type TransitionToActiveResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransitionToActiveResponseProto) Reset() {
	*x = TransitionToActiveResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionToActiveResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionToActiveResponseProto) ProtoMessage() {}

func (x *TransitionToActiveResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionToActiveResponseProto.ProtoReflect.Descriptor instead.
func (*TransitionToActiveResponseProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{4}
}

//*
// void request
type TransitionToStandbyRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqInfo *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
}

func (x *TransitionToStandbyRequestProto) Reset() {
	*x = TransitionToStandbyRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionToStandbyRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionToStandbyRequestProto) ProtoMessage() {}

func (x *TransitionToStandbyRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionToStandbyRequestProto.ProtoReflect.Descriptor instead.
func (*TransitionToStandbyRequestProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{5}
}

func (x *TransitionToStandbyRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if x != nil {
		return x.ReqInfo
	}
	return nil
}

//*
// void response
type TransitionToStandbyResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransitionToStandbyResponseProto) Reset() {
	*x = TransitionToStandbyResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionToStandbyResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionToStandbyResponseProto) ProtoMessage() {}

func (x *TransitionToStandbyResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionToStandbyResponseProto.ProtoReflect.Descriptor instead.
func (*TransitionToStandbyResponseProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{6}
}

//*
// void request
type TransitionToObserverRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqInfo *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
}

func (x *TransitionToObserverRequestProto) Reset() {
	*x = TransitionToObserverRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionToObserverRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionToObserverRequestProto) ProtoMessage() {}

func (x *TransitionToObserverRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionToObserverRequestProto.ProtoReflect.Descriptor instead.
func (*TransitionToObserverRequestProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{7}
}

func (x *TransitionToObserverRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if x != nil {
		return x.ReqInfo
	}
	return nil
}

//*
// void response
type TransitionToObserverResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransitionToObserverResponseProto) Reset() {
	*x = TransitionToObserverResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransitionToObserverResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransitionToObserverResponseProto) ProtoMessage() {}

func (x *TransitionToObserverResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransitionToObserverResponseProto.ProtoReflect.Descriptor instead.
func (*TransitionToObserverResponseProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{8}
}

//*
// void request
type GetServiceStatusRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServiceStatusRequestProto) Reset() {
	*x = GetServiceStatusRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceStatusRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceStatusRequestProto) ProtoMessage() {}

func (x *GetServiceStatusRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceStatusRequestProto.ProtoReflect.Descriptor instead.
func (*GetServiceStatusRequestProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{9}
}

//*
// Returns the state of the service
type GetServiceStatusResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *HAServiceStateProto `protobuf:"varint,1,req,name=state,enum=hadoop.common.HAServiceStateProto" json:"state,omitempty"`
	// If state is STANDBY, indicate whether it is
	// ready to become active.
	ReadyToBecomeActive *bool `protobuf:"varint,2,opt,name=readyToBecomeActive" json:"readyToBecomeActive,omitempty"`
	// If not ready to become active, a textual explanation of why not
	NotReadyReason *string `protobuf:"bytes,3,opt,name=notReadyReason" json:"notReadyReason,omitempty"`
}

func (x *GetServiceStatusResponseProto) Reset() {
	*x = GetServiceStatusResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HAServiceProtocol_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceStatusResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceStatusResponseProto) ProtoMessage() {}

func (x *GetServiceStatusResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_HAServiceProtocol_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceStatusResponseProto.ProtoReflect.Descriptor instead.
func (*GetServiceStatusResponseProto) Descriptor() ([]byte, []int) {
	return file_HAServiceProtocol_proto_rawDescGZIP(), []int{10}
}

func (x *GetServiceStatusResponseProto) GetState() HAServiceStateProto {
	if x != nil && x.State != nil {
		return *x.State
	}
	return HAServiceStateProto_INITIALIZING
}

func (x *GetServiceStatusResponseProto) GetReadyToBecomeActive() bool {
	if x != nil && x.ReadyToBecomeActive != nil {
		return *x.ReadyToBecomeActive
	}
	return false
}

func (x *GetServiceStatusResponseProto) GetNotReadyReason() string {
	if x != nil && x.NotReadyReason != nil {
		return *x.NotReadyReason
	}
	return ""
}

var File_HAServiceProtocol_proto protoreflect.FileDescriptor

var file_HAServiceProtocol_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x1d, 0x48, 0x41, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x72, 0x65, 0x71,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x41, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x68, 0x0a, 0x1e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x41, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x21, 0x0a, 0x1f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x69, 0x0a, 0x1f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x41, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x0a, 0x20, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a,
	0x0a, 0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x46, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x41, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x07, 0x72, 0x65, 0x71, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x0a, 0x21, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb3, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x48, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x42, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x61, 0x64, 0x79, 0x54,
	0x6f, 0x42, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x6e, 0x6f, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x4e, 0x0a, 0x13, 0x48, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54,
	0x41, 0x4e, 0x44, 0x42, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x42, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x57, 0x0a, 0x0f, 0x48, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x16, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x5a, 0x4b, 0x46, 0x43, 0x10, 0x02, 0x32, 0xd7,
	0x04, 0x0a, 0x18, 0x48, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x28, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x73, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x2d, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x76, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x12, 0x2e, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x79,
	0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x6d, 0x0a, 0x10, 0x67, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x2e, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x79, 0x0a, 0x1a, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x17, 0x48, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x69,
	0x6e, 0x6d, 0x61, 0x72, 0x63, 0x2f, 0x68, 0x64, 0x66, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0xa0, 0x01, 0x01,
}

var (
	file_HAServiceProtocol_proto_rawDescOnce sync.Once
	file_HAServiceProtocol_proto_rawDescData = file_HAServiceProtocol_proto_rawDesc
)

func file_HAServiceProtocol_proto_rawDescGZIP() []byte {
	file_HAServiceProtocol_proto_rawDescOnce.Do(func() {
		file_HAServiceProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_HAServiceProtocol_proto_rawDescData)
	})
	return file_HAServiceProtocol_proto_rawDescData
}

var file_HAServiceProtocol_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_HAServiceProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_HAServiceProtocol_proto_goTypes = []interface{}{
	(HAServiceStateProto)(0),                  // 0: hadoop.common.HAServiceStateProto
	(HARequestSource)(0),                      // 1: hadoop.common.HARequestSource
	(*HAStateChangeRequestInfoProto)(nil),     // 2: hadoop.common.HAStateChangeRequestInfoProto
	(*MonitorHealthRequestProto)(nil),         // 3: hadoop.common.MonitorHealthRequestProto
	(*MonitorHealthResponseProto)(nil),        // 4: hadoop.common.MonitorHealthResponseProto
	(*TransitionToActiveRequestProto)(nil),    // 5: hadoop.common.TransitionToActiveRequestProto
	(*TransitionToActiveResponseProto)(nil),   // 6: hadoop.common.TransitionToActiveResponseProto
	(*TransitionToStandbyRequestProto)(nil),   // 7: hadoop.common.TransitionToStandbyRequestProto
	(*TransitionToStandbyResponseProto)(nil),  // 8: hadoop.common.TransitionToStandbyResponseProto
	(*TransitionToObserverRequestProto)(nil),  // 9: hadoop.common.TransitionToObserverRequestProto
	(*TransitionToObserverResponseProto)(nil), // 10: hadoop.common.TransitionToObserverResponseProto
	(*GetServiceStatusRequestProto)(nil),      // 11: hadoop.common.GetServiceStatusRequestProto
	(*GetServiceStatusResponseProto)(nil),     // 12: hadoop.common.GetServiceStatusResponseProto
}
var file_HAServiceProtocol_proto_depIdxs = []int32{
	1,  // 0: hadoop.common.HAStateChangeRequestInfoProto.reqSource:type_name -> hadoop.common.HARequestSource
	2,  // 1: hadoop.common.TransitionToActiveRequestProto.reqInfo:type_name -> hadoop.common.HAStateChangeRequestInfoProto
	2,  // 2: hadoop.common.TransitionToStandbyRequestProto.reqInfo:type_name -> hadoop.common.HAStateChangeRequestInfoProto
	2,  // 3: hadoop.common.TransitionToObserverRequestProto.reqInfo:type_name -> hadoop.common.HAStateChangeRequestInfoProto
	0,  // 4: hadoop.common.GetServiceStatusResponseProto.state:type_name -> hadoop.common.HAServiceStateProto
	3,  // 5: hadoop.common.HAServiceProtocolService.monitorHealth:input_type -> hadoop.common.MonitorHealthRequestProto
	5,  // 6: hadoop.common.HAServiceProtocolService.transitionToActive:input_type -> hadoop.common.TransitionToActiveRequestProto
	7,  // 7: hadoop.common.HAServiceProtocolService.transitionToStandby:input_type -> hadoop.common.TransitionToStandbyRequestProto
	9,  // 8: hadoop.common.HAServiceProtocolService.transitionToObserver:input_type -> hadoop.common.TransitionToObserverRequestProto
	11, // 9: hadoop.common.HAServiceProtocolService.getServiceStatus:input_type -> hadoop.common.GetServiceStatusRequestProto
	4,  // 10: hadoop.common.HAServiceProtocolService.monitorHealth:output_type -> hadoop.common.MonitorHealthResponseProto
	6,  // 11: hadoop.common.HAServiceProtocolService.transitionToActive:output_type -> hadoop.common.TransitionToActiveResponseProto
	8,  // 12: hadoop.common.HAServiceProtocolService.transitionToStandby:output_type -> hadoop.common.TransitionToStandbyResponseProto
	10, // 13: hadoop.common.HAServiceProtocolService.transitionToObserver:output_type -> hadoop.common.TransitionToObserverResponseProto
	12, // 14: hadoop.common.HAServiceProtocolService.getServiceStatus:output_type -> hadoop.common.GetServiceStatusResponseProto
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_HAServiceProtocol_proto_init() }
func file_HAServiceProtocol_proto_init() {
	if File_HAServiceProtocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HAServiceProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HAStateChangeRequestInfoProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorHealthRequestProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorHealthResponseProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionToActiveRequestProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionToActiveResponseProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionToStandbyRequestProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionToStandbyResponseProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionToObserverRequestProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransitionToObserverResponseProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceStatusRequestProto); i {
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
		file_HAServiceProtocol_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceStatusResponseProto); i {
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
			RawDescriptor: file_HAServiceProtocol_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_HAServiceProtocol_proto_goTypes,
		DependencyIndexes: file_HAServiceProtocol_proto_depIdxs,
		EnumInfos:         file_HAServiceProtocol_proto_enumTypes,
		MessageInfos:      file_HAServiceProtocol_proto_msgTypes,
	}.Build()
	File_HAServiceProtocol_proto = out.File
	file_HAServiceProtocol_proto_rawDesc = nil
	file_HAServiceProtocol_proto_goTypes = nil
	file_HAServiceProtocol_proto_depIdxs = nil
}