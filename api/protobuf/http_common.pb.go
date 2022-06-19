// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: http_common.proto

package protobuf

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

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value      string                 `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Path       string                 `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	Domain     string                 `protobuf:"bytes,4,opt,name=Domain,proto3" json:"Domain,omitempty"`
	Expires    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=Expires,proto3" json:"Expires,omitempty"`
	RawExpires string                 `protobuf:"bytes,6,opt,name=RawExpires,proto3" json:"RawExpires,omitempty"`
	MaxAge     int32                  `protobuf:"varint,7,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"`
	Secure     bool                   `protobuf:"varint,8,opt,name=Secure,proto3" json:"Secure,omitempty"`
	HttpOnly   bool                   `protobuf:"varint,9,opt,name=HttpOnly,proto3" json:"HttpOnly,omitempty"`
	SameSite   int32                  `protobuf:"varint,10,opt,name=SameSite,proto3" json:"SameSite,omitempty"`
	Raw        string                 `protobuf:"bytes,11,opt,name=Raw,proto3" json:"Raw,omitempty"`
	Unparsed   []string               `protobuf:"bytes,12,rep,name=Unparsed,proto3" json:"Unparsed,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_http_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_http_common_proto_rawDescGZIP(), []int{0}
}

func (x *Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cookie) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Cookie) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Cookie) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

func (x *Cookie) GetRawExpires() string {
	if x != nil {
		return x.RawExpires
	}
	return ""
}

func (x *Cookie) GetMaxAge() int32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Cookie) GetSecure() bool {
	if x != nil {
		return x.Secure
	}
	return false
}

func (x *Cookie) GetHttpOnly() bool {
	if x != nil {
		return x.HttpOnly
	}
	return false
}

func (x *Cookie) GetSameSite() int32 {
	if x != nil {
		return x.SameSite
	}
	return 0
}

func (x *Cookie) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *Cookie) GetUnparsed() []string {
	if x != nil {
		return x.Unparsed
	}
	return nil
}

var File_http_common_proto protoreflect.FileDescriptor

var file_http_common_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a, 0x06, 0x43, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x61, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x61, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x48, 0x74, 0x74, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x48, 0x74, 0x74, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x53, 0x61, 0x6d, 0x65, 0x53, 0x69, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x52, 0x61, 0x77, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x55, 0x6e,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x6c, 0x61, 0x78, 0x79, 0x2d, 0x52, 0x61, 0x69, 0x6c,
	0x77, 0x61, 0x79, 0x2f, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_http_common_proto_rawDescOnce sync.Once
	file_http_common_proto_rawDescData = file_http_common_proto_rawDesc
)

func file_http_common_proto_rawDescGZIP() []byte {
	file_http_common_proto_rawDescOnce.Do(func() {
		file_http_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_common_proto_rawDescData)
	})
	return file_http_common_proto_rawDescData
}

var file_http_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_http_common_proto_goTypes = []interface{}{
	(*Cookie)(nil),                // 0: GGargantua.v1.proto.Cookie
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_http_common_proto_depIdxs = []int32{
	1, // 0: GGargantua.v1.proto.Cookie.Expires:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_http_common_proto_init() }
func file_http_common_proto_init() {
	if File_http_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
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
			RawDescriptor: file_http_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_common_proto_goTypes,
		DependencyIndexes: file_http_common_proto_depIdxs,
		MessageInfos:      file_http_common_proto_msgTypes,
	}.Build()
	File_http_common_proto = out.File
	file_http_common_proto_rawDesc = nil
	file_http_common_proto_goTypes = nil
	file_http_common_proto_depIdxs = nil
}
