// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: experiment/metastore/compactionpb/compaction.proto

package compactionpb

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
	return file_experiment_metastore_compactionpb_compaction_proto_enumTypes[0].Descriptor()
}

func (CompactionStatus) Type() protoreflect.EnumType {
	return &file_experiment_metastore_compactionpb_compaction_proto_enumTypes[0]
}

func (x CompactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompactionStatus.Descriptor instead.
func (CompactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_experiment_metastore_compactionpb_compaction_proto_rawDescGZIP(), []int{0}
}

type CompactionJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique name of the job.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of the input blocks.
	Blocks []string `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	// Compaction level (all blocks are the same)
	CompactionLevel uint32 `protobuf:"varint,3,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	// The index of the raft command that changed the status of the job.
	// Used as a fencing token in conjunction with the lease_expires_at
	// field to manage ownership of the compaction job. Any access to the
	// job must be guarded by the check: current_index >= raft_log_index.
	// If the check fails, the access should be denied.
	//
	// The index is updated every time the job is assigned to a worker.
	RaftLogIndex uint64 `protobuf:"varint,4,opt,name=raft_log_index,json=raftLogIndex,proto3" json:"raft_log_index,omitempty"`
	// Shard the blocks belong to.
	Shard uint32 `protobuf:"varint,5,opt,name=shard,proto3" json:"shard,omitempty"`
	// Optional, empty for compaction level 0.
	TenantId string           `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Status   CompactionStatus `protobuf:"varint,7,opt,name=status,proto3,enum=compaction.CompactionStatus" json:"status,omitempty"`
	// The time the compaction job lease expires. If a lease is expired, the
	// job is considered abandoned and can be picked up by another worker.
	// The expiration check should be done by comparing the timestamp of
	// the raft log entry (command that accesses the job) with the value of
	// this field.
	//
	// The lease is extended every time the owner reports a status update.
	LeaseExpiresAt int64 `protobuf:"varint,8,opt,name=lease_expires_at,json=leaseExpiresAt,proto3" json:"lease_expires_at,omitempty"`
	// The number of failures when processing this job. Used for retries.
	Failures          uint32 `protobuf:"varint,9,opt,name=failures,proto3" json:"failures,omitempty"`
	LastFailureReason string `protobuf:"bytes,10,opt,name=last_failure_reason,json=lastFailureReason,proto3" json:"last_failure_reason,omitempty"`
	// Timestamp in nanoseconds from Unix epoch when the job was added.
	AddedAt int64 `protobuf:"varint,11,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
}

func (x *CompactionJob) Reset() {
	*x = CompactionJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_metastore_compactionpb_compaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJob) ProtoMessage() {}

func (x *CompactionJob) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_metastore_compactionpb_compaction_proto_msgTypes[0]
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
	return file_experiment_metastore_compactionpb_compaction_proto_rawDescGZIP(), []int{0}
}

func (x *CompactionJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompactionJob) GetBlocks() []string {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *CompactionJob) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
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

func (x *CompactionJob) GetStatus() CompactionStatus {
	if x != nil {
		return x.Status
	}
	return CompactionStatus_COMPACTION_STATUS_UNSPECIFIED
}

func (x *CompactionJob) GetLeaseExpiresAt() int64 {
	if x != nil {
		return x.LeaseExpiresAt
	}
	return 0
}

func (x *CompactionJob) GetFailures() uint32 {
	if x != nil {
		return x.Failures
	}
	return 0
}

func (x *CompactionJob) GetLastFailureReason() string {
	if x != nil {
		return x.LastFailureReason
	}
	return ""
}

func (x *CompactionJob) GetAddedAt() int64 {
	if x != nil {
		return x.AddedAt
	}
	return 0
}

type CompactionJobBlockQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompactionLevel uint32   `protobuf:"varint,1,opt,name=compaction_level,json=compactionLevel,proto3" json:"compaction_level,omitempty"`
	Shard           uint32   `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	Tenant          string   `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Blocks          []string `protobuf:"bytes,4,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *CompactionJobBlockQueue) Reset() {
	*x = CompactionJobBlockQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_metastore_compactionpb_compaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactionJobBlockQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactionJobBlockQueue) ProtoMessage() {}

func (x *CompactionJobBlockQueue) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_metastore_compactionpb_compaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactionJobBlockQueue.ProtoReflect.Descriptor instead.
func (*CompactionJobBlockQueue) Descriptor() ([]byte, []int) {
	return file_experiment_metastore_compactionpb_compaction_proto_rawDescGZIP(), []int{1}
}

func (x *CompactionJobBlockQueue) GetCompactionLevel() uint32 {
	if x != nil {
		return x.CompactionLevel
	}
	return 0
}

func (x *CompactionJobBlockQueue) GetShard() uint32 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *CompactionJobBlockQueue) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CompactionJobBlockQueue) GetBlocks() []string {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_experiment_metastore_compactionpb_compaction_proto protoreflect.FileDescriptor

var file_experiment_metastore_compactionpb_compaction_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x86, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61, 0x66,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x72, 0x61, 0x66, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x73,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x17, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2a, 0xb7, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21,
	0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04,
	0x42, 0xad, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0a,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x16, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_experiment_metastore_compactionpb_compaction_proto_rawDescOnce sync.Once
	file_experiment_metastore_compactionpb_compaction_proto_rawDescData = file_experiment_metastore_compactionpb_compaction_proto_rawDesc
)

func file_experiment_metastore_compactionpb_compaction_proto_rawDescGZIP() []byte {
	file_experiment_metastore_compactionpb_compaction_proto_rawDescOnce.Do(func() {
		file_experiment_metastore_compactionpb_compaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_experiment_metastore_compactionpb_compaction_proto_rawDescData)
	})
	return file_experiment_metastore_compactionpb_compaction_proto_rawDescData
}

var file_experiment_metastore_compactionpb_compaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_experiment_metastore_compactionpb_compaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_experiment_metastore_compactionpb_compaction_proto_goTypes = []any{
	(CompactionStatus)(0),           // 0: compaction.CompactionStatus
	(*CompactionJob)(nil),           // 1: compaction.CompactionJob
	(*CompactionJobBlockQueue)(nil), // 2: compaction.CompactionJobBlockQueue
}
var file_experiment_metastore_compactionpb_compaction_proto_depIdxs = []int32{
	0, // 0: compaction.CompactionJob.status:type_name -> compaction.CompactionStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_experiment_metastore_compactionpb_compaction_proto_init() }
func file_experiment_metastore_compactionpb_compaction_proto_init() {
	if File_experiment_metastore_compactionpb_compaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_experiment_metastore_compactionpb_compaction_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_experiment_metastore_compactionpb_compaction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CompactionJobBlockQueue); i {
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
			RawDescriptor: file_experiment_metastore_compactionpb_compaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_experiment_metastore_compactionpb_compaction_proto_goTypes,
		DependencyIndexes: file_experiment_metastore_compactionpb_compaction_proto_depIdxs,
		EnumInfos:         file_experiment_metastore_compactionpb_compaction_proto_enumTypes,
		MessageInfos:      file_experiment_metastore_compactionpb_compaction_proto_msgTypes,
	}.Build()
	File_experiment_metastore_compactionpb_compaction_proto = out.File
	file_experiment_metastore_compactionpb_compaction_proto_rawDesc = nil
	file_experiment_metastore_compactionpb_compaction_proto_goTypes = nil
	file_experiment_metastore_compactionpb_compaction_proto_depIdxs = nil
}
