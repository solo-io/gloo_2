// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoint.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core_solo_io "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// @solo-kit:resource
// @solo-kit:resource.short_name=ep
// @solo-kit:resource.plural_name=endpoints
// @solo-kit:resource.group_name=gloo.solo.io
// @solo-kit:resource.version=v1
//
// Endpoints represent a single IP address or hostname + port where an application or service is listening
type Endpoint struct {
	// Address of the endpoint (ip or hostname)
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// listening port for the endpoint
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata *core_solo_io.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{0} }

func (m *Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetMetadata() *core_solo_io.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "gloo.api.v1.Endpoint")
}
func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("endpoint.proto", fileDescriptorEndpoint) }

var fileDescriptorEndpoint = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xbd, 0x4e, 0xc5, 0x30,
	0x0c, 0x85, 0x15, 0x84, 0xb8, 0x97, 0x5c, 0xc1, 0x10, 0x21, 0x14, 0x75, 0x40, 0x15, 0x53, 0x17,
	0x1c, 0xb5, 0x30, 0x31, 0x22, 0x31, 0xb2, 0x74, 0x64, 0x4b, 0x9b, 0x28, 0x84, 0xfe, 0x38, 0x4a,
	0x4c, 0x9f, 0x89, 0xe7, 0xe2, 0x49, 0x50, 0xff, 0x18, 0x99, 0x72, 0x62, 0x1f, 0xfb, 0x7c, 0xe6,
	0xd7, 0x76, 0x34, 0x01, 0xfd, 0x48, 0x10, 0x22, 0x12, 0x8a, 0x93, 0xeb, 0x11, 0x41, 0x07, 0x0f,
	0x53, 0x99, 0xdd, 0x38, 0x74, 0xb8, 0xd4, 0xd5, 0xac, 0x56, 0x4b, 0x56, 0x3a, 0x4f, 0x1f, 0x5f,
	0x0d, 0xb4, 0x38, 0xa8, 0x84, 0x3d, 0x3e, 0x78, 0x5c, 0xdf, 0xce, 0x93, 0xd2, 0xc1, 0xab, 0xa9,
	0x54, 0x83, 0x25, 0x6d, 0x34, 0xe9, 0x75, 0xe4, 0xbe, 0xe7, 0xc7, 0xd7, 0x2d, 0x47, 0x48, 0x7e,
	0xd0, 0xc6, 0x44, 0x9b, 0x92, 0x64, 0x39, 0x2b, 0x2e, 0xeb, 0xfd, 0x2b, 0x04, 0x3f, 0x0f, 0x18,
	0x49, 0x9e, 0xe5, 0xac, 0xb8, 0xaa, 0x17, 0x2d, 0x2a, 0x7e, 0xdc, 0x77, 0xc9, 0x43, 0xce, 0x8a,
	0x53, 0x75, 0x0b, 0x2d, 0x46, 0x0b, 0x73, 0x22, 0x78, 0x84, 0xb7, 0xad, 0x5b, 0xff, 0xf9, 0x5e,
	0x9e, 0xbf, 0x7f, 0xee, 0xd8, 0xfb, 0xd3, 0x7f, 0x98, 0x21, 0xe2, 0xa7, 0x6d, 0x29, 0xa9, 0xf9,
	0x5c, 0x15, 0x3a, 0xb7, 0x81, 0x37, 0x17, 0x0b, 0xf0, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7d, 0xcc, 0xf1, 0x4c, 0x18, 0x01, 0x00, 0x00,
}
