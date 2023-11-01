// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-projects/projects/gloo-fed/api/fed/core/v1/placement.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type PlacementStatus_State int32

const (
	// UNKNOWN indicates that the system does not know the placement status of the resource.
	PlacementStatus_UNKNOWN PlacementStatus_State = 0
	// PLACED indicates that the resource has been placed as specified.
	PlacementStatus_PLACED PlacementStatus_State = 1
	// FAILED indicates that the resource could not be placed in a specified destination.
	PlacementStatus_FAILED PlacementStatus_State = 2
	// STALE indicates that the resource continues to be present in a destination that is no longer specified.
	PlacementStatus_STALE PlacementStatus_State = 3
	// INVALID indicates that the resource cannot be placed as specified because one or more destinations do not exist.
	PlacementStatus_INVALID PlacementStatus_State = 4
	// PENDING indicates that the resource is waiting to be processed.
	PlacementStatus_PENDING PlacementStatus_State = 5
)

// Enum value maps for PlacementStatus_State.
var (
	PlacementStatus_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "PLACED",
		2: "FAILED",
		3: "STALE",
		4: "INVALID",
		5: "PENDING",
	}
	PlacementStatus_State_value = map[string]int32{
		"UNKNOWN": 0,
		"PLACED":  1,
		"FAILED":  2,
		"STALE":   3,
		"INVALID": 4,
		"PENDING": 5,
	}
)

func (x PlacementStatus_State) Enum() *PlacementStatus_State {
	p := new(PlacementStatus_State)
	*p = x
	return p
}

func (x PlacementStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlacementStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_enumTypes[0].Descriptor()
}

func (PlacementStatus_State) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_enumTypes[0]
}

func (x PlacementStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlacementStatus_State.Descriptor instead.
func (PlacementStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescGZIP(), []int{1, 0}
}

// Object Metadata to be written with the resource into the remote cluster
type TemplateMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Annotations map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels      map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Name        string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TemplateMetadata) Reset() {
	*x = TemplateMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateMetadata) ProtoMessage() {}

func (x *TemplateMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateMetadata.ProtoReflect.Descriptor instead.
func (*TemplateMetadata) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateMetadata) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *TemplateMetadata) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TemplateMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PlacementStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// map containing the name of the cluster, with the associated Cluster namespaces.
	Clusters map[string]*PlacementStatus_Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State    PlacementStatus_State               `protobuf:"varint,2,opt,name=state,proto3,enum=core.fed.solo.io.PlacementStatus_State" json:"state,omitempty"`
	Message  string                              `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// metadata.Generation of the resource which has been processed
	ObservedGeneration int64 `protobuf:"varint,4,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// A field indicating the entity responsible for writing this status.
	// This is useful for determining if the pod has been restarted since the resource was processed.
	// Typically this value will be set to metadata.name of the pod
	WrittenBy string `protobuf:"bytes,5,opt,name=written_by,json=writtenBy,proto3" json:"written_by,omitempty"`
}

func (x *PlacementStatus) Reset() {
	*x = PlacementStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementStatus) ProtoMessage() {}

func (x *PlacementStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementStatus.ProtoReflect.Descriptor instead.
func (*PlacementStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescGZIP(), []int{1}
}

func (x *PlacementStatus) GetClusters() map[string]*PlacementStatus_Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *PlacementStatus) GetState() PlacementStatus_State {
	if x != nil {
		return x.State
	}
	return PlacementStatus_UNKNOWN
}

func (x *PlacementStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PlacementStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *PlacementStatus) GetWrittenBy() string {
	if x != nil {
		return x.WrittenBy
	}
	return ""
}

type PlacementStatus_Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   PlacementStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=core.fed.solo.io.PlacementStatus_State" json:"state,omitempty"`
	Message string                `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PlacementStatus_Namespace) Reset() {
	*x = PlacementStatus_Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementStatus_Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementStatus_Namespace) ProtoMessage() {}

func (x *PlacementStatus_Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementStatus_Namespace.ProtoReflect.Descriptor instead.
func (*PlacementStatus_Namespace) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PlacementStatus_Namespace) GetState() PlacementStatus_State {
	if x != nil {
		return x.State
	}
	return PlacementStatus_UNKNOWN
}

func (x *PlacementStatus_Namespace) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PlacementStatus_Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// map containing the name of the namespace, with the associated status.
	Namespaces map[string]*PlacementStatus_Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlacementStatus_Cluster) Reset() {
	*x = PlacementStatus_Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementStatus_Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementStatus_Cluster) ProtoMessage() {}

func (x *PlacementStatus_Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementStatus_Cluster.ProtoReflect.Descriptor instead.
func (*PlacementStatus_Cluster) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PlacementStatus_Cluster) GetNamespaces() map[string]*PlacementStatus_Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

var File_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2d, 0x66, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x10, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x55, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65,
	0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfb, 0x05,
	0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x4b, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3d,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74,
	0x74, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x72,
	0x69, 0x74, 0x74, 0x65, 0x6e, 0x42, 0x79, 0x1a, 0x64, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xd0, 0x01,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x59, 0x0a, 0x0a, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x1a, 0x6a, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x66, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x51, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x4c, 0x45, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x42, 0x58, 0xb8, 0xf5, 0x04,
	0x01, 0xc0, 0xf5, 0x04, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescData = file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDesc
)

func file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDescData
}

var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_goTypes = []interface{}{
	(PlacementStatus_State)(0),        // 0: core.fed.solo.io.PlacementStatus.State
	(*TemplateMetadata)(nil),          // 1: core.fed.solo.io.TemplateMetadata
	(*PlacementStatus)(nil),           // 2: core.fed.solo.io.PlacementStatus
	nil,                               // 3: core.fed.solo.io.TemplateMetadata.AnnotationsEntry
	nil,                               // 4: core.fed.solo.io.TemplateMetadata.LabelsEntry
	(*PlacementStatus_Namespace)(nil), // 5: core.fed.solo.io.PlacementStatus.Namespace
	(*PlacementStatus_Cluster)(nil),   // 6: core.fed.solo.io.PlacementStatus.Cluster
	nil,                               // 7: core.fed.solo.io.PlacementStatus.ClustersEntry
	nil,                               // 8: core.fed.solo.io.PlacementStatus.Cluster.NamespacesEntry
}
var file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_depIdxs = []int32{
	3, // 0: core.fed.solo.io.TemplateMetadata.annotations:type_name -> core.fed.solo.io.TemplateMetadata.AnnotationsEntry
	4, // 1: core.fed.solo.io.TemplateMetadata.labels:type_name -> core.fed.solo.io.TemplateMetadata.LabelsEntry
	7, // 2: core.fed.solo.io.PlacementStatus.clusters:type_name -> core.fed.solo.io.PlacementStatus.ClustersEntry
	0, // 3: core.fed.solo.io.PlacementStatus.state:type_name -> core.fed.solo.io.PlacementStatus.State
	0, // 4: core.fed.solo.io.PlacementStatus.Namespace.state:type_name -> core.fed.solo.io.PlacementStatus.State
	8, // 5: core.fed.solo.io.PlacementStatus.Cluster.namespaces:type_name -> core.fed.solo.io.PlacementStatus.Cluster.NamespacesEntry
	6, // 6: core.fed.solo.io.PlacementStatus.ClustersEntry.value:type_name -> core.fed.solo.io.PlacementStatus.Cluster
	5, // 7: core.fed.solo.io.PlacementStatus.Cluster.NamespacesEntry.value:type_name -> core.fed.solo.io.PlacementStatus.Namespace
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_init()
}
func file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_init() {
	if File_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateMetadata); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementStatus); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementStatus_Namespace); i {
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
		file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacementStatus_Cluster); i {
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
			RawDescriptor: file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto = out.File
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_rawDesc = nil
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_goTypes = nil
	file_github_com_solo_io_solo_projects_projects_gloo_fed_api_fed_core_v1_placement_proto_depIdxs = nil
}
