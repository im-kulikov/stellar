// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/health/v1/health.proto

package health // import "github.com/ehazlett/stellar/api/services/health/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
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

type NodeHealth struct {
	OSName    string `protobuf:"bytes,1,opt,name=os_name,json=osName,proto3" json:"os_name,omitempty"`
	OSVersion string `protobuf:"bytes,2,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	// TODO: use gogoproto.stdtime (returning panic: message/group field time.Time:bytes without pointer when trying to use)
	StartedAt            *types.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	Cpus                 int64            `protobuf:"varint,4,opt,name=cpus,proto3" json:"cpus,omitempty"`
	MemoryTotal          int64            `protobuf:"varint,5,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	MemoryFree           int64            `protobuf:"varint,6,opt,name=memory_free,json=memoryFree,proto3" json:"memory_free,omitempty"`
	MemoryUsed           int64            `protobuf:"varint,7,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
	Peers                []*Peer          `protobuf:"bytes,8,rep,name=peers" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeHealth) Reset()         { *m = NodeHealth{} }
func (m *NodeHealth) String() string { return proto.CompactTextString(m) }
func (*NodeHealth) ProtoMessage()    {}
func (*NodeHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_6c0db3482130361c, []int{0}
}
func (m *NodeHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeHealth.Unmarshal(m, b)
}
func (m *NodeHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeHealth.Marshal(b, m, deterministic)
}
func (dst *NodeHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeHealth.Merge(dst, src)
}
func (m *NodeHealth) XXX_Size() int {
	return xxx_messageInfo_NodeHealth.Size(m)
}
func (m *NodeHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeHealth.DiscardUnknown(m)
}

var xxx_messageInfo_NodeHealth proto.InternalMessageInfo

func (m *NodeHealth) GetOSName() string {
	if m != nil {
		return m.OSName
	}
	return ""
}

func (m *NodeHealth) GetOSVersion() string {
	if m != nil {
		return m.OSVersion
	}
	return ""
}

func (m *NodeHealth) GetStartedAt() *types.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *NodeHealth) GetCpus() int64 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *NodeHealth) GetMemoryTotal() int64 {
	if m != nil {
		return m.MemoryTotal
	}
	return 0
}

func (m *NodeHealth) GetMemoryFree() int64 {
	if m != nil {
		return m.MemoryFree
	}
	return 0
}

func (m *NodeHealth) GetMemoryUsed() int64 {
	if m != nil {
		return m.MemoryUsed
	}
	return 0
}

func (m *NodeHealth) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type HealthResponse struct {
	Health               *NodeHealth `protobuf:"bytes,1,opt,name=health" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_6c0db3482130361c, []int{1}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (dst *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(dst, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetHealth() *NodeHealth {
	if m != nil {
		return m.Health
	}
	return nil
}

type Peer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_6c0db3482130361c, []int{2}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Peer) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeHealth)(nil), "stellar.services.health.v1.NodeHealth")
	proto.RegisterType((*HealthResponse)(nil), "stellar.services.health.v1.HealthResponse")
	proto.RegisterType((*Peer)(nil), "stellar.services.health.v1.Peer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/stellar.services.health.v1.Health/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Health(context.Context, *types.Empty) (*HealthResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.health.v1.Health/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Health_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/health/v1/health.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/health/v1/health.proto", fileDescriptor_health_6c0db3482130361c)
}

var fileDescriptor_health_6c0db3482130361c = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x49, 0x93, 0x6e, 0xcd, 0x8b, 0x7a, 0x18, 0x44, 0x96, 0xf5, 0x90, 0x35, 0x82, 0x04,
	0x91, 0x19, 0x1a, 0x41, 0x28, 0x05, 0xc5, 0x82, 0xe2, 0x41, 0xda, 0x32, 0xad, 0x3d, 0x78, 0x59,
	0x26, 0xd9, 0xd7, 0xcd, 0xc2, 0x4e, 0x66, 0x99, 0x37, 0x09, 0xd4, 0xaf, 0xe8, 0x77, 0xe8, 0xa1,
	0x9f, 0x44, 0x76, 0x66, 0x62, 0xab, 0x62, 0xc1, 0x53, 0xfe, 0xf3, 0x7f, 0xbf, 0xc9, 0x7b, 0xfb,
	0xfe, 0x03, 0xef, 0xab, 0xda, 0x2d, 0xd7, 0x73, 0xbe, 0x30, 0x5a, 0xe0, 0x52, 0x7d, 0x6f, 0xd0,
	0x39, 0x41, 0x0e, 0x9b, 0x46, 0x59, 0xa1, 0xda, 0x5a, 0x10, 0xda, 0x4d, 0xbd, 0x40, 0x12, 0x4b,
	0x54, 0x8d, 0x5b, 0x8a, 0xcd, 0x7e, 0x54, 0xbc, 0xb5, 0xc6, 0x19, 0x96, 0x45, 0x98, 0x6f, 0x41,
	0x1e, 0xcb, 0x9b, 0xfd, 0xec, 0x49, 0x65, 0x2a, 0xe3, 0x31, 0xd1, 0xa9, 0x70, 0x23, 0x7b, 0x56,
	0x19, 0x53, 0x35, 0x28, 0xfc, 0x69, 0xbe, 0xbe, 0x14, 0xa8, 0x5b, 0x77, 0x15, 0x8b, 0xe3, 0x3f,
	0x8b, 0xae, 0xd6, 0x48, 0x4e, 0xe9, 0x36, 0x00, 0x93, 0x1f, 0x3b, 0x00, 0xc7, 0xa6, 0xc4, 0xcf,
	0xbe, 0x0b, 0x7b, 0x01, 0x7b, 0x86, 0x8a, 0x95, 0xd2, 0x98, 0xf6, 0xf2, 0xde, 0x74, 0x78, 0x04,
	0x37, 0xd7, 0xe3, 0xe4, 0xe4, 0xec, 0x58, 0x69, 0x94, 0x89, 0xa1, 0xee, 0x97, 0xbd, 0x06, 0x30,
	0x54, 0x6c, 0xd0, 0x52, 0x6d, 0x56, 0xe9, 0x8e, 0xe7, 0x1e, 0xdd, 0x5c, 0x8f, 0x87, 0x27, 0x67,
	0x17, 0xc1, 0x94, 0x43, 0x43, 0x51, 0xb2, 0x03, 0x00, 0x72, 0xca, 0x3a, 0x2c, 0x0b, 0xe5, 0xd2,
	0x7e, 0xde, 0x9b, 0x8e, 0x66, 0x19, 0x0f, 0x73, 0xf1, 0xed, 0x5c, 0xfc, 0x7c, 0x3b, 0x97, 0x1c,
	0x46, 0xfa, 0x83, 0x63, 0x0c, 0x06, 0x8b, 0x76, 0x4d, 0xe9, 0x20, 0xef, 0x4d, 0xfb, 0xd2, 0x6b,
	0xf6, 0x1c, 0x1e, 0x6a, 0xd4, 0xc6, 0x5e, 0x15, 0xce, 0x38, 0xd5, 0xa4, 0xbb, 0xbe, 0x36, 0x0a,
	0xde, 0x79, 0x67, 0xb1, 0x31, 0xc4, 0x63, 0x71, 0x69, 0x11, 0xd3, 0xc4, 0x13, 0x10, 0xac, 0x4f,
	0x16, 0xf1, 0x0e, 0xb0, 0x26, 0x2c, 0xd3, 0xbd, 0xbb, 0xc0, 0x57, 0xc2, 0x92, 0xbd, 0x85, 0xdd,
	0x16, 0xd1, 0x52, 0xfa, 0x20, 0xef, 0x4f, 0x47, 0xb3, 0x9c, 0xff, 0x3b, 0x15, 0x7e, 0x8a, 0x68,
	0x65, 0xc0, 0x27, 0xa7, 0xf0, 0x38, 0x2c, 0x52, 0x22, 0xb5, 0x66, 0x45, 0xc8, 0xde, 0x41, 0x12,
	0x50, 0xbf, 0xcf, 0xd1, 0xec, 0xe5, 0x7d, 0x7f, 0x75, 0x1b, 0x84, 0x8c, 0xb7, 0x26, 0x1c, 0x06,
	0x5d, 0x83, 0x6e, 0x15, 0xb7, 0xa9, 0x48, 0xaf, 0x3b, 0x4f, 0x95, 0xa5, 0x0d, 0x09, 0x48, 0xaf,
	0x67, 0x17, 0x90, 0xc4, 0x28, 0xbf, 0xfc, 0x52, 0x4f, 0xff, 0xda, 0xf6, 0xc7, 0xee, 0x89, 0x64,
	0xaf, 0xee, 0x9b, 0xe5, 0xf7, 0xef, 0x38, 0x3a, 0xfc, 0x76, 0xf0, 0xff, 0x4f, 0xfb, 0x30, 0xa8,
	0x79, 0xe2, 0x1b, 0xbf, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x20, 0x5a, 0xb6, 0x51, 0x1e, 0x03,
	0x00, 0x00,
}
