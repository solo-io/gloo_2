// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/matching/custom_matchers/server_name/v3/server_name_matcher.proto

package v3

import (
	reflect "reflect"
	sync "sync"

	v3 "github.com/cncf/xds/go/xds/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Matches a specific server name provided in the client request against a set server names configured for the matcher to handle, with possible prefix wildcard.
type ServerNameMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Match server names. Order doesn't matter, the most specific server name is matched.
	ServerNameMatchers []*ServerNameMatcher_ServerNameSetMatcher `protobuf:"bytes,1,rep,name=server_name_matchers,json=serverNameMatchers,proto3" json:"server_name_matchers,omitempty"`
}

func (x *ServerNameMatcher) Reset() {
	*x = ServerNameMatcher{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerNameMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNameMatcher) ProtoMessage() {}

func (x *ServerNameMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNameMatcher.ProtoReflect.Descriptor instead.
func (*ServerNameMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescGZIP(), []int{0}
}

func (x *ServerNameMatcher) GetServerNameMatchers() []*ServerNameMatcher_ServerNameSetMatcher {
	if x != nil {
		return x.ServerNameMatchers
	}
	return nil
}

// Specifies a list of server names and a match action.
type ServerNameMatcher_ServerNameSetMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A non-empty set of server names.
	// Server name can start with a wildcard prefix, e.g. "*.example.com".
	ServerNames []string `protobuf:"bytes,1,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	// Match action to apply when the input matches the server name.
	OnMatch *v3.Matcher_OnMatch `protobuf:"bytes,2,opt,name=on_match,json=onMatch,proto3" json:"on_match,omitempty"`
}

func (x *ServerNameMatcher_ServerNameSetMatcher) Reset() {
	*x = ServerNameMatcher_ServerNameSetMatcher{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerNameMatcher_ServerNameSetMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNameMatcher_ServerNameSetMatcher) ProtoMessage() {}

func (x *ServerNameMatcher_ServerNameSetMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNameMatcher_ServerNameSetMatcher.ProtoReflect.Descriptor instead.
func (*ServerNameMatcher_ServerNameSetMatcher) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerNameMatcher_ServerNameSetMatcher) GetServerNames() []string {
	if x != nil {
		return x.ServerNames
	}
	return nil
}

func (x *ServerNameMatcher_ServerNameSetMatcher) GetOnMatch() *v3.Matcher_OnMatch {
	if x != nil {
		return x.OnMatch
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDesc = []byte{
	0x0a, 0x81, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x11, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x8e, 0x01, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x5c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x12, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x73, 0x1a, 0x84, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x6f, 0x6e, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x78, 0x64, 0x73, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x07, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0xd0, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0,
	0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_goTypes = []any{
	(*ServerNameMatcher)(nil),                      // 0: envoy.config.matching.custom_matchers.server_name.v3.ServerNameMatcher
	(*ServerNameMatcher_ServerNameSetMatcher)(nil), // 1: envoy.config.matching.custom_matchers.server_name.v3.ServerNameMatcher.ServerNameSetMatcher
	(*v3.Matcher_OnMatch)(nil),                     // 2: xds.type.matcher.v3.Matcher.OnMatch
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_depIdxs = []int32{
	1, // 0: envoy.config.matching.custom_matchers.server_name.v3.ServerNameMatcher.server_name_matchers:type_name -> envoy.config.matching.custom_matchers.server_name.v3.ServerNameMatcher.ServerNameSetMatcher
	2, // 1: envoy.config.matching.custom_matchers.server_name.v3.ServerNameMatcher.ServerNameSetMatcher.on_match:type_name -> xds.type.matcher.v3.Matcher.OnMatch
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_matching_custom_matchers_server_name_v3_server_name_matcher_proto_depIdxs = nil
}
