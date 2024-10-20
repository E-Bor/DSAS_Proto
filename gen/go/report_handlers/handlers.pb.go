// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: report_handlers/handlers.proto

package reportsv1

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

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatasourceName string `protobuf:"bytes,1,opt,name=datasource_name,json=datasourceName,proto3" json:"datasource_name,omitempty"`
	ReportTypeName string `protobuf:"bytes,2,opt,name=report_type_name,json=reportTypeName,proto3" json:"report_type_name,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	mi := &file_report_handlers_handlers_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_handlers_handlers_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_report_handlers_handlers_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetDatasourceName() string {
	if x != nil {
		return x.DatasourceName
	}
	return ""
}

func (x *StartRequest) GetReportTypeName() string {
	if x != nil {
		return x.ReportTypeName
	}
	return ""
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportStartId int64 `protobuf:"varint,1,opt,name=report_start_id,json=reportStartId,proto3" json:"report_start_id,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	mi := &file_report_handlers_handlers_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_handlers_handlers_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_report_handlers_handlers_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetReportStartId() int64 {
	if x != nil {
		return x.ReportStartId
	}
	return 0
}

var File_report_handlers_handlers_proto protoreflect.FileDescriptor

var file_report_handlers_handlers_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x73, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x61, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x32, 0x40, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x36, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x65, 0x2d, 0x62, 0x6f, 0x72,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_handlers_handlers_proto_rawDescOnce sync.Once
	file_report_handlers_handlers_proto_rawDescData = file_report_handlers_handlers_proto_rawDesc
)

func file_report_handlers_handlers_proto_rawDescGZIP() []byte {
	file_report_handlers_handlers_proto_rawDescOnce.Do(func() {
		file_report_handlers_handlers_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_handlers_handlers_proto_rawDescData)
	})
	return file_report_handlers_handlers_proto_rawDescData
}

var file_report_handlers_handlers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_report_handlers_handlers_proto_goTypes = []any{
	(*StartRequest)(nil),  // 0: reports.StartRequest
	(*StartResponse)(nil), // 1: reports.StartResponse
}
var file_report_handlers_handlers_proto_depIdxs = []int32{
	0, // 0: reports.Report.Start:input_type -> reports.StartRequest
	1, // 1: reports.Report.Start:output_type -> reports.StartResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_report_handlers_handlers_proto_init() }
func file_report_handlers_handlers_proto_init() {
	if File_report_handlers_handlers_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_report_handlers_handlers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_handlers_handlers_proto_goTypes,
		DependencyIndexes: file_report_handlers_handlers_proto_depIdxs,
		MessageInfos:      file_report_handlers_handlers_proto_msgTypes,
	}.Build()
	File_report_handlers_handlers_proto = out.File
	file_report_handlers_handlers_proto_rawDesc = nil
	file_report_handlers_handlers_proto_goTypes = nil
	file_report_handlers_handlers_proto_depIdxs = nil
}
