// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dns.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	dns.proto

It has these top-level messages:
	DnsPacket
	WatchRequest
	WatchCreateRequest
	WatchCancelRequest
	WatchResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
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

type DnsPacket struct {
	Msg []byte `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *DnsPacket) Reset()                    { *m = DnsPacket{} }
func (m *DnsPacket) String() string            { return proto.CompactTextString(m) }
func (*DnsPacket) ProtoMessage()               {}
func (*DnsPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DnsPacket) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type WatchRequest struct {
	// request_union is a request to either create a new watcher or cancel an existing watcher.
	//
	// Types that are valid to be assigned to RequestUnion:
	//	*WatchRequest_CreateRequest
	//	*WatchRequest_CancelRequest
	RequestUnion isWatchRequest_RequestUnion `protobuf_oneof:"request_union"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
}

type WatchRequest_CreateRequest struct {
	CreateRequest *WatchCreateRequest `protobuf:"bytes,1,opt,name=create_request,json=createRequest,oneof"`
}
type WatchRequest_CancelRequest struct {
	CancelRequest *WatchCancelRequest `protobuf:"bytes,2,opt,name=cancel_request,json=cancelRequest,oneof"`
}

func (*WatchRequest_CreateRequest) isWatchRequest_RequestUnion() {}
func (*WatchRequest_CancelRequest) isWatchRequest_RequestUnion() {}

func (m *WatchRequest) GetRequestUnion() isWatchRequest_RequestUnion {
	if m != nil {
		return m.RequestUnion
	}
	return nil
}

func (m *WatchRequest) GetCreateRequest() *WatchCreateRequest {
	if x, ok := m.GetRequestUnion().(*WatchRequest_CreateRequest); ok {
		return x.CreateRequest
	}
	return nil
}

func (m *WatchRequest) GetCancelRequest() *WatchCancelRequest {
	if x, ok := m.GetRequestUnion().(*WatchRequest_CancelRequest); ok {
		return x.CancelRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WatchRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WatchRequest_OneofMarshaler, _WatchRequest_OneofUnmarshaler, _WatchRequest_OneofSizer, []interface{}{
		(*WatchRequest_CreateRequest)(nil),
		(*WatchRequest_CancelRequest)(nil),
	}
}

func _WatchRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WatchRequest)
	// request_union
	switch x := m.RequestUnion.(type) {
	case *WatchRequest_CreateRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateRequest); err != nil {
			return err
		}
	case *WatchRequest_CancelRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CancelRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WatchRequest.RequestUnion has unexpected type %T", x)
	}
	return nil
}

func _WatchRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WatchRequest)
	switch tag {
	case 1: // request_union.create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WatchCreateRequest)
		err := b.DecodeMessage(msg)
		m.RequestUnion = &WatchRequest_CreateRequest{msg}
		return true, err
	case 2: // request_union.cancel_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WatchCancelRequest)
		err := b.DecodeMessage(msg)
		m.RequestUnion = &WatchRequest_CancelRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WatchRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WatchRequest)
	// request_union
	switch x := m.RequestUnion.(type) {
	case *WatchRequest_CreateRequest:
		s := proto.Size(x.CreateRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WatchRequest_CancelRequest:
		s := proto.Size(x.CancelRequest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WatchCreateRequest struct {
	Query *DnsPacket `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *WatchCreateRequest) Reset()                    { *m = WatchCreateRequest{} }
func (m *WatchCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchCreateRequest) ProtoMessage()               {}
func (*WatchCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WatchCreateRequest) GetQuery() *DnsPacket {
	if m != nil {
		return m.Query
	}
	return nil
}

type WatchCancelRequest struct {
	// watch_id is the watcher id to cancel
	WatchId int64 `protobuf:"varint,1,opt,name=watch_id,json=watchId" json:"watch_id,omitempty"`
}

func (m *WatchCancelRequest) Reset()                    { *m = WatchCancelRequest{} }
func (m *WatchCancelRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchCancelRequest) ProtoMessage()               {}
func (*WatchCancelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WatchCancelRequest) GetWatchId() int64 {
	if m != nil {
		return m.WatchId
	}
	return 0
}

type WatchResponse struct {
	// watch_id is the ID of the watcher that corresponds to the response.
	WatchId int64 `protobuf:"varint,1,opt,name=watch_id,json=watchId" json:"watch_id,omitempty"`
	// created is set to true if the response is for a create watch request.
	// The client should record the watch_id and expect to receive DNS replies
	// from the same stream.
	// All replies sent to the created watcher will attach with the same watch_id.
	Created bool `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	// canceled is set to true if the response is for a cancel watch request.
	// No further events will be sent to the canceled watcher.
	Canceled bool   `protobuf:"varint,3,opt,name=canceled" json:"canceled,omitempty"`
	Qname    string `protobuf:"bytes,4,opt,name=qname" json:"qname,omitempty"`
	Err      string `protobuf:"bytes,5,opt,name=err" json:"err,omitempty"`
}

func (m *WatchResponse) Reset()                    { *m = WatchResponse{} }
func (m *WatchResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()               {}
func (*WatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WatchResponse) GetWatchId() int64 {
	if m != nil {
		return m.WatchId
	}
	return 0
}

func (m *WatchResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *WatchResponse) GetCanceled() bool {
	if m != nil {
		return m.Canceled
	}
	return false
}

func (m *WatchResponse) GetQname() string {
	if m != nil {
		return m.Qname
	}
	return ""
}

func (m *WatchResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*DnsPacket)(nil), "coredns.dns.DnsPacket")
	proto.RegisterType((*WatchRequest)(nil), "coredns.dns.WatchRequest")
	proto.RegisterType((*WatchCreateRequest)(nil), "coredns.dns.WatchCreateRequest")
	proto.RegisterType((*WatchCancelRequest)(nil), "coredns.dns.WatchCancelRequest")
	proto.RegisterType((*WatchResponse)(nil), "coredns.dns.WatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DnsService service

type DnsServiceClient interface {
	Query(ctx context.Context, in *DnsPacket, opts ...grpc.CallOption) (*DnsPacket, error)
	Watch(ctx context.Context, opts ...grpc.CallOption) (DnsService_WatchClient, error)
}

type dnsServiceClient struct {
	cc *grpc.ClientConn
}

func NewDnsServiceClient(cc *grpc.ClientConn) DnsServiceClient {
	return &dnsServiceClient{cc}
}

func (c *dnsServiceClient) Query(ctx context.Context, in *DnsPacket, opts ...grpc.CallOption) (*DnsPacket, error) {
	out := new(DnsPacket)
	err := grpc.Invoke(ctx, "/coredns.dns.DnsService/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsServiceClient) Watch(ctx context.Context, opts ...grpc.CallOption) (DnsService_WatchClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DnsService_serviceDesc.Streams[0], c.cc, "/coredns.dns.DnsService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &dnsServiceWatchClient{stream}
	return x, nil
}

type DnsService_WatchClient interface {
	Send(*WatchRequest) error
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type dnsServiceWatchClient struct {
	grpc.ClientStream
}

func (x *dnsServiceWatchClient) Send(m *WatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dnsServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DnsService service

type DnsServiceServer interface {
	Query(context.Context, *DnsPacket) (*DnsPacket, error)
	Watch(DnsService_WatchServer) error
}

func RegisterDnsServiceServer(s *grpc.Server, srv DnsServiceServer) {
	s.RegisterService(&_DnsService_serviceDesc, srv)
}

func _DnsService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coredns.dns.DnsService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServiceServer).Query(ctx, req.(*DnsPacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DnsServiceServer).Watch(&dnsServiceWatchServer{stream})
}

type DnsService_WatchServer interface {
	Send(*WatchResponse) error
	Recv() (*WatchRequest, error)
	grpc.ServerStream
}

type dnsServiceWatchServer struct {
	grpc.ServerStream
}

func (x *dnsServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dnsServiceWatchServer) Recv() (*WatchRequest, error) {
	m := new(WatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DnsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coredns.dns.DnsService",
	HandlerType: (*DnsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DnsService_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _DnsService_Watch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dns.proto",
}

func init() { proto.RegisterFile("dns.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4f, 0x32, 0x31,
	0x14, 0xfc, 0x0a, 0xec, 0x07, 0x3c, 0x40, 0xcd, 0x8b, 0x31, 0xcb, 0x26, 0x46, 0xb2, 0x27, 0x0e,
	0x06, 0x0d, 0x1e, 0xbc, 0xaf, 0x1c, 0xf0, 0xa6, 0xf5, 0x60, 0xe2, 0x85, 0x2c, 0xdd, 0x17, 0x25,
	0x4a, 0x17, 0xda, 0x45, 0xe3, 0x3f, 0xd0, 0xdf, 0xe3, 0x1f, 0x34, 0xfb, 0xba, 0x90, 0x35, 0x88,
	0xb7, 0xce, 0x74, 0x3a, 0xed, 0xcc, 0x2b, 0x34, 0x13, 0x6d, 0x07, 0x0b, 0x93, 0x66, 0x29, 0xb6,
	0x54, 0x6a, 0x28, 0x87, 0x89, 0xb6, 0xe1, 0x31, 0x34, 0x47, 0xda, 0xde, 0xc4, 0xea, 0x99, 0x32,
	0x3c, 0x80, 0xea, 0xdc, 0x3e, 0xfa, 0xa2, 0x27, 0xfa, 0x6d, 0x99, 0x2f, 0xc3, 0x2f, 0x01, 0xed,
	0xfb, 0x38, 0x53, 0x4f, 0x92, 0x96, 0x2b, 0xb2, 0x19, 0x8e, 0x61, 0x4f, 0x19, 0x8a, 0x33, 0x9a,
	0x18, 0xc7, 0xb0, 0xba, 0x35, 0x3c, 0x19, 0x94, 0x5c, 0x07, 0x7c, 0xe4, 0x8a, 0x75, 0xc5, 0xc1,
	0xf1, 0x3f, 0xd9, 0x51, 0x65, 0x82, 0x9d, 0x62, 0xad, 0xe8, 0x65, 0xe3, 0x54, 0xd9, 0xe9, 0xc4,
	0xba, 0xb2, 0x53, 0x99, 0x88, 0xf6, 0xa1, 0x53, 0x58, 0x4c, 0x56, 0x7a, 0x96, 0xea, 0x30, 0x02,
	0xdc, 0x7e, 0x01, 0x9e, 0x82, 0xb7, 0x5c, 0x91, 0x79, 0x2f, 0x5e, 0x7c, 0xf4, 0xe3, 0x9e, 0x4d,
	0x09, 0xd2, 0x89, 0xc2, 0xb3, 0xb5, 0x47, 0xf9, 0x2a, 0xec, 0x42, 0xe3, 0x2d, 0x67, 0x27, 0xb3,
	0x84, 0x6d, 0xaa, 0xb2, 0xce, 0xf8, 0x3a, 0x09, 0x3f, 0x04, 0x74, 0x8a, 0xaa, 0xec, 0x22, 0xd5,
	0x96, 0xfe, 0x10, 0xa3, 0x0f, 0x75, 0xd7, 0x46, 0xc2, 0xa9, 0x1b, 0x72, 0x0d, 0x31, 0x80, 0x86,
	0x4b, 0x47, 0x89, 0x5f, 0xe5, 0xad, 0x0d, 0xc6, 0x43, 0xf0, 0x96, 0x3a, 0x9e, 0x93, 0x5f, 0xeb,
	0x89, 0x7e, 0x53, 0x3a, 0x90, 0x4f, 0x8d, 0x8c, 0xf1, 0x3d, 0xe6, 0xf2, 0xe5, 0xf0, 0x53, 0x00,
	0x8c, 0xb4, 0xbd, 0x23, 0xf3, 0x3a, 0x53, 0x84, 0x97, 0xe0, 0xdd, 0xe6, 0x99, 0x70, 0x47, 0xe4,
	0x60, 0x07, 0x8f, 0x11, 0x78, 0x9c, 0x08, 0xbb, 0xdb, 0x33, 0x29, 0x1a, 0x09, 0x82, 0xdf, 0xb6,
	0x5c, 0x01, 0x7d, 0x71, 0x2e, 0xa2, 0xda, 0x43, 0x65, 0x31, 0x9d, 0xfe, 0xe7, 0xaf, 0x77, 0xf1,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x5b, 0x8c, 0xe1, 0x87, 0x02, 0x00, 0x00,
}
