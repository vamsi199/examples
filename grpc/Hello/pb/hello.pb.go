// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	DuplexOut
	DuplexOut1
	HealthCheckRequest
	HealthCheckResponse
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type DuplexOut struct {
	Event []byte `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (m *DuplexOut) Reset()                    { *m = DuplexOut{} }
func (m *DuplexOut) String() string            { return proto.CompactTextString(m) }
func (*DuplexOut) ProtoMessage()               {}
func (*DuplexOut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DuplexOut) GetEvent() []byte {
	if m != nil {
		return m.Event
	}
	return nil
}

type DuplexOut1 struct {
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *DuplexOut1) Reset()                    { *m = DuplexOut1{} }
func (m *DuplexOut1) String() string            { return proto.CompactTextString(m) }
func (*DuplexOut1) ProtoMessage()               {}
func (*DuplexOut1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DuplexOut1) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=hello.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*DuplexOut)(nil), "hello.DuplexOut")
	proto.RegisterType((*DuplexOut1)(nil), "hello.DuplexOut1")
	proto.RegisterType((*HealthCheckRequest)(nil), "hello.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "hello.HealthCheckResponse")
	proto.RegisterEnum("hello.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	Duplexstream(ctx context.Context, opts ...grpc.CallOption) (Hello_DuplexstreamClient, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Duplexstream(ctx context.Context, opts ...grpc.CallOption) (Hello_DuplexstreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hello_serviceDesc.Streams[0], c.cc, "/hello.Hello/Duplexstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloDuplexstreamClient{stream}
	return x, nil
}

type Hello_DuplexstreamClient interface {
	Send(*DuplexOut) error
	Recv() (*DuplexOut1, error)
	grpc.ClientStream
}

type helloDuplexstreamClient struct {
	grpc.ClientStream
}

func (x *helloDuplexstreamClient) Send(m *DuplexOut) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloDuplexstreamClient) Recv() (*DuplexOut1, error) {
	m := new(DuplexOut1)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hello service

type HelloServer interface {
	Duplexstream(Hello_DuplexstreamServer) error
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Duplexstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).Duplexstream(&helloDuplexstreamServer{stream})
}

type Hello_DuplexstreamServer interface {
	Send(*DuplexOut1) error
	Recv() (*DuplexOut, error)
	grpc.ServerStream
}

type helloDuplexstreamServer struct {
	grpc.ServerStream
}

func (x *helloDuplexstreamServer) Send(m *DuplexOut1) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloDuplexstreamServer) Recv() (*DuplexOut, error) {
	m := new(DuplexOut)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Duplexstream",
			Handler:       _Hello_Duplexstream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}

// Client API for Health service

type HealthClient interface {
	// checks the health for the client and server connection
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/hello.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	// checks the health for the client and server connection
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0xbb, 0x26, 0xb4, 0x76, 0xa8, 0x8a, 0xa3, 0x87, 0xca, 0x49, 0xf7, 0xc4, 0x89, 0x58,
	0x3c, 0xe9, 0xc1, 0x18, 0xff, 0x44, 0x8c, 0x09, 0x24, 0x8b, 0x7f, 0x8e, 0x06, 0x9b, 0x89, 0x18,
	0x11, 0x90, 0x5d, 0x1a, 0xbf, 0x87, 0x5f, 0xd8, 0xb0, 0x5b, 0x48, 0x6a, 0xea, 0xf1, 0x37, 0xbc,
	0x37, 0xbc, 0x37, 0x0b, 0x76, 0x46, 0x79, 0x5e, 0xfa, 0x55, 0x5d, 0xaa, 0x12, 0x2d, 0x0d, 0xfc,
	0x08, 0xc6, 0xd7, 0x4d, 0x95, 0xd3, 0x77, 0xdc, 0x28, 0xdc, 0x07, 0x8b, 0x16, 0x54, 0xa8, 0x29,
	0x3b, 0x64, 0xde, 0x44, 0x18, 0xe0, 0x1e, 0x40, 0x2f, 0x99, 0xa1, 0x0b, 0x9b, 0x82, 0x64, 0x55,
	0x16, 0x92, 0xb4, 0x6c, 0x2c, 0x7a, 0xe6, 0x3e, 0x60, 0x48, 0x69, 0xae, 0xb2, 0xab, 0x8c, 0xe6,
	0x1f, 0x82, 0xbe, 0x1a, 0x92, 0x0a, 0xa7, 0x30, 0x92, 0x54, 0x2f, 0xde, 0xe7, 0x9d, 0xa1, 0x43,
	0xfe, 0xc3, 0x60, 0x6f, 0xc5, 0x60, 0xf6, 0xe0, 0x05, 0x0c, 0xa5, 0x4a, 0x55, 0x23, 0xb5, 0x61,
	0x3b, 0xf0, 0x7c, 0x93, 0x7c, 0x8d, 0xd6, 0x4f, 0xda, 0x5d, 0xc5, 0x5b, 0xa2, 0xf5, 0x62, 0xe9,
	0xe3, 0x67, 0xb0, 0xb5, 0xf2, 0x01, 0x6d, 0x18, 0x3d, 0x46, 0xf7, 0x51, 0xfc, 0x1c, 0x39, 0x83,
	0x16, 0x92, 0x1b, 0xf1, 0x74, 0x17, 0xdd, 0x3a, 0x0c, 0x77, 0xc0, 0x8e, 0xe2, 0x87, 0x97, 0x6e,
	0xb0, 0x11, 0x5c, 0x82, 0x15, 0xb6, 0xbf, 0xc3, 0x53, 0x98, 0x98, 0xe2, 0x52, 0xd5, 0x94, 0x7e,
	0xa2, 0xb3, 0x8c, 0xd1, 0x5f, 0xc3, 0xdd, 0xfd, 0x3b, 0x99, 0xf1, 0x81, 0xc7, 0x8e, 0x59, 0x10,
	0xc2, 0xd0, 0x84, 0xc5, 0x73, 0xb0, 0x74, 0x60, 0x3c, 0x58, 0x57, 0x42, 0x5f, 0xc8, 0x75, 0xff,
	0xef, 0xf7, 0x3a, 0xd4, 0xcf, 0x75, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x7e, 0x85, 0xaa,
	0xbd, 0x01, 0x00, 0x00,
}
