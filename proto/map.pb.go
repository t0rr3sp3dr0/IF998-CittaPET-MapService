// Code generated by protoc-gen-go. DO NOT EDIT.
// source: map.proto

package proto

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

type GetBusesLocationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBusesLocationRequest) Reset()         { *m = GetBusesLocationRequest{} }
func (m *GetBusesLocationRequest) String() string { return proto.CompactTextString(m) }
func (*GetBusesLocationRequest) ProtoMessage()    {}
func (*GetBusesLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_c7b3c2ba6c13e130, []int{0}
}
func (m *GetBusesLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusesLocationRequest.Unmarshal(m, b)
}
func (m *GetBusesLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusesLocationRequest.Marshal(b, m, deterministic)
}
func (dst *GetBusesLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusesLocationRequest.Merge(dst, src)
}
func (m *GetBusesLocationRequest) XXX_Size() int {
	return xxx_messageInfo_GetBusesLocationRequest.Size(m)
}
func (m *GetBusesLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusesLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusesLocationRequest proto.InternalMessageInfo

type GetBusesLocationResponse struct {
	BusesLocation        []*GetBusesLocationResponse_BusLocation `protobuf:"bytes,1,rep,name=buses_location,json=busesLocation,proto3" json:"buses_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *GetBusesLocationResponse) Reset()         { *m = GetBusesLocationResponse{} }
func (m *GetBusesLocationResponse) String() string { return proto.CompactTextString(m) }
func (*GetBusesLocationResponse) ProtoMessage()    {}
func (*GetBusesLocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_c7b3c2ba6c13e130, []int{1}
}
func (m *GetBusesLocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusesLocationResponse.Unmarshal(m, b)
}
func (m *GetBusesLocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusesLocationResponse.Marshal(b, m, deterministic)
}
func (dst *GetBusesLocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusesLocationResponse.Merge(dst, src)
}
func (m *GetBusesLocationResponse) XXX_Size() int {
	return xxx_messageInfo_GetBusesLocationResponse.Size(m)
}
func (m *GetBusesLocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusesLocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusesLocationResponse proto.InternalMessageInfo

func (m *GetBusesLocationResponse) GetBusesLocation() []*GetBusesLocationResponse_BusLocation {
	if m != nil {
		return m.BusesLocation
	}
	return nil
}

type GetBusesLocationResponse_BusLocation struct {
	Unit                 string   `protobuf:"bytes,1,opt,name=unit,proto3" json:"unit,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Coordinate           *LatLng  `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBusesLocationResponse_BusLocation) Reset()         { *m = GetBusesLocationResponse_BusLocation{} }
func (m *GetBusesLocationResponse_BusLocation) String() string { return proto.CompactTextString(m) }
func (*GetBusesLocationResponse_BusLocation) ProtoMessage()    {}
func (*GetBusesLocationResponse_BusLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_c7b3c2ba6c13e130, []int{1, 0}
}
func (m *GetBusesLocationResponse_BusLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBusesLocationResponse_BusLocation.Unmarshal(m, b)
}
func (m *GetBusesLocationResponse_BusLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBusesLocationResponse_BusLocation.Marshal(b, m, deterministic)
}
func (dst *GetBusesLocationResponse_BusLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBusesLocationResponse_BusLocation.Merge(dst, src)
}
func (m *GetBusesLocationResponse_BusLocation) XXX_Size() int {
	return xxx_messageInfo_GetBusesLocationResponse_BusLocation.Size(m)
}
func (m *GetBusesLocationResponse_BusLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBusesLocationResponse_BusLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GetBusesLocationResponse_BusLocation proto.InternalMessageInfo

func (m *GetBusesLocationResponse_BusLocation) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GetBusesLocationResponse_BusLocation) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *GetBusesLocationResponse_BusLocation) GetCoordinate() *LatLng {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

// An object representing a latitude/longitude pair. This is expressed as a pair
// of doubles representing degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude            float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatLng) Reset()         { *m = LatLng{} }
func (m *LatLng) String() string { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()    {}
func (*LatLng) Descriptor() ([]byte, []int) {
	return fileDescriptor_map_c7b3c2ba6c13e130, []int{2}
}
func (m *LatLng) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatLng.Unmarshal(m, b)
}
func (m *LatLng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatLng.Marshal(b, m, deterministic)
}
func (dst *LatLng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatLng.Merge(dst, src)
}
func (m *LatLng) XXX_Size() int {
	return xxx_messageInfo_LatLng.Size(m)
}
func (m *LatLng) XXX_DiscardUnknown() {
	xxx_messageInfo_LatLng.DiscardUnknown(m)
}

var xxx_messageInfo_LatLng proto.InternalMessageInfo

func (m *LatLng) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *LatLng) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*GetBusesLocationRequest)(nil), "proto.GetBusesLocationRequest")
	proto.RegisterType((*GetBusesLocationResponse)(nil), "proto.GetBusesLocationResponse")
	proto.RegisterType((*GetBusesLocationResponse_BusLocation)(nil), "proto.GetBusesLocationResponse.BusLocation")
	proto.RegisterType((*LatLng)(nil), "proto.LatLng")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BucketClient is the client API for Bucket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BucketClient interface {
	GetBusesLocation(ctx context.Context, in *GetBusesLocationRequest, opts ...grpc.CallOption) (*GetBusesLocationResponse, error)
}

type bucketClient struct {
	cc *grpc.ClientConn
}

func NewBucketClient(cc *grpc.ClientConn) BucketClient {
	return &bucketClient{cc}
}

func (c *bucketClient) GetBusesLocation(ctx context.Context, in *GetBusesLocationRequest, opts ...grpc.CallOption) (*GetBusesLocationResponse, error) {
	out := new(GetBusesLocationResponse)
	err := c.cc.Invoke(ctx, "/proto.Bucket/GetBusesLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BucketServer is the server API for Bucket service.
type BucketServer interface {
	GetBusesLocation(context.Context, *GetBusesLocationRequest) (*GetBusesLocationResponse, error)
}

func RegisterBucketServer(s *grpc.Server, srv BucketServer) {
	s.RegisterService(&_Bucket_serviceDesc, srv)
}

func _Bucket_GetBusesLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusesLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServer).GetBusesLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bucket/GetBusesLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServer).GetBusesLocation(ctx, req.(*GetBusesLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bucket_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Bucket",
	HandlerType: (*BucketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusesLocation",
			Handler:    _Bucket_GetBusesLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "map.proto",
}

func init() { proto.RegisterFile("map.proto", fileDescriptor_map_c7b3c2ba6c13e130) }

var fileDescriptor_map_c7b3c2ba6c13e130 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xcd, 0xae, 0x16, 0x3b, 0x65, 0x45, 0x72, 0x31, 0x16, 0xd1, 0xd2, 0x53, 0x41, 0xec,
	0xa1, 0xfe, 0x83, 0x5e, 0xbc, 0xf4, 0x14, 0xf0, 0xbc, 0xa4, 0xdd, 0x61, 0x09, 0xb6, 0x49, 0xdd,
	0x4c, 0xfe, 0xae, 0xbf, 0x45, 0x36, 0x5d, 0x6c, 0x51, 0x16, 0x4f, 0x49, 0xde, 0xf7, 0x92, 0x97,
	0x79, 0x10, 0x0f, 0x6a, 0x2c, 0xc7, 0x83, 0x25, 0xcb, 0xaf, 0xc2, 0x92, 0xdf, 0xc3, 0xdd, 0x1b,
	0x52, 0xed, 0x1d, 0xba, 0xc6, 0x76, 0x8a, 0xb4, 0x35, 0x12, 0x3f, 0x3d, 0x3a, 0xca, 0xbf, 0x18,
	0x88, 0xbf, 0xcc, 0x8d, 0xd6, 0x38, 0xe4, 0x12, 0x6e, 0xda, 0x23, 0xd8, 0xf6, 0x27, 0x22, 0x58,
	0xb6, 0x2e, 0x92, 0xea, 0x79, 0x7a, 0xbe, 0x3c, 0x77, 0xb1, 0xac, 0xfd, 0xac, 0x6d, 0xda, 0xa5,
	0x25, 0x35, 0x90, 0x2c, 0x28, 0xe7, 0x70, 0xe9, 0x8d, 0x26, 0xc1, 0x32, 0x56, 0xc4, 0x32, 0xec,
	0xf9, 0x03, 0xc4, 0xa4, 0x07, 0x74, 0xa4, 0x86, 0x51, 0xac, 0x02, 0x98, 0x05, 0xfe, 0x02, 0xd0,
	0x59, 0x7b, 0xd8, 0x69, 0xa3, 0x08, 0xc5, 0x3a, 0x63, 0x45, 0x52, 0x6d, 0x4e, 0x1f, 0x6a, 0x14,
	0x35, 0x66, 0x2f, 0x17, 0x86, 0xbc, 0x86, 0x68, 0x52, 0x79, 0x0a, 0xd7, 0xbd, 0x22, 0x4d, 0x7e,
	0x87, 0x21, 0x8e, 0xc9, 0x9f, 0xf3, 0x31, 0xb2, 0xb7, 0x66, 0x3f, 0xc1, 0x55, 0x80, 0xb3, 0x50,
	0x6d, 0x21, 0xaa, 0x7d, 0xf7, 0x81, 0xc4, 0xdf, 0xe1, 0xf6, 0xf7, 0xd0, 0xfc, 0xf1, 0x6c, 0x1b,
	0xa1, 0xe2, 0xf4, 0xe9, 0x9f, 0xb6, 0xf2, 0x8b, 0x36, 0x0a, 0x8e, 0xd7, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xde, 0x21, 0x52, 0x39, 0xbb, 0x01, 0x00, 0x00,
}
