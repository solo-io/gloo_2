// copied from https://github.com/envoyproxy/envoy/blob/ad89a587aa0177bfdad6b5c968a6aead5d9be7a4/api/envoy/config/common/mutation_rules/v3/mutation_rules.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/common/mutation_rules/v3/mutation_rules.proto

// manually updated pkg

package v3

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	v31 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/core/v3"
	v3 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/type/matcher/v3"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
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

// The HeaderMutationRules structure specifies what headers may be
// manipulated by a processing filter. This set of rules makes it
// possible to control which modifications a filter may make.
//
// By default, an external processing server may add, modify, or remove
// any header except for an "Envoy internal" header (which is typically
// denoted by an x-envoy prefix) or specific headers that may affect
// further filter processing:
//
// * “host“
// * “:authority“
// * “:scheme“
// * “:method“
//
// Every attempt to add, change, append, or remove a header will be
// tested against the rules here. Disallowed header mutations will be
// ignored unless “disallow_is_error“ is set to true.
//
// Attempts to remove headers are further constrained -- regardless of the
// settings, system-defined headers (that start with “:“) and the “host“
// header may never be removed.
//
// In addition, a counter will be incremented whenever a mutation is
// rejected. In the ext_proc filter, that counter is named
// “rejected_header_mutations“.
// [#next-free-field: 8]
type HeaderMutationRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// By default, certain headers that could affect processing of subsequent
	// filters or request routing cannot be modified. These headers are
	// “host“, “:authority“, “:scheme“, and “:method“. Setting this parameter
	// to true allows these headers to be modified as well.
	AllowAllRouting *wrappers.BoolValue `protobuf:"bytes,1,opt,name=allow_all_routing,json=allowAllRouting,proto3" json:"allow_all_routing,omitempty"`
	// If true, allow modification of envoy internal headers. By default, these
	// start with “x-envoy“ but this may be overridden in the “Bootstrap“
	// configuration using the
	// :ref:`header_prefix <envoy_v3_api_field_config.bootstrap.v3.Bootstrap.header_prefix>`
	// field. Default is false.
	AllowEnvoy *wrappers.BoolValue `protobuf:"bytes,2,opt,name=allow_envoy,json=allowEnvoy,proto3" json:"allow_envoy,omitempty"`
	// If true, prevent modification of any system header, defined as a header
	// that starts with a “:“ character, regardless of any other settings.
	// A processing server may still override the “:status“ of an HTTP response
	// using an “ImmediateResponse“ message. Default is false.
	DisallowSystem *wrappers.BoolValue `protobuf:"bytes,3,opt,name=disallow_system,json=disallowSystem,proto3" json:"disallow_system,omitempty"`
	// If true, prevent modifications of all header values, regardless of any
	// other settings. A processing server may still override the “:status“
	// of an HTTP response using an “ImmediateResponse“ message. Default is false.
	DisallowAll *wrappers.BoolValue `protobuf:"bytes,4,opt,name=disallow_all,json=disallowAll,proto3" json:"disallow_all,omitempty"`
	// If set, specifically allow any header that matches this regular
	// expression. This overrides all other settings except for
	// “disallow_expression“.
	AllowExpression *v3.RegexMatcher `protobuf:"bytes,5,opt,name=allow_expression,json=allowExpression,proto3" json:"allow_expression,omitempty"`
	// If set, specifically disallow any header that matches this regular
	// expression regardless of any other settings.
	DisallowExpression *v3.RegexMatcher `protobuf:"bytes,6,opt,name=disallow_expression,json=disallowExpression,proto3" json:"disallow_expression,omitempty"`
	// If true, and if the rules in this list cause a header mutation to be
	// disallowed, then the filter using this configuration will terminate the
	// request with a 500 error. In addition, regardless of the setting of this
	// parameter, any attempt to set, add, or modify a disallowed header will
	// cause the “rejected_header_mutations“ counter to be incremented.
	// Default is false.
	DisallowIsError *wrappers.BoolValue `protobuf:"bytes,7,opt,name=disallow_is_error,json=disallowIsError,proto3" json:"disallow_is_error,omitempty"`
}

func (x *HeaderMutationRules) Reset() {
	*x = HeaderMutationRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMutationRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMutationRules) ProtoMessage() {}

func (x *HeaderMutationRules) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMutationRules.ProtoReflect.Descriptor instead.
func (*HeaderMutationRules) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderMutationRules) GetAllowAllRouting() *wrappers.BoolValue {
	if x != nil {
		return x.AllowAllRouting
	}
	return nil
}

func (x *HeaderMutationRules) GetAllowEnvoy() *wrappers.BoolValue {
	if x != nil {
		return x.AllowEnvoy
	}
	return nil
}

func (x *HeaderMutationRules) GetDisallowSystem() *wrappers.BoolValue {
	if x != nil {
		return x.DisallowSystem
	}
	return nil
}

func (x *HeaderMutationRules) GetDisallowAll() *wrappers.BoolValue {
	if x != nil {
		return x.DisallowAll
	}
	return nil
}

func (x *HeaderMutationRules) GetAllowExpression() *v3.RegexMatcher {
	if x != nil {
		return x.AllowExpression
	}
	return nil
}

func (x *HeaderMutationRules) GetDisallowExpression() *v3.RegexMatcher {
	if x != nil {
		return x.DisallowExpression
	}
	return nil
}

func (x *HeaderMutationRules) GetDisallowIsError() *wrappers.BoolValue {
	if x != nil {
		return x.DisallowIsError
	}
	return nil
}

// The HeaderMutation structure specifies an action that may be taken on HTTP
// headers.
type HeaderMutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*HeaderMutation_Remove
	//	*HeaderMutation_Append
	Action isHeaderMutation_Action `protobuf_oneof:"action"`
}

func (x *HeaderMutation) Reset() {
	*x = HeaderMutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMutation) ProtoMessage() {}

func (x *HeaderMutation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMutation.ProtoReflect.Descriptor instead.
func (*HeaderMutation) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescGZIP(), []int{1}
}

func (m *HeaderMutation) GetAction() isHeaderMutation_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *HeaderMutation) GetRemove() string {
	if x, ok := x.GetAction().(*HeaderMutation_Remove); ok {
		return x.Remove
	}
	return ""
}

func (x *HeaderMutation) GetAppend() *v31.HeaderValueOption {
	if x, ok := x.GetAction().(*HeaderMutation_Append); ok {
		return x.Append
	}
	return nil
}

type isHeaderMutation_Action interface {
	isHeaderMutation_Action()
}

type HeaderMutation_Remove struct {
	// Remove the specified header if it exists.
	Remove string `protobuf:"bytes,1,opt,name=remove,proto3,oneof"`
}

type HeaderMutation_Append struct {
	// Append new header by the specified HeaderValueOption.
	Append *v31.HeaderValueOption `protobuf:"bytes,2,opt,name=append,proto3,oneof"`
}

func (*HeaderMutation_Remove) isHeaderMutation_Action() {}

func (*HeaderMutation_Append) isHeaderMutation_Action() {}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDesc = []byte{
	0x0a, 0x6d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2d, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x75, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x04, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x11,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x45, 0x6e, 0x76, 0x6f,
	0x79, 0x12, 0x43, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x12, 0x56, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e,
	0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0f, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a,
	0x13, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x12, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x73, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xc8, 0x01, 0x00,
	0xc0, 0x01, 0x02, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x49, 0x0a,
	0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x42, 0x0d, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xbd, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5,
	0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a, 0x33,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x33, 0x42, 0x12, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_goTypes = []interface{}{
	(*HeaderMutationRules)(nil),   // 0: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules
	(*HeaderMutation)(nil),        // 1: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation
	(*wrappers.BoolValue)(nil),    // 2: google.protobuf.BoolValue
	(*v3.RegexMatcher)(nil),       // 3: solo.io.envoy.type.matcher.v3.RegexMatcher
	(*v31.HeaderValueOption)(nil), // 4: solo.io.envoy.config.core.v3.HeaderValueOption
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_depIdxs = []int32{
	2, // 0: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.allow_all_routing:type_name -> google.protobuf.BoolValue
	2, // 1: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.allow_envoy:type_name -> google.protobuf.BoolValue
	2, // 2: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.disallow_system:type_name -> google.protobuf.BoolValue
	2, // 3: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.disallow_all:type_name -> google.protobuf.BoolValue
	3, // 4: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.allow_expression:type_name -> solo.io.envoy.type.matcher.v3.RegexMatcher
	3, // 5: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.disallow_expression:type_name -> solo.io.envoy.type.matcher.v3.RegexMatcher
	2, // 6: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.disallow_is_error:type_name -> google.protobuf.BoolValue
	4, // 7: solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.append:type_name -> solo.io.envoy.config.core.v3.HeaderValueOption
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMutationRules); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMutation); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HeaderMutation_Remove)(nil),
		(*HeaderMutation_Append)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_common_mutation_rules_v3_mutation_rules_proto_depIdxs = nil
}
