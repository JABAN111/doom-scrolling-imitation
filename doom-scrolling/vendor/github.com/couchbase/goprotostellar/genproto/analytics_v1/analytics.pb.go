// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: couchbase/analytics/v1/analytics.proto

package analytics_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AnalyticsQueryRequest_ScanConsistency int32

const (
	AnalyticsQueryRequest_SCAN_CONSISTENCY_NOT_BOUNDED  AnalyticsQueryRequest_ScanConsistency = 0
	AnalyticsQueryRequest_SCAN_CONSISTENCY_REQUEST_PLUS AnalyticsQueryRequest_ScanConsistency = 1
)

// Enum value maps for AnalyticsQueryRequest_ScanConsistency.
var (
	AnalyticsQueryRequest_ScanConsistency_name = map[int32]string{
		0: "SCAN_CONSISTENCY_NOT_BOUNDED",
		1: "SCAN_CONSISTENCY_REQUEST_PLUS",
	}
	AnalyticsQueryRequest_ScanConsistency_value = map[string]int32{
		"SCAN_CONSISTENCY_NOT_BOUNDED":  0,
		"SCAN_CONSISTENCY_REQUEST_PLUS": 1,
	}
)

func (x AnalyticsQueryRequest_ScanConsistency) Enum() *AnalyticsQueryRequest_ScanConsistency {
	p := new(AnalyticsQueryRequest_ScanConsistency)
	*p = x
	return p
}

func (x AnalyticsQueryRequest_ScanConsistency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnalyticsQueryRequest_ScanConsistency) Descriptor() protoreflect.EnumDescriptor {
	return file_couchbase_analytics_v1_analytics_proto_enumTypes[0].Descriptor()
}

func (AnalyticsQueryRequest_ScanConsistency) Type() protoreflect.EnumType {
	return &file_couchbase_analytics_v1_analytics_proto_enumTypes[0]
}

func (x AnalyticsQueryRequest_ScanConsistency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnalyticsQueryRequest_ScanConsistency.Descriptor instead.
func (AnalyticsQueryRequest_ScanConsistency) EnumDescriptor() ([]byte, []int) {
	return file_couchbase_analytics_v1_analytics_proto_rawDescGZIP(), []int{0, 0}
}

type AnalyticsQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName           *string                                `protobuf:"bytes,8,opt,name=bucket_name,json=bucketName,proto3,oneof" json:"bucket_name,omitempty"`
	ScopeName            *string                                `protobuf:"bytes,9,opt,name=scope_name,json=scopeName,proto3,oneof" json:"scope_name,omitempty"`
	Statement            string                                 `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
	ReadOnly             *bool                                  `protobuf:"varint,2,opt,name=read_only,json=readOnly,proto3,oneof" json:"read_only,omitempty"`
	ClientContextId      *string                                `protobuf:"bytes,3,opt,name=client_context_id,json=clientContextId,proto3,oneof" json:"client_context_id,omitempty"`
	Priority             *bool                                  `protobuf:"varint,4,opt,name=priority,proto3,oneof" json:"priority,omitempty"`
	ScanConsistency      *AnalyticsQueryRequest_ScanConsistency `protobuf:"varint,5,opt,name=scan_consistency,json=scanConsistency,proto3,enum=couchbase.analytics.v1.AnalyticsQueryRequest_ScanConsistency,oneof" json:"scan_consistency,omitempty"`
	PositionalParameters [][]byte                               `protobuf:"bytes,6,rep,name=positional_parameters,json=positionalParameters,proto3" json:"positional_parameters,omitempty"`
	NamedParameters      map[string][]byte                      `protobuf:"bytes,7,rep,name=named_parameters,json=namedParameters,proto3" json:"named_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AnalyticsQueryRequest) Reset() {
	*x = AnalyticsQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsQueryRequest) ProtoMessage() {}

func (x *AnalyticsQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsQueryRequest.ProtoReflect.Descriptor instead.
func (*AnalyticsQueryRequest) Descriptor() ([]byte, []int) {
	return file_couchbase_analytics_v1_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *AnalyticsQueryRequest) GetBucketName() string {
	if x != nil && x.BucketName != nil {
		return *x.BucketName
	}
	return ""
}

func (x *AnalyticsQueryRequest) GetScopeName() string {
	if x != nil && x.ScopeName != nil {
		return *x.ScopeName
	}
	return ""
}

func (x *AnalyticsQueryRequest) GetStatement() string {
	if x != nil {
		return x.Statement
	}
	return ""
}

func (x *AnalyticsQueryRequest) GetReadOnly() bool {
	if x != nil && x.ReadOnly != nil {
		return *x.ReadOnly
	}
	return false
}

func (x *AnalyticsQueryRequest) GetClientContextId() string {
	if x != nil && x.ClientContextId != nil {
		return *x.ClientContextId
	}
	return ""
}

func (x *AnalyticsQueryRequest) GetPriority() bool {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return false
}

func (x *AnalyticsQueryRequest) GetScanConsistency() AnalyticsQueryRequest_ScanConsistency {
	if x != nil && x.ScanConsistency != nil {
		return *x.ScanConsistency
	}
	return AnalyticsQueryRequest_SCAN_CONSISTENCY_NOT_BOUNDED
}

func (x *AnalyticsQueryRequest) GetPositionalParameters() [][]byte {
	if x != nil {
		return x.PositionalParameters
	}
	return nil
}

func (x *AnalyticsQueryRequest) GetNamedParameters() map[string][]byte {
	if x != nil {
		return x.NamedParameters
	}
	return nil
}

type AnalyticsQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows     [][]byte                         `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	MetaData *AnalyticsQueryResponse_MetaData `protobuf:"bytes,2,opt,name=meta_data,json=metaData,proto3,oneof" json:"meta_data,omitempty"`
}

func (x *AnalyticsQueryResponse) Reset() {
	*x = AnalyticsQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsQueryResponse) ProtoMessage() {}

func (x *AnalyticsQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsQueryResponse.ProtoReflect.Descriptor instead.
func (*AnalyticsQueryResponse) Descriptor() ([]byte, []int) {
	return file_couchbase_analytics_v1_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *AnalyticsQueryResponse) GetRows() [][]byte {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *AnalyticsQueryResponse) GetMetaData() *AnalyticsQueryResponse_MetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

type AnalyticsQueryResponse_Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElapsedTime      *durationpb.Duration `protobuf:"bytes,1,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	ExecutionTime    *durationpb.Duration `protobuf:"bytes,2,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	ResultCount      uint64               `protobuf:"varint,3,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	ResultSize       uint64               `protobuf:"varint,4,opt,name=result_size,json=resultSize,proto3" json:"result_size,omitempty"`
	MutationCount    uint64               `protobuf:"varint,5,opt,name=mutation_count,json=mutationCount,proto3" json:"mutation_count,omitempty"`
	SortCount        uint64               `protobuf:"varint,6,opt,name=sort_count,json=sortCount,proto3" json:"sort_count,omitempty"`
	ErrorCount       uint64               `protobuf:"varint,7,opt,name=error_count,json=errorCount,proto3" json:"error_count,omitempty"`
	WarningCount     uint64               `protobuf:"varint,8,opt,name=warning_count,json=warningCount,proto3" json:"warning_count,omitempty"`
	ProcessedObjects uint64               `protobuf:"varint,9,opt,name=processed_objects,json=processedObjects,proto3" json:"processed_objects,omitempty"`
}

func (x *AnalyticsQueryResponse_Metrics) Reset() {
	*x = AnalyticsQueryResponse_Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsQueryResponse_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsQueryResponse_Metrics) ProtoMessage() {}

func (x *AnalyticsQueryResponse_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsQueryResponse_Metrics.ProtoReflect.Descriptor instead.
func (*AnalyticsQueryResponse_Metrics) Descriptor() ([]byte, []int) {
	return file_couchbase_analytics_v1_analytics_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AnalyticsQueryResponse_Metrics) GetElapsedTime() *durationpb.Duration {
	if x != nil {
		return x.ElapsedTime
	}
	return nil
}

func (x *AnalyticsQueryResponse_Metrics) GetExecutionTime() *durationpb.Duration {
	if x != nil {
		return x.ExecutionTime
	}
	return nil
}

func (x *AnalyticsQueryResponse_Metrics) GetResultCount() uint64 {
	if x != nil {
		return x.ResultCount
	}
	return 0
}

func (x *AnalyticsQueryResponse_Metrics) GetResultSize() uint64 {
	if x != nil {
		return x.ResultSize
	}
	return 0
}

func (x *AnalyticsQueryResponse_Metrics) GetMutationCount() uint64 {
	if x != nil {
		return x.MutationCount
	}
	return 0
}

func (x *AnalyticsQueryResponse_Metrics) GetSortCount() uint64 {
	if x != nil {
		return x.SortCount
	}
	return 0
}

func (x *AnalyticsQueryResponse_Metrics) GetErrorCount() uint64 {
	if x != nil {
		return x.ErrorCount
	}
	return 0
}

func (x *AnalyticsQueryResponse_Metrics) GetWarningCount() uint64 {
	if x != nil {
		return x.WarningCount
	}
	return 0
}

func (x *AnalyticsQueryResponse_Metrics) GetProcessedObjects() uint64 {
	if x != nil {
		return x.ProcessedObjects
	}
	return 0
}

type AnalyticsQueryResponse_MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId       string                                     `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ClientContextId string                                     `protobuf:"bytes,2,opt,name=client_context_id,json=clientContextId,proto3" json:"client_context_id,omitempty"`
	Metrics         *AnalyticsQueryResponse_Metrics            `protobuf:"bytes,3,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Signature       []byte                                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Warnings        []*AnalyticsQueryResponse_MetaData_Warning `protobuf:"bytes,5,rep,name=warnings,proto3" json:"warnings,omitempty"`
	Status          string                                     `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AnalyticsQueryResponse_MetaData) Reset() {
	*x = AnalyticsQueryResponse_MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsQueryResponse_MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsQueryResponse_MetaData) ProtoMessage() {}

func (x *AnalyticsQueryResponse_MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsQueryResponse_MetaData.ProtoReflect.Descriptor instead.
func (*AnalyticsQueryResponse_MetaData) Descriptor() ([]byte, []int) {
	return file_couchbase_analytics_v1_analytics_proto_rawDescGZIP(), []int{1, 1}
}

func (x *AnalyticsQueryResponse_MetaData) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AnalyticsQueryResponse_MetaData) GetClientContextId() string {
	if x != nil {
		return x.ClientContextId
	}
	return ""
}

func (x *AnalyticsQueryResponse_MetaData) GetMetrics() *AnalyticsQueryResponse_Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *AnalyticsQueryResponse_MetaData) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *AnalyticsQueryResponse_MetaData) GetWarnings() []*AnalyticsQueryResponse_MetaData_Warning {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *AnalyticsQueryResponse_MetaData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type AnalyticsQueryResponse_MetaData_Warning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AnalyticsQueryResponse_MetaData_Warning) Reset() {
	*x = AnalyticsQueryResponse_MetaData_Warning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyticsQueryResponse_MetaData_Warning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyticsQueryResponse_MetaData_Warning) ProtoMessage() {}

func (x *AnalyticsQueryResponse_MetaData_Warning) ProtoReflect() protoreflect.Message {
	mi := &file_couchbase_analytics_v1_analytics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyticsQueryResponse_MetaData_Warning.ProtoReflect.Descriptor instead.
func (*AnalyticsQueryResponse_MetaData_Warning) Descriptor() ([]byte, []int) {
	return file_couchbase_analytics_v1_analytics_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *AnalyticsQueryResponse_MetaData_Warning) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AnalyticsQueryResponse_MetaData_Warning) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_couchbase_analytics_v1_analytics_proto protoreflect.FileDescriptor

var file_couchbase_analytics_v1_analytics_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x06, 0x0a, 0x15, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x03, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x6d, 0x0a, 0x10, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x3d, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x48,
	0x05, 0x52, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x15, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x14, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x6d, 0x0a, 0x10, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x42, 0x0a, 0x14, 0x4e, 0x61, 0x6d,
	0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a,
	0x0f, 0x53, 0x63, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x20, 0x0a, 0x1c, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54,
	0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x49,
	0x53, 0x54, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x50,
	0x4c, 0x55, 0x53, 0x10, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x94, 0x07, 0x0a, 0x16, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x59, 0x0a, 0x09, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63,
	0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x88, 0x01, 0x01, 0x1a, 0x86, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x3c, 0x0a, 0x0c, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x77,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a, 0xf3, 0x02,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x37, 0x0a, 0x07, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x87, 0x01, 0x0a, 0x10, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x84, 0x02, 0x0a, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61,
	0x72, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x75,
	0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x74,
	0x65, 0x6c, 0x6c, 0x61, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0x3b, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x76, 0x31, 0xaa, 0x02, 0x23, 0x43, 0x6f, 0x75, 0x63, 0x68,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61,
	0x72, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x2d, 0x43, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x5c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x5c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x31, 0x43, 0x6f, 0x75, 0x63, 0x68, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x3a, 0x3a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x3a, 0x3a, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_couchbase_analytics_v1_analytics_proto_rawDescOnce sync.Once
	file_couchbase_analytics_v1_analytics_proto_rawDescData = file_couchbase_analytics_v1_analytics_proto_rawDesc
)

func file_couchbase_analytics_v1_analytics_proto_rawDescGZIP() []byte {
	file_couchbase_analytics_v1_analytics_proto_rawDescOnce.Do(func() {
		file_couchbase_analytics_v1_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_couchbase_analytics_v1_analytics_proto_rawDescData)
	})
	return file_couchbase_analytics_v1_analytics_proto_rawDescData
}

var file_couchbase_analytics_v1_analytics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_couchbase_analytics_v1_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_couchbase_analytics_v1_analytics_proto_goTypes = []interface{}{
	(AnalyticsQueryRequest_ScanConsistency)(0),      // 0: couchbase.analytics.v1.AnalyticsQueryRequest.ScanConsistency
	(*AnalyticsQueryRequest)(nil),                   // 1: couchbase.analytics.v1.AnalyticsQueryRequest
	(*AnalyticsQueryResponse)(nil),                  // 2: couchbase.analytics.v1.AnalyticsQueryResponse
	nil,                                             // 3: couchbase.analytics.v1.AnalyticsQueryRequest.NamedParametersEntry
	(*AnalyticsQueryResponse_Metrics)(nil),          // 4: couchbase.analytics.v1.AnalyticsQueryResponse.Metrics
	(*AnalyticsQueryResponse_MetaData)(nil),         // 5: couchbase.analytics.v1.AnalyticsQueryResponse.MetaData
	(*AnalyticsQueryResponse_MetaData_Warning)(nil), // 6: couchbase.analytics.v1.AnalyticsQueryResponse.MetaData.Warning
	(*durationpb.Duration)(nil),                     // 7: google.protobuf.Duration
}
var file_couchbase_analytics_v1_analytics_proto_depIdxs = []int32{
	0, // 0: couchbase.analytics.v1.AnalyticsQueryRequest.scan_consistency:type_name -> couchbase.analytics.v1.AnalyticsQueryRequest.ScanConsistency
	3, // 1: couchbase.analytics.v1.AnalyticsQueryRequest.named_parameters:type_name -> couchbase.analytics.v1.AnalyticsQueryRequest.NamedParametersEntry
	5, // 2: couchbase.analytics.v1.AnalyticsQueryResponse.meta_data:type_name -> couchbase.analytics.v1.AnalyticsQueryResponse.MetaData
	7, // 3: couchbase.analytics.v1.AnalyticsQueryResponse.Metrics.elapsed_time:type_name -> google.protobuf.Duration
	7, // 4: couchbase.analytics.v1.AnalyticsQueryResponse.Metrics.execution_time:type_name -> google.protobuf.Duration
	4, // 5: couchbase.analytics.v1.AnalyticsQueryResponse.MetaData.metrics:type_name -> couchbase.analytics.v1.AnalyticsQueryResponse.Metrics
	6, // 6: couchbase.analytics.v1.AnalyticsQueryResponse.MetaData.warnings:type_name -> couchbase.analytics.v1.AnalyticsQueryResponse.MetaData.Warning
	1, // 7: couchbase.analytics.v1.AnalyticsService.AnalyticsQuery:input_type -> couchbase.analytics.v1.AnalyticsQueryRequest
	2, // 8: couchbase.analytics.v1.AnalyticsService.AnalyticsQuery:output_type -> couchbase.analytics.v1.AnalyticsQueryResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_couchbase_analytics_v1_analytics_proto_init() }
func file_couchbase_analytics_v1_analytics_proto_init() {
	if File_couchbase_analytics_v1_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_couchbase_analytics_v1_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsQueryRequest); i {
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
		file_couchbase_analytics_v1_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsQueryResponse); i {
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
		file_couchbase_analytics_v1_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsQueryResponse_Metrics); i {
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
		file_couchbase_analytics_v1_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsQueryResponse_MetaData); i {
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
		file_couchbase_analytics_v1_analytics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyticsQueryResponse_MetaData_Warning); i {
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
	file_couchbase_analytics_v1_analytics_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_couchbase_analytics_v1_analytics_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_couchbase_analytics_v1_analytics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_couchbase_analytics_v1_analytics_proto_goTypes,
		DependencyIndexes: file_couchbase_analytics_v1_analytics_proto_depIdxs,
		EnumInfos:         file_couchbase_analytics_v1_analytics_proto_enumTypes,
		MessageInfos:      file_couchbase_analytics_v1_analytics_proto_msgTypes,
	}.Build()
	File_couchbase_analytics_v1_analytics_proto = out.File
	file_couchbase_analytics_v1_analytics_proto_rawDesc = nil
	file_couchbase_analytics_v1_analytics_proto_goTypes = nil
	file_couchbase_analytics_v1_analytics_proto_depIdxs = nil
}
