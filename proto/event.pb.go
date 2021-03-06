// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BusEvent struct {
	Unit                 string   `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Coordinate           *UTM     `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusEvent) Reset()         { *m = BusEvent{} }
func (m *BusEvent) String() string { return proto.CompactTextString(m) }
func (*BusEvent) ProtoMessage()    {}
func (*BusEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_5f4cd52dda648a81, []int{0}
}
func (m *BusEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusEvent.Unmarshal(m, b)
}
func (m *BusEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusEvent.Marshal(b, m, deterministic)
}
func (dst *BusEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusEvent.Merge(dst, src)
}
func (m *BusEvent) XXX_Size() int {
	return xxx_messageInfo_BusEvent.Size(m)
}
func (m *BusEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BusEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BusEvent proto.InternalMessageInfo

func (m *BusEvent) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *BusEvent) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *BusEvent) GetCoordinate() *UTM {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

type UTM struct {
	Easting              float64  `protobuf:"fixed64,1,opt,name=easting,proto3" json:"easting,omitempty"`
	Northing             float64  `protobuf:"fixed64,2,opt,name=northing,proto3" json:"northing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTM) Reset()         { *m = UTM{} }
func (m *UTM) String() string { return proto.CompactTextString(m) }
func (*UTM) ProtoMessage()    {}
func (*UTM) Descriptor() ([]byte, []int) {
	return fileDescriptor_event_5f4cd52dda648a81, []int{1}
}
func (m *UTM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTM.Unmarshal(m, b)
}
func (m *UTM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTM.Marshal(b, m, deterministic)
}
func (dst *UTM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTM.Merge(dst, src)
}
func (m *UTM) XXX_Size() int {
	return xxx_messageInfo_UTM.Size(m)
}
func (m *UTM) XXX_DiscardUnknown() {
	xxx_messageInfo_UTM.DiscardUnknown(m)
}

var xxx_messageInfo_UTM proto.InternalMessageInfo

func (m *UTM) GetEasting() float64 {
	if m != nil {
		return m.Easting
	}
	return 0
}

func (m *UTM) GetNorthing() float64 {
	if m != nil {
		return m.Northing
	}
	return 0
}

func init() {
	proto.RegisterType((*BusEvent)(nil), "proto.BusEvent")
	proto.RegisterType((*UTM)(nil), "proto.UTM")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_event_5f4cd52dda648a81) }

var fileDescriptor_event_5f4cd52dda648a81 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x19, 0x5c, 0x1c, 0x4e,
	0xa5, 0xc5, 0xae, 0x20, 0x09, 0x21, 0x21, 0x2e, 0x96, 0xd2, 0xbc, 0xcc, 0x12, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x86, 0x8b, 0xb3, 0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24,
	0x31, 0xb7, 0x40, 0x82, 0x09, 0x2c, 0x81, 0x10, 0x10, 0xd2, 0xe2, 0xe2, 0x4a, 0xce, 0xcf, 0x2f,
	0x4a, 0xc9, 0xcc, 0x4b, 0x2c, 0x49, 0x95, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0xe2, 0x82, 0x58,
	0xa0, 0x17, 0x1a, 0xe2, 0x1b, 0x84, 0x24, 0xab, 0x64, 0xcd, 0xc5, 0x1c, 0x1a, 0xe2, 0x2b, 0x24,
	0xc1, 0xc5, 0x9e, 0x9a, 0x58, 0x5c, 0x92, 0x99, 0x97, 0x0e, 0xb6, 0x87, 0x31, 0x08, 0xc6, 0x15,
	0x92, 0xe2, 0xe2, 0xc8, 0xcb, 0x2f, 0x2a, 0xc9, 0x00, 0x49, 0x31, 0x81, 0xa5, 0xe0, 0xfc, 0x24,
	0x36, 0xb0, 0x99, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x7e, 0x49, 0x57, 0xc3, 0x00,
	0x00, 0x00,
}
