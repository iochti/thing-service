// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thing-service/proto/thing.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	thing-service/proto/thing.proto

It has these top-level messages:
	ThingIDArray
	Thing
	ThingIDRequest
	GroupRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ThingIDArray struct {
	Things []int32 `protobuf:"varint,1,rep,packed,name=things" json:"things,omitempty"`
}

func (m *ThingIDArray) Reset()                    { *m = ThingIDArray{} }
func (m *ThingIDArray) String() string            { return proto1.CompactTextString(m) }
func (*ThingIDArray) ProtoMessage()               {}
func (*ThingIDArray) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ThingIDArray) GetThings() []int32 {
	if m != nil {
		return m.Things
	}
	return nil
}

type Thing struct {
	ID        int32                      `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	GroupID   int32                      `protobuf:"varint,2,opt,name=GroupID" json:"GroupID,omitempty"`
	Name      string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	MAC       string                     `protobuf:"bytes,4,opt,name=MAC" json:"MAC,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Thing) Reset()                    { *m = Thing{} }
func (m *Thing) String() string            { return proto1.CompactTextString(m) }
func (*Thing) ProtoMessage()               {}
func (*Thing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Thing) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Thing) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *Thing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Thing) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func (m *Thing) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Thing) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ThingIDRequest struct {
	ID int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *ThingIDRequest) Reset()                    { *m = ThingIDRequest{} }
func (m *ThingIDRequest) String() string            { return proto1.CompactTextString(m) }
func (*ThingIDRequest) ProtoMessage()               {}
func (*ThingIDRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ThingIDRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GroupRequest struct {
	ID int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *GroupRequest) Reset()                    { *m = GroupRequest{} }
func (m *GroupRequest) String() string            { return proto1.CompactTextString(m) }
func (*GroupRequest) ProtoMessage()               {}
func (*GroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GroupRequest) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto1.RegisterType((*ThingIDArray)(nil), "proto.ThingIDArray")
	proto1.RegisterType((*Thing)(nil), "proto.Thing")
	proto1.RegisterType((*ThingIDRequest)(nil), "proto.ThingIDRequest")
	proto1.RegisterType((*GroupRequest)(nil), "proto.GroupRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ThingSvc service

type ThingSvcClient interface {
	GetThing(ctx context.Context, in *ThingIDRequest, opts ...grpc.CallOption) (*Thing, error)
	ListGroupThings(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (ThingSvc_ListGroupThingsClient, error)
	CreateThing(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error)
	UpdateThing(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error)
	DeleteThing(ctx context.Context, in *ThingIDRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	BulkDeleteThing(ctx context.Context, in *ThingIDArray, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type thingSvcClient struct {
	cc *grpc.ClientConn
}

func NewThingSvcClient(cc *grpc.ClientConn) ThingSvcClient {
	return &thingSvcClient{cc}
}

func (c *thingSvcClient) GetThing(ctx context.Context, in *ThingIDRequest, opts ...grpc.CallOption) (*Thing, error) {
	out := new(Thing)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/GetThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingSvcClient) ListGroupThings(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (ThingSvc_ListGroupThingsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ThingSvc_serviceDesc.Streams[0], c.cc, "/proto.ThingSvc/ListGroupThings", opts...)
	if err != nil {
		return nil, err
	}
	x := &thingSvcListGroupThingsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ThingSvc_ListGroupThingsClient interface {
	Recv() (*Thing, error)
	grpc.ClientStream
}

type thingSvcListGroupThingsClient struct {
	grpc.ClientStream
}

func (x *thingSvcListGroupThingsClient) Recv() (*Thing, error) {
	m := new(Thing)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *thingSvcClient) CreateThing(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error) {
	out := new(Thing)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/CreateThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingSvcClient) UpdateThing(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error) {
	out := new(Thing)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/UpdateThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingSvcClient) DeleteThing(ctx context.Context, in *ThingIDRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/DeleteThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingSvcClient) BulkDeleteThing(ctx context.Context, in *ThingIDArray, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/BulkDeleteThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ThingSvc service

type ThingSvcServer interface {
	GetThing(context.Context, *ThingIDRequest) (*Thing, error)
	ListGroupThings(*GroupRequest, ThingSvc_ListGroupThingsServer) error
	CreateThing(context.Context, *Thing) (*Thing, error)
	UpdateThing(context.Context, *Thing) (*Thing, error)
	DeleteThing(context.Context, *ThingIDRequest) (*google_protobuf1.Empty, error)
	BulkDeleteThing(context.Context, *ThingIDArray) (*google_protobuf1.Empty, error)
}

func RegisterThingSvcServer(s *grpc.Server, srv ThingSvcServer) {
	s.RegisterService(&_ThingSvc_serviceDesc, srv)
}

func _ThingSvc_GetThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingSvcServer).GetThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingSvc/GetThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingSvcServer).GetThing(ctx, req.(*ThingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingSvc_ListGroupThings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GroupRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ThingSvcServer).ListGroupThings(m, &thingSvcListGroupThingsServer{stream})
}

type ThingSvc_ListGroupThingsServer interface {
	Send(*Thing) error
	grpc.ServerStream
}

type thingSvcListGroupThingsServer struct {
	grpc.ServerStream
}

func (x *thingSvcListGroupThingsServer) Send(m *Thing) error {
	return x.ServerStream.SendMsg(m)
}

func _ThingSvc_CreateThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingSvcServer).CreateThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingSvc/CreateThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingSvcServer).CreateThing(ctx, req.(*Thing))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingSvc_UpdateThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingSvcServer).UpdateThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingSvc/UpdateThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingSvcServer).UpdateThing(ctx, req.(*Thing))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingSvc_DeleteThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThingIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingSvcServer).DeleteThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingSvc/DeleteThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingSvcServer).DeleteThing(ctx, req.(*ThingIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThingSvc_BulkDeleteThing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThingIDArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingSvcServer).BulkDeleteThing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThingSvc/BulkDeleteThing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingSvcServer).BulkDeleteThing(ctx, req.(*ThingIDArray))
	}
	return interceptor(ctx, in, info, handler)
}

var _ThingSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ThingSvc",
	HandlerType: (*ThingSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThing",
			Handler:    _ThingSvc_GetThing_Handler,
		},
		{
			MethodName: "CreateThing",
			Handler:    _ThingSvc_CreateThing_Handler,
		},
		{
			MethodName: "UpdateThing",
			Handler:    _ThingSvc_UpdateThing_Handler,
		},
		{
			MethodName: "DeleteThing",
			Handler:    _ThingSvc_DeleteThing_Handler,
		},
		{
			MethodName: "BulkDeleteThing",
			Handler:    _ThingSvc_BulkDeleteThing_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListGroupThings",
			Handler:       _ThingSvc_ListGroupThings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "thing-service/proto/thing.proto",
}

func init() { proto1.RegisterFile("thing-service/proto/thing.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0x85, 0x22, 0x4c, 0x09, 0x98, 0x25, 0x92, 0x4d, 0x4d, 0xa4, 0xe9, 0xc1, 0x34, 0x21,
	0x16, 0xc5, 0x8b, 0x5e, 0x8c, 0x95, 0x1a, 0xd2, 0x44, 0x2f, 0x15, 0xcf, 0xa6, 0xc0, 0x8a, 0x8d,
	0x94, 0xd6, 0x76, 0x4b, 0xc2, 0x8f, 0xf4, 0xe4, 0x1f, 0x32, 0x9d, 0xb6, 0x09, 0x1f, 0x21, 0x7a,
	0xda, 0xd9, 0xb7, 0x6f, 0xde, 0xdb, 0x37, 0x19, 0xe8, 0xf2, 0x0f, 0x6f, 0x39, 0xbf, 0x88, 0x59,
	0xb4, 0xf2, 0xa6, 0xac, 0x1f, 0x46, 0x01, 0x0f, 0xfa, 0x88, 0x19, 0x58, 0x13, 0x09, 0x0f, 0xa5,
	0x3b, 0x0f, 0x82, 0xf9, 0x22, 0x27, 0x4c, 0x92, 0xf7, 0x3e, 0xf7, 0x7c, 0x16, 0x73, 0xd7, 0x0f,
	0x33, 0x9e, 0x72, 0xba, 0x4b, 0x60, 0x7e, 0xc8, 0xd7, 0xd9, 0xa3, 0x76, 0x0e, 0x8d, 0x71, 0xaa,
	0x69, 0x5b, 0x66, 0x14, 0xb9, 0x6b, 0xd2, 0x81, 0x2a, 0x7a, 0xc4, 0x54, 0x50, 0xcb, 0xba, 0xe4,
	0xe4, 0x37, 0xed, 0x5b, 0x00, 0x09, 0x89, 0xa4, 0x09, 0xa2, 0x6d, 0x51, 0x41, 0x15, 0x74, 0xc9,
	0x11, 0x6d, 0x8b, 0x50, 0x38, 0x1a, 0x45, 0x41, 0x12, 0xda, 0x16, 0x15, 0x11, 0x2c, 0xae, 0x84,
	0x40, 0x65, 0xe9, 0xfa, 0x8c, 0x96, 0x55, 0x41, 0xaf, 0x3b, 0x58, 0x93, 0x63, 0x28, 0x3f, 0x9b,
	0x43, 0x5a, 0x41, 0x28, 0x2d, 0xc9, 0x2d, 0xc0, 0x34, 0x62, 0x2e, 0x67, 0xb3, 0x37, 0x97, 0x53,
	0x49, 0x15, 0x74, 0x79, 0xa0, 0x18, 0xd9, 0x9f, 0x8d, 0xe2, 0xcf, 0xc6, 0xb8, 0x08, 0xe5, 0xd4,
	0x73, 0xb6, 0xc9, 0xd3, 0xd6, 0x24, 0x9c, 0x15, 0xad, 0xd5, 0xbf, 0x5b, 0x73, 0xb6, 0xc9, 0x35,
	0x15, 0x9a, 0x79, 0x6e, 0x87, 0x7d, 0x25, 0x2c, 0xe6, 0xbb, 0xb9, 0xb4, 0x33, 0x68, 0x60, 0x90,
	0x03, 0xef, 0x83, 0x1f, 0x11, 0x6a, 0x28, 0xf1, 0xb2, 0x9a, 0x92, 0x2b, 0xa8, 0x8d, 0x18, 0xcf,
	0x06, 0x74, 0x92, 0x59, 0x1b, 0xdb, 0xfa, 0x4a, 0x63, 0x13, 0xd6, 0x4a, 0xe4, 0x06, 0x5a, 0x4f,
	0x5e, 0xcc, 0xd1, 0x03, 0xb1, 0x98, 0xb4, 0x73, 0xca, 0xa6, 0xef, 0x6e, 0xdf, 0xa5, 0x40, 0x7a,
	0x20, 0x0f, 0x71, 0x06, 0x99, 0xdf, 0x16, 0x61, 0xcf, 0xa6, 0x07, 0xf2, 0x2b, 0xa6, 0xfe, 0x0f,
	0xf9, 0x0e, 0x64, 0x8b, 0x2d, 0x58, 0x41, 0x3e, 0x90, 0xa4, 0xb3, 0x37, 0xe2, 0xc7, 0x74, 0xa3,
	0xb4, 0x12, 0xb9, 0x87, 0xd6, 0x43, 0xb2, 0xf8, 0xdc, 0xd4, 0x68, 0x6f, 0x6b, 0xe0, 0x96, 0x1d,
	0x56, 0x98, 0x54, 0x11, 0xb9, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x75, 0xc3, 0x45, 0xfe,
	0x02, 0x00, 0x00,
}
