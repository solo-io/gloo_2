// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gateway.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	gateway.proto

It has these top-level messages:
	Gateway
	VirtualService
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core_solo_io "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
import core_solo_io1 "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
import gloo_solo_io1 "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
import gloo_solo_io "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"

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
// @solo-kit:resource
// @solo-kit:resource.short_name=gw
// @solo-kit:resource.plural_name=gateways
// @solo-kit:resource.group_name=gateway.solo.io
// @solo-kit:resource.version=v1
//
// A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.
type Gateway struct {
	// names of the the virtual services, which contain the actual routes for the gateway
	// if this list is empty, the gateway will be considered invalid by Gloo
	VirtualServices []string `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices" json:"virtual_services,omitempty"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways in a namespace
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// If provided, the Gateway will serve TLS/SSL traffic on this port
	SslConfig *gloo_solo_io1.SslConfig `protobuf:"bytes,1,opt,name=ssl_config,json=sslConfig" json:"ssl_config,omitempty"`
	// top level plugin configuration for all routes on the gateway
	Plugins map[string]*gloo_solo_io.ListenerPlugin `protobuf:"bytes,5,rep,name=plugins" json:"plugins,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core_solo_io1.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core_solo_io.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
}

func (m *Gateway) Reset()                    { *m = Gateway{} }
func (m *Gateway) String() string            { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()               {}
func (*Gateway) Descriptor() ([]byte, []int) { return fileDescriptorGateway, []int{0} }

func (m *Gateway) GetVirtualServices() []string {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Gateway) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *Gateway) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *Gateway) GetSslConfig() *gloo_solo_io1.SslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *Gateway) GetPlugins() map[string]*gloo_solo_io.ListenerPlugin {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *Gateway) GetStatus() core_solo_io1.Status {
	if m != nil {
		return m.Status
	}
	return core_solo_io1.Status{}
}

func (m *Gateway) GetMetadata() core_solo_io.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core_solo_io.Metadata{}
}

//
// @solo-kit:resource
// @solo-kit:resource.short_name=vs
// @solo-kit:resource.plural_name=virtualservices
// @solo-kit:resource.group_name=gateway.solo.io
// @solo-kit:resource.version=v1
//
// A virtual service describes the set of routes to match for a set of domains.
//
// Domains must be unique across all virtual services within a gateway (i.e. no overlap between sets).
type VirtualService struct {
	VirtualHost *gloo_solo_io1.VirtualHost `protobuf:"bytes,1,opt,name=virtual_host,json=virtualHost" json:"virtual_host,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core_solo_io1.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core_solo_io.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
}

func (m *VirtualService) Reset()                    { *m = VirtualService{} }
func (m *VirtualService) String() string            { return proto.CompactTextString(m) }
func (*VirtualService) ProtoMessage()               {}
func (*VirtualService) Descriptor() ([]byte, []int) { return fileDescriptorGateway, []int{1} }

func (m *VirtualService) GetVirtualHost() *gloo_solo_io1.VirtualHost {
	if m != nil {
		return m.VirtualHost
	}
	return nil
}

func (m *VirtualService) GetStatus() core_solo_io1.Status {
	if m != nil {
		return m.Status
	}
	return core_solo_io1.Status{}
}

func (m *VirtualService) GetMetadata() core_solo_io.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core_solo_io.Metadata{}
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.Gateway")
	proto.RegisterType((*VirtualService)(nil), "gateway.solo.io.VirtualService")
}
func (this *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
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
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if this.VirtualServices[i] != that1.VirtualServices[i] {
			return false
		}
	}
	if this.BindAddress != that1.BindAddress {
		return false
	}
	if this.BindPort != that1.BindPort {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if len(this.Plugins) != len(that1.Plugins) {
		return false
	}
	for i := range this.Plugins {
		if !this.Plugins[i].Equal(that1.Plugins[i]) {
			return false
		}
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	return true
}
func (this *VirtualService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualService)
	if !ok {
		that2, ok := that.(VirtualService)
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
	if !this.VirtualHost.Equal(that1.VirtualHost) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptorGateway) }

var fileDescriptorGateway = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x9b, 0xd4, 0xeb, 0x84, 0x96, 0x55, 0x05, 0x6e, 0x40, 0xd4, 0x44, 0x42, 0x32,
	0x07, 0x6c, 0x35, 0x48, 0x50, 0x55, 0x20, 0x44, 0x10, 0x2a, 0x07, 0x90, 0xaa, 0xad, 0x84, 0x10,
	0x97, 0x68, 0x63, 0x6f, 0xdc, 0x25, 0xae, 0xc7, 0xda, 0x1d, 0x1b, 0xf2, 0x47, 0x7c, 0x0a, 0xe2,
	0xc6, 0x0f, 0xf4, 0xc0, 0x27, 0xf0, 0x05, 0xc8, 0xf6, 0x9a, 0x12, 0x0e, 0x88, 0x63, 0x4f, 0x9e,
	0x9d, 0x79, 0xef, 0xed, 0xf8, 0xcd, 0x2c, 0x19, 0x26, 0x1c, 0xc5, 0x27, 0xbe, 0x0a, 0x72, 0x05,
	0x08, 0x74, 0xbb, 0x3d, 0x6a, 0x48, 0x21, 0x90, 0x30, 0xda, 0x4d, 0x20, 0x81, 0xba, 0x16, 0x56,
	0x51, 0x03, 0x1b, 0x1d, 0x24, 0x12, 0xcf, 0x8a, 0x79, 0x10, 0xc1, 0x79, 0x58, 0x21, 0x1f, 0x4a,
	0x68, 0xbe, 0x4b, 0x89, 0x21, 0xcf, 0x65, 0x58, 0x1e, 0x84, 0xe7, 0x02, 0x79, 0xcc, 0x91, 0x1b,
	0x4a, 0xf8, 0x1f, 0x14, 0x8d, 0x1c, 0x0b, 0x6d, 0x08, 0x4e, 0xae, 0xe0, 0xb3, 0xe9, 0x6b, 0x34,
	0xcc, 0xd3, 0x22, 0x91, 0x99, 0xa9, 0x8d, 0xbf, 0x75, 0x49, 0xff, 0xb8, 0xe9, 0x94, 0x3e, 0x20,
	0x3b, 0xa5, 0x54, 0x58, 0xf0, 0x74, 0xa6, 0x85, 0x2a, 0x65, 0x24, 0xb4, 0xdb, 0xf1, 0xba, 0xbe,
	0xcd, 0xb6, 0x4d, 0xfe, 0xd4, 0xa4, 0xe9, 0x3d, 0x32, 0x98, 0xcb, 0x2c, 0x9e, 0xf1, 0x38, 0x56,
	0x42, 0x6b, 0xb7, 0xeb, 0x59, 0xbe, 0xcd, 0x9c, 0x2a, 0xf7, 0xa2, 0x49, 0xd1, 0xdb, 0xc4, 0xae,
	0x21, 0x39, 0x28, 0x74, 0x37, 0x3c, 0xcb, 0x1f, 0xb2, 0xad, 0x2a, 0x71, 0x02, 0x0a, 0xe9, 0x63,
	0x42, 0xb4, 0x4e, 0x67, 0x11, 0x64, 0x0b, 0x99, 0xb8, 0x96, 0x67, 0xf9, 0xce, 0xe4, 0x56, 0x90,
	0xa4, 0x00, 0xad, 0x5f, 0xc1, 0xa9, 0x4e, 0x5f, 0xd6, 0x65, 0x66, 0xeb, 0x36, 0xa4, 0xcf, 0x49,
	0xdf, 0xf4, 0xef, 0x6e, 0x7a, 0x5d, 0xdf, 0x99, 0xdc, 0x0f, 0xfe, 0xf2, 0x39, 0x30, 0x7f, 0x13,
	0x9c, 0x34, 0xb8, 0x57, 0x19, 0xaa, 0x15, 0x6b, 0x59, 0xf4, 0x98, 0xf4, 0x1a, 0x6f, 0xdc, 0x5e,
	0x7d, 0xe9, 0x6e, 0x10, 0x81, 0x12, 0x97, 0x97, 0xd6, 0xb5, 0xe9, 0xde, 0xd7, 0x8b, 0xfd, 0x6b,
	0x3f, 0x2f, 0xf6, 0x6f, 0xa0, 0xd0, 0x18, 0xcb, 0xc5, 0xe2, 0x68, 0x2c, 0x93, 0x0c, 0x94, 0x18,
	0x33, 0x43, 0xa7, 0x87, 0x64, 0xab, 0x9d, 0x8b, 0xdb, 0xaf, 0xa5, 0x6e, 0xae, 0x4b, 0xbd, 0x35,
	0xd5, 0xe9, 0x46, 0x25, 0xc6, 0x7e, 0xa3, 0x47, 0xef, 0xc9, 0xe0, 0xcf, 0xde, 0xe8, 0x0e, 0xe9,
	0x2e, 0xc5, 0xaa, 0x36, 0xc1, 0x66, 0x55, 0x48, 0x27, 0x64, 0xb3, 0xe4, 0x69, 0x21, 0xdc, 0x4e,
	0x2d, 0x7c, 0x67, 0xdd, 0x98, 0x37, 0x52, 0xa3, 0xc8, 0x84, 0x6a, 0x44, 0x58, 0x03, 0x3d, 0xea,
	0x1c, 0x5a, 0xe3, 0xef, 0x16, 0xb9, 0xfe, 0x6e, 0x6d, 0x52, 0xf4, 0x29, 0x19, 0xb4, 0x33, 0x3d,
	0x03, 0x8d, 0xc6, 0xea, 0xbd, 0x75, 0x45, 0xc3, 0x79, 0x0d, 0x1a, 0x99, 0x53, 0x5e, 0x1e, 0xae,
	0x80, 0x5b, 0xd3, 0x67, 0x5f, 0x7e, 0xdc, 0xb5, 0x3e, 0x3c, 0xf9, 0xd7, 0xce, 0xe7, 0x0a, 0x3e,
	0x8a, 0x08, 0x75, 0x68, 0x56, 0x21, 0xcc, 0x97, 0x89, 0x79, 0x08, 0xf3, 0x5e, 0xbd, 0xe6, 0x8f,
	0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x25, 0x45, 0xa2, 0x8a, 0x9e, 0x03, 0x00, 0x00,
}
