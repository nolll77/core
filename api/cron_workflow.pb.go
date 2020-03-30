// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cron_workflow.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCronWorkflowRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCronWorkflowRequest) Reset()         { *m = GetCronWorkflowRequest{} }
func (m *GetCronWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*GetCronWorkflowRequest) ProtoMessage()    {}
func (*GetCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_989cccaad551a50c, []int{0}
}

func (m *GetCronWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCronWorkflowRequest.Unmarshal(m, b)
}
func (m *GetCronWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCronWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *GetCronWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCronWorkflowRequest.Merge(m, src)
}
func (m *GetCronWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_GetCronWorkflowRequest.Size(m)
}
func (m *GetCronWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCronWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCronWorkflowRequest proto.InternalMessageInfo

func (m *GetCronWorkflowRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetCronWorkflowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateCronWorkflowRequest struct {
	Namespace            string        `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CronWorkflow         *CronWorkflow `protobuf:"bytes,2,opt,name=cronWorkflow,proto3" json:"cronWorkflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateCronWorkflowRequest) Reset()         { *m = CreateCronWorkflowRequest{} }
func (m *CreateCronWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCronWorkflowRequest) ProtoMessage()    {}
func (*CreateCronWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_989cccaad551a50c, []int{1}
}

func (m *CreateCronWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCronWorkflowRequest.Unmarshal(m, b)
}
func (m *CreateCronWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCronWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *CreateCronWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCronWorkflowRequest.Merge(m, src)
}
func (m *CreateCronWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCronWorkflowRequest.Size(m)
}
func (m *CreateCronWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCronWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCronWorkflowRequest proto.InternalMessageInfo

func (m *CreateCronWorkflowRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateCronWorkflowRequest) GetCronWorkflow() *CronWorkflow {
	if m != nil {
		return m.CronWorkflow
	}
	return nil
}

type CronWorkflow struct {
	Schedule                   string             `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Timezone                   string             `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Suspend                    bool               `protobuf:"varint,3,opt,name=suspend,proto3" json:"suspend,omitempty"`
	ConcurrencyPolicy          string             `protobuf:"bytes,4,opt,name=concurrencyPolicy,proto3" json:"concurrencyPolicy,omitempty"`
	StartingDeadlineSeconds    int64              `protobuf:"varint,5,opt,name=startingDeadlineSeconds,proto3" json:"startingDeadlineSeconds,omitempty"`
	SuccessfulJobsHistoryLimit int32              `protobuf:"varint,6,opt,name=successfulJobsHistoryLimit,proto3" json:"successfulJobsHistoryLimit,omitempty"`
	FailedJobsHistoryLimit     int32              `protobuf:"varint,7,opt,name=failedJobsHistoryLimit,proto3" json:"failedJobsHistoryLimit,omitempty"`
	WorkflowExecution          *WorkflowExecution `protobuf:"bytes,8,opt,name=workflowExecution,proto3" json:"workflowExecution,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}           `json:"-"`
	XXX_unrecognized           []byte             `json:"-"`
	XXX_sizecache              int32              `json:"-"`
}

func (m *CronWorkflow) Reset()         { *m = CronWorkflow{} }
func (m *CronWorkflow) String() string { return proto.CompactTextString(m) }
func (*CronWorkflow) ProtoMessage()    {}
func (*CronWorkflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_989cccaad551a50c, []int{2}
}

func (m *CronWorkflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronWorkflow.Unmarshal(m, b)
}
func (m *CronWorkflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronWorkflow.Marshal(b, m, deterministic)
}
func (m *CronWorkflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronWorkflow.Merge(m, src)
}
func (m *CronWorkflow) XXX_Size() int {
	return xxx_messageInfo_CronWorkflow.Size(m)
}
func (m *CronWorkflow) XXX_DiscardUnknown() {
	xxx_messageInfo_CronWorkflow.DiscardUnknown(m)
}

var xxx_messageInfo_CronWorkflow proto.InternalMessageInfo

func (m *CronWorkflow) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *CronWorkflow) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *CronWorkflow) GetSuspend() bool {
	if m != nil {
		return m.Suspend
	}
	return false
}

func (m *CronWorkflow) GetConcurrencyPolicy() string {
	if m != nil {
		return m.ConcurrencyPolicy
	}
	return ""
}

func (m *CronWorkflow) GetStartingDeadlineSeconds() int64 {
	if m != nil {
		return m.StartingDeadlineSeconds
	}
	return 0
}

func (m *CronWorkflow) GetSuccessfulJobsHistoryLimit() int32 {
	if m != nil {
		return m.SuccessfulJobsHistoryLimit
	}
	return 0
}

func (m *CronWorkflow) GetFailedJobsHistoryLimit() int32 {
	if m != nil {
		return m.FailedJobsHistoryLimit
	}
	return 0
}

func (m *CronWorkflow) GetWorkflowExecution() *WorkflowExecution {
	if m != nil {
		return m.WorkflowExecution
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCronWorkflowRequest)(nil), "api.GetCronWorkflowRequest")
	proto.RegisterType((*CreateCronWorkflowRequest)(nil), "api.CreateCronWorkflowRequest")
	proto.RegisterType((*CronWorkflow)(nil), "api.CronWorkflow")
}

func init() {
	proto.RegisterFile("cron_workflow.proto", fileDescriptor_989cccaad551a50c)
}

var fileDescriptor_989cccaad551a50c = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0xa4, 0x1f, 0xe9, 0x12, 0x81, 0x32, 0x95, 0x52, 0xe3, 0x56, 0x28, 0xf2, 0x85,
	0x1c, 0x50, 0x4c, 0x8b, 0x28, 0x08, 0x09, 0x2e, 0x2d, 0x02, 0x55, 0x1c, 0x90, 0x7b, 0xe8, 0xb1,
	0xda, 0xac, 0x27, 0x61, 0x85, 0xbd, 0xbb, 0xec, 0xae, 0x5b, 0x0c, 0xea, 0x85, 0x43, 0x5f, 0x80,
	0x13, 0xcf, 0xc5, 0x2b, 0xf0, 0x20, 0x28, 0x1b, 0x3b, 0xc4, 0x75, 0x83, 0xe0, 0xe6, 0xbf, 0x7f,
	0xf3, 0xa5, 0x9d, 0xff, 0x90, 0x6d, 0xa6, 0xa5, 0x38, 0xbf, 0x94, 0xfa, 0xe3, 0x24, 0x95, 0x97,
	0x23, 0xa5, 0xa5, 0x95, 0xd0, 0xa6, 0x8a, 0x07, 0x7b, 0x53, 0x29, 0xa7, 0x29, 0x46, 0x54, 0xf1,
	0x88, 0x0a, 0x21, 0x2d, 0xb5, 0x5c, 0x0a, 0x33, 0x0f, 0x09, 0x76, 0x4b, 0xea, 0xd4, 0x38, 0x9f,
	0x44, 0x98, 0x29, 0x5b, 0x94, 0x70, 0xa7, 0xaa, 0x77, 0x6e, 0x31, 0x53, 0x29, 0xb5, 0x58, 0x82,
	0xbb, 0xf5, 0x46, 0x41, 0x37, 0x43, 0xab, 0x39, 0x9b, 0xab, 0xf0, 0x84, 0xf4, 0xdf, 0xa0, 0x3d,
	0xd2, 0x52, 0x9c, 0x95, 0x61, 0x31, 0x7e, 0xca, 0xd1, 0x58, 0xd8, 0x23, 0x5b, 0x82, 0x66, 0x68,
	0x14, 0x65, 0xe8, 0x7b, 0x03, 0x6f, 0xb8, 0x15, 0xff, 0xf9, 0x01, 0x40, 0xd6, 0x66, 0xc2, 0x6f,
	0x39, 0xe0, 0xbe, 0x43, 0x45, 0xee, 0x1f, 0x69, 0xa4, 0x16, 0xff, 0xbf, 0xdc, 0x53, 0xd2, 0x65,
	0x4b, 0x49, 0xae, 0xec, 0x9d, 0x83, 0xde, 0x88, 0x2a, 0x3e, 0xaa, 0x55, 0xab, 0x85, 0x85, 0xd7,
	0x6d, 0xd2, 0x5d, 0xc6, 0x10, 0x90, 0x8e, 0x61, 0x1f, 0x30, 0xc9, 0xd3, 0xaa, 0xc9, 0x42, 0xcf,
	0x98, 0xe5, 0x19, 0x7e, 0x91, 0xa2, 0x1a, 0x7b, 0xa1, 0xc1, 0x27, 0x9b, 0x26, 0x37, 0x0a, 0x45,
	0xe2, 0xb7, 0x07, 0xde, 0xb0, 0x13, 0x57, 0x12, 0x1e, 0x91, 0x1e, 0x93, 0x82, 0xe5, 0x5a, 0xa3,
	0x60, 0xc5, 0x7b, 0x99, 0x72, 0x56, 0xf8, 0x6b, 0x2e, 0xbd, 0x09, 0xe0, 0x39, 0xd9, 0x31, 0x96,
	0x6a, 0xcb, 0xc5, 0xf4, 0x18, 0x69, 0x92, 0x72, 0x81, 0xa7, 0xc8, 0xa4, 0x48, 0x8c, 0xbf, 0x3e,
	0xf0, 0x86, 0xed, 0x78, 0x15, 0x86, 0x57, 0x24, 0x30, 0x39, 0x63, 0x68, 0xcc, 0x24, 0x4f, 0x4f,
	0xe4, 0xd8, 0xbc, 0xe5, 0xc6, 0x4a, 0x5d, 0xbc, 0xe3, 0x19, 0xb7, 0xfe, 0xc6, 0xc0, 0x1b, 0xae,
	0xc7, 0x7f, 0x89, 0x80, 0x43, 0xd2, 0x9f, 0x50, 0x9e, 0x62, 0xd2, 0xc8, 0xdd, 0x74, 0xb9, 0x2b,
	0x28, 0x1c, 0x93, 0x5e, 0x65, 0x90, 0xd7, 0x9f, 0x91, 0xe5, 0x33, 0xc3, 0xf9, 0x1d, 0xf7, 0xfc,
	0x7d, 0xf7, 0xfc, 0x67, 0x37, 0x69, 0xdc, 0x4c, 0x38, 0xf8, 0xd1, 0x22, 0xdb, 0xcb, 0x8b, 0x38,
	0x45, 0x7d, 0xc1, 0x19, 0xc2, 0xb5, 0x47, 0xa0, 0xe9, 0x09, 0x78, 0x50, 0x2e, 0x76, 0x85, 0x59,
	0x82, 0xe6, 0xe2, 0xc3, 0x97, 0xdf, 0x7e, 0xfe, 0xfa, 0xde, 0x7a, 0x16, 0x3e, 0x9c, 0x1d, 0x87,
	0x89, 0x2e, 0xf6, 0xc7, 0x68, 0xe9, 0x7e, 0xf4, 0x75, 0xe1, 0xa1, 0xab, 0xa8, 0x76, 0x56, 0x2f,
	0x6a, 0x4e, 0x81, 0x82, 0xdc, 0xbb, 0xe1, 0x73, 0xd8, 0x75, 0x4d, 0x6e, 0x77, 0xff, 0x6d, 0x13,
	0x1c, 0xba, 0x09, 0x1e, 0xc3, 0xe8, 0x1f, 0x27, 0x98, 0x93, 0xab, 0xf1, 0x86, 0xbb, 0xb4, 0x27,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x22, 0xa8, 0xb9, 0xf7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CronWorkflowServiceClient is the client API for CronWorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CronWorkflowServiceClient interface {
	CreateCronWorkflow(ctx context.Context, in *CreateCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error)
	GetCronWorkflow(ctx context.Context, in *GetCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error)
}

type cronWorkflowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCronWorkflowServiceClient(cc grpc.ClientConnInterface) CronWorkflowServiceClient {
	return &cronWorkflowServiceClient{cc}
}

func (c *cronWorkflowServiceClient) CreateCronWorkflow(ctx context.Context, in *CreateCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error) {
	out := new(CronWorkflow)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/CreateCronWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronWorkflowServiceClient) GetCronWorkflow(ctx context.Context, in *GetCronWorkflowRequest, opts ...grpc.CallOption) (*CronWorkflow, error) {
	out := new(CronWorkflow)
	err := c.cc.Invoke(ctx, "/api.CronWorkflowService/GetCronWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronWorkflowServiceServer is the server API for CronWorkflowService service.
type CronWorkflowServiceServer interface {
	CreateCronWorkflow(context.Context, *CreateCronWorkflowRequest) (*CronWorkflow, error)
	GetCronWorkflow(context.Context, *GetCronWorkflowRequest) (*CronWorkflow, error)
}

// UnimplementedCronWorkflowServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCronWorkflowServiceServer struct {
}

func (*UnimplementedCronWorkflowServiceServer) CreateCronWorkflow(ctx context.Context, req *CreateCronWorkflowRequest) (*CronWorkflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCronWorkflow not implemented")
}
func (*UnimplementedCronWorkflowServiceServer) GetCronWorkflow(ctx context.Context, req *GetCronWorkflowRequest) (*CronWorkflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCronWorkflow not implemented")
}

func RegisterCronWorkflowServiceServer(s *grpc.Server, srv CronWorkflowServiceServer) {
	s.RegisterService(&_CronWorkflowService_serviceDesc, srv)
}

func _CronWorkflowService_CreateCronWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).CreateCronWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/CreateCronWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).CreateCronWorkflow(ctx, req.(*CreateCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronWorkflowService_GetCronWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCronWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronWorkflowServiceServer).GetCronWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CronWorkflowService/GetCronWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronWorkflowServiceServer).GetCronWorkflow(ctx, req.(*GetCronWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CronWorkflowService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CronWorkflowService",
	HandlerType: (*CronWorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCronWorkflow",
			Handler:    _CronWorkflowService_CreateCronWorkflow_Handler,
		},
		{
			MethodName: "GetCronWorkflow",
			Handler:    _CronWorkflowService_GetCronWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cron_workflow.proto",
}
