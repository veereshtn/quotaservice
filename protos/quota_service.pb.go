// Code generated by protoc-gen-go.
// source: protos/quota_service.proto
// DO NOT EDIT!

/*
Package quotaservice is a generated protocol buffer package.

It is generated from these files:
	protos/quota_service.proto

It has these top-level messages:
	AllowRequest
	AllowResponse
*/
package quotaservice

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

type AllowResponse_Status int32

const (
	AllowResponse_OK                                 AllowResponse_Status = 0
	AllowResponse_REJECTED_TIMEOUT                   AllowResponse_Status = 1
	AllowResponse_REJECTED_NO_BUCKET                 AllowResponse_Status = 2
	AllowResponse_REJECTED_TOO_MANY_BUCKETS          AllowResponse_Status = 3
	AllowResponse_REJECTED_TOO_MANY_TOKENS_REQUESTED AllowResponse_Status = 4
	AllowResponse_REJECTED_INVALID_REQUEST           AllowResponse_Status = 5
	AllowResponse_REJECTED_SERVER_ERROR              AllowResponse_Status = 6
)

var AllowResponse_Status_name = map[int32]string{
	0: "OK",
	1: "REJECTED_TIMEOUT",
	2: "REJECTED_NO_BUCKET",
	3: "REJECTED_TOO_MANY_BUCKETS",
	4: "REJECTED_TOO_MANY_TOKENS_REQUESTED",
	5: "REJECTED_INVALID_REQUEST",
	6: "REJECTED_SERVER_ERROR",
}
var AllowResponse_Status_value = map[string]int32{
	"OK":                                 0,
	"REJECTED_TIMEOUT":                   1,
	"REJECTED_NO_BUCKET":                 2,
	"REJECTED_TOO_MANY_BUCKETS":          3,
	"REJECTED_TOO_MANY_TOKENS_REQUESTED": 4,
	"REJECTED_INVALID_REQUEST":           5,
	"REJECTED_SERVER_ERROR":              6,
}

func (x AllowResponse_Status) String() string {
	return proto.EnumName(AllowResponse_Status_name, int32(x))
}
func (AllowResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type AllowRequest struct {
	Namespace  string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	BucketName string `protobuf:"bytes,2,opt,name=bucket_name" json:"bucket_name,omitempty"`
	// *
	// Number of tokens requested. Defaults to 1, cannot be 0.
	TokensRequested int64 `protobuf:"varint,3,opt,name=tokens_requested" json:"tokens_requested,omitempty"`
	// *
	// Max wait time, in millis. Defaults to 0, which assumes no waiting.
	MaxWaitMillisOverride int64 `protobuf:"varint,4,opt,name=max_wait_millis_override" json:"max_wait_millis_override,omitempty"`
	// *
	// Whether to override max wait time with the above value.
	// Defaults to false, which falls back to the bucket's configured value.
	MaxWaitTimeOverride bool `protobuf:"varint,5,opt,name=max_wait_time_override" json:"max_wait_time_override,omitempty"`
}

func (m *AllowRequest) Reset()                    { *m = AllowRequest{} }
func (m *AllowRequest) String() string            { return proto.CompactTextString(m) }
func (*AllowRequest) ProtoMessage()               {}
func (*AllowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllowResponse struct {
	Status AllowResponse_Status `protobuf:"varint,1,opt,name=status,enum=quotaservice.AllowResponse_Status" json:"status,omitempty"`
	// *
	// Number of tokens granted, if status == OK
	TokensGranted int64 `protobuf:"varint,2,opt,name=tokens_granted" json:"tokens_granted,omitempty"`
	// *
	// Wait for this many millis before proceeding, if status == OK. 0 if no waiting is required.
	WaitMillis int64 `protobuf:"varint,3,opt,name=wait_millis" json:"wait_millis,omitempty"`
}

func (m *AllowResponse) Reset()                    { *m = AllowResponse{} }
func (m *AllowResponse) String() string            { return proto.CompactTextString(m) }
func (*AllowResponse) ProtoMessage()               {}
func (*AllowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*AllowRequest)(nil), "quotaservice.AllowRequest")
	proto.RegisterType((*AllowResponse)(nil), "quotaservice.AllowResponse")
	proto.RegisterEnum("quotaservice.AllowResponse_Status", AllowResponse_Status_name, AllowResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for QuotaService service

type QuotaServiceClient interface {
	Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error)
}

type quotaServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuotaServiceClient(cc *grpc.ClientConn) QuotaServiceClient {
	return &quotaServiceClient{cc}
}

func (c *quotaServiceClient) Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*AllowResponse, error) {
	out := new(AllowResponse)
	err := grpc.Invoke(ctx, "/quotaservice.QuotaService/Allow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QuotaService service

type QuotaServiceServer interface {
	Allow(context.Context, *AllowRequest) (*AllowResponse, error)
}

func RegisterQuotaServiceServer(s *grpc.Server, srv QuotaServiceServer) {
	s.RegisterService(&_QuotaService_serviceDesc, srv)
}

func _QuotaService_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaServiceServer).Allow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quotaservice.QuotaService/Allow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaServiceServer).Allow(ctx, req.(*AllowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuotaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quotaservice.QuotaService",
	HandlerType: (*QuotaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allow",
			Handler:    _QuotaService_Allow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("protos/quota_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0x5b, 0xee, 0xd2, 0x40,
	0x18, 0xc5, 0xff, 0x2d, 0xd0, 0xc8, 0x07, 0x92, 0x71, 0x50, 0x52, 0xf0, 0x92, 0xa6, 0x0f, 0x86,
	0xa7, 0x9a, 0xe0, 0x0a, 0xb8, 0xcc, 0x03, 0x22, 0x6d, 0x98, 0x16, 0x12, 0x9f, 0x26, 0x05, 0x26,
	0xa6, 0xa1, 0x17, 0xe8, 0x4c, 0xc1, 0x75, 0xb8, 0x13, 0x37, 0xe0, 0xda, 0x0c, 0x43, 0x2d, 0x18,
	0x8d, 0xaf, 0xe7, 0xfc, 0xbe, 0xf4, 0xd7, 0xd3, 0xc2, 0xe0, 0x98, 0x67, 0x32, 0x13, 0x1f, 0x4e,
	0x45, 0x26, 0x43, 0x26, 0x78, 0x7e, 0x8e, 0x76, 0xdc, 0x51, 0x21, 0x6e, 0xab, 0xb0, 0xcc, 0xec,
	0xef, 0x1a, 0xb4, 0xc7, 0x71, 0x9c, 0x5d, 0x28, 0x3f, 0x15, 0x5c, 0x48, 0xfc, 0x02, 0x9a, 0x69,
	0x98, 0x70, 0x71, 0x0c, 0x77, 0xdc, 0xd4, 0x2c, 0x6d, 0xd8, 0xc4, 0x5d, 0x68, 0x6d, 0x8b, 0xdd,
	0x81, 0x4b, 0x76, 0x6d, 0x4c, 0x5d, 0x85, 0x26, 0x20, 0x99, 0x1d, 0x78, 0x2a, 0x58, 0x7e, 0xbb,
	0xe4, 0x7b, 0xb3, 0x66, 0x69, 0xc3, 0x1a, 0xb6, 0xc0, 0x4c, 0xc2, 0x6f, 0xec, 0x12, 0x46, 0x92,
	0x25, 0x51, 0x1c, 0x47, 0x82, 0x65, 0x67, 0x9e, 0xe7, 0xd1, 0x9e, 0x9b, 0x75, 0x45, 0xbc, 0x83,
	0x5e, 0x45, 0xc8, 0x28, 0xe1, 0xf7, 0xbe, 0x61, 0x69, 0xc3, 0x67, 0xf6, 0x0f, 0x1d, 0x9e, 0x97,
	0x52, 0xe2, 0x98, 0xa5, 0x82, 0xe3, 0x11, 0x18, 0x42, 0x86, 0xb2, 0x10, 0x4a, 0xa9, 0x33, 0xb2,
	0x9d, 0xc7, 0xb7, 0x70, 0xfe, 0x80, 0x1d, 0x5f, 0x91, 0xb8, 0x07, 0x9d, 0xd2, 0xf0, 0x6b, 0x1e,
	0xa6, 0x57, 0x3f, 0x5d, 0x3d, 0xbd, 0x0b, 0xad, 0x07, 0xb7, 0x9b, 0xb4, 0xfd, 0x53, 0x03, 0xa3,
	0xbc, 0x33, 0x40, 0xf7, 0x16, 0xe8, 0x09, 0xbf, 0x04, 0x44, 0xc9, 0x27, 0x32, 0x0d, 0xc8, 0x8c,
	0x05, 0xf3, 0x25, 0xf1, 0xd6, 0x01, 0xd2, 0x70, 0x0f, 0x70, 0x95, 0xba, 0x1e, 0x9b, 0xac, 0xa7,
	0x0b, 0x12, 0x20, 0x1d, 0xbf, 0x85, 0xfe, 0x9d, 0xf6, 0x3c, 0xb6, 0x1c, 0xbb, 0x5f, 0xca, 0xd6,
	0x47, 0x35, 0xfc, 0x1e, 0xec, 0xbf, 0xeb, 0xc0, 0x5b, 0x10, 0xd7, 0x67, 0x94, 0xac, 0xd6, 0xc4,
	0x0f, 0xc8, 0x0c, 0xd5, 0xf1, 0x1b, 0x30, 0x2b, 0x6e, 0xee, 0x6e, 0xc6, 0x9f, 0xe7, 0xb3, 0xdf,
	0x3d, 0x6a, 0xe0, 0x3e, 0xbc, 0xaa, 0x5a, 0x9f, 0xd0, 0x0d, 0xa1, 0x8c, 0x50, 0xea, 0x51, 0x64,
	0x8c, 0x28, 0xb4, 0x57, 0xd7, 0x49, 0xfc, 0xdb, 0x24, 0x78, 0x02, 0x0d, 0xb5, 0x0a, 0x1e, 0xfc,
	0x73, 0x2a, 0xf5, 0xc9, 0x06, 0xaf, 0xff, 0x33, 0xa3, 0xfd, 0xb4, 0x35, 0xd4, 0x1f, 0xf3, 0xf1,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x84, 0xae, 0x79, 0x4f, 0x02, 0x00, 0x00,
}
