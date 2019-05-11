// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/topic_view_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Request message for [TopicViewService.GetTopicView][google.ads.googleads.v1.services.TopicViewService.GetTopicView].
type GetTopicViewRequest struct {
	// The resource name of the topic view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicViewRequest) Reset()         { *m = GetTopicViewRequest{} }
func (m *GetTopicViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicViewRequest) ProtoMessage()    {}
func (*GetTopicViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_topic_view_service_42250129c59eb567, []int{0}
}
func (m *GetTopicViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicViewRequest.Unmarshal(m, b)
}
func (m *GetTopicViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicViewRequest.Marshal(b, m, deterministic)
}
func (dst *GetTopicViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicViewRequest.Merge(dst, src)
}
func (m *GetTopicViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicViewRequest.Size(m)
}
func (m *GetTopicViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicViewRequest proto.InternalMessageInfo

func (m *GetTopicViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTopicViewRequest)(nil), "google.ads.googleads.v1.services.GetTopicViewRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicViewServiceClient is the client API for TopicViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicViewServiceClient interface {
	// Returns the requested topic view in full detail.
	GetTopicView(ctx context.Context, in *GetTopicViewRequest, opts ...grpc.CallOption) (*resources.TopicView, error)
}

type topicViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopicViewServiceClient(cc *grpc.ClientConn) TopicViewServiceClient {
	return &topicViewServiceClient{cc}
}

func (c *topicViewServiceClient) GetTopicView(ctx context.Context, in *GetTopicViewRequest, opts ...grpc.CallOption) (*resources.TopicView, error) {
	out := new(resources.TopicView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.TopicViewService/GetTopicView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicViewServiceServer is the server API for TopicViewService service.
type TopicViewServiceServer interface {
	// Returns the requested topic view in full detail.
	GetTopicView(context.Context, *GetTopicViewRequest) (*resources.TopicView, error)
}

func RegisterTopicViewServiceServer(s *grpc.Server, srv TopicViewServiceServer) {
	s.RegisterService(&_TopicViewService_serviceDesc, srv)
}

func _TopicViewService_GetTopicView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicViewServiceServer).GetTopicView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.TopicViewService/GetTopicView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicViewServiceServer).GetTopicView(ctx, req.(*GetTopicViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.TopicViewService",
	HandlerType: (*TopicViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopicView",
			Handler:    _TopicViewService_GetTopicView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/topic_view_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/topic_view_service.proto", fileDescriptor_topic_view_service_42250129c59eb567)
}

var fileDescriptor_topic_view_service_42250129c59eb567 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x43, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0x92, 0xfc, 0x82, 0xcc, 0xe4, 0xf8, 0xb2, 0xcc, 0xd4,
	0xf2, 0x78, 0xa8, 0x98, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x02, 0x44, 0xbd, 0x5e, 0x62,
	0x4a, 0xb1, 0x1e, 0x5c, 0xab, 0x5e, 0x99, 0xa1, 0x1e, 0x4c, 0xab, 0x94, 0x11, 0x2e, 0xc3, 0x8b,
	0x52, 0x8b, 0xf3, 0x4b, 0x8b, 0x50, 0x4d, 0x87, 0x98, 0x2a, 0x25, 0x03, 0xd3, 0x53, 0x90, 0xa9,
	0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x55, 0xb2, 0xe2,
	0x12, 0x76, 0x4f, 0x2d, 0x09, 0x01, 0x69, 0x0a, 0xcb, 0x4c, 0x2d, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x52, 0xe6, 0xe2, 0x85, 0x19, 0x19, 0x9f, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x03, 0x13, 0xf4, 0x4b, 0xcc, 0x4d, 0x35, 0xda, 0xc7, 0xc8, 0x25,
	0x00, 0xd7, 0x19, 0x0c, 0x71, 0xa3, 0xd0, 0x4a, 0x46, 0x2e, 0x1e, 0x64, 0x13, 0x85, 0x4c, 0xf5,
	0x08, 0x79, 0x4b, 0x0f, 0x8b, 0x0b, 0xa4, 0x74, 0x70, 0x6a, 0x83, 0xfb, 0x55, 0x0f, 0xae, 0x49,
	0xc9, 0xa4, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x7a, 0x42, 0x3a, 0xa0, 0xc0, 0xa8, 0x46, 0x71, 0xba,
	0x6d, 0x72, 0x69, 0x71, 0x49, 0x7e, 0x6e, 0x6a, 0x51, 0xb1, 0xbe, 0x16, 0x24, 0x74, 0x40, 0x3a,
	0x8a, 0xf5, 0xb5, 0x6a, 0x9d, 0xfe, 0x30, 0x72, 0xa9, 0x24, 0xe7, 0xe7, 0x12, 0x74, 0xa0, 0x93,
	0x28, 0xba, 0x37, 0x03, 0x40, 0x81, 0x17, 0xc0, 0x18, 0xe5, 0x01, 0xd5, 0x9a, 0x9e, 0x9f, 0x93,
	0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0x0e, 0x5a, 0x58, 0x04, 0x15,
	0x64, 0x16, 0xe3, 0x4e, 0x0c, 0xd6, 0x30, 0xc6, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c,
	0x0a, 0xee, 0x10, 0x03, 0x1d, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x50, 0x0f, 0x6a,
	0x71, 0xf1, 0x29, 0x98, 0x92, 0x18, 0xc7, 0x94, 0xe2, 0x18, 0xb8, 0x92, 0x98, 0x30, 0xc3, 0x18,
	0x98, 0x92, 0x57, 0x4c, 0x2a, 0x10, 0x71, 0x2b, 0x2b, 0xc7, 0x94, 0x62, 0x2b, 0x2b, 0xb8, 0x22,
	0x2b, 0xab, 0x30, 0x43, 0x2b, 0x2b, 0x98, 0xb2, 0x24, 0x36, 0xb0, 0x3b, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0xc0, 0x75, 0x26, 0xb3, 0x02, 0x00, 0x00,
}
