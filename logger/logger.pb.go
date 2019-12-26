// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logger.proto

package logger

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Event struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Activity             string   `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Event) GetActivity() string {
	if m != nil {
		return m.Activity
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "logger.Event")
}

func init() { proto.RegisterFile("logger.proto", fileDescriptor_d43b7bfc6b6f7b16) }

var fileDescriptor_d43b7bfc6b6f7b16 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xc9, 0x4f, 0x4f,
	0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0xa4, 0xd3, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xa2, 0x49, 0xa5, 0x69, 0xfa, 0xae, 0xb9, 0x05, 0x25, 0x95,
	0x10, 0x45, 0x4a, 0x36, 0x5c, 0xac, 0xae, 0x65, 0xa9, 0x79, 0x25, 0x42, 0xe2, 0x5c, 0xec, 0xa5,
	0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x6c, 0x20, 0xae,
	0x67, 0x8a, 0x90, 0x14, 0x17, 0x47, 0x62, 0x72, 0x49, 0x66, 0x59, 0x66, 0x49, 0xa5, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x6f, 0x64, 0xcb, 0xc5, 0xe6, 0x03, 0xb6, 0x44, 0xc8, 0x98,
	0x8b, 0xc3, 0x27, 0x3f, 0x1d, 0x62, 0x14, 0xaf, 0x1e, 0xd4, 0x1d, 0x60, 0xae, 0x94, 0x98, 0x1e,
	0xc4, 0x01, 0x7a, 0x30, 0x07, 0xe8, 0x81, 0x1d, 0xa0, 0xc4, 0x90, 0xc4, 0x06, 0x16, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x10, 0xa8, 0x4a, 0xb8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggerClient interface {
	LogEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
}

type loggerClient struct {
	cc *grpc.ClientConn
}

func NewLoggerClient(cc *grpc.ClientConn) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) LogEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/logger.Logger/LogEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
type LoggerServer interface {
	LogEvent(context.Context, *Event) (*empty.Empty, error)
}

// UnimplementedLoggerServer can be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (*UnimplementedLoggerServer) LogEvent(ctx context.Context, req *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogEvent not implemented")
}

func RegisterLoggerServer(s *grpc.Server, srv LoggerServer) {
	s.RegisterService(&_Logger_serviceDesc, srv)
}

func _Logger_LogEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).LogEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logger.Logger/LogEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).LogEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logger.Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogEvent",
			Handler:    _Logger_LogEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logger.proto",
}
