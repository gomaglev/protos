// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/proto/order/item/message/types.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/gomaglev/protos/v2/pkg/proto/common"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// OrderItemMessage
type OrderItemMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: faker:"uuid_hyphenated"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: faker:"paragraph"
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// @inject_tag: faker:"uuid_hyphenated"
	ItemId    string                 `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OrderItemMessage) Reset() {
	*x = OrderItemMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_order_item_message_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemMessage) ProtoMessage() {}

func (x *OrderItemMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_order_item_message_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemMessage.ProtoReflect.Descriptor instead.
func (*OrderItemMessage) Descriptor() ([]byte, []int) {
	return file_pkg_proto_order_item_message_types_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItemMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderItemMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderItemMessage) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *OrderItemMessage) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OrderItemMessage) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type OrderItemMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List       []*OrderItemMessage      `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pagination *common.PaginationResult `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *OrderItemMessages) Reset() {
	*x = OrderItemMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_order_item_message_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemMessages) ProtoMessage() {}

func (x *OrderItemMessages) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_order_item_message_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemMessages.ProtoReflect.Descriptor instead.
func (*OrderItemMessages) Descriptor() ([]byte, []int) {
	return file_pkg_proto_order_item_message_types_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemMessages) GetList() []*OrderItemMessage {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *OrderItemMessages) GetPagination() *common.PaginationResult {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_pkg_proto_order_item_message_types_proto protoreflect.FileDescriptor

var file_pkg_proto_order_item_message_types_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x42, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x67, 0x6c, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_proto_order_item_message_types_proto_rawDescOnce sync.Once
	file_pkg_proto_order_item_message_types_proto_rawDescData = file_pkg_proto_order_item_message_types_proto_rawDesc
)

func file_pkg_proto_order_item_message_types_proto_rawDescGZIP() []byte {
	file_pkg_proto_order_item_message_types_proto_rawDescOnce.Do(func() {
		file_pkg_proto_order_item_message_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_order_item_message_types_proto_rawDescData)
	})
	return file_pkg_proto_order_item_message_types_proto_rawDescData
}

var file_pkg_proto_order_item_message_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_order_item_message_types_proto_goTypes = []interface{}{
	(*OrderItemMessage)(nil),        // 0: pkg.proto.order.item.message.OrderItemMessage
	(*OrderItemMessages)(nil),       // 1: pkg.proto.order.item.message.OrderItemMessages
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
	(*common.PaginationResult)(nil), // 3: pkg.proto.common.PaginationResult
}
var file_pkg_proto_order_item_message_types_proto_depIdxs = []int32{
	2, // 0: pkg.proto.order.item.message.OrderItemMessage.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: pkg.proto.order.item.message.OrderItemMessage.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: pkg.proto.order.item.message.OrderItemMessages.list:type_name -> pkg.proto.order.item.message.OrderItemMessage
	3, // 3: pkg.proto.order.item.message.OrderItemMessages.pagination:type_name -> pkg.proto.common.PaginationResult
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_proto_order_item_message_types_proto_init() }
func file_pkg_proto_order_item_message_types_proto_init() {
	if File_pkg_proto_order_item_message_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_order_item_message_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemMessage); i {
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
		file_pkg_proto_order_item_message_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemMessages); i {
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
			RawDescriptor: file_pkg_proto_order_item_message_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_order_item_message_types_proto_goTypes,
		DependencyIndexes: file_pkg_proto_order_item_message_types_proto_depIdxs,
		MessageInfos:      file_pkg_proto_order_item_message_types_proto_msgTypes,
	}.Build()
	File_pkg_proto_order_item_message_types_proto = out.File
	file_pkg_proto_order_item_message_types_proto_rawDesc = nil
	file_pkg_proto_order_item_message_types_proto_goTypes = nil
	file_pkg_proto_order_item_message_types_proto_depIdxs = nil
}
