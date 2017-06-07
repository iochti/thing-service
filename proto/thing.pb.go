// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thing.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	thing.proto

It has these top-level messages:
	ThingIDArray
	Thing
	ThingIDRequest
	ThingDeletion
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
	ID          int32  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
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

func (m *Thing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Thing) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
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

type ThingDeletion struct {
	ID int32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *ThingDeletion) Reset()                    { *m = ThingDeletion{} }
func (m *ThingDeletion) String() string            { return proto1.CompactTextString(m) }
func (*ThingDeletion) ProtoMessage()               {}
func (*ThingDeletion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ThingDeletion) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto1.RegisterType((*ThingIDArray)(nil), "proto.ThingIDArray")
	proto1.RegisterType((*Thing)(nil), "proto.Thing")
	proto1.RegisterType((*ThingIDRequest)(nil), "proto.ThingIDRequest")
	proto1.RegisterType((*ThingDeletion)(nil), "proto.ThingDeletion")
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
	CreateThing(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error)
	UpdateThing(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error)
	DeleteThing(ctx context.Context, in *ThingIDRequest, opts ...grpc.CallOption) (*ThingDeletion, error)
	BulkDeleteThing(ctx context.Context, in *ThingIDArray, opts ...grpc.CallOption) (*ThingDeletion, error)
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

func (c *thingSvcClient) DeleteThing(ctx context.Context, in *ThingIDRequest, opts ...grpc.CallOption) (*ThingDeletion, error) {
	out := new(ThingDeletion)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/DeleteThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingSvcClient) BulkDeleteThing(ctx context.Context, in *ThingIDArray, opts ...grpc.CallOption) (*ThingDeletion, error) {
	out := new(ThingDeletion)
	err := grpc.Invoke(ctx, "/proto.ThingSvc/BulkDeleteThing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ThingSvc service

type ThingSvcServer interface {
	GetThing(context.Context, *ThingIDRequest) (*Thing, error)
	CreateThing(context.Context, *Thing) (*Thing, error)
	UpdateThing(context.Context, *Thing) (*Thing, error)
	DeleteThing(context.Context, *ThingIDRequest) (*ThingDeletion, error)
	BulkDeleteThing(context.Context, *ThingIDArray) (*ThingDeletion, error)
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
	Streams:  []grpc.StreamDesc{},
	Metadata: "thing.proto",
}

func init() { proto1.RegisterFile("thing.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc9, 0xc8, 0xcc,
	0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6a, 0x5c, 0x3c, 0x21,
	0x20, 0x51, 0x4f, 0x17, 0xc7, 0xa2, 0xa2, 0xc4, 0x4a, 0x21, 0x31, 0x2e, 0x36, 0xb0, 0xaa, 0x62,
	0x09, 0x46, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x28, 0x4f, 0xc9, 0x97, 0x8b, 0x15, 0xac, 0x4e, 0x88,
	0x8f, 0x8b, 0xc9, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0xc9, 0xd3, 0x45, 0x48,
	0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16,
	0x52, 0xe0, 0xe2, 0x4e, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0xcc, 0xcf, 0x93, 0x60,
	0x06, 0x4b, 0x21, 0x0b, 0x29, 0x29, 0x70, 0xf1, 0x41, 0xad, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x41, 0x37, 0x57, 0x49, 0x9e, 0x8b, 0x17, 0xac, 0xc2, 0x25, 0x35, 0x27, 0x15, 0xa4, 0x05,
	0x5d, 0x81, 0xd1, 0x54, 0x26, 0x2e, 0x0e, 0xb0, 0x8a, 0xe0, 0xb2, 0x64, 0x21, 0x43, 0x2e, 0x0e,
	0xf7, 0xd4, 0x12, 0x88, 0x0b, 0x45, 0x21, 0x3e, 0xd4, 0x43, 0xb5, 0x40, 0x8a, 0x07, 0x59, 0x58,
	0x89, 0x41, 0x48, 0x9b, 0x8b, 0xdb, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x15, 0xa2, 0x0b, 0x45, 0x1a,
	0x9b, 0xe2, 0xd0, 0x82, 0x14, 0x22, 0x15, 0xdb, 0x70, 0x71, 0x83, 0x5d, 0x9d, 0x8a, 0xd7, 0x3d,
	0x22, 0xc8, 0xc2, 0x30, 0x5f, 0x2a, 0x31, 0x08, 0xd9, 0x71, 0xf1, 0x3b, 0x95, 0xe6, 0x64, 0x23,
	0x9b, 0x20, 0x8c, 0x6a, 0x02, 0x38, 0xa6, 0x70, 0xe9, 0x4f, 0x62, 0x03, 0x0b, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa2, 0xc4, 0x82, 0x2c, 0xee, 0x01, 0x00, 0x00,
}
