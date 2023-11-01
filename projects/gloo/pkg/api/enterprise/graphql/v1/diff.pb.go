// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-projects/projects/gloo/api/enterprise/graphql/v1/diff.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GraphQLInspectorDiffOutput_CriticalityLevel int32

const (
	GraphQLInspectorDiffOutput_NON_BREAKING GraphQLInspectorDiffOutput_CriticalityLevel = 0
	GraphQLInspectorDiffOutput_DANGEROUS    GraphQLInspectorDiffOutput_CriticalityLevel = 1
	GraphQLInspectorDiffOutput_BREAKING     GraphQLInspectorDiffOutput_CriticalityLevel = 2
)

// Enum value maps for GraphQLInspectorDiffOutput_CriticalityLevel.
var (
	GraphQLInspectorDiffOutput_CriticalityLevel_name = map[int32]string{
		0: "NON_BREAKING",
		1: "DANGEROUS",
		2: "BREAKING",
	}
	GraphQLInspectorDiffOutput_CriticalityLevel_value = map[string]int32{
		"NON_BREAKING": 0,
		"DANGEROUS":    1,
		"BREAKING":     2,
	}
)

func (x GraphQLInspectorDiffOutput_CriticalityLevel) Enum() *GraphQLInspectorDiffOutput_CriticalityLevel {
	p := new(GraphQLInspectorDiffOutput_CriticalityLevel)
	*p = x
	return p
}

func (x GraphQLInspectorDiffOutput_CriticalityLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphQLInspectorDiffOutput_CriticalityLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_enumTypes[0].Descriptor()
}

func (GraphQLInspectorDiffOutput_CriticalityLevel) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_enumTypes[0]
}

func (x GraphQLInspectorDiffOutput_CriticalityLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphQLInspectorDiffOutput_CriticalityLevel.Descriptor instead.
func (GraphQLInspectorDiffOutput_CriticalityLevel) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescGZIP(), []int{1, 0}
}

// This message is used to pass data from the control plane / apiserver to the graphql-inspector diff js script.
type GraphQLInspectorDiffInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldSchema string                                                           `protobuf:"bytes,1,opt,name=old_schema,json=oldSchema,proto3" json:"old_schema,omitempty"`
	NewSchema string                                                           `protobuf:"bytes,2,opt,name=new_schema,json=newSchema,proto3" json:"new_schema,omitempty"`
	Rules     []v1.GraphqlOptions_SchemaChangeValidationOptions_ProcessingRule `protobuf:"varint,3,rep,packed,name=rules,proto3,enum=gloo.solo.io.GraphqlOptions_SchemaChangeValidationOptions_ProcessingRule" json:"rules,omitempty"`
}

func (x *GraphQLInspectorDiffInput) Reset() {
	*x = GraphQLInspectorDiffInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLInspectorDiffInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLInspectorDiffInput) ProtoMessage() {}

func (x *GraphQLInspectorDiffInput) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLInspectorDiffInput.ProtoReflect.Descriptor instead.
func (*GraphQLInspectorDiffInput) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescGZIP(), []int{0}
}

func (x *GraphQLInspectorDiffInput) GetOldSchema() string {
	if x != nil {
		return x.OldSchema
	}
	return ""
}

func (x *GraphQLInspectorDiffInput) GetNewSchema() string {
	if x != nil {
		return x.NewSchema
	}
	return ""
}

func (x *GraphQLInspectorDiffInput) GetRules() []v1.GraphqlOptions_SchemaChangeValidationOptions_ProcessingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// This message is used to pass data from the graphql-inspector diff js script back to the control plane / apiserver.
type GraphQLInspectorDiffOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*GraphQLInspectorDiffOutput_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *GraphQLInspectorDiffOutput) Reset() {
	*x = GraphQLInspectorDiffOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLInspectorDiffOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLInspectorDiffOutput) ProtoMessage() {}

func (x *GraphQLInspectorDiffOutput) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLInspectorDiffOutput.ProtoReflect.Descriptor instead.
func (*GraphQLInspectorDiffOutput) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescGZIP(), []int{1}
}

func (x *GraphQLInspectorDiffOutput) GetChanges() []*GraphQLInspectorDiffOutput_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type GraphQLInspectorDiffOutput_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string                                  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Path        string                                  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	ChangeType  string                                  `protobuf:"bytes,3,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	Criticality *GraphQLInspectorDiffOutput_Criticality `protobuf:"bytes,4,opt,name=criticality,proto3" json:"criticality,omitempty"`
}

func (x *GraphQLInspectorDiffOutput_Change) Reset() {
	*x = GraphQLInspectorDiffOutput_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLInspectorDiffOutput_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLInspectorDiffOutput_Change) ProtoMessage() {}

func (x *GraphQLInspectorDiffOutput_Change) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLInspectorDiffOutput_Change.ProtoReflect.Descriptor instead.
func (*GraphQLInspectorDiffOutput_Change) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GraphQLInspectorDiffOutput_Change) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GraphQLInspectorDiffOutput_Change) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GraphQLInspectorDiffOutput_Change) GetChangeType() string {
	if x != nil {
		return x.ChangeType
	}
	return ""
}

func (x *GraphQLInspectorDiffOutput_Change) GetCriticality() *GraphQLInspectorDiffOutput_Criticality {
	if x != nil {
		return x.Criticality
	}
	return nil
}

type GraphQLInspectorDiffOutput_Criticality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level  GraphQLInspectorDiffOutput_CriticalityLevel `protobuf:"varint,1,opt,name=level,proto3,enum=enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput_CriticalityLevel" json:"level,omitempty"`
	Reason string                                      `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *GraphQLInspectorDiffOutput_Criticality) Reset() {
	*x = GraphQLInspectorDiffOutput_Criticality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphQLInspectorDiffOutput_Criticality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphQLInspectorDiffOutput_Criticality) ProtoMessage() {}

func (x *GraphQLInspectorDiffOutput_Criticality) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphQLInspectorDiffOutput_Criticality.ProtoReflect.Descriptor instead.
func (*GraphQLInspectorDiffOutput_Criticality) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GraphQLInspectorDiffOutput_Criticality) GetLevel() GraphQLInspectorDiffOutput_CriticalityLevel {
	if x != nil {
		return x.Level
	}
	return GraphQLInspectorDiffOutput_NON_BREAKING
}

func (x *GraphQLInspectorDiffOutput_Criticality) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDesc = []byte{
	0x0a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x66, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x67, 0x6c,
	0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x19, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c,
	0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x66, 0x66, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x5f, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x49, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x22, 0x8e, 0x04, 0x0a, 0x1a, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x49, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x66, 0x66, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x5c, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f,
	0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x66, 0x66, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x1a, 0xc2,
	0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x69, 0x0a, 0x0b, 0x63, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x44, 0x69, 0x66, 0x66, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x1a, 0x89, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x62, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2e,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x49, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x66, 0x66, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e,
	0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x41, 0x0a, 0x10, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x4e, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x41, 0x4e, 0x47, 0x45, 0x52, 0x4f,
	0x55, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescData = file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDesc
)

func file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDescData
}

var file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_goTypes = []interface{}{
	(GraphQLInspectorDiffOutput_CriticalityLevel)(0),                    // 0: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.CriticalityLevel
	(*GraphQLInspectorDiffInput)(nil),                                   // 1: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffInput
	(*GraphQLInspectorDiffOutput)(nil),                                  // 2: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput
	(*GraphQLInspectorDiffOutput_Change)(nil),                           // 3: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.Change
	(*GraphQLInspectorDiffOutput_Criticality)(nil),                      // 4: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.Criticality
	(v1.GraphqlOptions_SchemaChangeValidationOptions_ProcessingRule)(0), // 5: gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule
}
var file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_depIdxs = []int32{
	5, // 0: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffInput.rules:type_name -> gloo.solo.io.GraphqlOptions.SchemaChangeValidationOptions.ProcessingRule
	3, // 1: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.changes:type_name -> enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.Change
	4, // 2: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.Change.criticality:type_name -> enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.Criticality
	0, // 3: enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.Criticality.level:type_name -> enterprise.graphql.gloo.solo.io.GraphQLInspectorDiffOutput.CriticalityLevel
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_init()
}
func file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_init() {
	if File_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLInspectorDiffInput); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLInspectorDiffOutput); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLInspectorDiffOutput_Change); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphQLInspectorDiffOutput_Criticality); i {
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
			RawDescriptor: file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto = out.File
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_rawDesc = nil
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_goTypes = nil
	file_github_com_solo_io_solo_projects_projects_gloo_api_enterprise_graphql_v1_diff_proto_depIdxs = nil
}
