// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gateway/api/v1/matchable_tcp_gateway.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	ssl "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/ssl"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A MatchableTcpGateway describes a single FilterChain configured with the TcpProxy network filter and a matcher.
//
// A Gateway CR may select one or more MatchableTcpGateways on a single listener.
// This enables separate teams to own Listener configuration (Gateway CR)
// and FilterChain configuration (MatchableTcpGateway CR).
type MatchableTcpGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NamespacedStatuses indicates the validation status of this resource.
	// NamespacedStatuses is read-only by clients, and set by gateway during validation
	NamespacedStatuses *core.NamespacedStatuses `protobuf:"bytes,1,opt,name=namespaced_statuses,json=namespacedStatuses,proto3" json:"namespaced_statuses,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata *core.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Matcher creates a FilterChainMatch and TransportSocket for a FilterChain
	// For each MatchableTcpGateway on a Gateway CR, the matcher must be unique.
	// If there are any identical matchers, the Gateway will be rejected.
	// An empty matcher will produce an empty FilterChainMatch (https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#envoy-v3-api-msg-config-listener-v3-filterchainmatch)
	// effectively matching all incoming connections
	Matcher *MatchableTcpGateway_Matcher `protobuf:"bytes,3,opt,name=matcher,proto3" json:"matcher,omitempty"`
	// TcpGateway creates a FilterChain with a TcpProxy.
	TcpGateway *TcpGateway `protobuf:"bytes,4,opt,name=tcp_gateway,json=tcpGateway,proto3" json:"tcp_gateway,omitempty"`
}

func (x *MatchableTcpGateway) Reset() {
	*x = MatchableTcpGateway{}
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchableTcpGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableTcpGateway) ProtoMessage() {}

func (x *MatchableTcpGateway) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableTcpGateway.ProtoReflect.Descriptor instead.
func (*MatchableTcpGateway) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *MatchableTcpGateway) GetNamespacedStatuses() *core.NamespacedStatuses {
	if x != nil {
		return x.NamespacedStatuses
	}
	return nil
}

func (x *MatchableTcpGateway) GetMetadata() *core.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MatchableTcpGateway) GetMatcher() *MatchableTcpGateway_Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *MatchableTcpGateway) GetTcpGateway() *TcpGateway {
	if x != nil {
		return x.TcpGateway
	}
	return nil
}

type MatchableTcpGateway_Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CidrRange specifies an IP Address and a prefix length to construct the subnet mask for a CIDR range.
	// See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/address.proto#envoy-v3-api-msg-config-core-v3-cidrrange
	SourcePrefixRanges []*v3.CidrRange `protobuf:"bytes,1,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// Ssl configuration applied to the FilterChain, if using passthrough should not include secrets :
	//   - FilterChainMatch: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener_components.proto#config-listener-v3-filterchainmatch)
	//   - TransportSocket: https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/base.proto#envoy-v3-api-msg-config-core-v3-transportsocket
	SslConfig *ssl.SslConfig `protobuf:"bytes,2,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Enterprise-only: Passthrough cipher suites is an allow-list of OpenSSL cipher suite names for which TLS passthrough will be enabled.
	// If a client does not support any ciphers that are natively supported by Envoy, but does support one of the ciphers in the passthrough list,
	// then traffic will be routed via TCP Proxy to a destination specified by the TcpGateway, where TLS can then be terminated.
	PassthroughCipherSuites []string `protobuf:"bytes,3,rep,name=passthrough_cipher_suites,json=passthroughCipherSuites,proto3" json:"passthrough_cipher_suites,omitempty"`
}

func (x *MatchableTcpGateway_Matcher) Reset() {
	*x = MatchableTcpGateway_Matcher{}
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchableTcpGateway_Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchableTcpGateway_Matcher) ProtoMessage() {}

func (x *MatchableTcpGateway_Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchableTcpGateway_Matcher.ProtoReflect.Descriptor instead.
func (*MatchableTcpGateway_Matcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MatchableTcpGateway_Matcher) GetSourcePrefixRanges() []*v3.CidrRange {
	if x != nil {
		return x.SourcePrefixRanges
	}
	return nil
}

func (x *MatchableTcpGateway_Matcher) GetSslConfig() *ssl.SslConfig {
	if x != nil {
		return x.SslConfig
	}
	return nil
}

func (x *MatchableTcpGateway_Matcher) GetPassthroughCipherSuites() []string {
	if x != nil {
		return x.PassthroughCipherSuites
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x63, 0x70, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b,
	0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x6c, 0x2f, 0x73, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x13, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x57, 0x0a, 0x13, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x42, 0x04, 0xb8, 0xf5, 0x04, 0x01, 0x52, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46,
	0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x63, 0x70, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x63, 0x70, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x63,
	0x70, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x0a, 0x74, 0x63, 0x70, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x1a, 0xd8, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x59, 0x0a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69,
	0x64, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x73,
	0x73, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53,
	0x73, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x73, 0x73, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f, 0x75,
	0x67, 0x68, 0x5f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x70, 0x61, 0x73, 0x73, 0x74, 0x68, 0x72, 0x6f,
	0x75, 0x67, 0x68, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x3a,
	0x17, 0x82, 0xf1, 0x04, 0x13, 0x0a, 0x03, 0x74, 0x67, 0x77, 0x12, 0x0c, 0x74, 0x63, 0x70, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x42, 0x41, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_goTypes = []any{
	(*MatchableTcpGateway)(nil),         // 0: gateway.solo.io.MatchableTcpGateway
	(*MatchableTcpGateway_Matcher)(nil), // 1: gateway.solo.io.MatchableTcpGateway.Matcher
	(*core.NamespacedStatuses)(nil),     // 2: core.solo.io.NamespacedStatuses
	(*core.Metadata)(nil),               // 3: core.solo.io.Metadata
	(*TcpGateway)(nil),                  // 4: gateway.solo.io.TcpGateway
	(*v3.CidrRange)(nil),                // 5: solo.io.envoy.config.core.v3.CidrRange
	(*ssl.SslConfig)(nil),               // 6: gloo.solo.io.SslConfig
}
var file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_depIdxs = []int32{
	2, // 0: gateway.solo.io.MatchableTcpGateway.namespaced_statuses:type_name -> core.solo.io.NamespacedStatuses
	3, // 1: gateway.solo.io.MatchableTcpGateway.metadata:type_name -> core.solo.io.Metadata
	1, // 2: gateway.solo.io.MatchableTcpGateway.matcher:type_name -> gateway.solo.io.MatchableTcpGateway.Matcher
	4, // 3: gateway.solo.io.MatchableTcpGateway.tcp_gateway:type_name -> gateway.solo.io.TcpGateway
	5, // 4: gateway.solo.io.MatchableTcpGateway.Matcher.source_prefix_ranges:type_name -> solo.io.envoy.config.core.v3.CidrRange
	6, // 5: gateway.solo.io.MatchableTcpGateway.Matcher.ssl_config:type_name -> gloo.solo.io.SslConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_init() }
func file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gateway_api_v1_gateway_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway_api_v1_matchable_tcp_gateway_proto_depIdxs = nil
}
