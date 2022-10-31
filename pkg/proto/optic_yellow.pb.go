// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: optic_yellow.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BookingResponse_Status int32

const (
	BookingResponse_PENDING    BookingResponse_Status = 0
	BookingResponse_SUCCESSFUL BookingResponse_Status = 1
	BookingResponse_FAILED     BookingResponse_Status = 2
	BookingResponse_CANCELLED  BookingResponse_Status = 3
)

// Enum value maps for BookingResponse_Status.
var (
	BookingResponse_Status_name = map[int32]string{
		0: "PENDING",
		1: "SUCCESSFUL",
		2: "FAILED",
		3: "CANCELLED",
	}
	BookingResponse_Status_value = map[string]int32{
		"PENDING":    0,
		"SUCCESSFUL": 1,
		"FAILED":     2,
		"CANCELLED":  3,
	}
)

func (x BookingResponse_Status) Enum() *BookingResponse_Status {
	p := new(BookingResponse_Status)
	*p = x
	return p
}

func (x BookingResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookingResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_optic_yellow_proto_enumTypes[0].Descriptor()
}

func (BookingResponse_Status) Type() protoreflect.EnumType {
	return &file_optic_yellow_proto_enumTypes[0]
}

func (x BookingResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookingResponse_Status.Descriptor instead.
func (BookingResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_optic_yellow_proto_rawDescGZIP(), []int{4, 0}
}

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*RunRequest_Login
	//	*RunRequest_List
	//	*RunRequest_Booking
	//	*RunRequest_Cancel
	Request isRunRequest_Request `protobuf_oneof:"Request"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optic_yellow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_optic_yellow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_optic_yellow_proto_rawDescGZIP(), []int{0}
}

func (m *RunRequest) GetRequest() isRunRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *RunRequest) GetLogin() *Login {
	if x, ok := x.GetRequest().(*RunRequest_Login); ok {
		return x.Login
	}
	return nil
}

func (x *RunRequest) GetList() *Login {
	if x, ok := x.GetRequest().(*RunRequest_List); ok {
		return x.List
	}
	return nil
}

func (x *RunRequest) GetBooking() *BookingRequest {
	if x, ok := x.GetRequest().(*RunRequest_Booking); ok {
		return x.Booking
	}
	return nil
}

func (x *RunRequest) GetCancel() *BookingRequest {
	if x, ok := x.GetRequest().(*RunRequest_Cancel); ok {
		return x.Cancel
	}
	return nil
}

type isRunRequest_Request interface {
	isRunRequest_Request()
}

type RunRequest_Login struct {
	Login *Login `protobuf:"bytes,1,opt,name=login,proto3,oneof"`
}

type RunRequest_List struct {
	List *Login `protobuf:"bytes,2,opt,name=list,proto3,oneof"`
}

type RunRequest_Booking struct {
	Booking *BookingRequest `protobuf:"bytes,3,opt,name=booking,proto3,oneof"`
}

type RunRequest_Cancel struct {
	Cancel *BookingRequest `protobuf:"bytes,4,opt,name=cancel,proto3,oneof"`
}

func (*RunRequest_Login) isRunRequest_Request() {}

func (*RunRequest_List) isRunRequest_Request() {}

func (*RunRequest_Booking) isRunRequest_Request() {}

func (*RunRequest_Cancel) isRunRequest_Request() {}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Booking []*BookingResponse `protobuf:"bytes,1,rep,name=booking,proto3" json:"booking,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optic_yellow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_optic_yellow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_optic_yellow_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetBooking() []*BookingResponse {
	if x != nil {
		return x.Booking
	}
	return nil
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optic_yellow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_optic_yellow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_optic_yellow_proto_rawDescGZIP(), []int{2}
}

func (x *Login) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Login) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    *Login                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Datetime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optic_yellow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_optic_yellow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_optic_yellow_proto_rawDescGZIP(), []int{3}
}

func (x *BookingRequest) GetLogin() *Login {
	if x != nil {
		return x.Login
	}
	return nil
}

func (x *BookingRequest) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   BookingResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=optic_yellow.BookingResponse_Status" json:"status,omitempty"`
	Datetime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_optic_yellow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_optic_yellow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_optic_yellow_proto_rawDescGZIP(), []int{4}
}

func (x *BookingResponse) GetStatus() BookingResponse_Status {
	if x != nil {
		return x.Status
	}
	return BookingResponse_PENDING
}

func (x *BookingResponse) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

var File_optic_yellow_proto protoreflect.FileDescriptor

var file_optic_yellow_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x09, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f,
	0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22,
	0x3f, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x73, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x36, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x40, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x46, 0x55, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10,
	0x03, 0x32, 0x50, 0x0a, 0x12, 0x4f, 0x70, 0x74, 0x69, 0x63, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x18,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x63, 0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x63,
	0x5f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x6e, 0x2d, 0x62, 0x61, 0x6c, 0x6c, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x63, 0x2d, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_optic_yellow_proto_rawDescOnce sync.Once
	file_optic_yellow_proto_rawDescData = file_optic_yellow_proto_rawDesc
)

func file_optic_yellow_proto_rawDescGZIP() []byte {
	file_optic_yellow_proto_rawDescOnce.Do(func() {
		file_optic_yellow_proto_rawDescData = protoimpl.X.CompressGZIP(file_optic_yellow_proto_rawDescData)
	})
	return file_optic_yellow_proto_rawDescData
}

var file_optic_yellow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_optic_yellow_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_optic_yellow_proto_goTypes = []interface{}{
	(BookingResponse_Status)(0),   // 0: optic_yellow.BookingResponse.Status
	(*RunRequest)(nil),            // 1: optic_yellow.RunRequest
	(*RunResponse)(nil),           // 2: optic_yellow.RunResponse
	(*Login)(nil),                 // 3: optic_yellow.Login
	(*BookingRequest)(nil),        // 4: optic_yellow.BookingRequest
	(*BookingResponse)(nil),       // 5: optic_yellow.BookingResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_optic_yellow_proto_depIdxs = []int32{
	3,  // 0: optic_yellow.RunRequest.login:type_name -> optic_yellow.Login
	3,  // 1: optic_yellow.RunRequest.list:type_name -> optic_yellow.Login
	4,  // 2: optic_yellow.RunRequest.booking:type_name -> optic_yellow.BookingRequest
	4,  // 3: optic_yellow.RunRequest.cancel:type_name -> optic_yellow.BookingRequest
	5,  // 4: optic_yellow.RunResponse.booking:type_name -> optic_yellow.BookingResponse
	3,  // 5: optic_yellow.BookingRequest.login:type_name -> optic_yellow.Login
	6,  // 6: optic_yellow.BookingRequest.datetime:type_name -> google.protobuf.Timestamp
	0,  // 7: optic_yellow.BookingResponse.status:type_name -> optic_yellow.BookingResponse.Status
	6,  // 8: optic_yellow.BookingResponse.datetime:type_name -> google.protobuf.Timestamp
	1,  // 9: optic_yellow.OpticYellowService.Run:input_type -> optic_yellow.RunRequest
	2,  // 10: optic_yellow.OpticYellowService.Run:output_type -> optic_yellow.RunResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_optic_yellow_proto_init() }
func file_optic_yellow_proto_init() {
	if File_optic_yellow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_optic_yellow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_optic_yellow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
		file_optic_yellow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
		file_optic_yellow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_optic_yellow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResponse); i {
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
	file_optic_yellow_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RunRequest_Login)(nil),
		(*RunRequest_List)(nil),
		(*RunRequest_Booking)(nil),
		(*RunRequest_Cancel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_optic_yellow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_optic_yellow_proto_goTypes,
		DependencyIndexes: file_optic_yellow_proto_depIdxs,
		EnumInfos:         file_optic_yellow_proto_enumTypes,
		MessageInfos:      file_optic_yellow_proto_msgTypes,
	}.Build()
	File_optic_yellow_proto = out.File
	file_optic_yellow_proto_rawDesc = nil
	file_optic_yellow_proto_goTypes = nil
	file_optic_yellow_proto_depIdxs = nil
}