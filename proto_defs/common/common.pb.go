// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: proto_defs/common/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasicInfo *AssetRegistration `protobuf:"bytes,1,opt,name=basicInfo,proto3" json:"basicInfo,omitempty"`
	Task      *Task              `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Alive     bool               `protobuf:"varint,3,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	mi := &file_proto_defs_common_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetBasicInfo() *AssetRegistration {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

func (x *Session) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Session) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	mi := &file_proto_defs_common_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{1}
}

type AssetRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MagicNb   int32    `protobuf:"fixed32,1,opt,name=MagicNb,proto3" json:"MagicNb,omitempty"`
	SessionId string   `protobuf:"bytes,2,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	Hostname  string   `protobuf:"bytes,3,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	Username  string   `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	OS        string   `protobuf:"bytes,5,opt,name=OS,proto3" json:"OS,omitempty"`
	IP        []string `protobuf:"bytes,6,rep,name=IP,proto3" json:"IP,omitempty"`
}

func (x *AssetRegistration) Reset() {
	*x = AssetRegistration{}
	mi := &file_proto_defs_common_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetRegistration) ProtoMessage() {}

func (x *AssetRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetRegistration.ProtoReflect.Descriptor instead.
func (*AssetRegistration) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *AssetRegistration) GetMagicNb() int32 {
	if x != nil {
		return x.MagicNb
	}
	return 0
}

func (x *AssetRegistration) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *AssetRegistration) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *AssetRegistration) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AssetRegistration) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *AssetRegistration) GetIP() []string {
	if x != nil {
		return x.IP
	}
	return nil
}

type ServerOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string  `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	In        *TaskIO `protobuf:"bytes,2,opt,name=in,proto3" json:"in,omitempty"`
}

func (x *ServerOrder) Reset() {
	*x = ServerOrder{}
	mi := &file_proto_defs_common_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerOrder) ProtoMessage() {}

func (x *ServerOrder) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerOrder.ProtoReflect.Descriptor instead.
func (*ServerOrder) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *ServerOrder) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ServerOrder) GetIn() *TaskIO {
	if x != nil {
		return x.In
	}
	return nil
}

type RegistrationConfirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmed bool `protobuf:"varint,1,opt,name=Confirmed,proto3" json:"Confirmed,omitempty"`
}

func (x *RegistrationConfirmation) Reset() {
	*x = RegistrationConfirmation{}
	mi := &file_proto_defs_common_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrationConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationConfirmation) ProtoMessage() {}

func (x *RegistrationConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationConfirmation.ProtoReflect.Descriptor instead.
func (*RegistrationConfirmation) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *RegistrationConfirmation) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

type AssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string  `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	Out       *TaskIO `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *AssetResponse) Reset() {
	*x = AssetResponse{}
	mi := &file_proto_defs_common_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetResponse) ProtoMessage() {}

func (x *AssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetResponse.ProtoReflect.Descriptor instead.
func (*AssetResponse) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *AssetResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *AssetResponse) GetOut() *TaskIO {
	if x != nil {
		return x.Out
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	In  *TaskIO `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	Out *TaskIO `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_proto_defs_common_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *Task) GetIn() *TaskIO {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *Task) GetOut() *TaskIO {
	if x != nil {
		return x.Out
	}
	return nil
}

type TaskIO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text   string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	Binary []byte `protobuf:"bytes,2,opt,name=Binary,proto3" json:"Binary,omitempty"`
}

func (x *TaskIO) Reset() {
	*x = TaskIO{}
	mi := &file_proto_defs_common_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskIO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskIO) ProtoMessage() {}

func (x *TaskIO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_defs_common_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskIO.ProtoReflect.Descriptor instead.
func (*TaskIO) Descriptor() ([]byte, []int) {
	return file_proto_defs_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *TaskIO) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TaskIO) GetBinary() []byte {
	if x != nil {
		return x.Binary
	}
	return nil
}

var File_proto_defs_common_common_proto protoreflect.FileDescriptor

var file_proto_defs_common_common_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x66, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x67,
	0x69, 0x63, 0x4e, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x07, 0x4d, 0x61, 0x67, 0x69,
	0x63, 0x4e, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x53, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x53, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x22, 0x44, 0x0a, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f, 0x52, 0x02, 0x69, 0x6e, 0x22,
	0x38, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x48, 0x0a, 0x0d, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f, 0x52, 0x03,
	0x6f, 0x75, 0x74, 0x22, 0x3a, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f,
	0x52, 0x02, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x22,
	0x34, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x32, 0xaf, 0x01, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x12, 0x0e, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x1e, 0x5a, 0x1c, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x47, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x66, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_defs_common_common_proto_rawDescOnce sync.Once
	file_proto_defs_common_common_proto_rawDescData = file_proto_defs_common_common_proto_rawDesc
)

func file_proto_defs_common_common_proto_rawDescGZIP() []byte {
	file_proto_defs_common_common_proto_rawDescOnce.Do(func() {
		file_proto_defs_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_defs_common_common_proto_rawDescData)
	})
	return file_proto_defs_common_common_proto_rawDescData
}

var file_proto_defs_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_defs_common_common_proto_goTypes = []any{
	(*Session)(nil),                  // 0: Session
	(*None)(nil),                     // 1: None
	(*AssetRegistration)(nil),        // 2: AssetRegistration
	(*ServerOrder)(nil),              // 3: ServerOrder
	(*RegistrationConfirmation)(nil), // 4: RegistrationConfirmation
	(*AssetResponse)(nil),            // 5: AssetResponse
	(*Task)(nil),                     // 6: Task
	(*TaskIO)(nil),                   // 7: TaskIO
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_proto_defs_common_common_proto_depIdxs = []int32{
	2, // 0: Session.basicInfo:type_name -> AssetRegistration
	6, // 1: Session.task:type_name -> Task
	7, // 2: ServerOrder.in:type_name -> TaskIO
	7, // 3: AssetResponse.out:type_name -> TaskIO
	7, // 4: Task.in:type_name -> TaskIO
	7, // 5: Task.out:type_name -> TaskIO
	2, // 6: AssetService.RegisterAsset:input_type -> AssetRegistration
	5, // 7: AssetService.SendResponse:input_type -> AssetResponse
	5, // 8: AssetService.CheckIn:input_type -> AssetResponse
	4, // 9: AssetService.RegisterAsset:output_type -> RegistrationConfirmation
	8, // 10: AssetService.SendResponse:output_type -> google.protobuf.Empty
	3, // 11: AssetService.CheckIn:output_type -> ServerOrder
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_defs_common_common_proto_init() }
func file_proto_defs_common_common_proto_init() {
	if File_proto_defs_common_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_defs_common_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_defs_common_common_proto_goTypes,
		DependencyIndexes: file_proto_defs_common_common_proto_depIdxs,
		MessageInfos:      file_proto_defs_common_common_proto_msgTypes,
	}.Build()
	File_proto_defs_common_common_proto = out.File
	file_proto_defs_common_common_proto_rawDesc = nil
	file_proto_defs_common_common_proto_goTypes = nil
	file_proto_defs_common_common_proto_depIdxs = nil
}
