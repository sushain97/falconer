// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jacquard.proto

/*
Package jacquard is a generated protocol buffer package.

It is generated from these files:
	jacquard.proto

It has these top-level messages:
	FindSpanRequest
	SpanBatch
	SpanResponse
	TraceRequest
*/
package jacquard

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ssf "github.com/stripe/veneur/ssf"

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

type FindSpanRequest struct {
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FindSpanRequest) Reset()                    { *m = FindSpanRequest{} }
func (m *FindSpanRequest) String() string            { return proto.CompactTextString(m) }
func (*FindSpanRequest) ProtoMessage()               {}
func (*FindSpanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FindSpanRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SpanBatch struct {
	Spans []*ssf.SSFSpan `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
}

func (m *SpanBatch) Reset()                    { *m = SpanBatch{} }
func (m *SpanBatch) String() string            { return proto.CompactTextString(m) }
func (*SpanBatch) ProtoMessage()               {}
func (*SpanBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SpanBatch) GetSpans() []*ssf.SSFSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

type SpanResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *SpanResponse) Reset()                    { *m = SpanResponse{} }
func (m *SpanResponse) String() string            { return proto.CompactTextString(m) }
func (*SpanResponse) ProtoMessage()               {}
func (*SpanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpanResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type TraceRequest struct {
	TraceID int64 `protobuf:"varint,1,opt,name=traceID" json:"traceID,omitempty"`
}

func (m *TraceRequest) Reset()                    { *m = TraceRequest{} }
func (m *TraceRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()               {}
func (*TraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TraceRequest) GetTraceID() int64 {
	if m != nil {
		return m.TraceID
	}
	return 0
}

func init() {
	proto.RegisterType((*FindSpanRequest)(nil), "jacquard.FindSpanRequest")
	proto.RegisterType((*SpanBatch)(nil), "jacquard.SpanBatch")
	proto.RegisterType((*SpanResponse)(nil), "jacquard.SpanResponse")
	proto.RegisterType((*TraceRequest)(nil), "jacquard.TraceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Jacquard service

type JacquardClient interface {
	SendSpans(ctx context.Context, opts ...grpc.CallOption) (Jacquard_SendSpansClient, error)
	FindSpans(ctx context.Context, in *FindSpanRequest, opts ...grpc.CallOption) (Jacquard_FindSpansClient, error)
	GetTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Jacquard_GetTraceClient, error)
}

type jacquardClient struct {
	cc *grpc.ClientConn
}

func NewJacquardClient(cc *grpc.ClientConn) JacquardClient {
	return &jacquardClient{cc}
}

func (c *jacquardClient) SendSpans(ctx context.Context, opts ...grpc.CallOption) (Jacquard_SendSpansClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Jacquard_serviceDesc.Streams[0], c.cc, "/jacquard.Jacquard/SendSpans", opts...)
	if err != nil {
		return nil, err
	}
	x := &jacquardSendSpansClient{stream}
	return x, nil
}

type Jacquard_SendSpansClient interface {
	Send(*SpanBatch) error
	CloseAndRecv() (*SpanResponse, error)
	grpc.ClientStream
}

type jacquardSendSpansClient struct {
	grpc.ClientStream
}

func (x *jacquardSendSpansClient) Send(m *SpanBatch) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jacquardSendSpansClient) CloseAndRecv() (*SpanResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SpanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jacquardClient) FindSpans(ctx context.Context, in *FindSpanRequest, opts ...grpc.CallOption) (Jacquard_FindSpansClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Jacquard_serviceDesc.Streams[1], c.cc, "/jacquard.Jacquard/FindSpans", opts...)
	if err != nil {
		return nil, err
	}
	x := &jacquardFindSpansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Jacquard_FindSpansClient interface {
	Recv() (*ssf.SSFSpan, error)
	grpc.ClientStream
}

type jacquardFindSpansClient struct {
	grpc.ClientStream
}

func (x *jacquardFindSpansClient) Recv() (*ssf.SSFSpan, error) {
	m := new(ssf.SSFSpan)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jacquardClient) GetTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Jacquard_GetTraceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Jacquard_serviceDesc.Streams[2], c.cc, "/jacquard.Jacquard/GetTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &jacquardGetTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Jacquard_GetTraceClient interface {
	Recv() (*ssf.SSFSpan, error)
	grpc.ClientStream
}

type jacquardGetTraceClient struct {
	grpc.ClientStream
}

func (x *jacquardGetTraceClient) Recv() (*ssf.SSFSpan, error) {
	m := new(ssf.SSFSpan)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Jacquard service

type JacquardServer interface {
	SendSpans(Jacquard_SendSpansServer) error
	FindSpans(*FindSpanRequest, Jacquard_FindSpansServer) error
	GetTrace(*TraceRequest, Jacquard_GetTraceServer) error
}

func RegisterJacquardServer(s *grpc.Server, srv JacquardServer) {
	s.RegisterService(&_Jacquard_serviceDesc, srv)
}

func _Jacquard_SendSpans_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JacquardServer).SendSpans(&jacquardSendSpansServer{stream})
}

type Jacquard_SendSpansServer interface {
	SendAndClose(*SpanResponse) error
	Recv() (*SpanBatch, error)
	grpc.ServerStream
}

type jacquardSendSpansServer struct {
	grpc.ServerStream
}

func (x *jacquardSendSpansServer) SendAndClose(m *SpanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jacquardSendSpansServer) Recv() (*SpanBatch, error) {
	m := new(SpanBatch)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Jacquard_FindSpans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindSpanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JacquardServer).FindSpans(m, &jacquardFindSpansServer{stream})
}

type Jacquard_FindSpansServer interface {
	Send(*ssf.SSFSpan) error
	grpc.ServerStream
}

type jacquardFindSpansServer struct {
	grpc.ServerStream
}

func (x *jacquardFindSpansServer) Send(m *ssf.SSFSpan) error {
	return x.ServerStream.SendMsg(m)
}

func _Jacquard_GetTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JacquardServer).GetTrace(m, &jacquardGetTraceServer{stream})
}

type Jacquard_GetTraceServer interface {
	Send(*ssf.SSFSpan) error
	grpc.ServerStream
}

type jacquardGetTraceServer struct {
	grpc.ServerStream
}

func (x *jacquardGetTraceServer) Send(m *ssf.SSFSpan) error {
	return x.ServerStream.SendMsg(m)
}

var _Jacquard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jacquard.Jacquard",
	HandlerType: (*JacquardServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendSpans",
			Handler:       _Jacquard_SendSpans_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindSpans",
			Handler:       _Jacquard_FindSpans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTrace",
			Handler:       _Jacquard_GetTrace_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "jacquard.proto",
}

func init() { proto.RegisterFile("jacquard.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x4e, 0xea, 0x40,
	0x14, 0x66, 0xe0, 0x72, 0x6f, 0x7b, 0x2e, 0x51, 0x33, 0x1a, 0x53, 0xbb, 0x22, 0x75, 0x53, 0x5d,
	0xb4, 0x06, 0x4d, 0x20, 0xc6, 0x95, 0x51, 0x8c, 0x2e, 0x5b, 0x5e, 0x60, 0x28, 0x87, 0x52, 0x85,
	0x69, 0x99, 0x33, 0x25, 0x61, 0xed, 0x4b, 0xf9, 0x78, 0x86, 0x96, 0x56, 0xd0, 0xb8, 0x9b, 0xef,
	0xcc, 0xf9, 0xe6, 0xfb, 0x19, 0x38, 0x78, 0x15, 0xd1, 0x32, 0x17, 0x6a, 0xe2, 0x65, 0x2a, 0xd5,
	0x29, 0x37, 0x2a, 0x6c, 0x5f, 0xc4, 0x89, 0x9e, 0xe5, 0x63, 0x2f, 0x4a, 0x17, 0x3e, 0x69, 0x95,
	0x64, 0xe8, 0xaf, 0x50, 0x62, 0xae, 0x7c, 0xa2, 0xa9, 0x4f, 0x62, 0x91, 0xcd, 0xb1, 0x24, 0x39,
	0xef, 0x0c, 0x0e, 0x87, 0x89, 0x9c, 0x84, 0x99, 0x90, 0x01, 0x2e, 0x73, 0x24, 0xcd, 0xfb, 0xf0,
	0x47, 0x8b, 0x98, 0x2c, 0xd6, 0x6d, 0xb9, 0xff, 0x7b, 0xe7, 0x5e, 0xad, 0xf3, 0x6d, 0xd1, 0x1b,
	0x89, 0x98, 0x1e, 0xa5, 0x56, 0xeb, 0xa0, 0x20, 0xd8, 0x7d, 0x30, 0xeb, 0x11, 0x3f, 0x82, 0xd6,
	0x1b, 0xae, 0x2d, 0xd6, 0x65, 0xae, 0x19, 0x6c, 0x8e, 0xfc, 0x04, 0xda, 0x2b, 0x31, 0xcf, 0xd1,
	0x6a, 0x16, 0xb3, 0x12, 0xdc, 0x36, 0x07, 0xcc, 0xf1, 0xc1, 0xdc, 0xbc, 0x7b, 0x2f, 0x74, 0x34,
	0xe3, 0x0e, 0xb4, 0x29, 0x13, 0xb2, 0xd2, 0xef, 0x78, 0x44, 0x53, 0x2f, 0x0c, 0x87, 0x85, 0x72,
	0x79, 0xe5, 0x5c, 0x42, 0xa7, 0x34, 0x42, 0x59, 0x2a, 0x09, 0xb9, 0x0d, 0x46, 0xac, 0x10, 0x75,
	0x22, 0xe3, 0xad, 0x62, 0x8d, 0x1d, 0x17, 0x3a, 0x23, 0x25, 0x22, 0xac, 0xe2, 0x59, 0xf0, 0x4f,
	0x6f, 0xf0, 0xf3, 0x43, 0xb1, 0xda, 0x0a, 0x2a, 0xd8, 0xfb, 0x60, 0x60, 0xbc, 0x6c, 0xc3, 0xf2,
	0x3b, 0x30, 0x43, 0x2c, 0xf3, 0x12, 0x3f, 0xfe, 0x2a, 0xa1, 0x36, 0x6a, 0x9f, 0xee, 0x0f, 0x2b,
	0x33, 0x4e, 0xc3, 0x65, 0x7c, 0x00, 0x66, 0xd5, 0x16, 0xf1, 0xb3, 0x5f, 0x2b, 0xb4, 0xf7, 0xd2,
	0x39, 0x8d, 0x2b, 0xc6, 0x6f, 0xc0, 0x78, 0x42, 0x5d, 0x38, 0xe6, 0x3b, 0x0a, 0xbb, 0x11, 0x7e,
	0xb2, 0xc6, 0x7f, 0x8b, 0xef, 0xbc, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x52, 0x7e, 0x99, 0xcf,
	0x15, 0x02, 0x00, 0x00,
}
