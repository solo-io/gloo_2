// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/stateful_session/stateful_session.proto

package stateful_session

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This configures the Envoy [Stateful Session](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/stateful_session_filter) filter for a listener
type StatefulSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SessionState:
	//
	//	*StatefulSession_CookieBased
	//	*StatefulSession_HeaderBased
	SessionState isStatefulSession_SessionState `protobuf_oneof:"SessionState"`
	// If set to True, the HTTP request must be routed to the requested destination. If the requested destination is not available, Envoy returns 503. Defaults to False.
	Strict bool `protobuf:"varint,3,opt,name=strict,proto3" json:"strict,omitempty"`
}

func (x *StatefulSession) Reset() {
	*x = StatefulSession{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatefulSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSession) ProtoMessage() {}

func (x *StatefulSession) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSession.ProtoReflect.Descriptor instead.
func (*StatefulSession) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{0}
}

func (m *StatefulSession) GetSessionState() isStatefulSession_SessionState {
	if m != nil {
		return m.SessionState
	}
	return nil
}

func (x *StatefulSession) GetCookieBased() *CookieBasedSessionState {
	if x, ok := x.GetSessionState().(*StatefulSession_CookieBased); ok {
		return x.CookieBased
	}
	return nil
}

func (x *StatefulSession) GetHeaderBased() *HeaderBasedSessionState {
	if x, ok := x.GetSessionState().(*StatefulSession_HeaderBased); ok {
		return x.HeaderBased
	}
	return nil
}

func (x *StatefulSession) GetStrict() bool {
	if x != nil {
		return x.Strict
	}
	return false
}

type isStatefulSession_SessionState interface {
	isStatefulSession_SessionState()
}

type StatefulSession_CookieBased struct {
	// Configure a cookie based session state - https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/cookie/v3/cookie.proto#envoy-v3-api-msg-extensions-http-stateful-session-cookie-v3-cookiebasedsessionstate
	// Exactly one of `cookie_based` or `header_based` must be set.
	CookieBased *CookieBasedSessionState `protobuf:"bytes,1,opt,name=cookie_based,json=cookieBased,proto3,oneof"`
}

type StatefulSession_HeaderBased struct {
	// Configure a header based session state - https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/cookie/v3/cookie.proto#envoy-v3-api-msg-extensions-http-stateful-session-cookie-v3-cookiebasedsessionstate
	// Exactly one of `cookie_based` or `header_based` must be set.
	HeaderBased *HeaderBasedSessionState `protobuf:"bytes,2,opt,name=header_based,json=headerBased,proto3,oneof"`
}

func (*StatefulSession_CookieBased) isStatefulSession_SessionState() {}

func (*StatefulSession_HeaderBased) isStatefulSession_SessionState() {}

// Configuration for the [cookie-based session state](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/cookie/v3/cookie.proto#envoy-v3-api-msg-extensions-http-stateful-session-cookie-v3-cookiebasedsessionstate) filter
type CookieBasedSessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required, the cookie configuration used to track session state.
	Cookie *CookieBasedSessionState_Cookie `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
}

func (x *CookieBasedSessionState) Reset() {
	*x = CookieBasedSessionState{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CookieBasedSessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookieBasedSessionState) ProtoMessage() {}

func (x *CookieBasedSessionState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookieBasedSessionState.ProtoReflect.Descriptor instead.
func (*CookieBasedSessionState) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{1}
}

func (x *CookieBasedSessionState) GetCookie() *CookieBasedSessionState_Cookie {
	if x != nil {
		return x.Cookie
	}
	return nil
}

// Configuration for the [header-based session state](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/header/v3/header.proto#extension-envoy-http-stateful-session-header) filter
type HeaderBasedSessionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required, the header used to track session state.
	HeaderName string `protobuf:"bytes,1,opt,name=header_name,json=headerName,proto3" json:"header_name,omitempty"`
}

func (x *HeaderBasedSessionState) Reset() {
	*x = HeaderBasedSessionState{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeaderBasedSessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderBasedSessionState) ProtoMessage() {}

func (x *HeaderBasedSessionState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderBasedSessionState.ProtoReflect.Descriptor instead.
func (*HeaderBasedSessionState) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderBasedSessionState) GetHeaderName() string {
	if x != nil {
		return x.HeaderName
	}
	return ""
}

type CookieBasedSessionState_Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required, the name that will be used to obtain cookie value from downstream HTTP request or generate new cookie for downstream.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Path of cookie. This will be used to set the path of a new cookie when it is generated. If no path is specified here, no path will be set for the cookie.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Duration of cookie. This will be used to set the expiry time of a new cookie when it is generated. Set this to 0s to use a session cookie and disable cookie expiration.
	Ttl *duration.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *CookieBasedSessionState_Cookie) Reset() {
	*x = CookieBasedSessionState_Cookie{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CookieBasedSessionState_Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookieBasedSessionState_Cookie) ProtoMessage() {}

func (x *CookieBasedSessionState_Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookieBasedSessionState_Cookie.ProtoReflect.Descriptor instead.
func (*CookieBasedSessionState_Cookie) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CookieBasedSessionState_Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CookieBasedSessionState_Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CookieBasedSessionState_Cookie) GetTtl() *duration.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc = []byte{
	0x0a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x0f, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x63, 0x0a,
	0x0c, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42, 0x61, 0x73,
	0x65, 0x64, 0x12, 0x63, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x42,
	0x13, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x22, 0xf7, 0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x67, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x1a, 0x73, 0x0a, 0x06, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0xaa, 0x01, 0x04, 0x08, 0x01, 0x32, 0x00, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x44,
	0x0a, 0x17, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x64, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x66, 0x75, 0x6c, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_goTypes = []any{
	(*StatefulSession)(nil),                // 0: stateful_session.options.gloo.solo.io.StatefulSession
	(*CookieBasedSessionState)(nil),        // 1: stateful_session.options.gloo.solo.io.CookieBasedSessionState
	(*HeaderBasedSessionState)(nil),        // 2: stateful_session.options.gloo.solo.io.HeaderBasedSessionState
	(*CookieBasedSessionState_Cookie)(nil), // 3: stateful_session.options.gloo.solo.io.CookieBasedSessionState.Cookie
	(*duration.Duration)(nil),              // 4: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_depIdxs = []int32{
	1, // 0: stateful_session.options.gloo.solo.io.StatefulSession.cookie_based:type_name -> stateful_session.options.gloo.solo.io.CookieBasedSessionState
	2, // 1: stateful_session.options.gloo.solo.io.StatefulSession.header_based:type_name -> stateful_session.options.gloo.solo.io.HeaderBasedSessionState
	3, // 2: stateful_session.options.gloo.solo.io.CookieBasedSessionState.cookie:type_name -> stateful_session.options.gloo.solo.io.CookieBasedSessionState.Cookie
	4, // 3: stateful_session.options.gloo.solo.io.CookieBasedSessionState.Cookie.ttl:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[0].OneofWrappers = []any{
		(*StatefulSession_CookieBased)(nil),
		(*StatefulSession_HeaderBased)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_depIdxs = nil
}
