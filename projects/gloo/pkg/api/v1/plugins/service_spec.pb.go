// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins/service_spec.proto

package plugins // import "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import grpc1 "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/grpc"
import rest "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/rest"
import sqoop "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/sqoop"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Describes APIs and application-level information for services
// Gloo routes to. ServiceSpec is contained within the UpstreamSpec for certain types
// of upstreams, including Kubernetes, Consul, and Static.
// ServiceSpec configuration is opaque to Gloo and handled by Service Plugins.
type ServiceSpec struct {
	// Note to developers: new Service Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to PluginType:
	//	*ServiceSpec_Rest
	//	*ServiceSpec_Grpc
	//	*ServiceSpec_Sqoop
	PluginType           isServiceSpec_PluginType `protobuf_oneof:"plugin_type"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_spec_7d88b16e17b650f3, []int{0}
}
func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (dst *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(dst, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

type isServiceSpec_PluginType interface {
	isServiceSpec_PluginType()
	Equal(interface{}) bool
}

type ServiceSpec_Rest struct {
	Rest *rest.ServiceSpec `protobuf:"bytes,1,opt,name=rest,oneof"`
}
type ServiceSpec_Grpc struct {
	Grpc *grpc1.ServiceSpec `protobuf:"bytes,2,opt,name=grpc,oneof"`
}
type ServiceSpec_Sqoop struct {
	Sqoop *sqoop.ServiceSpec `protobuf:"bytes,3,opt,name=sqoop,oneof"`
}

func (*ServiceSpec_Rest) isServiceSpec_PluginType()  {}
func (*ServiceSpec_Grpc) isServiceSpec_PluginType()  {}
func (*ServiceSpec_Sqoop) isServiceSpec_PluginType() {}

func (m *ServiceSpec) GetPluginType() isServiceSpec_PluginType {
	if m != nil {
		return m.PluginType
	}
	return nil
}

func (m *ServiceSpec) GetRest() *rest.ServiceSpec {
	if x, ok := m.GetPluginType().(*ServiceSpec_Rest); ok {
		return x.Rest
	}
	return nil
}

func (m *ServiceSpec) GetGrpc() *grpc1.ServiceSpec {
	if x, ok := m.GetPluginType().(*ServiceSpec_Grpc); ok {
		return x.Grpc
	}
	return nil
}

func (m *ServiceSpec) GetSqoop() *sqoop.ServiceSpec {
	if x, ok := m.GetPluginType().(*ServiceSpec_Sqoop); ok {
		return x.Sqoop
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServiceSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServiceSpec_OneofMarshaler, _ServiceSpec_OneofUnmarshaler, _ServiceSpec_OneofSizer, []interface{}{
		(*ServiceSpec_Rest)(nil),
		(*ServiceSpec_Grpc)(nil),
		(*ServiceSpec_Sqoop)(nil),
	}
}

func _ServiceSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServiceSpec)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *ServiceSpec_Rest:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Rest); err != nil {
			return err
		}
	case *ServiceSpec_Grpc:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Grpc); err != nil {
			return err
		}
	case *ServiceSpec_Sqoop:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Sqoop); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServiceSpec.PluginType has unexpected type %T", x)
	}
	return nil
}

func _ServiceSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServiceSpec)
	switch tag {
	case 1: // plugin_type.rest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(rest.ServiceSpec)
		err := b.DecodeMessage(msg)
		m.PluginType = &ServiceSpec_Rest{msg}
		return true, err
	case 2: // plugin_type.grpc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(grpc1.ServiceSpec)
		err := b.DecodeMessage(msg)
		m.PluginType = &ServiceSpec_Grpc{msg}
		return true, err
	case 3: // plugin_type.sqoop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(sqoop.ServiceSpec)
		err := b.DecodeMessage(msg)
		m.PluginType = &ServiceSpec_Sqoop{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServiceSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServiceSpec)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *ServiceSpec_Rest:
		s := proto.Size(x.Rest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServiceSpec_Grpc:
		s := proto.Size(x.Grpc)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServiceSpec_Sqoop:
		s := proto.Size(x.Sqoop)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ServiceSpec)(nil), "plugins.gloo.solo.io.ServiceSpec")
}
func (this *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
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
	if that1.PluginType == nil {
		if this.PluginType != nil {
			return false
		}
	} else if this.PluginType == nil {
		return false
	} else if !this.PluginType.Equal(that1.PluginType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSpec_Rest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_Rest)
	if !ok {
		that2, ok := that.(ServiceSpec_Rest)
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
	if !this.Rest.Equal(that1.Rest) {
		return false
	}
	return true
}
func (this *ServiceSpec_Grpc) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_Grpc)
	if !ok {
		that2, ok := that.(ServiceSpec_Grpc)
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
	if !this.Grpc.Equal(that1.Grpc) {
		return false
	}
	return true
}
func (this *ServiceSpec_Sqoop) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSpec_Sqoop)
	if !ok {
		that2, ok := that.(ServiceSpec_Sqoop)
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
	if !this.Sqoop.Equal(that1.Sqoop) {
		return false
	}
	return true
}

func init() {
	proto.RegisterFile("plugins/service_spec.proto", fileDescriptor_service_spec_7d88b16e17b650f3)
}

var fileDescriptor_service_spec_7d88b16e17b650f3 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xbd, 0x4a, 0x05, 0x31,
	0x10, 0x85, 0x8d, 0x7f, 0x45, 0x16, 0x9b, 0xe5, 0x16, 0xcb, 0x16, 0x22, 0x16, 0x6a, 0xe3, 0x04,
	0xb5, 0x15, 0xc4, 0xdb, 0x68, 0x23, 0x82, 0xb7, 0xb3, 0xb9, 0xb8, 0x21, 0xc4, 0xe8, 0xea, 0x8c,
	0x49, 0x76, 0xc1, 0x37, 0xf2, 0x95, 0x6c, 0x7d, 0x12, 0x49, 0x26, 0x8a, 0x88, 0xa0, 0xa2, 0xcd,
	0x64, 0x20, 0xe7, 0x3b, 0x39, 0x39, 0xb2, 0xa5, 0x7e, 0xb0, 0xee, 0x3e, 0xa8, 0x60, 0xfc, 0xe8,
	0xb4, 0x99, 0x07, 0x32, 0x1a, 0xc8, 0x63, 0xc4, 0x7a, 0x52, 0xee, 0xc0, 0xf6, 0x88, 0x10, 0xb0,
	0x47, 0x70, 0xd8, 0x4e, 0x2c, 0x5a, 0xcc, 0x02, 0x95, 0x36, 0xd6, 0xb6, 0x67, 0xd6, 0xc5, 0xeb,
	0xa1, 0x03, 0x8d, 0x77, 0x2a, 0x29, 0x77, 0x1d, 0xf2, 0x49, 0x1e, 0x6f, 0x8c, 0x8e, 0x41, 0xbd,
	0x2f, 0xc9, 0x4d, 0x5d, 0x91, 0x53, 0xe3, 0x9e, 0x7a, 0x7b, 0xdd, 0x9b, 0x10, 0xf3, 0xf8, 0x1f,
	0x3b, 0xeb, 0x49, 0xe7, 0x51, 0xec, 0xce, 0xff, 0x66, 0x17, 0x1e, 0x10, 0x89, 0x27, 0x1b, 0x6e,
	0x3e, 0x0b, 0x59, 0xcd, 0xb8, 0xb1, 0x19, 0x19, 0x5d, 0x1f, 0xca, 0xe5, 0x94, 0xbe, 0x11, 0x1b,
	0x62, 0xa7, 0xda, 0xdf, 0x02, 0xfe, 0xca, 0x17, 0xf5, 0xc1, 0x07, 0xea, 0x74, 0xe1, 0x22, 0x53,
	0x89, 0x4e, 0x61, 0x9b, 0xc5, 0x42, 0x73, 0xf2, 0x1f, 0xd0, 0x49, 0x58, 0x1f, 0xc9, 0x95, 0x1c,
	0xad, 0x59, 0xca, 0xf8, 0x36, 0x94, 0xa0, 0xdf, 0xf3, 0xcc, 0x4d, 0xd7, 0x64, 0xc5, 0xe2, 0x79,
	0x7c, 0x24, 0x33, 0x3d, 0x79, 0x7a, 0x59, 0x17, 0x97, 0xc7, 0xbf, 0xac, 0x8c, 0x6e, 0xed, 0xa7,
	0xda, 0xba, 0xd5, 0xdc, 0xd5, 0xc1, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0xd2, 0x55, 0x2d,
	0x64, 0x02, 0x00, 0x00,
}
