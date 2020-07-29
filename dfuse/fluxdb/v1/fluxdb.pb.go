// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/fluxdb/v1/fluxdb.proto

package pbfluxdb

import (
	fmt "fmt"
	v1 "github.com/dfuse-io/pbgo/dfuse/bstream/v1"
	proto "github.com/golang/protobuf/proto"
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

type TabletIndex struct {
	SquelchedCount       uint64              `protobuf:"varint,1,opt,name=squelched_count,json=squelchedCount,proto3" json:"squelched_count,omitempty"`
	Entries              []*TabletIndexEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TabletIndex) Reset()         { *m = TabletIndex{} }
func (m *TabletIndex) String() string { return proto.CompactTextString(m) }
func (*TabletIndex) ProtoMessage()    {}
func (*TabletIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_74afbad2e99b142c, []int{0}
}

func (m *TabletIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TabletIndex.Unmarshal(m, b)
}
func (m *TabletIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TabletIndex.Marshal(b, m, deterministic)
}
func (m *TabletIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TabletIndex.Merge(m, src)
}
func (m *TabletIndex) XXX_Size() int {
	return xxx_messageInfo_TabletIndex.Size(m)
}
func (m *TabletIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_TabletIndex.DiscardUnknown(m)
}

var xxx_messageInfo_TabletIndex proto.InternalMessageInfo

func (m *TabletIndex) GetSquelchedCount() uint64 {
	if m != nil {
		return m.SquelchedCount
	}
	return 0
}

func (m *TabletIndex) GetEntries() []*TabletIndexEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TabletIndexEntry struct {
	PrimaryKey           []byte   `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TabletIndexEntry) Reset()         { *m = TabletIndexEntry{} }
func (m *TabletIndexEntry) String() string { return proto.CompactTextString(m) }
func (*TabletIndexEntry) ProtoMessage()    {}
func (*TabletIndexEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_74afbad2e99b142c, []int{1}
}

func (m *TabletIndexEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TabletIndexEntry.Unmarshal(m, b)
}
func (m *TabletIndexEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TabletIndexEntry.Marshal(b, m, deterministic)
}
func (m *TabletIndexEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TabletIndexEntry.Merge(m, src)
}
func (m *TabletIndexEntry) XXX_Size() int {
	return xxx_messageInfo_TabletIndexEntry.Size(m)
}
func (m *TabletIndexEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TabletIndexEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TabletIndexEntry proto.InternalMessageInfo

func (m *TabletIndexEntry) GetPrimaryKey() []byte {
	if m != nil {
		return m.PrimaryKey
	}
	return nil
}

func (m *TabletIndexEntry) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// Row represents a FluxDB storage row as seen by the underlying storage
// engine which is a Key-Value store.
//
// A FluxDB row is always of the form:
// ```
// <collection>/<tablet>/<height>/<primaryKey> => <payload>
// ```
type Row struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	TabletKey            string   `protobuf:"bytes,2,opt,name=tablet_key,json=tabletKey,proto3" json:"tablet_key,omitempty"`
	HeightKey            string   `protobuf:"bytes,3,opt,name=height_key,json=heightKey,proto3" json:"height_key,omitempty"`
	PrimKey              string   `protobuf:"bytes,4,opt,name=prim_key,json=primKey,proto3" json:"prim_key,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_74afbad2e99b142c, []int{2}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *Row) GetTabletKey() string {
	if m != nil {
		return m.TabletKey
	}
	return ""
}

func (m *Row) GetHeightKey() string {
	if m != nil {
		return m.HeightKey
	}
	return ""
}

func (m *Row) GetPrimKey() string {
	if m != nil {
		return m.PrimKey
	}
	return ""
}

func (m *Row) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Checkpoint represents where FluxDB did its last write
// information to the database. It mainly contains the latest
// height value written as well as the block reference where it
// was.
type Checkpoint struct {
	Height               uint64       `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Block                *v1.BlockRef `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_74afbad2e99b142c, []int{3}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Checkpoint) GetBlock() *v1.BlockRef {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*TabletIndex)(nil), "dfuse.fluxdb.v1.TabletIndex")
	proto.RegisterType((*TabletIndexEntry)(nil), "dfuse.fluxdb.v1.TabletIndexEntry")
	proto.RegisterType((*Row)(nil), "dfuse.fluxdb.v1.Row")
	proto.RegisterType((*Checkpoint)(nil), "dfuse.fluxdb.v1.Checkpoint")
}

func init() { proto.RegisterFile("dfuse/fluxdb/v1/fluxdb.proto", fileDescriptor_74afbad2e99b142c) }

var fileDescriptor_74afbad2e99b142c = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4d, 0x6f, 0xaa, 0x40,
	0x14, 0x86, 0x83, 0x9f, 0xd7, 0xc3, 0xcd, 0xf5, 0x66, 0x16, 0x0d, 0x35, 0xd5, 0x5a, 0x36, 0x75,
	0x53, 0x28, 0xba, 0x74, 0xa7, 0xe9, 0xa2, 0x71, 0x47, 0x9a, 0x2e, 0xba, 0x31, 0x0c, 0x8c, 0x42,
	0x1c, 0x19, 0x0a, 0x83, 0x95, 0x7f, 0xd2, 0x9f, 0xdb, 0xcc, 0x19, 0x6c, 0x88, 0xbb, 0x73, 0xde,
	0xe7, 0x7c, 0xbc, 0xc3, 0x01, 0xee, 0xa2, 0x5d, 0x59, 0x30, 0x77, 0xc7, 0xcb, 0x73, 0x44, 0xdd,
	0x93, 0x57, 0x47, 0x4e, 0x96, 0x0b, 0x29, 0xc8, 0x10, 0xa9, 0x53, 0x6b, 0x27, 0x6f, 0x34, 0xd1,
	0xe5, 0xb4, 0x90, 0x39, 0x0b, 0x8e, 0xaa, 0xbe, 0x0e, 0x75, 0x83, 0x5d, 0x80, 0xf9, 0x16, 0x50,
	0xce, 0xe4, 0x6b, 0x1a, 0xb1, 0x33, 0x79, 0x84, 0x61, 0xf1, 0x59, 0x32, 0x1e, 0xc6, 0x2c, 0xda,
	0x86, 0xa2, 0x4c, 0xa5, 0x65, 0x4c, 0x8d, 0x59, 0xc7, 0xff, 0xf7, 0x2b, 0xaf, 0x95, 0x4a, 0x96,
	0xd0, 0x67, 0xa9, 0xcc, 0x13, 0x56, 0x58, 0xad, 0x69, 0x7b, 0x66, 0xce, 0x1f, 0x9c, 0xab, 0xd5,
	0x4e, 0x63, 0xee, 0x4b, 0x2a, 0xf3, 0xca, 0xbf, 0x74, 0xd8, 0x1b, 0xf8, 0x7f, 0x0d, 0xc9, 0x3d,
	0x98, 0x59, 0x9e, 0x1c, 0x83, 0xbc, 0xda, 0x1e, 0x58, 0x85, 0x5b, 0xff, 0xfa, 0x50, 0x4b, 0x1b,
	0x56, 0x91, 0x1b, 0xe8, 0xc5, 0x2c, 0xd9, 0xc7, 0xd2, 0x6a, 0xa1, 0xa3, 0x3a, 0xb3, 0xbf, 0x0d,
	0x68, 0xfb, 0xe2, 0x8b, 0x4c, 0x00, 0x42, 0xc1, 0x39, 0x0b, 0x65, 0x22, 0x52, 0xec, 0x1f, 0xf8,
	0x0d, 0x85, 0x8c, 0x01, 0x24, 0x2e, 0xc5, 0xf9, 0x2d, 0xe4, 0x03, 0xad, 0xa8, 0xf1, 0x63, 0x00,
	0x3d, 0x10, 0x71, 0x5b, 0x63, 0xad, 0x28, 0x7c, 0x0b, 0x7f, 0x94, 0x17, 0x84, 0x1d, 0x84, 0x7d,
	0x95, 0x2b, 0x64, 0x41, 0x3f, 0x0b, 0x2a, 0x2e, 0x82, 0xc8, 0xea, 0xa2, 0xeb, 0x4b, 0x6a, 0xbf,
	0x03, 0xac, 0x63, 0x16, 0x1e, 0x32, 0x91, 0xa4, 0xb2, 0xf1, 0x00, 0xa3, 0xf9, 0x00, 0xf2, 0x0c,
	0x5d, 0xca, 0x45, 0x78, 0x40, 0x4f, 0xe6, 0x7c, 0x54, 0x7f, 0xc8, 0xcb, 0x9d, 0x4e, 0x9e, 0xb3,
	0x52, 0xd8, 0x67, 0x3b, 0x5f, 0x17, 0xae, 0x16, 0x1f, 0xde, 0x3e, 0x91, 0x71, 0x49, 0x9d, 0x50,
	0x1c, 0x5d, 0x2c, 0x7f, 0x4a, 0x84, 0x9b, 0xd1, 0xbd, 0x70, 0xaf, 0x7e, 0x8f, 0x65, 0x46, 0x75,
	0x4c, 0x7b, 0x78, 0xf0, 0xc5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x98, 0x2f, 0x24, 0x41,
	0x02, 0x00, 0x00,
}
