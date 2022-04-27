// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/proto_files/single_request_sender.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SchemaType int32

const (
	SchemaType_HTTP  SchemaType = 0
	SchemaType_HTTPS SchemaType = 1
)

// Enum value maps for SchemaType.
var (
	SchemaType_name = map[int32]string{
		0: "HTTP",
		1: "HTTPS",
	}
	SchemaType_value = map[string]int32{
		"HTTP":  0,
		"HTTPS": 1,
	}
)

func (x SchemaType) Enum() *SchemaType {
	p := new(SchemaType)
	*p = x
	return p
}

func (x SchemaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SchemaType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_files_single_request_sender_proto_enumTypes[0].Descriptor()
}

func (SchemaType) Type() protoreflect.EnumType {
	return &file_api_proto_files_single_request_sender_proto_enumTypes[0]
}

func (x SchemaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SchemaType.Descriptor instead.
func (SchemaType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_files_single_request_sender_proto_rawDescGZIP(), []int{0}
}

type SingleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestSchema  SchemaType `protobuf:"varint,1,opt,name=RequestSchema,proto3,enum=GGargantua.v1.proto.SchemaType" json:"RequestSchema,omitempty"`
	RequestContent []byte     `protobuf:"bytes,2,opt,name=RequestContent,proto3" json:"RequestContent,omitempty"`
}

func (x *SingleRequest) Reset() {
	*x = SingleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_files_single_request_sender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleRequest) ProtoMessage() {}

func (x *SingleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_files_single_request_sender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleRequest.ProtoReflect.Descriptor instead.
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_files_single_request_sender_proto_rawDescGZIP(), []int{0}
}

func (x *SingleRequest) GetRequestSchema() SchemaType {
	if x != nil {
		return x.RequestSchema
	}
	return SchemaType_HTTP
}

func (x *SingleRequest) GetRequestContent() []byte {
	if x != nil {
		return x.RequestContent
	}
	return nil
}

type SingleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseSchema  SchemaType `protobuf:"varint,1,opt,name=ResponseSchema,proto3,enum=GGargantua.v1.proto.SchemaType" json:"ResponseSchema,omitempty"`
	ResponseContent []byte     `protobuf:"bytes,2,opt,name=ResponseContent,proto3" json:"ResponseContent,omitempty"`
}

func (x *SingleResponse) Reset() {
	*x = SingleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_files_single_request_sender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleResponse) ProtoMessage() {}

func (x *SingleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_files_single_request_sender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleResponse.ProtoReflect.Descriptor instead.
func (*SingleResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_files_single_request_sender_proto_rawDescGZIP(), []int{1}
}

func (x *SingleResponse) GetResponseSchema() SchemaType {
	if x != nil {
		return x.ResponseSchema
	}
	return SchemaType_HTTP
}

func (x *SingleResponse) GetResponseContent() []byte {
	if x != nil {
		return x.ResponseContent
	}
	return nil
}

var File_api_proto_files_single_request_sender_proto protoreflect.FileDescriptor

var file_api_proto_files_single_request_sender_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x47,
	0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0d, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x47, 0x47, 0x61,
	0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x28,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x21, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x53, 0x10, 0x01, 0x32, 0x73, 0x0a, 0x13, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x5c, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61,
	0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x47, 0x47,
	0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_files_single_request_sender_proto_rawDescOnce sync.Once
	file_api_proto_files_single_request_sender_proto_rawDescData = file_api_proto_files_single_request_sender_proto_rawDesc
)

func file_api_proto_files_single_request_sender_proto_rawDescGZIP() []byte {
	file_api_proto_files_single_request_sender_proto_rawDescOnce.Do(func() {
		file_api_proto_files_single_request_sender_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_files_single_request_sender_proto_rawDescData)
	})
	return file_api_proto_files_single_request_sender_proto_rawDescData
}

var file_api_proto_files_single_request_sender_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_files_single_request_sender_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_files_single_request_sender_proto_goTypes = []interface{}{
	(SchemaType)(0),        // 0: GGargantua.v1.proto.SchemaType
	(*SingleRequest)(nil),  // 1: GGargantua.v1.proto.SingleRequest
	(*SingleResponse)(nil), // 2: GGargantua.v1.proto.SingleResponse
}
var file_api_proto_files_single_request_sender_proto_depIdxs = []int32{
	0, // 0: GGargantua.v1.proto.SingleRequest.RequestSchema:type_name -> GGargantua.v1.proto.SchemaType
	0, // 1: GGargantua.v1.proto.SingleResponse.ResponseSchema:type_name -> GGargantua.v1.proto.SchemaType
	1, // 2: GGargantua.v1.proto.SingleRequestSender.SendSingleRequest:input_type -> GGargantua.v1.proto.SingleRequest
	2, // 3: GGargantua.v1.proto.SingleRequestSender.SendSingleRequest:output_type -> GGargantua.v1.proto.SingleResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_files_single_request_sender_proto_init() }
func file_api_proto_files_single_request_sender_proto_init() {
	if File_api_proto_files_single_request_sender_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_files_single_request_sender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleRequest); i {
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
		file_api_proto_files_single_request_sender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleResponse); i {
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
			RawDescriptor: file_api_proto_files_single_request_sender_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_files_single_request_sender_proto_goTypes,
		DependencyIndexes: file_api_proto_files_single_request_sender_proto_depIdxs,
		EnumInfos:         file_api_proto_files_single_request_sender_proto_enumTypes,
		MessageInfos:      file_api_proto_files_single_request_sender_proto_msgTypes,
	}.Build()
	File_api_proto_files_single_request_sender_proto = out.File
	file_api_proto_files_single_request_sender_proto_rawDesc = nil
	file_api_proto_files_single_request_sender_proto_goTypes = nil
	file_api_proto_files_single_request_sender_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SingleRequestSenderClient is the client API for SingleRequestSender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SingleRequestSenderClient interface {
	SendSingleRequest(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*SingleResponse, error)
}

type singleRequestSenderClient struct {
	cc grpc.ClientConnInterface
}

func NewSingleRequestSenderClient(cc grpc.ClientConnInterface) SingleRequestSenderClient {
	return &singleRequestSenderClient{cc}
}

func (c *singleRequestSenderClient) SendSingleRequest(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*SingleResponse, error) {
	out := new(SingleResponse)
	err := c.cc.Invoke(ctx, "/GGargantua.v1.proto.SingleRequestSender/SendSingleRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SingleRequestSenderServer is the server API for SingleRequestSender service.
type SingleRequestSenderServer interface {
	SendSingleRequest(context.Context, *SingleRequest) (*SingleResponse, error)
}

// UnimplementedSingleRequestSenderServer can be embedded to have forward compatible implementations.
type UnimplementedSingleRequestSenderServer struct {
}

func (*UnimplementedSingleRequestSenderServer) SendSingleRequest(context.Context, *SingleRequest) (*SingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSingleRequest not implemented")
}

func RegisterSingleRequestSenderServer(s *grpc.Server, srv SingleRequestSenderServer) {
	s.RegisterService(&_SingleRequestSender_serviceDesc, srv)
}

func _SingleRequestSender_SendSingleRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SingleRequestSenderServer).SendSingleRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GGargantua.v1.proto.SingleRequestSender/SendSingleRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SingleRequestSenderServer).SendSingleRequest(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SingleRequestSender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GGargantua.v1.proto.SingleRequestSender",
	HandlerType: (*SingleRequestSenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSingleRequest",
			Handler:    _SingleRequestSender_SendSingleRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto_files/single_request_sender.proto",
}