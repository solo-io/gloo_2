// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/als/als.proto

package als

import (
	reflect "reflect"
	sync "sync"

	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Contains various settings for Envoy's access logging service.
// See here for more information: https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/accesslog/v2/accesslog.proto#envoy-api-msg-config-filter-accesslog-v2-accesslog
type AccessLoggingService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessLog []*AccessLog `protobuf:"bytes,1,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
}

func (x *AccessLoggingService) Reset() {
	*x = AccessLoggingService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLoggingService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLoggingService) ProtoMessage() {}

func (x *AccessLoggingService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLoggingService.ProtoReflect.Descriptor instead.
func (*AccessLoggingService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescGZIP(), []int{0}
}

func (x *AccessLoggingService) GetAccessLog() []*AccessLog {
	if x != nil {
		return x.AccessLog
	}
	return nil
}

type AccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// type of Access Logging service to implement
	//
	// Types that are assignable to OutputDestination:
	//
	//	*AccessLog_FileSink
	//	*AccessLog_GrpcService
	OutputDestination isAccessLog_OutputDestination `protobuf_oneof:"OutputDestination"`
}

func (x *AccessLog) Reset() {
	*x = AccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessLog) ProtoMessage() {}

func (x *AccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessLog.ProtoReflect.Descriptor instead.
func (*AccessLog) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescGZIP(), []int{1}
}

func (m *AccessLog) GetOutputDestination() isAccessLog_OutputDestination {
	if m != nil {
		return m.OutputDestination
	}
	return nil
}

func (x *AccessLog) GetFileSink() *FileSink {
	if x, ok := x.GetOutputDestination().(*AccessLog_FileSink); ok {
		return x.FileSink
	}
	return nil
}

func (x *AccessLog) GetGrpcService() *GrpcService {
	if x, ok := x.GetOutputDestination().(*AccessLog_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

type isAccessLog_OutputDestination interface {
	isAccessLog_OutputDestination()
}

type AccessLog_FileSink struct {
	// Output access logs to local file
	FileSink *FileSink `protobuf:"bytes,2,opt,name=file_sink,json=fileSink,proto3,oneof"`
}

type AccessLog_GrpcService struct {
	// Send access logs to gRPC service
	GrpcService *GrpcService `protobuf:"bytes,3,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*AccessLog_FileSink) isAccessLog_OutputDestination() {}

func (*AccessLog_GrpcService) isAccessLog_OutputDestination() {}

type FileSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the file path to which the file access logging service will sink
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// the format which the logs should be outputted by
	//
	// Types that are assignable to OutputFormat:
	//
	//	*FileSink_StringFormat
	//	*FileSink_JsonFormat
	OutputFormat isFileSink_OutputFormat `protobuf_oneof:"output_format"`
}

func (x *FileSink) Reset() {
	*x = FileSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSink) ProtoMessage() {}

func (x *FileSink) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSink.ProtoReflect.Descriptor instead.
func (*FileSink) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescGZIP(), []int{2}
}

func (x *FileSink) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (m *FileSink) GetOutputFormat() isFileSink_OutputFormat {
	if m != nil {
		return m.OutputFormat
	}
	return nil
}

func (x *FileSink) GetStringFormat() string {
	if x, ok := x.GetOutputFormat().(*FileSink_StringFormat); ok {
		return x.StringFormat
	}
	return ""
}

func (x *FileSink) GetJsonFormat() *_struct.Struct {
	if x, ok := x.GetOutputFormat().(*FileSink_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

type isFileSink_OutputFormat interface {
	isFileSink_OutputFormat()
}

type FileSink_StringFormat struct {
	// the format string by which envoy will format the log lines
	// https://www.envoyproxy.io/docs/envoy/v1.14.1/configuration/observability/access_log#config-access-log-format-strings
	StringFormat string `protobuf:"bytes,2,opt,name=string_format,json=stringFormat,proto3,oneof"`
}

type FileSink_JsonFormat struct {
	// the format object by which to envoy will emit the logs in a structured way.
	// https://www.envoyproxy.io/docs/envoy/v1.14.1/configuration/observability/access_log#format-dictionaries
	JsonFormat *_struct.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof"`
}

func (*FileSink_StringFormat) isFileSink_OutputFormat() {}

func (*FileSink_JsonFormat) isFileSink_OutputFormat() {}

type GrpcService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of log stream
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The static cluster defined in bootstrap config to route to
	//
	// Types that are assignable to ServiceRef:
	//
	//	*GrpcService_StaticClusterName
	ServiceRef                      isGrpcService_ServiceRef `protobuf_oneof:"service_ref"`
	AdditionalRequestHeadersToLog   []string                 `protobuf:"bytes,4,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	AdditionalResponseHeadersToLog  []string                 `protobuf:"bytes,5,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	AdditionalResponseTrailersToLog []string                 `protobuf:"bytes,6,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
}

func (x *GrpcService) Reset() {
	*x = GrpcService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcService) ProtoMessage() {}

func (x *GrpcService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcService.ProtoReflect.Descriptor instead.
func (*GrpcService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescGZIP(), []int{3}
}

func (x *GrpcService) GetLogName() string {
	if x != nil {
		return x.LogName
	}
	return ""
}

func (m *GrpcService) GetServiceRef() isGrpcService_ServiceRef {
	if m != nil {
		return m.ServiceRef
	}
	return nil
}

func (x *GrpcService) GetStaticClusterName() string {
	if x, ok := x.GetServiceRef().(*GrpcService_StaticClusterName); ok {
		return x.StaticClusterName
	}
	return ""
}

func (x *GrpcService) GetAdditionalRequestHeadersToLog() []string {
	if x != nil {
		return x.AdditionalRequestHeadersToLog
	}
	return nil
}

func (x *GrpcService) GetAdditionalResponseHeadersToLog() []string {
	if x != nil {
		return x.AdditionalResponseHeadersToLog
	}
	return nil
}

func (x *GrpcService) GetAdditionalResponseTrailersToLog() []string {
	if x != nil {
		return x.AdditionalResponseTrailersToLog
	}
	return nil
}

type isGrpcService_ServiceRef interface {
	isGrpcService_ServiceRef()
}

type GrpcService_StaticClusterName struct {
	StaticClusterName string `protobuf:"bytes,2,opt,name=static_cluster_name,json=staticClusterName,proto3,oneof"`
}

func (*GrpcService_StaticClusterName) isGrpcService_ServiceRef() {}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6c, 0x73, 0x2f, 0x61, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x6c, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a,
	0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6c, 0x73,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x22, 0xaf, 0x01, 0x0a, 0x09, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x41, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x6c,
	0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x48,
	0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x4a, 0x0a, 0x0c, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a,
	0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x42, 0x0f, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x22, 0xcd, 0x02, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48,
	0x0a, 0x21, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x12, 0x4a, 0x0a, 0x22, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x1e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54,
	0x6f, 0x4c, 0x6f, 0x67, 0x12, 0x4c, 0x0a, 0x23, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x1f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x4c,
	0x6f, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x42, 0x4a, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6c,
	0x73, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_goTypes = []interface{}{
	(*AccessLoggingService)(nil), // 0: als.options.gloo.solo.io.AccessLoggingService
	(*AccessLog)(nil),            // 1: als.options.gloo.solo.io.AccessLog
	(*FileSink)(nil),             // 2: als.options.gloo.solo.io.FileSink
	(*GrpcService)(nil),          // 3: als.options.gloo.solo.io.GrpcService
	(*_struct.Struct)(nil),       // 4: google.protobuf.Struct
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_depIdxs = []int32{
	1, // 0: als.options.gloo.solo.io.AccessLoggingService.access_log:type_name -> als.options.gloo.solo.io.AccessLog
	2, // 1: als.options.gloo.solo.io.AccessLog.file_sink:type_name -> als.options.gloo.solo.io.FileSink
	3, // 2: als.options.gloo.solo.io.AccessLog.grpc_service:type_name -> als.options.gloo.solo.io.GrpcService
	4, // 3: als.options.gloo.solo.io.FileSink.json_format:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLoggingService); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessLog); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSink); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcService); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AccessLog_FileSink)(nil),
		(*AccessLog_GrpcService)(nil),
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FileSink_StringFormat)(nil),
		(*FileSink_JsonFormat)(nil),
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GrpcService_StaticClusterName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_als_als_proto_depIdxs = nil
}
