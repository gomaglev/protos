// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/proto/common/common.proto

package common

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrderBy int32

const (
	OrderBy_UNKNOWN OrderBy = 0
	OrderBy_ASC     OrderBy = 1
	OrderBy_DESC    OrderBy = 2
)

// Enum value maps for OrderBy.
var (
	OrderBy_name = map[int32]string{
		0: "UNKNOWN",
		1: "ASC",
		2: "DESC",
	}
	OrderBy_value = map[string]int32{
		"UNKNOWN": 0,
		"ASC":     1,
		"DESC":    2,
	}
)

func (x OrderBy) Enum() *OrderBy {
	p := new(OrderBy)
	*p = x
	return p
}

func (x OrderBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderBy) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_common_common_proto_enumTypes[0].Descriptor()
}

func (OrderBy) Type() protoreflect.EnumType {
	return &file_pkg_proto_common_common_proto_enumTypes[0]
}

func (x OrderBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderBy.Descriptor instead.
func (OrderBy) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{0}
}

type PaginationParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination bool   `protobuf:"varint,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	CountOnly  bool   `protobuf:"varint,2,opt,name=count_only,json=countOnly,proto3" json:"count_only,omitempty"`
	Page       int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Cursor     string `protobuf:"bytes,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *PaginationParam) Reset() {
	*x = PaginationParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationParam) ProtoMessage() {}

func (x *PaginationParam) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationParam.ProtoReflect.Descriptor instead.
func (*PaginationParam) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationParam) GetPagination() bool {
	if x != nil {
		return x.Pagination
	}
	return false
}

func (x *PaginationParam) GetCountOnly() bool {
	if x != nil {
		return x.CountOnly
	}
	return false
}

func (x *PaginationParam) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationParam) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationParam) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type PaginationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Cursor   string `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *PaginationResult) Reset() {
	*x = PaginationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResult) ProtoMessage() {}

func (x *PaginationResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResult.ProtoReflect.Descriptor instead.
func (*PaginationResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationResult) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PaginationResult) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationResult) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationResult) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type StatusResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusResult) Reset() {
	*x = StatusResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResult) ProtoMessage() {}

func (x *StatusResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResult.ProtoReflect.Descriptor instead.
func (*StatusResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type IDResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IDResult) Reset() {
	*x = IDResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDResult) ProtoMessage() {}

func (x *IDResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDResult.ProtoReflect.Descriptor instead.
func (*IDResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *IDResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ErrorItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorItem) Reset() {
	*x = ErrorItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorItem) ProtoMessage() {}

func (x *ErrorItem) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorItem.ProtoReflect.Descriptor instead.
func (*ErrorItem) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorItem) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ErrorResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *ErrorItem `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorResult) Reset() {
	*x = ErrorResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResult) ProtoMessage() {}

func (x *ErrorResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResult.ProtoReflect.Descriptor instead.
func (*ErrorResult) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorResult) GetError() *ErrorItem {
	if x != nil {
		return x.Error
	}
	return nil
}

type OrderByField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string  `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Direction OrderBy `protobuf:"varint,2,opt,name=Direction,proto3,enum=pkg.proto.common.OrderBy" json:"Direction,omitempty"`
}

func (x *OrderByField) Reset() {
	*x = OrderByField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderByField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByField) ProtoMessage() {}

func (x *OrderByField) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByField.ProtoReflect.Descriptor instead.
func (*OrderByField) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *OrderByField) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OrderByField) GetDirection() OrderBy {
	if x != nil {
		return x.Direction
	}
	return OrderBy_UNKNOWN
}

type QueryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderByFields []*OrderByField `protobuf:"bytes,1,rep,name=OrderByFields,proto3" json:"OrderByFields,omitempty"`
	SelectFields  []string        `protobuf:"bytes,2,rep,name=SelectFields,proto3" json:"SelectFields,omitempty"`
}

func (x *QueryOptions) Reset() {
	*x = QueryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_common_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOptions) ProtoMessage() {}

func (x *QueryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_common_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOptions.ProtoReflect.Descriptor instead.
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return file_pkg_proto_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *QueryOptions) GetOrderByFields() []*OrderByField {
	if x != nil {
		return x.OrderByFields
	}
	return nil
}

func (x *QueryOptions) GetSelectFields() []string {
	if x != nil {
		return x.SelectFields
	}
	return nil
}

var File_pkg_proto_common_common_proto protoreflect.FileDescriptor

var file_pkg_proto_common_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0f, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x1a, 0x03, 0x18, 0x90, 0x4e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x1a, 0x03, 0x18, 0x90, 0x4e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x1a, 0x0a, 0x08, 0x49, 0x44, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a,
	0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x59, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x09,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x78, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2a,
	0x29, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x67, 0x6c, 0x65,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_common_common_proto_rawDescOnce sync.Once
	file_pkg_proto_common_common_proto_rawDescData = file_pkg_proto_common_common_proto_rawDesc
)

func file_pkg_proto_common_common_proto_rawDescGZIP() []byte {
	file_pkg_proto_common_common_proto_rawDescOnce.Do(func() {
		file_pkg_proto_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_common_common_proto_rawDescData)
	})
	return file_pkg_proto_common_common_proto_rawDescData
}

var file_pkg_proto_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_common_common_proto_goTypes = []interface{}{
	(OrderBy)(0),             // 0: pkg.proto.common.OrderBy
	(*PaginationParam)(nil),  // 1: pkg.proto.common.PaginationParam
	(*PaginationResult)(nil), // 2: pkg.proto.common.PaginationResult
	(*StatusResult)(nil),     // 3: pkg.proto.common.StatusResult
	(*IDResult)(nil),         // 4: pkg.proto.common.IDResult
	(*ErrorItem)(nil),        // 5: pkg.proto.common.ErrorItem
	(*ErrorResult)(nil),      // 6: pkg.proto.common.ErrorResult
	(*OrderByField)(nil),     // 7: pkg.proto.common.OrderByField
	(*QueryOptions)(nil),     // 8: pkg.proto.common.QueryOptions
}
var file_pkg_proto_common_common_proto_depIdxs = []int32{
	5, // 0: pkg.proto.common.ErrorResult.error:type_name -> pkg.proto.common.ErrorItem
	0, // 1: pkg.proto.common.OrderByField.Direction:type_name -> pkg.proto.common.OrderBy
	7, // 2: pkg.proto.common.QueryOptions.OrderByFields:type_name -> pkg.proto.common.OrderByField
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_common_common_proto_init() }
func file_pkg_proto_common_common_proto_init() {
	if File_pkg_proto_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationParam); i {
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
		file_pkg_proto_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationResult); i {
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
		file_pkg_proto_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResult); i {
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
		file_pkg_proto_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDResult); i {
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
		file_pkg_proto_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorItem); i {
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
		file_pkg_proto_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResult); i {
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
		file_pkg_proto_common_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderByField); i {
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
		file_pkg_proto_common_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOptions); i {
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
			RawDescriptor: file_pkg_proto_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_common_common_proto_goTypes,
		DependencyIndexes: file_pkg_proto_common_common_proto_depIdxs,
		EnumInfos:         file_pkg_proto_common_common_proto_enumTypes,
		MessageInfos:      file_pkg_proto_common_common_proto_msgTypes,
	}.Build()
	File_pkg_proto_common_common_proto = out.File
	file_pkg_proto_common_common_proto_rawDesc = nil
	file_pkg_proto_common_common_proto_goTypes = nil
	file_pkg_proto_common_common_proto_depIdxs = nil
}
