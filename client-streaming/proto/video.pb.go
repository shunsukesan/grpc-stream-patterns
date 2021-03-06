// Code generated by protoc-gen-go. DO NOT EDIT.
// source: video.proto

package video

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

type UploadRequest struct {
	VideoData            []byte   `protobuf:"bytes,1,opt,name=video_data,json=videoData,proto3" json:"video_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{0}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetVideoData() []byte {
	if m != nil {
		return m.VideoData
	}
	return nil
}

type UploadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad4ea8866efb1e3, []int{1}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UploadRequest)(nil), "video.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "video.UploadResponse")
}

func init() { proto.RegisterFile("video.proto", fileDescriptor_0ad4ea8866efb1e3) }

var fileDescriptor_0ad4ea8866efb1e3 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xcb, 0x4c, 0x49,
	0xcd, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xf4, 0xb8, 0x78, 0x43,
	0x0b, 0x72, 0xf2, 0x13, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb9, 0xb8,
	0xc0, 0x32, 0xf1, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x9c, 0x60,
	0x11, 0x97, 0xc4, 0x92, 0x44, 0x25, 0x01, 0x2e, 0x3e, 0x98, 0xfa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x23, 0x4f, 0x2e, 0x9e, 0x30, 0x90, 0x74, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90,
	0x25, 0x17, 0x1b, 0x44, 0x85, 0x90, 0x88, 0x1e, 0xc4, 0x42, 0x14, 0x0b, 0xa4, 0x44, 0xd1, 0x44,
	0x21, 0xc6, 0x28, 0x31, 0x68, 0x30, 0x26, 0xb1, 0x81, 0x9d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x52, 0x55, 0xb6, 0x16, 0xa9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (VideoService_UploadClient, error)
}

type videoServiceClient struct {
	cc *grpc.ClientConn
}

func NewVideoServiceClient(cc *grpc.ClientConn) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (VideoService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VideoService_serviceDesc.Streams[0], "/video.VideoService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoServiceUploadClient{stream}
	return x, nil
}

type VideoService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type videoServiceUploadClient struct {
	grpc.ClientStream
}

func (x *videoServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoServiceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoServiceServer is the server API for VideoService service.
type VideoServiceServer interface {
	Upload(VideoService_UploadServer) error
}

func RegisterVideoServiceServer(s *grpc.Server, srv VideoServiceServer) {
	s.RegisterService(&_VideoService_serviceDesc, srv)
}

func _VideoService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoServiceServer).Upload(&videoServiceUploadServer{stream})
}

type VideoService_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type videoServiceUploadServer struct {
	grpc.ServerStream
}

func (x *videoServiceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VideoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _VideoService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "video.proto",
}
