// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: raalgo.proto

package raalgo

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

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Id     int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	mi := &file_raalgo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_raalgo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_raalgo_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Amount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack bool `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_raalgo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_raalgo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_raalgo_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

var File_raalgo_proto protoreflect.FileDescriptor

var file_raalgo_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x61, 0x61, 0x6c, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x61, 0x61, 0x6c, 0x67, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x32, 0x5f,
	0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x03, 0x42, 0x69, 0x64,
	0x12, 0x0e, 0x2e, 0x72, 0x61, 0x61, 0x6c, 0x67, 0x6f, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x0b, 0x2e, 0x72, 0x61, 0x61, 0x6c, 0x67, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x30, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0e, 0x2e, 0x72, 0x61, 0x61, 0x6c, 0x67, 0x6f, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x61, 0x6c, 0x67, 0x6f, 0x2f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x61,
	0x6c, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raalgo_proto_rawDescOnce sync.Once
	file_raalgo_proto_rawDescData = file_raalgo_proto_rawDesc
)

func file_raalgo_proto_rawDescGZIP() []byte {
	file_raalgo_proto_rawDescOnce.Do(func() {
		file_raalgo_proto_rawDescData = protoimpl.X.CompressGZIP(file_raalgo_proto_rawDescData)
	})
	return file_raalgo_proto_rawDescData
}

var file_raalgo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_raalgo_proto_goTypes = []any{
	(*Amount)(nil),        // 0: raalgo.Amount
	(*Ack)(nil),           // 1: raalgo.Ack
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_raalgo_proto_depIdxs = []int32{
	0, // 0: raalgo.Auction.Bid:input_type -> raalgo.Amount
	2, // 1: raalgo.Auction.Result:input_type -> google.protobuf.Empty
	1, // 2: raalgo.Auction.Bid:output_type -> raalgo.Ack
	0, // 3: raalgo.Auction.Result:output_type -> raalgo.Amount
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_raalgo_proto_init() }
func file_raalgo_proto_init() {
	if File_raalgo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raalgo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raalgo_proto_goTypes,
		DependencyIndexes: file_raalgo_proto_depIdxs,
		MessageInfos:      file_raalgo_proto_msgTypes,
	}.Build()
	File_raalgo_proto = out.File
	file_raalgo_proto_rawDesc = nil
	file_raalgo_proto_goTypes = nil
	file_raalgo_proto_depIdxs = nil
}