// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/merger/v1/merger.proto

package pbmerger

import (
	context "context"
	fmt "fmt"
	v1 "github.com/dfuse-io/pbgo/dfuse/bstream/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	LowBlockNum          uint64   `protobuf:"varint,1,opt,name=lowBlockNum,proto3" json:"lowBlockNum,omitempty"`
	HighBlockID          string   `protobuf:"bytes,2,opt,name=highBlockID,proto3" json:"highBlockID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f1327e68c6db98d, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetLowBlockNum() uint64 {
	if m != nil {
		return m.LowBlockNum
	}
	return 0
}

func (m *Request) GetHighBlockID() string {
	if m != nil {
		return m.HighBlockID
	}
	return ""
}

type Response struct {
	Found                bool      `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Block                *v1.Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f1327e68c6db98d, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *Response) GetBlock() *v1.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "dfuse.merger.v1.Request")
	proto.RegisterType((*Response)(nil), "dfuse.merger.v1.Response")
}

func init() { proto.RegisterFile("dfuse/merger/v1/merger.proto", fileDescriptor_0f1327e68c6db98d) }

var fileDescriptor_0f1327e68c6db98d = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0x23, 0x88, 0xc3, 0x81, 0x64, 0x63, 0x62, 0x25, 0xc6, 0x34, 0x9c, 0xb8, 0xb0,
	0x6b, 0xe1, 0xe8, 0x8d, 0x78, 0xd0, 0x03, 0x6a, 0xf6, 0xe8, 0xcd, 0xa5, 0x43, 0xdb, 0x48, 0x99,
	0xba, 0xdb, 0xad, 0x7f, 0xdf, 0x74, 0xb6, 0x26, 0x84, 0x70, 0x9b, 0x79, 0xf3, 0xcd, 0xe4, 0xbd,
	0x81, 0xfb, 0x6c, 0xe7, 0x1d, 0xaa, 0x0a, 0x6d, 0x8e, 0x56, 0xb5, 0x69, 0x5f, 0xc9, 0xda, 0x52,
	0x43, 0x62, 0xc2, 0x53, 0xd9, 0x6b, 0x6d, 0x3a, 0x7d, 0x08, 0xb8, 0x71, 0x8d, 0xc5, 0xaf, 0xaa,
	0xe3, 0xfb, 0x32, 0x2c, 0xcc, 0x36, 0x70, 0xa5, 0xf1, 0xc7, 0xa3, 0x6b, 0x44, 0x02, 0xe3, 0x3d,
	0xfd, 0xae, 0xf7, 0xb4, 0xfd, 0x7e, 0xf3, 0x55, 0x1c, 0x25, 0xd1, 0xfc, 0x52, 0x1f, 0x4b, 0x1d,
	0x51, 0x94, 0x79, 0xc1, 0xfd, 0xeb, 0x73, 0x7c, 0x91, 0x44, 0xf3, 0x6b, 0x7d, 0x2c, 0xcd, 0xde,
	0x61, 0xa4, 0xd1, 0xd5, 0x74, 0x70, 0x28, 0x6e, 0x60, 0xb0, 0x23, 0x7f, 0xc8, 0xf8, 0xd2, 0x48,
	0x87, 0x46, 0x2c, 0x60, 0x60, 0x3a, 0x98, 0xb7, 0xc7, 0xcb, 0x5b, 0x19, 0x1c, 0xff, 0xbb, 0x6a,
	0x53, 0xc9, 0xb7, 0x74, 0xa0, 0x96, 0x1a, 0x86, 0x1b, 0x0e, 0x23, 0x5e, 0x60, 0xf2, 0x61, 0x91,
	0x9b, 0x8c, 0x11, 0x27, 0x62, 0x79, 0x12, 0x57, 0xf6, 0x59, 0xa6, 0x77, 0x67, 0x26, 0xc1, 0xd6,
	0x63, 0xb4, 0x5e, 0x7d, 0xa6, 0x79, 0xd9, 0x14, 0xde, 0xc8, 0x2d, 0x55, 0x8a, 0xc1, 0x45, 0x49,
	0xaa, 0x36, 0x39, 0xa9, 0x93, 0xef, 0x3e, 0xd5, 0x26, 0xd4, 0x66, 0xc8, 0xff, 0x5a, 0xfd, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x06, 0xfb, 0xcb, 0xb4, 0x80, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MergerClient is the client API for Merger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MergerClient interface {
	PreMergedBlocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Merger_PreMergedBlocksClient, error)
}

type mergerClient struct {
	cc *grpc.ClientConn
}

func NewMergerClient(cc *grpc.ClientConn) MergerClient {
	return &mergerClient{cc}
}

func (c *mergerClient) PreMergedBlocks(ctx context.Context, in *Request, opts ...grpc.CallOption) (Merger_PreMergedBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Merger_serviceDesc.Streams[0], "/dfuse.merger.v1.Merger/PreMergedBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &mergerPreMergedBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Merger_PreMergedBlocksClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type mergerPreMergedBlocksClient struct {
	grpc.ClientStream
}

func (x *mergerPreMergedBlocksClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MergerServer is the server API for Merger service.
type MergerServer interface {
	PreMergedBlocks(*Request, Merger_PreMergedBlocksServer) error
}

// UnimplementedMergerServer can be embedded to have forward compatible implementations.
type UnimplementedMergerServer struct {
}

func (*UnimplementedMergerServer) PreMergedBlocks(req *Request, srv Merger_PreMergedBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method PreMergedBlocks not implemented")
}

func RegisterMergerServer(s *grpc.Server, srv MergerServer) {
	s.RegisterService(&_Merger_serviceDesc, srv)
}

func _Merger_PreMergedBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MergerServer).PreMergedBlocks(m, &mergerPreMergedBlocksServer{stream})
}

type Merger_PreMergedBlocksServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type mergerPreMergedBlocksServer struct {
	grpc.ServerStream
}

func (x *mergerPreMergedBlocksServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Merger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.merger.v1.Merger",
	HandlerType: (*MergerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PreMergedBlocks",
			Handler:       _Merger_PreMergedBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/merger/v1/merger.proto",
}
