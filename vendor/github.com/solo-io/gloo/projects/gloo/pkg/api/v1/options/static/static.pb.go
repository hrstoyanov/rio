// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto

package static

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Static upstreams are used to route request to services listening at fixed IP/Host & Port pairs.
// Static upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Static Upstreams must be created manually by users
type UpstreamSpec struct {
	// A list of addresses and ports
	// at least one must be specified
	Hosts []*Host `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Attempt to use outbound TLS
	// Gloo will automatically set this to true for port 443
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec          *options.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b3c87c0f36512, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetHosts() []*Host {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *UpstreamSpec) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

// Represents a single instance of an upstream
type Host struct {
	// Address (hostname or IP)
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Port the instance is listening on
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_c08b3c87c0f36512, []int{1}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (m *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(m, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Host) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "static.options.gloo.solo.io.UpstreamSpec")
	proto.RegisterType((*Host)(nil), "static.options.gloo.solo.io.Host")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto", fileDescriptor_c08b3c87c0f36512)
}

var fileDescriptor_c08b3c87c0f36512 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x94, 0xe9, 0x03, 0x70, 0xcb, 0xc5, 0x42, 0x22, 0x2a, 0x52, 0x15, 0x7a, 0xca, 0x05, 0x5b,
	0x3c, 0x24, 0x8e, 0x48, 0x88, 0x43, 0xb9, 0xa6, 0x70, 0xe1, 0x52, 0xa5, 0xa9, 0x71, 0x0d, 0x29,
	0x6b, 0x79, 0x37, 0xa5, 0x5f, 0x84, 0xf8, 0x04, 0xbe, 0x87, 0x7f, 0xe0, 0x8e, 0x92, 0x58, 0x82,
	0x43, 0x85, 0x7a, 0xca, 0xcc, 0x64, 0x76, 0x76, 0xe4, 0xe5, 0x63, 0x63, 0x69, 0x51, 0xce, 0x64,
	0x0e, 0x4b, 0x85, 0x50, 0xc0, 0xa9, 0x05, 0x65, 0x0a, 0x00, 0xe5, 0x3c, 0x3c, 0xeb, 0x9c, 0xb0,
	0x61, 0x99, 0xb3, 0x6a, 0x75, 0xa6, 0xc0, 0x91, 0x85, 0x57, 0x54, 0x48, 0x19, 0xd9, 0x3c, 0x7c,
	0xa4, 0xf3, 0x40, 0x20, 0x8e, 0x03, 0x0b, 0x1e, 0x59, 0xcd, 0xc9, 0x2a, 0x52, 0x5a, 0x18, 0x1c,
	0x1a, 0x30, 0x50, 0xfb, 0x54, 0x85, 0x9a, 0x91, 0x81, 0xd0, 0x6b, 0x6a, 0x44, 0xbd, 0xa6, 0xa0,
	0x0d, 0x0d, 0x80, 0x29, 0xb4, 0xaa, 0xd9, 0xac, 0x7c, 0x52, 0x6f, 0x3e, 0x73, 0x4e, 0x7b, 0x0c,
	0xff, 0x2f, 0xb7, 0x68, 0xa7, 0xfd, 0xca, 0xe6, 0x7a, 0x8a, 0x4e, 0x87, 0x72, 0xa3, 0x77, 0xc6,
	0xfb, 0x0f, 0x0e, 0xc9, 0xeb, 0x6c, 0x39, 0x71, 0x3a, 0x17, 0x57, 0xbc, 0xb3, 0x00, 0x24, 0x8c,
	0x58, 0xdc, 0x4a, 0x7a, 0xe7, 0x27, 0xf2, 0x9f, 0xf6, 0x72, 0x0c, 0x48, 0x69, 0xe3, 0x17, 0x47,
	0x7c, 0xb7, 0x44, 0x3d, 0xa5, 0x02, 0xa3, 0x56, 0xcc, 0x92, 0xbd, 0xb4, 0x5b, 0xa2, 0xbe, 0x2f,
	0x50, 0xdc, 0xf2, 0xfe, 0xdf, 0xc5, 0x51, 0x27, 0x66, 0x75, 0xf0, 0xc6, 0xc4, 0x49, 0xe3, 0xac,
	0xaa, 0xa4, 0x3d, 0xfc, 0x25, 0x23, 0xc9, 0xdb, 0xd5, 0x36, 0x21, 0x78, 0x3b, 0x9b, 0xcf, 0x7d,
	0xc4, 0x62, 0x96, 0xec, 0xa7, 0x35, 0xae, 0x34, 0x07, 0x9e, 0xa2, 0x9d, 0x98, 0x25, 0x07, 0x69,
	0x8d, 0x6f, 0xee, 0x3e, 0xbf, 0xdb, 0xec, 0xe3, 0x6b, 0xc8, 0x1e, 0xaf, 0xb7, 0xbb, 0xa4, 0x7b,
	0x31, 0x9b, 0xaf, 0x39, 0xeb, 0xd6, 0x4f, 0x75, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x39,
	0xc5, 0x9b, 0x13, 0x02, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if !this.Hosts[i].Equal(that1.Hosts[i]) {
			return false
		}
	}
	if this.UseTls != that1.UseTls {
		return false
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Host) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Host)
	if !ok {
		that2, ok := that.(Host)
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
	if this.Addr != that1.Addr {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
