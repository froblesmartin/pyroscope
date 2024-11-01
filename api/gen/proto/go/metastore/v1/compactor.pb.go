// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: metastore/v1/compactor.proto

package metastorev1

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

type CompactionStatus int32

const (
	CompactionStatus_COMPACTION_STATUS_UNSPECIFIED CompactionStatus = 0
	CompactionStatus_COMPACTION_STATUS_IN_PROGRESS CompactionStatus = 1
	CompactionStatus_COMPACTION_STATUS_SUCCESS     CompactionStatus = 2
	CompactionStatus_COMPACTION_STATUS_FAILURE     CompactionStatus = 3
	CompactionStatus_COMPACTION_STATUS_CANCELLED   CompactionStatus = 4
)

// Enum value maps for CompactionStatus.
var (
	CompactionStatus_name = map[int32]string{
		0: "COMPACTION_STATUS_UNSPECIFIED",
		1: "COMPACTION_STATUS_IN_PROGRESS",
		2: "COMPACTION_STATUS_SUCCESS",
		3: "COMPACTION_STATUS_FAILURE",
		4: "COMPACTION_STATUS_CANCELLED",
	}
	CompactionStatus_value = map[string]int32{
		"COMPACTION_STATUS_UNSPECIFIED": 0,
		"COMPACTION_STATUS_IN_PROGRESS": 1,
		"COMPACTION_STATUS_SUCCESS":     2,
		"COMPACTION_STATUS_FAILURE":     3,
		"COMPACTION_STATUS_CANCELLED":   4,
	}
)

func (x CompactionStatus) Enum() *CompactionStatus {
	p := new(CompactionStatus)
	*p = x
	return p
}

func (x CompactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_metastore_v1_compactor_proto_enumTypes[0].Descriptor()
}

func (CompactionStatus) Type() protoreflect.EnumType {
	return &file_metastore_v1_compactor_proto_enumTypes[0]
}

func (x CompactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompactionStatus.Descriptor instead.
func (CompactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{0}
}

type PollCompactionJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A batch of status updates for in-progress jobs from a worker.
	JobStatusUpdates []*CompactionJobStatus `protobuf:"bytes,1,rep,name=job_status_updates,json=jobStatusUpdates,proto3" json:"job_status_updates,omitempty"`
	// How many new jobs a worker can be assigned to.
	JobCapacity uint32 `protobuf:"varint,2,opt,name=job_capacity,json=jobCapacity,proto3" json:"job_capacity,omitempty"`
	// Max number of block tombstones a worker can clean up
	// in a single compaction job.
	CleanupCapacity uint32 `protobuf:"varint,3,opt,name=cleanup_capacity,json=cleanupCapacity,proto3" json:"cleanup_capacity,omitempty"`
}

func (x *PollCompactionJobsRequest) Reset() {
	*x = PollCompactionJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollCompactionJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollCompactionJobsRequest) ProtoMessage() {}

func (x *PollCompactionJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollCompactionJobsRequest.ProtoReflect.Descriptor instead.
func (*PollCompactionJobsRequest) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{0}
}

func (x *PollCompactionJobsRequest) GetJobStatusUpdates() []*CompactionJobStatus {
	if x != nil {
		return x.JobStatusUpdates
	}
	return nil
}

func (x *PollCompactionJobsRequest) GetJobCapacity() uint32 {
	if x != nil {
		return x.JobCapacity
	}
	return 0
}

func (x *PollCompactionJobsRequest) GetCleanupCapacity() uint32 {
	if x != nil {
		return x.CleanupCapacity
	}
	return 0
}

type PollCompactionJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactionJobs []*CompactionJob `protobuf:"bytes,1,rep,name=compaction_jobs,json=compactionJobs,proto3" json:"compaction_jobs,omitempty"`
	CancelledJobs  []string         `protobuf:"bytes,2,rep,name=cancelled_jobs,json=cancelledJobs,proto3" json:"cancelled_jobs,omitempty"`
}

func (x *PollCompactionJobsResponse) Reset() {
	*x = PollCompactionJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollCompactionJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollCompactionJobsResponse) ProtoMessage() {}

func (x *PollCompactionJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollCompactionJobsResponse.ProtoReflect.Descriptor instead.
func (*PollCompactionJobsResponse) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{1}
}

func (x *PollCompactionJobsResponse) GetCompactionJobs() []*CompactionJob {
	if x != nil {
		return x.CompactionJobs
	}
	return nil
}

func (x *PollCompactionJobsResponse) GetCancelledJobs() []string {
	if x != nil {
		return x.CancelledJobs
	}
	return nil
}

// One compaction job may result in multiple output blocks.
type CompactionJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique name of the job.
	Name    string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options *CompactionOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	// List of the input blocks.
	Blocks []*BlockMeta         `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Status *CompactionJobStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// Fencing token.
	RaftLogIndex uint64 `protobuf:"varint,5,opt,name=raft_log_index,json=raftLogIndex,proto3" json:"raft_log_index,omitempty"`
	// Shard the blocks belong to.
	Shard uint32 `protobuf:"varint,6,opt,name=shard,proto3" json:"shard,omitempty"`
	// Optional, empty for compaction level 0.
	TenantId        string `protobuf:"bytes,7,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	CompactionLevel uint32 `protobuf:"varint,8,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	// Blocks deleted from the index but present
	// in the storage. The full block path is used.
	Tombstones []string `protobuf:"bytes,9,rep,name=tombstones,proto3" json:"tombstones,omitempty"`
}

func (x *CompactionJob) Reset() {
	*x = CompactionJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJob) ProtoMessage() {}

func (x *CompactionJob) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionJob.ProtoReflect.Descriptor instead.
func (*CompactionJob) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{2}
}

func (x *CompactionJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompactionJob) GetOptions() *CompactionOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CompactionJob) GetBlocks() []*BlockMeta {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *CompactionJob) GetStatus() *CompactionJobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CompactionJob) GetRaftLogIndex() uint64 {
	if x != nil {
		return x.RaftLogIndex
	}
	return 0
}

func (x *CompactionJob) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *CompactionJob) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *CompactionJob) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *CompactionJob) GetTombstones() []string {
	if x != nil {
		return x.Tombstones
	}
	return nil
}

type CompactionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How often the compaction worker should update
	// the job status. If overdue, the job ownership
	// is revoked.
	StatusUpdateIntervalSeconds uint64 `protobuf:"varint,1,opt,name=status_update_interval_seconds,json=statusUpdateIntervalSeconds,proto3" json:"status_update_interval_seconds,omitempty"`
}

func (x *CompactionOptions) Reset() {
	*x = CompactionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionOptions) ProtoMessage() {}

func (x *CompactionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionOptions.ProtoReflect.Descriptor instead.
func (*CompactionOptions) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{3}
}

func (x *CompactionOptions) GetStatusUpdateIntervalSeconds() uint64 {
	if x != nil {
		return x.StatusUpdateIntervalSeconds
	}
	return 0
}

type CompactionJobStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	// Status update allows the planner to keep
	// track of the job ownership and compaction
	// progress:
	//   - If the job status is other than IN_PROGRESS,
	//     the ownership of the job is revoked.
	//   - FAILURE must only be sent if the failure is
	//     persistent and the compaction can't be accomplished.
	//   - completed_job must be empty if the status is
	//     other than SUCCESS, and vice-versa.
	//   - UNSPECIFIED must be sent if the worker rejects
	//     or cancels the compaction job.
	//
	// Partial results/status is not allowed.
	Status          CompactionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=metastore.v1.CompactionStatus" json:"status,omitempty"`
	CompactedBlocks []*BlockMeta     `protobuf:"bytes,3,rep,name=compacted_blocks,json=compactedBlocks,proto3" json:"compacted_blocks,omitempty"`
	// Fencing token.
	RaftLogIndex uint64 `protobuf:"varint,4,opt,name=raft_log_index,json=raftLogIndex,proto3" json:"raft_log_index,omitempty"`
	// Shard the blocks belong to.
	Shard uint32 `protobuf:"varint,5,opt,name=shard,proto3" json:"shard,omitempty"`
	// Optional, empty for compaction level 0.
	TenantId string `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Blocks deleted from the storage.
	DeletedBlocks []string `protobuf:"bytes,7,rep,name=deleted_blocks,json=deletedBlocks,proto3" json:"deleted_blocks,omitempty"`
}

func (x *CompactionJobStatus) Reset() {
	*x = CompactionJobStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metastore_v1_compactor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJobStatus) ProtoMessage() {}

func (x *CompactionJobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_metastore_v1_compactor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionJobStatus.ProtoReflect.Descriptor instead.
func (*CompactionJobStatus) Descriptor() ([]byte, []int) {
	return file_metastore_v1_compactor_proto_rawDescGZIP(), []int{4}
}

func (x *CompactionJobStatus) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *CompactionJobStatus) GetStatus() CompactionStatus {
	if x != nil {
		return x.Status
	}
	return CompactionStatus_COMPACTION_STATUS_UNSPECIFIED
}

func (x *CompactionJobStatus) GetCompactedBlocks() []*BlockMeta {
	if x != nil {
		return x.CompactedBlocks
	}
	return nil
}

func (x *CompactionJobStatus) GetRaftLogIndex() uint64 {
	if x != nil {
		return x.RaftLogIndex
	}
	return 0
}

func (x *CompactionJobStatus) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *CompactionJobStatus) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *CompactionJobStatus) GetDeletedBlocks() []string {
	if x != nil {
		return x.DeletedBlocks
	}
	return nil
}

var File_metastore_v1_compactor_proto protoreflect.FileDescriptor

var file_metastore_v1_compactor_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x19, 0x50,
	0x6f, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x6f, 0x62,
	0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10,
	0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x50, 0x6f, 0x6c, 0x6c,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x4a,
	0x6f, 0x62, 0x73, 0x22, 0xee, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6d, 0x62, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x1b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xac,
	0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x0a,
	0x0e, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2a, 0xb7, 0x01,
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0x7e, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x12,
	0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f,
	0x62, 0x73, 0x12, 0x27, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xbb, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61,
	0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x4d, 0x65,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x65, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metastore_v1_compactor_proto_rawDescOnce sync.Once
	file_metastore_v1_compactor_proto_rawDescData = file_metastore_v1_compactor_proto_rawDesc
)

func file_metastore_v1_compactor_proto_rawDescGZIP() []byte {
	file_metastore_v1_compactor_proto_rawDescOnce.Do(func() {
		file_metastore_v1_compactor_proto_rawDescData = protoimpl.X.CompressGZIP(file_metastore_v1_compactor_proto_rawDescData)
	})
	return file_metastore_v1_compactor_proto_rawDescData
}

var file_metastore_v1_compactor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metastore_v1_compactor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metastore_v1_compactor_proto_goTypes = []any{
	(CompactionStatus)(0),              // 0: metastore.v1.CompactionStatus
	(*PollCompactionJobsRequest)(nil),  // 1: metastore.v1.PollCompactionJobsRequest
	(*PollCompactionJobsResponse)(nil), // 2: metastore.v1.PollCompactionJobsResponse
	(*CompactionJob)(nil),              // 3: metastore.v1.CompactionJob
	(*CompactionOptions)(nil),          // 4: metastore.v1.CompactionOptions
	(*CompactionJobStatus)(nil),        // 5: metastore.v1.CompactionJobStatus
	(*BlockMeta)(nil),                  // 6: metastore.v1.BlockMeta
}
var file_metastore_v1_compactor_proto_depIdxs = []int32{
	5, // 0: metastore.v1.PollCompactionJobsRequest.job_status_updates:type_name -> metastore.v1.CompactionJobStatus
	3, // 1: metastore.v1.PollCompactionJobsResponse.compaction_jobs:type_name -> metastore.v1.CompactionJob
	4, // 2: metastore.v1.CompactionJob.options:type_name -> metastore.v1.CompactionOptions
	6, // 3: metastore.v1.CompactionJob.blocks:type_name -> metastore.v1.BlockMeta
	5, // 4: metastore.v1.CompactionJob.status:type_name -> metastore.v1.CompactionJobStatus
	0, // 5: metastore.v1.CompactionJobStatus.status:type_name -> metastore.v1.CompactionStatus
	6, // 6: metastore.v1.CompactionJobStatus.compacted_blocks:type_name -> metastore.v1.BlockMeta
	1, // 7: metastore.v1.CompactionPlanner.PollCompactionJobs:input_type -> metastore.v1.PollCompactionJobsRequest
	2, // 8: metastore.v1.CompactionPlanner.PollCompactionJobs:output_type -> metastore.v1.PollCompactionJobsResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_metastore_v1_compactor_proto_init() }
func file_metastore_v1_compactor_proto_init() {
	if File_metastore_v1_compactor_proto != nil {
		return
	}
	file_metastore_v1_metastore_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metastore_v1_compactor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PollCompactionJobsRequest); i {
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
		file_metastore_v1_compactor_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PollCompactionJobsResponse); i {
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
		file_metastore_v1_compactor_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionJob); i {
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
		file_metastore_v1_compactor_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionOptions); i {
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
		file_metastore_v1_compactor_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionJobStatus); i {
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
			RawDescriptor: file_metastore_v1_compactor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metastore_v1_compactor_proto_goTypes,
		DependencyIndexes: file_metastore_v1_compactor_proto_depIdxs,
		EnumInfos:         file_metastore_v1_compactor_proto_enumTypes,
		MessageInfos:      file_metastore_v1_compactor_proto_msgTypes,
	}.Build()
	File_metastore_v1_compactor_proto = out.File
	file_metastore_v1_compactor_proto_rawDesc = nil
	file_metastore_v1_compactor_proto_goTypes = nil
	file_metastore_v1_compactor_proto_depIdxs = nil
}
