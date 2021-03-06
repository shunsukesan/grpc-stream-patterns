// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notification.proto

package notification

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type NotificationRequest struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationRequest) Reset()         { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string { return proto.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()    {}
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{0}
}

func (m *NotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationRequest.Unmarshal(m, b)
}
func (m *NotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationRequest.Marshal(b, m, deterministic)
}
func (m *NotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationRequest.Merge(m, src)
}
func (m *NotificationRequest) XXX_Size() int {
	return xxx_messageInfo_NotificationRequest.Size(m)
}
func (m *NotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationRequest proto.InternalMessageInfo

func (m *NotificationRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type NotificationResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationResponse) Reset()         { *m = NotificationResponse{} }
func (m *NotificationResponse) String() string { return proto.CompactTextString(m) }
func (*NotificationResponse) ProtoMessage()    {}
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_736a457d4a5efa07, []int{1}
}

func (m *NotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationResponse.Unmarshal(m, b)
}
func (m *NotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationResponse.Marshal(b, m, deterministic)
}
func (m *NotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationResponse.Merge(m, src)
}
func (m *NotificationResponse) XXX_Size() int {
	return xxx_messageInfo_NotificationResponse.Size(m)
}
func (m *NotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationResponse proto.InternalMessageInfo

func (m *NotificationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*NotificationRequest)(nil), "notification.NotificationRequest")
	proto.RegisterType((*NotificationResponse)(nil), "notification.NotificationResponse")
}

func init() { proto.RegisterFile("notification.proto", fileDescriptor_736a457d4a5efa07) }

var fileDescriptor_736a457d4a5efa07 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcb, 0x2f, 0xc9,
	0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x41, 0x16, 0x53, 0x52, 0xe7, 0x12, 0xf6, 0x43, 0xe2, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x70, 0x31, 0xe7, 0x95, 0xe6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98,
	0x4a, 0x06, 0x5c, 0x22, 0xa8, 0x0a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xc1, 0xaa, 0x39, 0x83, 0x60, 0x5c, 0xa3, 0x02, 0x54,
	0xa3, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x22, 0xb9, 0x78, 0x90, 0x85, 0x85, 0x14,
	0xf5, 0x50, 0x1c, 0x89, 0xc5, 0x35, 0x52, 0x4a, 0xf8, 0x94, 0x40, 0xdc, 0xa1, 0xc4, 0x60, 0xc0,
	0x98, 0xc4, 0x06, 0xf6, 0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xd8, 0x03, 0x6c, 0xf7,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Notification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (NotificationService_NotificationClient, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Notification(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (NotificationService_NotificationClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NotificationService_serviceDesc.Streams[0], "/notification.NotificationService/Notification", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceNotificationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_NotificationClient interface {
	Recv() (*NotificationResponse, error)
	grpc.ClientStream
}

type notificationServiceNotificationClient struct {
	grpc.ClientStream
}

func (x *notificationServiceNotificationClient) Recv() (*NotificationResponse, error) {
	m := new(NotificationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Notification(*NotificationRequest, NotificationService_NotificationServer) error
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Notification_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NotificationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).Notification(m, &notificationServiceNotificationServer{stream})
}

type NotificationService_NotificationServer interface {
	Send(*NotificationResponse) error
	grpc.ServerStream
}

type notificationServiceNotificationServer struct {
	grpc.ServerStream
}

func (x *notificationServiceNotificationServer) Send(m *NotificationResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Notification",
			Handler:       _NotificationService_Notification_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "notification.proto",
}
