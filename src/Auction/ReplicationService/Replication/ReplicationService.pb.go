// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: ReplicationService.proto

package Replication

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

type NewBid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount    int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Bidderid  int64 `protobuf:"varint,2,opt,name=bidderid,proto3" json:"bidderid,omitempty"`
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *NewBid) Reset() {
	*x = NewBid{}
	mi := &file_ReplicationService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewBid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBid) ProtoMessage() {}

func (x *NewBid) ProtoReflect() protoreflect.Message {
	mi := &file_ReplicationService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBid.ProtoReflect.Descriptor instead.
func (*NewBid) Descriptor() ([]byte, []int) {
	return file_ReplicationService_proto_rawDescGZIP(), []int{0}
}

func (x *NewBid) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *NewBid) GetBidderid() int64 {
	if x != nil {
		return x.Bidderid
	}
	return 0
}

func (x *NewBid) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type NewLeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NewLeader) Reset() {
	*x = NewLeader{}
	mi := &file_ReplicationService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewLeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLeader) ProtoMessage() {}

func (x *NewLeader) ProtoReflect() protoreflect.Message {
	mi := &file_ReplicationService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLeader.ProtoReflect.Descriptor instead.
func (*NewLeader) Descriptor() ([]byte, []int) {
	return file_ReplicationService_proto_rawDescGZIP(), []int{1}
}

func (x *NewLeader) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack      bool  `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Bidderid int64 `protobuf:"varint,2,opt,name=bidderid,proto3" json:"bidderid,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_ReplicationService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_ReplicationService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_ReplicationService_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetAck() bool {
	if x != nil {
		return x.Ack
	}
	return false
}

func (x *Response) GetBidderid() int64 {
	if x != nil {
		return x.Bidderid
	}
	return 0
}

var File_ReplicationService_proto protoreflect.FileDescriptor

var file_ReplicationService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5a,
	0x0a, 0x06, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x25, 0x0a, 0x09, 0x4e, 0x65,
	0x77, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x69, 0x64, 0x32, 0xfb, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x64, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x64, 0x1a, 0x1c,
	0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x64, 0x1a, 0x1c, 0x2e, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ReplicationService_proto_rawDescOnce sync.Once
	file_ReplicationService_proto_rawDescData = file_ReplicationService_proto_rawDesc
)

func file_ReplicationService_proto_rawDescGZIP() []byte {
	file_ReplicationService_proto_rawDescOnce.Do(func() {
		file_ReplicationService_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReplicationService_proto_rawDescData)
	})
	return file_ReplicationService_proto_rawDescData
}

var file_ReplicationService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ReplicationService_proto_goTypes = []any{
	(*NewBid)(nil),    // 0: replicationservice.NewBid
	(*NewLeader)(nil), // 1: replicationservice.NewLeader
	(*Response)(nil),  // 2: replicationservice.Response
}
var file_ReplicationService_proto_depIdxs = []int32{
	0, // 0: replicationservice.ReplicationService.ReplicateBid:input_type -> replicationservice.NewBid
	1, // 1: replicationservice.ReplicationService.ConfirmLeader:input_type -> replicationservice.NewLeader
	0, // 2: replicationservice.ReplicationService.PropagateToLeader:input_type -> replicationservice.NewBid
	2, // 3: replicationservice.ReplicationService.ReplicateBid:output_type -> replicationservice.Response
	2, // 4: replicationservice.ReplicationService.ConfirmLeader:output_type -> replicationservice.Response
	2, // 5: replicationservice.ReplicationService.PropagateToLeader:output_type -> replicationservice.Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReplicationService_proto_init() }
func file_ReplicationService_proto_init() {
	if File_ReplicationService_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ReplicationService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ReplicationService_proto_goTypes,
		DependencyIndexes: file_ReplicationService_proto_depIdxs,
		MessageInfos:      file_ReplicationService_proto_msgTypes,
	}.Build()
	File_ReplicationService_proto = out.File
	file_ReplicationService_proto_rawDesc = nil
	file_ReplicationService_proto_goTypes = nil
	file_ReplicationService_proto_depIdxs = nil
}
