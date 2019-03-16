// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/gloo/api/v1/ratelimit.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	ratelimit "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/ratelimit"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
//@solo-kit:xds-service=RateLimitDiscoveryService
//@solo-kit:resource.no_references
type RateLimitConfig struct {
	// @solo-kit:resource.name
	Domain               string                  `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*ratelimit.Descriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RateLimitConfig) Reset()         { *m = RateLimitConfig{} }
func (m *RateLimitConfig) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfig) ProtoMessage()    {}
func (*RateLimitConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f54dd703bc8e28ff, []int{0}
}
func (m *RateLimitConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfig.Unmarshal(m, b)
}
func (m *RateLimitConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfig.Marshal(b, m, deterministic)
}
func (m *RateLimitConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfig.Merge(m, src)
}
func (m *RateLimitConfig) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfig.Size(m)
}
func (m *RateLimitConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfig proto.InternalMessageInfo

func (m *RateLimitConfig) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitConfig) GetDescriptors() []*ratelimit.Descriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimitConfig)(nil), "gloo.solo.io.RateLimitConfig")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/gloo/api/v1/ratelimit.proto", fileDescriptor_f54dd703bc8e28ff)
}

var fileDescriptor_f54dd703bc8e28ff = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x35, 0x15, 0x0a, 0x4e, 0x05, 0x21, 0x54, 0xa9, 0xa1, 0xd6, 0x92, 0x53, 0x2c, 0x38, 0xd1,
	0x78, 0xb2, 0x17, 0x41, 0xab, 0x20, 0xf4, 0x94, 0x1e, 0x04, 0x0f, 0xc2, 0x74, 0xfa, 0x39, 0x1d,
	0x4d, 0xe6, 0x1b, 0x67, 0xa6, 0x91, 0x1e, 0xf5, 0x2f, 0xf8, 0x27, 0xfc, 0x3f, 0xfb, 0x0f, 0x96,
	0xfd, 0x15, 0x7b, 0x5a, 0x92, 0x74, 0xb3, 0x69, 0x97, 0x85, 0xdd, 0x65, 0x4f, 0xf9, 0x92, 0xf7,
	0x5e, 0xde, 0x9b, 0x6f, 0x1e, 0xf9, 0x28, 0xa4, 0x5b, 0x6f, 0x96, 0x94, 0x63, 0x1e, 0x5b, 0xcc,
	0xf0, 0xa5, 0xc4, 0xfa, 0xa9, 0x0d, 0xfe, 0x00, 0xee, 0x6c, 0xdc, 0x0c, 0x22, 0x43, 0x8c, 0x99,
	0x96, 0x71, 0xf1, 0x3a, 0x36, 0xcc, 0x41, 0x26, 0x73, 0xe9, 0xa8, 0x36, 0xe8, 0xd0, 0x7f, 0x58,
	0x82, 0xb4, 0x14, 0x52, 0x89, 0xc1, 0x10, 0x54, 0x81, 0xdb, 0x9a, 0x9b, 0xc4, 0x2b, 0x69, 0x39,
	0x16, 0x60, 0xb6, 0x35, 0x37, 0x18, 0x0a, 0x44, 0x91, 0x41, 0x05, 0x33, 0xa5, 0xd0, 0x31, 0x27,
	0x51, 0xd9, 0x1d, 0xfa, 0xe5, 0x76, 0x81, 0x74, 0xb6, 0x11, 0x52, 0xd9, 0x8b, 0x60, 0x87, 0x11,
	0x83, 0xbe, 0x40, 0x81, 0xd5, 0x18, 0x97, 0x53, 0xfd, 0x35, 0xfc, 0x4d, 0x1e, 0xa5, 0xcc, 0xc1,
	0xbc, 0x24, 0x7e, 0x40, 0xf5, 0x5d, 0x0a, 0xff, 0x09, 0xe9, 0xae, 0x30, 0x67, 0x52, 0x0d, 0xbc,
	0xb1, 0x17, 0x3d, 0x48, 0x77, 0x6f, 0xfe, 0x9c, 0xf4, 0x56, 0x60, 0xb9, 0x91, 0xda, 0xa1, 0xb1,
	0x83, 0xce, 0xf8, 0x7e, 0xd4, 0x4b, 0x26, 0xb4, 0xe5, 0x53, 0x67, 0xa0, 0xed, 0x5d, 0xd0, 0x59,
	0x23, 0x49, 0xdb, 0xf2, 0xe4, 0xb4, 0x43, 0x9e, 0x36, 0xce, 0xb3, 0xf3, 0x15, 0x2d, 0xc0, 0x14,
	0x92, 0x83, 0xff, 0x8d, 0x3c, 0x5e, 0x38, 0x03, 0x2c, 0x3f, 0x0c, 0x37, 0xa2, 0xd5, 0x6e, 0x29,
	0xd3, 0x92, 0x16, 0x09, 0x6d, 0x84, 0x29, 0xfc, 0xda, 0x80, 0x75, 0xc1, 0xf3, 0x2b, 0x71, 0xab,
	0x51, 0x59, 0x08, 0xef, 0x45, 0xde, 0x2b, 0xcf, 0xdf, 0x92, 0xe0, 0xb3, 0xe2, 0x06, 0x72, 0x50,
	0x8e, 0x65, 0x87, 0x26, 0x2f, 0xf6, 0x7f, 0xd2, 0x62, 0x5e, 0xf2, 0x9b, 0x5c, 0x87, 0xba, 0x67,
	0xfd, 0xc7, 0x23, 0xfd, 0x4f, 0xe0, 0xf8, 0xfa, 0xce, 0x8f, 0x16, 0xfd, 0x3d, 0x3a, 0xf9, 0xd7,
	0x09, 0xc3, 0x67, 0x7b, 0xad, 0x9b, 0x36, 0x17, 0xc4, 0x2b, 0x9f, 0xa9, 0x37, 0x79, 0xff, 0xee,
	0xff, 0xf1, 0xc8, 0xfb, 0xfa, 0xf6, 0x86, 0x55, 0xd3, 0x3f, 0xc5, 0xae, 0x6e, 0xcb, 0x6e, 0xd5,
	0x9e, 0x37, 0x67, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x69, 0x38, 0x75, 0x3f, 0x03, 0x00, 0x00,
}

func (this *RateLimitConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitConfig)
	if !ok {
		that2, ok := that.(RateLimitConfig)
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
	if this.Domain != that1.Domain {
		return false
	}
	if len(this.Descriptors) != len(that1.Descriptors) {
		return false
	}
	for i := range this.Descriptors {
		if !this.Descriptors[i].Equal(that1.Descriptors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitDiscoveryServiceClient is the client API for RateLimitDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitDiscoveryServiceClient interface {
	StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error)
	IncrementalRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_IncrementalRateLimitConfigClient, error)
	FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type rateLimitDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitDiscoveryServiceClient(cc *grpc.ClientConn) RateLimitDiscoveryServiceClient {
	return &rateLimitDiscoveryServiceClient{cc}
}

func (c *rateLimitDiscoveryServiceClient) StreamRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_StreamRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[0], "/gloo.solo.io.RateLimitDiscoveryService/StreamRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceStreamRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_StreamRateLimitConfigClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) IncrementalRateLimitConfig(ctx context.Context, opts ...grpc.CallOption) (RateLimitDiscoveryService_IncrementalRateLimitConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RateLimitDiscoveryService_serviceDesc.Streams[1], "/gloo.solo.io.RateLimitDiscoveryService/IncrementalRateLimitConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &rateLimitDiscoveryServiceIncrementalRateLimitConfigClient{stream}
	return x, nil
}

type RateLimitDiscoveryService_IncrementalRateLimitConfigClient interface {
	Send(*v2.IncrementalDiscoveryRequest) error
	Recv() (*v2.IncrementalDiscoveryResponse, error)
	grpc.ClientStream
}

type rateLimitDiscoveryServiceIncrementalRateLimitConfigClient struct {
	grpc.ClientStream
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigClient) Send(m *v2.IncrementalDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigClient) Recv() (*v2.IncrementalDiscoveryResponse, error) {
	m := new(v2.IncrementalDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rateLimitDiscoveryServiceClient) FetchRateLimitConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/gloo.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitDiscoveryServiceServer is the server API for RateLimitDiscoveryService service.
type RateLimitDiscoveryServiceServer interface {
	StreamRateLimitConfig(RateLimitDiscoveryService_StreamRateLimitConfigServer) error
	IncrementalRateLimitConfig(RateLimitDiscoveryService_IncrementalRateLimitConfigServer) error
	FetchRateLimitConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterRateLimitDiscoveryServiceServer(s *grpc.Server, srv RateLimitDiscoveryServiceServer) {
	s.RegisterService(&_RateLimitDiscoveryService_serviceDesc, srv)
}

func _RateLimitDiscoveryService_StreamRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).StreamRateLimitConfig(&rateLimitDiscoveryServiceStreamRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_StreamRateLimitConfigServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceStreamRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceStreamRateLimitConfigServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_IncrementalRateLimitConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RateLimitDiscoveryServiceServer).IncrementalRateLimitConfig(&rateLimitDiscoveryServiceIncrementalRateLimitConfigServer{stream})
}

type RateLimitDiscoveryService_IncrementalRateLimitConfigServer interface {
	Send(*v2.IncrementalDiscoveryResponse) error
	Recv() (*v2.IncrementalDiscoveryRequest, error)
	grpc.ServerStream
}

type rateLimitDiscoveryServiceIncrementalRateLimitConfigServer struct {
	grpc.ServerStream
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigServer) Send(m *v2.IncrementalDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rateLimitDiscoveryServiceIncrementalRateLimitConfigServer) Recv() (*v2.IncrementalDiscoveryRequest, error) {
	m := new(v2.IncrementalDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RateLimitDiscoveryService_FetchRateLimitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gloo.solo.io.RateLimitDiscoveryService/FetchRateLimitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitDiscoveryServiceServer).FetchRateLimitConfig(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gloo.solo.io.RateLimitDiscoveryService",
	HandlerType: (*RateLimitDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRateLimitConfig",
			Handler:    _RateLimitDiscoveryService_FetchRateLimitConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_StreamRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "IncrementalRateLimitConfig",
			Handler:       _RateLimitDiscoveryService_IncrementalRateLimitConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/solo-io/solo-projects/projects/gloo/api/v1/ratelimit.proto",
}
