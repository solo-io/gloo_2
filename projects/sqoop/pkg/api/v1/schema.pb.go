// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/sqoop/api/v1/schema.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//
// The Schema object wraps the user's GraphQL Schema, which is stored as an inline string.
// The Schema Object contains a Status field which is used by SQooP to validate the user's input schema.
//
// Schemas are matched to resolver maps in the same namespace with the same name
type Schema struct {
	// inline the entire graphql schema as a string here
	InlineSchema string `protobuf:"bytes,3,opt,name=inline_schema,json=inlineSchema,proto3" json:"inline_schema,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_00cb508776857c4d, []int{0}
}
func (m *Schema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schema.Unmarshal(m, b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return xxx_messageInfo_Schema.Size(m)
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetInlineSchema() string {
	if m != nil {
		return m.InlineSchema
	}
	return ""
}

func (m *Schema) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Schema) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Schema)(nil), "sqoop.solo.io.Schema")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/sqoop/api/v1/schema.proto", fileDescriptor_00cb508776857c4d)
}

var fileDescriptor_00cb508776857c4d = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0x87, 0xd0,
	0x05, 0x45, 0xf9, 0x59, 0xa9, 0xc9, 0x25, 0xc5, 0xfa, 0x70, 0x46, 0x71, 0x61, 0x7e, 0x7e, 0x81,
	0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x99, 0xa1, 0x7e, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0xa2, 0x5e, 0x41,
	0x51, 0x7e, 0x49, 0xbe, 0x10, 0x2f, 0x58, 0x4a, 0x0f, 0xa4, 0x4f, 0x2f, 0x33, 0x5f, 0x4a, 0x24,
	0x3d, 0x3f, 0x3d, 0x1f, 0x2c, 0xa3, 0x0f, 0x62, 0x41, 0x14, 0x49, 0x19, 0xe2, 0xb2, 0x29, 0x3b,
	0xb3, 0x04, 0x66, 0x6a, 0x6e, 0x6a, 0x49, 0x62, 0x4a, 0x62, 0x09, 0xd4, 0x5c, 0x29, 0x7d, 0x22,
	0xb4, 0x14, 0x97, 0x24, 0x96, 0x94, 0x16, 0x93, 0x60, 0x07, 0x8c, 0x0f, 0xd1, 0xa2, 0x74, 0x96,
	0x91, 0x8b, 0x2d, 0x18, 0xec, 0x19, 0x21, 0x65, 0x2e, 0xde, 0xcc, 0xbc, 0x9c, 0xcc, 0xbc, 0xd4,
	0x78, 0x88, 0xef, 0x24, 0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0x20, 0x82, 0x50, 0x45, 0xee,
	0x5c, 0x6c, 0x10, 0x2b, 0x25, 0xd8, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf4, 0x92, 0xf3, 0x8b,
	0x52, 0x61, 0x7e, 0xd7, 0x0b, 0x06, 0xcb, 0x39, 0x49, 0x9e, 0xb8, 0x27, 0xcf, 0xf0, 0xe9, 0x9e,
	0xbc, 0x60, 0x49, 0x6a, 0x71, 0x49, 0x4a, 0x66, 0x5a, 0x9a, 0x95, 0x52, 0x66, 0x7a, 0x5e, 0x7e,
	0x51, 0xaa, 0x52, 0x10, 0x54, 0xbb, 0x90, 0x05, 0x17, 0x07, 0xcc, 0xbb, 0x12, 0xec, 0x60, 0xa3,
	0xc4, 0x50, 0x8d, 0xf2, 0x85, 0xca, 0x3a, 0xb1, 0x80, 0x0c, 0x0b, 0x82, 0xab, 0xb6, 0x92, 0x6e,
	0xfa, 0xc8, 0xc2, 0xc2, 0xc5, 0x54, 0x9c, 0xdc, 0xf4, 0x91, 0x85, 0x53, 0x88, 0x1d, 0xe2, 0xd8,
	0xe2, 0xa6, 0x8f, 0x2c, 0x4c, 0x0a, 0x8c, 0x4e, 0x0e, 0x2b, 0x1e, 0xc9, 0x31, 0x46, 0x59, 0x91,
	0x1a, 0xad, 0x05, 0xd9, 0xe9, 0xd0, 0x00, 0x4a, 0x62, 0x03, 0x07, 0x8c, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x96, 0xd6, 0x04, 0x3a, 0x1b, 0x02, 0x00, 0x00,
}

func (this *Schema) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Schema)
	if !ok {
		that2, ok := that.(Schema)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InlineSchema != that1.InlineSchema {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
