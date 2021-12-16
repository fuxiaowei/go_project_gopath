// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/cloud/bigquery/storage/v1alpha2/stream.proto

package storage

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type WriteStream_Type int32

const (
	// Unknown type.
	WriteStream_TYPE_UNSPECIFIED WriteStream_Type = 0
	// Data will commit automatically and appear as soon as the write is
	// acknowledged.
	WriteStream_COMMITTED WriteStream_Type = 1
	// Data is invisible until the stream is committed.
	WriteStream_PENDING WriteStream_Type = 2
	// Data is only visible up to the offset to which it was flushed.
	WriteStream_BUFFERED WriteStream_Type = 3
)

// Enum value maps for WriteStream_Type.
var (
	WriteStream_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "COMMITTED",
		2: "PENDING",
		3: "BUFFERED",
	}
	WriteStream_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"COMMITTED":        1,
		"PENDING":          2,
		"BUFFERED":         3,
	}
)

func (x WriteStream_Type) Enum() *WriteStream_Type {
	p := new(WriteStream_Type)
	*p = x
	return p
}

func (x WriteStream_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WriteStream_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_storage_v1alpha2_stream_proto_enumTypes[0].Descriptor()
}

func (WriteStream_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_storage_v1alpha2_stream_proto_enumTypes[0]
}

func (x WriteStream_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WriteStream_Type.Descriptor instead.
func (WriteStream_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescGZIP(), []int{0, 0}
}

// Information about a single stream that gets data inside the storage system.
type WriteStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Name of the stream, in the form
	// `projects/{project}/datasets/{dataset}/tables/{table}/streams/{stream}`.
	Name string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type WriteStream_Type `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.bigquery.storage.v1alpha2.WriteStream_Type" json:"type,omitempty"`
	// Output only. Create time of the stream.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Commit time of the stream.
	// If a stream is of `COMMITTED` type, then it will have a commit_time same as
	// `create_time`. If the stream is of `PENDING` type, commit_time being empty
	// means it is not committed.
	CommitTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=commit_time,json=commitTime,proto3" json:"commit_time,omitempty"`
	// Output only. The schema of the destination table. It is only returned in
	// `CreateWriteStream` response. Caller should generate data that's
	// compatible with this schema to send in initial `AppendRowsRequest`.
	// The table schema could go out of date during the life time of the stream.
	TableSchema *TableSchema `protobuf:"bytes,5,opt,name=table_schema,json=tableSchema,proto3" json:"table_schema,omitempty"`
	// Id set by client to annotate its identity.
	ExternalId string `protobuf:"bytes,6,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *WriteStream) Reset() {
	*x = WriteStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1alpha2_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteStream) ProtoMessage() {}

func (x *WriteStream) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1alpha2_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteStream.ProtoReflect.Descriptor instead.
func (*WriteStream) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescGZIP(), []int{0}
}

func (x *WriteStream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WriteStream) GetType() WriteStream_Type {
	if x != nil {
		return x.Type
	}
	return WriteStream_TYPE_UNSPECIFIED
}

func (x *WriteStream) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *WriteStream) GetCommitTime() *timestamp.Timestamp {
	if x != nil {
		return x.CommitTime
	}
	return nil
}

func (x *WriteStream) GetTableSchema() *TableSchema {
	if x != nil {
		return x.TableSchema
	}
	return nil
}

func (x *WriteStream) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

var File_google_cloud_bigquery_storage_v1alpha2_stream_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb,
	0x04, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5b,
	0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f,
	0x4d, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52,
	0x45, 0x44, 0x10, 0x03, 0x3a, 0x76, 0xea, 0x41, 0x73, 0x0a, 0x2a, 0x62, 0x69, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x45, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x7d, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x7d, 0x42, 0xda, 0x01, 0x0a,
	0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x5a, 0x4d, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0xea, 0x41, 0x5c, 0x0a, 0x24, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x34, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x2f, 0x7b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescData = file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDesc
)

func file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescData)
	})
	return file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDescData
}

var file_google_cloud_bigquery_storage_v1alpha2_stream_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_bigquery_storage_v1alpha2_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_bigquery_storage_v1alpha2_stream_proto_goTypes = []interface{}{
	(WriteStream_Type)(0),       // 0: google.cloud.bigquery.storage.v1alpha2.WriteStream.Type
	(*WriteStream)(nil),         // 1: google.cloud.bigquery.storage.v1alpha2.WriteStream
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*TableSchema)(nil),         // 3: google.cloud.bigquery.storage.v1alpha2.TableSchema
}
var file_google_cloud_bigquery_storage_v1alpha2_stream_proto_depIdxs = []int32{
	0, // 0: google.cloud.bigquery.storage.v1alpha2.WriteStream.type:type_name -> google.cloud.bigquery.storage.v1alpha2.WriteStream.Type
	2, // 1: google.cloud.bigquery.storage.v1alpha2.WriteStream.create_time:type_name -> google.protobuf.Timestamp
	2, // 2: google.cloud.bigquery.storage.v1alpha2.WriteStream.commit_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.bigquery.storage.v1alpha2.WriteStream.table_schema:type_name -> google.cloud.bigquery.storage.v1alpha2.TableSchema
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_storage_v1alpha2_stream_proto_init() }
func file_google_cloud_bigquery_storage_v1alpha2_stream_proto_init() {
	if File_google_cloud_bigquery_storage_v1alpha2_stream_proto != nil {
		return
	}
	file_google_cloud_bigquery_storage_v1alpha2_table_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_storage_v1alpha2_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteStream); i {
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
			RawDescriptor: file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_storage_v1alpha2_stream_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_storage_v1alpha2_stream_proto_depIdxs,
		EnumInfos:         file_google_cloud_bigquery_storage_v1alpha2_stream_proto_enumTypes,
		MessageInfos:      file_google_cloud_bigquery_storage_v1alpha2_stream_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_storage_v1alpha2_stream_proto = out.File
	file_google_cloud_bigquery_storage_v1alpha2_stream_proto_rawDesc = nil
	file_google_cloud_bigquery_storage_v1alpha2_stream_proto_goTypes = nil
	file_google_cloud_bigquery_storage_v1alpha2_stream_proto_depIdxs = nil
}
