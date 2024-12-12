// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/proxylatency/proxylatency.proto

package proxylatency

import (
	reflect "reflect"
	sync "sync"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// How to perform the latency measurement. Given an incoming request from downstream and
// outging request to upstream; or incoming response from upstream and outgoing repsonse to
// downstream, This outlines how to measure the latency used by the proxy.
type ProxyLatency_Measurement int32

const (
	// Count from the last byte of the incoming request\response to the first byte of the outgoing request\response.
	ProxyLatency_LAST_INCOMING_FIRST_OUTGOING ProxyLatency_Measurement = 0
	// Count from the first byte of the incoming request\response to the first byte of the outgoing request\response.
	ProxyLatency_FIRST_INCOMING_FIRST_OUTGOING ProxyLatency_Measurement = 1
	// Count from the last byte of the incoming request\response to the last byte of the outgoing request\response.
	ProxyLatency_LAST_INCOMING_LAST_OUTGOING ProxyLatency_Measurement = 2
	// Count from the first byte of the incoming request\response to the last byte of the outgoing request\response.
	ProxyLatency_FIRST_INCOMING_LAST_OUTGOING ProxyLatency_Measurement = 3
)

// Enum value maps for ProxyLatency_Measurement.
var (
	ProxyLatency_Measurement_name = map[int32]string{
		0: "LAST_INCOMING_FIRST_OUTGOING",
		1: "FIRST_INCOMING_FIRST_OUTGOING",
		2: "LAST_INCOMING_LAST_OUTGOING",
		3: "FIRST_INCOMING_LAST_OUTGOING",
	}
	ProxyLatency_Measurement_value = map[string]int32{
		"LAST_INCOMING_FIRST_OUTGOING":  0,
		"FIRST_INCOMING_FIRST_OUTGOING": 1,
		"LAST_INCOMING_LAST_OUTGOING":   2,
		"FIRST_INCOMING_LAST_OUTGOING":  3,
	}
)

func (x ProxyLatency_Measurement) Enum() *ProxyLatency_Measurement {
	p := new(ProxyLatency_Measurement)
	*p = x
	return p
}

func (x ProxyLatency_Measurement) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyLatency_Measurement) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_enumTypes[0].Descriptor()
}

func (ProxyLatency_Measurement) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_enumTypes[0]
}

func (x ProxyLatency_Measurement) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyLatency_Measurement.Descriptor instead.
func (ProxyLatency_Measurement) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescGZIP(), []int{0, 0}
}

// Configure the proxy latency filter. This filter measures the latency
// incurred by the filter chain in a histogram.
type ProxyLatency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// How to measure the request.
	Request ProxyLatency_Measurement `protobuf:"varint,1,opt,name=request,proto3,enum=envoy.config.filter.http.proxylatency.v2.ProxyLatency_Measurement" json:"request,omitempty"`
	// When FIRST_OUTGOING (i.e. LAST_INCOMING_FIRST_OUTGOING or FIRST_INCOMING_FIRST_OUTGOING) is
	// instead of when the first byte is sent upstream. This has the advantage of not measuring the time
	// selected for request measurment, finish measuring proxy latency when decodeHeader for this
	// it takes a connection to form, which may skew the P99.
	// filter is hit instead of when the first byte is sent upstream. This has the advantage of not
	// for this to work the filter should be inserted last, just before the router filter.
	// measuring the time it takes a connection to form, which may skew the P99. For this to work
	// this filter should be inserted last, just before the router filter. This has no effect if
	// other measurement type is selected, and has no effect on how response is measured.
	MeasureRequestInternally bool `protobuf:"varint,5,opt,name=measure_request_internally,json=measureRequestInternally,proto3" json:"measure_request_internally,omitempty"`
	// How measure the response.
	Response ProxyLatency_Measurement `protobuf:"varint,2,opt,name=response,proto3,enum=envoy.config.filter.http.proxylatency.v2.ProxyLatency_Measurement" json:"response,omitempty"`
	// Charge a stat per upstream cluster. If not specified, defaults to true.
	ChargeClusterStat *wrappers.BoolValue `protobuf:"bytes,3,opt,name=charge_cluster_stat,json=chargeClusterStat,proto3" json:"charge_cluster_stat,omitempty"`
	// Charge a stat per listener. If not specified, defaults to true.
	ChargeListenerStat *wrappers.BoolValue `protobuf:"bytes,4,opt,name=charge_listener_stat,json=chargeListenerStat,proto3" json:"charge_listener_stat,omitempty"`
	// Should we emit request timing to dynamic metadata. defaults to true.
	EmitDynamicMetadata *wrappers.BoolValue `protobuf:"bytes,6,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
}

func (x *ProxyLatency) Reset() {
	*x = ProxyLatency{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProxyLatency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyLatency) ProtoMessage() {}

func (x *ProxyLatency) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyLatency.ProtoReflect.Descriptor instead.
func (*ProxyLatency) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyLatency) GetRequest() ProxyLatency_Measurement {
	if x != nil {
		return x.Request
	}
	return ProxyLatency_LAST_INCOMING_FIRST_OUTGOING
}

func (x *ProxyLatency) GetMeasureRequestInternally() bool {
	if x != nil {
		return x.MeasureRequestInternally
	}
	return false
}

func (x *ProxyLatency) GetResponse() ProxyLatency_Measurement {
	if x != nil {
		return x.Response
	}
	return ProxyLatency_LAST_INCOMING_FIRST_OUTGOING
}

func (x *ProxyLatency) GetChargeClusterStat() *wrappers.BoolValue {
	if x != nil {
		return x.ChargeClusterStat
	}
	return nil
}

func (x *ProxyLatency) GetChargeListenerStat() *wrappers.BoolValue {
	if x != nil {
		return x.ChargeListenerStat
	}
	return nil
}

func (x *ProxyLatency) GetEmitDynamicMetadata() *wrappers.BoolValue {
	if x != nil {
		return x.EmitDynamicMetadata
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDesc = []byte{
	0x0a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x05, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x5c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x1a, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x6c, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x6c, 0x79,
	0x12, 0x5e, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x4c, 0x0a, 0x14,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x4e, 0x0a, 0x15, 0x65, 0x6d,
	0x69, 0x74, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x65, 0x6d, 0x69, 0x74, 0x44, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x0b, 0x4d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x1c, 0x4c, 0x41,
	0x53, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x52, 0x53,
	0x54, 0x5f, 0x4f, 0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d,
	0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x46,
	0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x1f, 0x0a, 0x1b, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47,
	0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x20, 0x0a, 0x1c, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x49,
	0x4e, 0x47, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x4f, 0x55, 0x54, 0x47, 0x4f, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x42, 0x62, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01,
	0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6c,
	0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_goTypes = []any{
	(ProxyLatency_Measurement)(0), // 0: envoy.config.filter.http.proxylatency.v2.ProxyLatency.Measurement
	(*ProxyLatency)(nil),          // 1: envoy.config.filter.http.proxylatency.v2.ProxyLatency
	(*wrappers.BoolValue)(nil),    // 2: google.protobuf.BoolValue
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_depIdxs = []int32{
	0, // 0: envoy.config.filter.http.proxylatency.v2.ProxyLatency.request:type_name -> envoy.config.filter.http.proxylatency.v2.ProxyLatency.Measurement
	0, // 1: envoy.config.filter.http.proxylatency.v2.ProxyLatency.response:type_name -> envoy.config.filter.http.proxylatency.v2.ProxyLatency.Measurement
	2, // 2: envoy.config.filter.http.proxylatency.v2.ProxyLatency.charge_cluster_stat:type_name -> google.protobuf.BoolValue
	2, // 3: envoy.config.filter.http.proxylatency.v2.ProxyLatency.charge_listener_stat:type_name -> google.protobuf.BoolValue
	2, // 4: envoy.config.filter.http.proxylatency.v2.ProxyLatency.emit_dynamic_metadata:type_name -> google.protobuf.BoolValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_proxylatency_proxylatency_proto_depIdxs = nil
}
