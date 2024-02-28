// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.25.2
// source: keyset.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Keyset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tink_keyset contains the serialized tink keyset this message wraps.
	TinkKeyset       *anypb.Any                        `protobuf:"bytes,1,opt,name=tink_keyset,json=tinkKeyset,proto3" json:"tink_keyset,omitempty"`
	KeyCreationDates map[uint32]*timestamppb.Timestamp `protobuf:"bytes,2,rep,name=key_creation_dates,json=keyCreationDates,proto3" json:"key_creation_dates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Keyset) Reset() {
	*x = Keyset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keyset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyset) ProtoMessage() {}

func (x *Keyset) ProtoReflect() protoreflect.Message {
	mi := &file_keyset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyset.ProtoReflect.Descriptor instead.
func (*Keyset) Descriptor() ([]byte, []int) {
	return file_keyset_proto_rawDescGZIP(), []int{0}
}

func (x *Keyset) GetTinkKeyset() *anypb.Any {
	if x != nil {
		return x.TinkKeyset
	}
	return nil
}

func (x *Keyset) GetKeyCreationDates() map[uint32]*timestamppb.Timestamp {
	if x != nil {
		return x.KeyCreationDates
	}
	return nil
}

var File_keyset_proto protoreflect.FileDescriptor

var file_keyset_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6c, 0x73, 0x74, 0x6f, 0x6c, 0x6c, 0x2e, 0x65, 0x6e, 0x76, 0x6b, 0x6d, 0x73, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x06, 0x4b, 0x65,
	0x79, 0x73, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x74, 0x69, 0x6e, 0x6b, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0a, 0x74, 0x69, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x12, 0x6b,
	0x65, 0x79, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x73, 0x74, 0x6f, 0x6c, 0x6c,
	0x2e, 0x65, 0x6e, 0x76, 0x6b, 0x6d, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x65, 0x74, 0x2e, 0x4b,
	0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x6b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x5f, 0x0a, 0x15, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x73, 0x74, 0x6f, 0x6c, 0x6c, 0x2f, 0x65, 0x6e, 0x76,
	0x6b, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_keyset_proto_rawDescOnce sync.Once
	file_keyset_proto_rawDescData = file_keyset_proto_rawDesc
)

func file_keyset_proto_rawDescGZIP() []byte {
	file_keyset_proto_rawDescOnce.Do(func() {
		file_keyset_proto_rawDescData = protoimpl.X.CompressGZIP(file_keyset_proto_rawDescData)
	})
	return file_keyset_proto_rawDescData
}

var file_keyset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_keyset_proto_goTypes = []interface{}{
	(*Keyset)(nil),                // 0: lstoll.envkms.Keyset
	nil,                           // 1: lstoll.envkms.Keyset.KeyCreationDatesEntry
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_keyset_proto_depIdxs = []int32{
	2, // 0: lstoll.envkms.Keyset.tink_keyset:type_name -> google.protobuf.Any
	1, // 1: lstoll.envkms.Keyset.key_creation_dates:type_name -> lstoll.envkms.Keyset.KeyCreationDatesEntry
	3, // 2: lstoll.envkms.Keyset.KeyCreationDatesEntry.value:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_keyset_proto_init() }
func file_keyset_proto_init() {
	if File_keyset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keyset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyset); i {
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
			RawDescriptor: file_keyset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keyset_proto_goTypes,
		DependencyIndexes: file_keyset_proto_depIdxs,
		MessageInfos:      file_keyset_proto_msgTypes,
	}.Build()
	File_keyset_proto = out.File
	file_keyset_proto_rawDesc = nil
	file_keyset_proto_goTypes = nil
	file_keyset_proto_depIdxs = nil
}