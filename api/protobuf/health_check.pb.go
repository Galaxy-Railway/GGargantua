// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: health_check.proto

package protobuf

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

type HealthCondition_HealthStatus int32

const (
	HealthCondition_HEALTHY   HealthCondition_HealthStatus = 0
	HealthCondition_UNHEALTHY HealthCondition_HealthStatus = 1
)

// Enum value maps for HealthCondition_HealthStatus.
var (
	HealthCondition_HealthStatus_name = map[int32]string{
		0: "HEALTHY",
		1: "UNHEALTHY",
	}
	HealthCondition_HealthStatus_value = map[string]int32{
		"HEALTHY":   0,
		"UNHEALTHY": 1,
	}
)

func (x HealthCondition_HealthStatus) Enum() *HealthCondition_HealthStatus {
	p := new(HealthCondition_HealthStatus)
	*p = x
	return p
}

func (x HealthCondition_HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCondition_HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_check_proto_enumTypes[0].Descriptor()
}

func (HealthCondition_HealthStatus) Type() protoreflect.EnumType {
	return &file_health_check_proto_enumTypes[0]
}

func (x HealthCondition_HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCondition_HealthStatus.Descriptor instead.
func (HealthCondition_HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_check_proto_rawDescGZIP(), []int{0, 0}
}

type HealthCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCondition_HealthStatus `protobuf:"varint,1,opt,name=status,proto3,enum=GGargantua.v1.proto.HealthCondition_HealthStatus" json:"status,omitempty"`
	Reason string                       `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Error  string                       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *HealthCondition) Reset() {
	*x = HealthCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCondition) ProtoMessage() {}

func (x *HealthCondition) ProtoReflect() protoreflect.Message {
	mi := &file_health_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCondition.ProtoReflect.Descriptor instead.
func (*HealthCondition) Descriptor() ([]byte, []int) {
	return file_health_check_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCondition) GetStatus() HealthCondition_HealthStatus {
	if x != nil {
		return x.Status
	}
	return HealthCondition_HEALTHY
}

func (x *HealthCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *HealthCondition) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_health_check_proto protoreflect.FileDescriptor

var file_health_check_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x47, 0x47, 0x61,
	0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x2a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x32,
	0x5c, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x12, 0x4b, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61,
	0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x61, 0x6c, 0x61,
	0x78, 0x79, 0x2d, 0x52, 0x61, 0x69, 0x6c, 0x77, 0x61, 0x79, 0x2f, 0x47, 0x47, 0x61, 0x72, 0x67,
	0x61, 0x6e, 0x74, 0x75, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_check_proto_rawDescOnce sync.Once
	file_health_check_proto_rawDescData = file_health_check_proto_rawDesc
)

func file_health_check_proto_rawDescGZIP() []byte {
	file_health_check_proto_rawDescOnce.Do(func() {
		file_health_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_check_proto_rawDescData)
	})
	return file_health_check_proto_rawDescData
}

var file_health_check_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_health_check_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_health_check_proto_goTypes = []interface{}{
	(HealthCondition_HealthStatus)(0), // 0: GGargantua.v1.proto.HealthCondition.HealthStatus
	(*HealthCondition)(nil),           // 1: GGargantua.v1.proto.HealthCondition
	(*emptypb.Empty)(nil),             // 2: google.protobuf.Empty
}
var file_health_check_proto_depIdxs = []int32{
	0, // 0: GGargantua.v1.proto.HealthCondition.status:type_name -> GGargantua.v1.proto.HealthCondition.HealthStatus
	2, // 1: GGargantua.v1.proto.HealthChecker.CheckHealth:input_type -> google.protobuf.Empty
	1, // 2: GGargantua.v1.proto.HealthChecker.CheckHealth:output_type -> GGargantua.v1.proto.HealthCondition
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_health_check_proto_init() }
func file_health_check_proto_init() {
	if File_health_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCondition); i {
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
			RawDescriptor: file_health_check_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_check_proto_goTypes,
		DependencyIndexes: file_health_check_proto_depIdxs,
		EnumInfos:         file_health_check_proto_enumTypes,
		MessageInfos:      file_health_check_proto_msgTypes,
	}.Build()
	File_health_check_proto = out.File
	file_health_check_proto_rawDesc = nil
	file_health_check_proto_goTypes = nil
	file_health_check_proto_depIdxs = nil
}
