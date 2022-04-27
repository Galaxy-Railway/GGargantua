// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: api/proto_files/job.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type JobProgress int32

const (
	JobProgress_CREATED    JobProgress = 0
	JobProgress_PROCESSING JobProgress = 1
	JobProgress_FINISHED   JobProgress = 2
	JobProgress_FAILED     JobProgress = 3
	JobProgress_CANCELING  JobProgress = 4
	JobProgress_CANCELED   JobProgress = 5
)

// Enum value maps for JobProgress.
var (
	JobProgress_name = map[int32]string{
		0: "CREATED",
		1: "PROCESSING",
		2: "FINISHED",
		3: "FAILED",
		4: "CANCELING",
		5: "CANCELED",
	}
	JobProgress_value = map[string]int32{
		"CREATED":    0,
		"PROCESSING": 1,
		"FINISHED":   2,
		"FAILED":     3,
		"CANCELING":  4,
		"CANCELED":   5,
	}
)

func (x JobProgress) Enum() *JobProgress {
	p := new(JobProgress)
	*p = x
	return p
}

func (x JobProgress) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobProgress) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_files_job_proto_enumTypes[0].Descriptor()
}

func (JobProgress) Type() protoreflect.EnumType {
	return &file_api_proto_files_job_proto_enumTypes[0]
}

func (x JobProgress) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobProgress.Descriptor instead.
func (JobProgress) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_files_job_proto_rawDescGZIP(), []int{0}
}

type JobUuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
}

func (x *JobUuid) Reset() {
	*x = JobUuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_files_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobUuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobUuid) ProtoMessage() {}

func (x *JobUuid) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_files_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobUuid.ProtoReflect.Descriptor instead.
func (*JobUuid) Descriptor() ([]byte, []int) {
	return file_api_proto_files_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobUuid) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type JobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string      `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Progress JobProgress `protobuf:"varint,2,opt,name=Progress,proto3,enum=GGargantua.v1.proto.JobProgress" json:"Progress,omitempty"`
	Reason   string      `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty"`
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_files_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_files_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_files_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobStatus) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *JobStatus) GetProgress() JobProgress {
	if x != nil {
		return x.Progress
	}
	return JobProgress_CREATED
}

func (x *JobStatus) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type UpdateJobContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     *JobUuid `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	MainStep *Step    `protobuf:"bytes,2,opt,name=MainStep,proto3" json:"MainStep,omitempty"`
}

func (x *UpdateJobContent) Reset() {
	*x = UpdateJobContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_files_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobContent) ProtoMessage() {}

func (x *UpdateJobContent) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_files_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobContent.ProtoReflect.Descriptor instead.
func (*UpdateJobContent) Descriptor() ([]byte, []int) {
	return file_api_proto_files_job_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateJobContent) GetUuid() *JobUuid {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *UpdateJobContent) GetMainStep() *Step {
	if x != nil {
		return x.MainStep
	}
	return nil
}

type JobResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *JobStatus  `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Result *StepResult `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *JobResult) Reset() {
	*x = JobResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_files_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResult) ProtoMessage() {}

func (x *JobResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_files_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResult.ProtoReflect.Descriptor instead.
func (*JobResult) Descriptor() ([]byte, []int) {
	return file_api_proto_files_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobResult) GetStatus() *JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *JobResult) GetResult() *StepResult {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_proto_files_job_proto protoreflect.FileDescriptor

var file_api_proto_files_job_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x47, 0x47, 0x61,
	0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73,
	0x74, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x4a, 0x6f, 0x62,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x47, 0x47,
	0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x7b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x55, 0x75, 0x69, 0x64, 0x52,
	0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61,
	0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x65, 0x70, 0x52, 0x08, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x22, 0x7c, 0x0a, 0x09,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x47, 0x47, 0x61, 0x72,
	0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x37, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x61, 0x0a, 0x0b, 0x4a, 0x6f,
	0x62, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x32, 0xae, 0x02,
	0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x4a, 0x6f, 0x62, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x4a, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x4a, 0x6f, 0x62, 0x12, 0x25, 0x2e,
	0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0a,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x4a, 0x6f, 0x62, 0x12, 0x1c, 0x2e, 0x47, 0x47, 0x61,
	0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4a, 0x6f, 0x62, 0x55, 0x75, 0x69, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x4c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1c, 0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x55, 0x75, 0x69, 0x64, 0x1a, 0x1e,
	0x2e, 0x47, 0x47, 0x61, 0x72, 0x67, 0x61, 0x6e, 0x74, 0x75, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0e,
	0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_files_job_proto_rawDescOnce sync.Once
	file_api_proto_files_job_proto_rawDescData = file_api_proto_files_job_proto_rawDesc
)

func file_api_proto_files_job_proto_rawDescGZIP() []byte {
	file_api_proto_files_job_proto_rawDescOnce.Do(func() {
		file_api_proto_files_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_files_job_proto_rawDescData)
	})
	return file_api_proto_files_job_proto_rawDescData
}

var file_api_proto_files_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_files_job_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_files_job_proto_goTypes = []interface{}{
	(JobProgress)(0),         // 0: GGargantua.v1.proto.JobProgress
	(*JobUuid)(nil),          // 1: GGargantua.v1.proto.JobUuid
	(*JobStatus)(nil),        // 2: GGargantua.v1.proto.JobStatus
	(*UpdateJobContent)(nil), // 3: GGargantua.v1.proto.UpdateJobContent
	(*JobResult)(nil),        // 4: GGargantua.v1.proto.JobResult
	(*Step)(nil),             // 5: GGargantua.v1.proto.Step
	(*StepResult)(nil),       // 6: GGargantua.v1.proto.StepResult
	(*emptypb.Empty)(nil),    // 7: google.protobuf.Empty
}
var file_api_proto_files_job_proto_depIdxs = []int32{
	0, // 0: GGargantua.v1.proto.JobStatus.Progress:type_name -> GGargantua.v1.proto.JobProgress
	1, // 1: GGargantua.v1.proto.UpdateJobContent.Uuid:type_name -> GGargantua.v1.proto.JobUuid
	5, // 2: GGargantua.v1.proto.UpdateJobContent.MainStep:type_name -> GGargantua.v1.proto.Step
	2, // 3: GGargantua.v1.proto.JobResult.Status:type_name -> GGargantua.v1.proto.JobStatus
	6, // 4: GGargantua.v1.proto.JobResult.Result:type_name -> GGargantua.v1.proto.StepResult
	7, // 5: GGargantua.v1.proto.JobService.CreateAJob:input_type -> google.protobuf.Empty
	3, // 6: GGargantua.v1.proto.JobService.StartAJob:input_type -> GGargantua.v1.proto.UpdateJobContent
	1, // 7: GGargantua.v1.proto.JobService.CancelAJob:input_type -> GGargantua.v1.proto.JobUuid
	1, // 8: GGargantua.v1.proto.JobService.GetJobResult:input_type -> GGargantua.v1.proto.JobUuid
	1, // 9: GGargantua.v1.proto.JobService.CreateAJob:output_type -> GGargantua.v1.proto.JobUuid
	7, // 10: GGargantua.v1.proto.JobService.StartAJob:output_type -> google.protobuf.Empty
	7, // 11: GGargantua.v1.proto.JobService.CancelAJob:output_type -> google.protobuf.Empty
	4, // 12: GGargantua.v1.proto.JobService.GetJobResult:output_type -> GGargantua.v1.proto.JobResult
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_files_job_proto_init() }
func file_api_proto_files_job_proto_init() {
	if File_api_proto_files_job_proto != nil {
		return
	}
	file_api_proto_files_step_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_files_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobUuid); i {
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
		file_api_proto_files_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatus); i {
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
		file_api_proto_files_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobContent); i {
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
		file_api_proto_files_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobResult); i {
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
			RawDescriptor: file_api_proto_files_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_files_job_proto_goTypes,
		DependencyIndexes: file_api_proto_files_job_proto_depIdxs,
		EnumInfos:         file_api_proto_files_job_proto_enumTypes,
		MessageInfos:      file_api_proto_files_job_proto_msgTypes,
	}.Build()
	File_api_proto_files_job_proto = out.File
	file_api_proto_files_job_proto_rawDesc = nil
	file_api_proto_files_job_proto_goTypes = nil
	file_api_proto_files_job_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	CreateAJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*JobUuid, error)
	StartAJob(ctx context.Context, in *UpdateJobContent, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelAJob(ctx context.Context, in *JobUuid, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetJobResult(ctx context.Context, in *JobUuid, opts ...grpc.CallOption) (*JobResult, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) CreateAJob(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*JobUuid, error) {
	out := new(JobUuid)
	err := c.cc.Invoke(ctx, "/GGargantua.v1.proto.JobService/CreateAJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) StartAJob(ctx context.Context, in *UpdateJobContent, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/GGargantua.v1.proto.JobService/StartAJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) CancelAJob(ctx context.Context, in *JobUuid, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/GGargantua.v1.proto.JobService/CancelAJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetJobResult(ctx context.Context, in *JobUuid, opts ...grpc.CallOption) (*JobResult, error) {
	out := new(JobResult)
	err := c.cc.Invoke(ctx, "/GGargantua.v1.proto.JobService/GetJobResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	CreateAJob(context.Context, *emptypb.Empty) (*JobUuid, error)
	StartAJob(context.Context, *UpdateJobContent) (*emptypb.Empty, error)
	CancelAJob(context.Context, *JobUuid) (*emptypb.Empty, error)
	GetJobResult(context.Context, *JobUuid) (*JobResult, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) CreateAJob(context.Context, *emptypb.Empty) (*JobUuid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAJob not implemented")
}
func (*UnimplementedJobServiceServer) StartAJob(context.Context, *UpdateJobContent) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAJob not implemented")
}
func (*UnimplementedJobServiceServer) CancelAJob(context.Context, *JobUuid) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAJob not implemented")
}
func (*UnimplementedJobServiceServer) GetJobResult(context.Context, *JobUuid) (*JobResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobResult not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_CreateAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CreateAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GGargantua.v1.proto.JobService/CreateAJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CreateAJob(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_StartAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).StartAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GGargantua.v1.proto.JobService/StartAJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).StartAJob(ctx, req.(*UpdateJobContent))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_CancelAJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).CancelAJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GGargantua.v1.proto.JobService/CancelAJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).CancelAJob(ctx, req.(*JobUuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetJobResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetJobResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GGargantua.v1.proto.JobService/GetJobResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetJobResult(ctx, req.(*JobUuid))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GGargantua.v1.proto.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAJob",
			Handler:    _JobService_CreateAJob_Handler,
		},
		{
			MethodName: "StartAJob",
			Handler:    _JobService_StartAJob_Handler,
		},
		{
			MethodName: "CancelAJob",
			Handler:    _JobService_CancelAJob_Handler,
		},
		{
			MethodName: "GetJobResult",
			Handler:    _JobService_GetJobResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto_files/job.proto",
}