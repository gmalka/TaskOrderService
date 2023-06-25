// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: build/proto/order_service.proto

package proto

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

type None struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *None) Reset() {
	*x = None{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{0}
}

type UsernameAndId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Id       int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UsernameAndId) Reset() {
	*x = UsernameAndId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsernameAndId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernameAndId) ProtoMessage() {}

func (x *UsernameAndId) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernameAndId.ProtoReflect.Descriptor instead.
func (*UsernameAndId) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{1}
}

func (x *UsernameAndId) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsernameAndId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskForUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Price int64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *TaskForUpdate) Reset() {
	*x = TaskForUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskForUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskForUpdate) ProtoMessage() {}

func (x *TaskForUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskForUpdate.ProtoReflect.Descriptor instead.
func (*TaskForUpdate) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{2}
}

func (x *TaskForUpdate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskForUpdate) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type OrderTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderTask) Reset() {
	*x = OrderTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderTask) ProtoMessage() {}

func (x *OrderTask) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderTask.ProtoReflect.Descriptor instead.
func (*OrderTask) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{3}
}

func (x *OrderTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{4}
}

func (x *UserId) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type TaskOrderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer int64 `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	Price  int64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *TaskOrderInfo) Reset() {
	*x = TaskOrderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskOrderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskOrderInfo) ProtoMessage() {}

func (x *TaskOrderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskOrderInfo.ProtoReflect.Descriptor instead.
func (*TaskOrderInfo) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{5}
}

func (x *TaskOrderInfo) GetAnswer() int64 {
	if x != nil {
		return x.Answer
	}
	return 0
}

func (x *TaskOrderInfo) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type TaskWithoutAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Count  int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Height []int64 `protobuf:"varint,3,rep,packed,name=height,proto3" json:"height,omitempty"`
	Price  int64   `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *TaskWithoutAnswer) Reset() {
	*x = TaskWithoutAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskWithoutAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskWithoutAnswer) ProtoMessage() {}

func (x *TaskWithoutAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskWithoutAnswer.ProtoReflect.Descriptor instead.
func (*TaskWithoutAnswer) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{6}
}

func (x *TaskWithoutAnswer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskWithoutAnswer) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *TaskWithoutAnswer) GetHeight() []int64 {
	if x != nil {
		return x.Height
	}
	return nil
}

func (x *TaskWithoutAnswer) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Count  int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Height []int64 `protobuf:"varint,3,rep,packed,name=height,proto3" json:"height,omitempty"`
	Price  int64   `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Answer int64   `protobuf:"varint,5,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{7}
}

func (x *Task) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Task) GetHeight() []int64 {
	if x != nil {
		return x.Height
	}
	return nil
}

func (x *Task) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Task) GetAnswer() int64 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type UserOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Page     int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *UserOrders) Reset() {
	*x = UserOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserOrders) ProtoMessage() {}

func (x *UserOrders) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserOrders.ProtoReflect.Descriptor instead.
func (*UserOrders) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{8}
}

func (x *UserOrders) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserOrders) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_build_proto_order_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_build_proto_order_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_build_proto_order_service_proto_rawDescGZIP(), []int{9}
}

func (x *Page) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

var File_build_proto_order_service_proto protoreflect.FileDescriptor

var file_build_proto_order_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65,
	0x22, 0x3b, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a,
	0x0d, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x24, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69,
	0x74, 0x68, 0x6f, 0x75, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x72, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x91, 0x04,
	0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x30, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75,
	0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x0f, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64,
	0x49, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x54,
	0x61, 0x73, 0x6b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x49, 0x64, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x0d,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x36, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4f, 0x66, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x12,
	0x31, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f,
	0x6e, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x61,
	0x73, 0x6b, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e, 0x65, 0x12,
	0x20, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x6e, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x6e,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_build_proto_order_service_proto_rawDescOnce sync.Once
	file_build_proto_order_service_proto_rawDescData = file_build_proto_order_service_proto_rawDesc
)

func file_build_proto_order_service_proto_rawDescGZIP() []byte {
	file_build_proto_order_service_proto_rawDescOnce.Do(func() {
		file_build_proto_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_build_proto_order_service_proto_rawDescData)
	})
	return file_build_proto_order_service_proto_rawDescData
}

var file_build_proto_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_build_proto_order_service_proto_goTypes = []interface{}{
	(*None)(nil),              // 0: proto.None
	(*UsernameAndId)(nil),     // 1: proto.UsernameAndId
	(*TaskForUpdate)(nil),     // 2: proto.TaskForUpdate
	(*OrderTask)(nil),         // 3: proto.OrderTask
	(*UserId)(nil),            // 4: proto.UserId
	(*TaskOrderInfo)(nil),     // 5: proto.TaskOrderInfo
	(*TaskWithoutAnswer)(nil), // 6: proto.TaskWithoutAnswer
	(*Task)(nil),              // 7: proto.Task
	(*UserOrders)(nil),        // 8: proto.UserOrders
	(*Page)(nil),              // 9: proto.Page
}
var file_build_proto_order_service_proto_depIdxs = []int32{
	8,  // 0: proto.TaskOrderService.getOrdersForUser:input_type -> proto.UserOrders
	0,  // 1: proto.TaskOrderService.getAllTasks:input_type -> proto.None
	9,  // 2: proto.TaskOrderService.getAllTasksWithoutAnswers:input_type -> proto.Page
	1,  // 3: proto.TaskOrderService.CheckAndGetTask:input_type -> proto.UsernameAndId
	1,  // 4: proto.TaskOrderService.buyTaskAnswer:input_type -> proto.UsernameAndId
	7,  // 5: proto.TaskOrderService.createNewTask:input_type -> proto.Task
	2,  // 6: proto.TaskOrderService.updatePriceOfTask:input_type -> proto.TaskForUpdate
	4,  // 7: proto.TaskOrderService.deleteOrdersForUser:input_type -> proto.UserId
	3,  // 8: proto.TaskOrderService.deleteTask:input_type -> proto.OrderTask
	0,  // 9: proto.TaskOrderService.ping:input_type -> proto.None
	7,  // 10: proto.TaskOrderService.getOrdersForUser:output_type -> proto.Task
	7,  // 11: proto.TaskOrderService.getAllTasks:output_type -> proto.Task
	6,  // 12: proto.TaskOrderService.getAllTasksWithoutAnswers:output_type -> proto.TaskWithoutAnswer
	5,  // 13: proto.TaskOrderService.CheckAndGetTask:output_type -> proto.TaskOrderInfo
	0,  // 14: proto.TaskOrderService.buyTaskAnswer:output_type -> proto.None
	0,  // 15: proto.TaskOrderService.createNewTask:output_type -> proto.None
	0,  // 16: proto.TaskOrderService.updatePriceOfTask:output_type -> proto.None
	0,  // 17: proto.TaskOrderService.deleteOrdersForUser:output_type -> proto.None
	0,  // 18: proto.TaskOrderService.deleteTask:output_type -> proto.None
	0,  // 19: proto.TaskOrderService.ping:output_type -> proto.None
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_build_proto_order_service_proto_init() }
func file_build_proto_order_service_proto_init() {
	if File_build_proto_order_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_build_proto_order_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*None); i {
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
		file_build_proto_order_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsernameAndId); i {
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
		file_build_proto_order_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskForUpdate); i {
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
		file_build_proto_order_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderTask); i {
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
		file_build_proto_order_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_build_proto_order_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskOrderInfo); i {
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
		file_build_proto_order_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskWithoutAnswer); i {
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
		file_build_proto_order_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_build_proto_order_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserOrders); i {
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
		file_build_proto_order_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_build_proto_order_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_build_proto_order_service_proto_goTypes,
		DependencyIndexes: file_build_proto_order_service_proto_depIdxs,
		MessageInfos:      file_build_proto_order_service_proto_msgTypes,
	}.Build()
	File_build_proto_order_service_proto = out.File
	file_build_proto_order_service_proto_rawDesc = nil
	file_build_proto_order_service_proto_goTypes = nil
	file_build_proto_order_service_proto_depIdxs = nil
}
