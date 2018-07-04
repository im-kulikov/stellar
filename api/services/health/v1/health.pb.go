// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/health/v1/health.proto

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/health/v1/health.proto

It has these top-level messages:
	HealthResponse
*/
package health

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type HealthResponse struct {
	OSName    string `protobuf:"bytes,1,opt,name=os_name,json=osName,proto3" json:"os_name,omitempty"`
	OSVersion string `protobuf:"bytes,2,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	// TODO: use gogoproto.stdtime (returning panic: message/group field time.Time:bytes without pointer when trying to use)
	StartedAt   *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt" json:"started_at,omitempty"`
	Cpus        int64                       `protobuf:"varint,4,opt,name=cpus,proto3" json:"cpus,omitempty"`
	MemoryTotal int64                       `protobuf:"varint,5,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	MemoryFree  int64                       `protobuf:"varint,6,opt,name=memory_free,json=memoryFree,proto3" json:"memory_free,omitempty"`
	MemoryUsed  int64                       `protobuf:"varint,7,opt,name=memory_used,json=memoryUsed,proto3" json:"memory_used,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptorHealth, []int{0} }

func (m *HealthResponse) GetOSName() string {
	if m != nil {
		return m.OSName
	}
	return ""
}

func (m *HealthResponse) GetOSVersion() string {
	if m != nil {
		return m.OSVersion
	}
	return ""
}

func (m *HealthResponse) GetStartedAt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *HealthResponse) GetCpus() int64 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *HealthResponse) GetMemoryTotal() int64 {
	if m != nil {
		return m.MemoryTotal
	}
	return 0
}

func (m *HealthResponse) GetMemoryFree() int64 {
	if m != nil {
		return m.MemoryFree
	}
	return 0
}

func (m *HealthResponse) GetMemoryUsed() int64 {
	if m != nil {
		return m.MemoryUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*HealthResponse)(nil), "stellar.services.health.v1.HealthResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Health(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Health(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := grpc.Invoke(ctx, "/stellar.services.health.v1.Health/Health", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Health(context.Context, *google_protobuf1.Empty) (*HealthResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
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
		return srv.(HealthServer).Health(ctx, req.(*google_protobuf1.Empty))
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
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/health/v1/health.proto", fileDescriptorHealth)
}

var fileDescriptorHealth = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6b, 0xe2, 0x40,
	0x14, 0xc6, 0x89, 0xba, 0x91, 0x8c, 0xbb, 0x7b, 0x18, 0x96, 0x25, 0x64, 0x0f, 0x71, 0x77, 0x2f,
	0xb2, 0x2c, 0x33, 0xe8, 0x9e, 0xc4, 0xc3, 0x52, 0xa1, 0xa5, 0x87, 0x52, 0x21, 0x5a, 0x0f, 0xbd,
	0x84, 0x51, 0x9f, 0x49, 0x20, 0xe3, 0x84, 0x99, 0x49, 0xc0, 0xfe, 0x23, 0xfd, 0xef, 0x3c, 0xf8,
	0x97, 0x94, 0xcc, 0xc4, 0x62, 0x5b, 0x7a, 0xe8, 0x29, 0xdf, 0xfb, 0xbe, 0x5f, 0x78, 0x1f, 0x2f,
	0x41, 0xff, 0x93, 0x4c, 0xa7, 0xe5, 0x8a, 0xac, 0x05, 0xa7, 0x90, 0xb2, 0x87, 0x1c, 0xb4, 0xa6,
	0x4a, 0x43, 0x9e, 0x33, 0x49, 0x59, 0x91, 0x51, 0x05, 0xb2, 0xca, 0xd6, 0xa0, 0x68, 0x0a, 0x2c,
	0xd7, 0x29, 0xad, 0x86, 0x8d, 0x22, 0x85, 0x14, 0x5a, 0xe0, 0xa0, 0x81, 0xc9, 0x09, 0x24, 0x4d,
	0x5c, 0x0d, 0x83, 0x6f, 0x89, 0x48, 0x84, 0xc1, 0x68, 0xad, 0xec, 0x1b, 0xc1, 0x8f, 0x44, 0x88,
	0x24, 0x07, 0x6a, 0xa6, 0x55, 0xb9, 0xa5, 0xc0, 0x0b, 0xbd, 0x6f, 0xc2, 0xf0, 0x75, 0xa8, 0x33,
	0x0e, 0x4a, 0x33, 0x5e, 0x58, 0xe0, 0xd7, 0x63, 0x0b, 0x7d, 0xbd, 0x36, 0x1b, 0x22, 0x50, 0x85,
	0xd8, 0x29, 0xc0, 0xbf, 0x51, 0x57, 0xa8, 0x78, 0xc7, 0x38, 0xf8, 0x4e, 0xdf, 0x19, 0x78, 0x53,
	0x74, 0x3c, 0x84, 0xee, 0x6c, 0x7e, 0xcb, 0x38, 0x44, 0xae, 0x50, 0xf5, 0x13, 0xff, 0x45, 0x48,
	0xa8, 0xb8, 0x02, 0xa9, 0x32, 0xb1, 0xf3, 0x5b, 0x86, 0xfb, 0x72, 0x3c, 0x84, 0xde, 0x6c, 0xbe,
	0xb4, 0x66, 0xe4, 0x09, 0xd5, 0x48, 0x3c, 0x46, 0x48, 0x69, 0x26, 0x35, 0x6c, 0x62, 0xa6, 0xfd,
	0x76, 0xdf, 0x19, 0xf4, 0x46, 0x01, 0xb1, 0xdd, 0xc8, 0xa9, 0x1b, 0x59, 0x9c, 0xba, 0x45, 0x5e,
	0x43, 0x5f, 0x68, 0x8c, 0x51, 0x67, 0x5d, 0x94, 0xca, 0xef, 0xf4, 0x9d, 0x41, 0x3b, 0x32, 0x1a,
	0xff, 0x44, 0x9f, 0x39, 0x70, 0x21, 0xf7, 0xb1, 0x16, 0x9a, 0xe5, 0xfe, 0x27, 0x93, 0xf5, 0xac,
	0xb7, 0xa8, 0x2d, 0x1c, 0xa2, 0x66, 0x8c, 0xb7, 0x12, 0xc0, 0x77, 0x0d, 0x81, 0xac, 0x75, 0x25,
	0x01, 0xce, 0x80, 0x52, 0xc1, 0xc6, 0xef, 0x9e, 0x03, 0x77, 0x0a, 0x36, 0xa3, 0x25, 0x72, 0xed,
	0x61, 0xf0, 0xcd, 0xb3, 0xfa, 0xfe, 0xa6, 0xf3, 0x65, 0x7d, 0xec, 0xe0, 0x0f, 0x79, 0xff, 0xb3,
	0x91, 0x97, 0xe7, 0x9d, 0x4e, 0xee, 0xc7, 0x1f, 0xff, 0x49, 0x26, 0x56, 0xad, 0x5c, 0xb3, 0xf8,
	0xdf, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xad, 0x11, 0x87, 0x68, 0x02, 0x00, 0x00,
}
