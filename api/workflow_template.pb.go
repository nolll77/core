// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow_template.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateWorkflowTemplateRequest struct {
	Namespace            string            `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkflowTemplate     *WorkflowTemplate `protobuf:"bytes,2,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateWorkflowTemplateRequest) Reset()         { *m = CreateWorkflowTemplateRequest{} }
func (m *CreateWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWorkflowTemplateRequest) ProtoMessage()    {}
func (*CreateWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{0}
}

func (m *CreateWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *CreateWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *CreateWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWorkflowTemplateRequest.Merge(m, src)
}
func (m *CreateWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWorkflowTemplateRequest.Size(m)
}
func (m *CreateWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWorkflowTemplateRequest proto.InternalMessageInfo

func (m *CreateWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateWorkflowTemplateRequest) GetWorkflowTemplate() *WorkflowTemplate {
	if m != nil {
		return m.WorkflowTemplate
	}
	return nil
}

type GetWorkflowTemplateRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWorkflowTemplateRequest) Reset()         { *m = GetWorkflowTemplateRequest{} }
func (m *GetWorkflowTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*GetWorkflowTemplateRequest) ProtoMessage()    {}
func (*GetWorkflowTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{1}
}

func (m *GetWorkflowTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Unmarshal(m, b)
}
func (m *GetWorkflowTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Marshal(b, m, deterministic)
}
func (m *GetWorkflowTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWorkflowTemplateRequest.Merge(m, src)
}
func (m *GetWorkflowTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_GetWorkflowTemplateRequest.Size(m)
}
func (m *GetWorkflowTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWorkflowTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWorkflowTemplateRequest proto.InternalMessageInfo

func (m *GetWorkflowTemplateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GetWorkflowTemplateRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type WorkflowTemplate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Manifest             string   `protobuf:"bytes,4,opt,name=manifest,proto3" json:"manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a07547748a96e8, []int{2}
}

func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (m *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(m, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *WorkflowTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowTemplate) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *WorkflowTemplate) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateWorkflowTemplateRequest)(nil), "api.CreateWorkflowTemplateRequest")
	proto.RegisterType((*GetWorkflowTemplateRequest)(nil), "api.GetWorkflowTemplateRequest")
	proto.RegisterType((*WorkflowTemplate)(nil), "api.WorkflowTemplate")
}

func init() { proto.RegisterFile("workflow_template.proto", fileDescriptor_b9a07547748a96e8) }

var fileDescriptor_b9a07547748a96e8 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xcf, 0x2f, 0xca,
	0x4e, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x6a, 0x60, 0xe4, 0x92, 0x75, 0x2e, 0x4a,
	0x4d, 0x2c, 0x49, 0x0d, 0x87, 0x2a, 0x0b, 0x81, 0xaa, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e,
	0x11, 0x92, 0xe1, 0xe2, 0xcc, 0x4b, 0xcc, 0x4d, 0x2d, 0x2e, 0x48, 0x4c, 0x4e, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x39, 0x72, 0x09, 0x94, 0xa3, 0x69, 0x94, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0xc3, 0x30, 0x15, 0x43, 0xb9, 0x92,
	0x0f, 0x97, 0x94, 0x7b, 0x6a, 0x09, 0x79, 0xd6, 0x0b, 0x70, 0x31, 0x97, 0x66, 0xa6, 0x80, 0x6d,
	0xe4, 0x0c, 0x02, 0x31, 0x95, 0xf2, 0xb8, 0x04, 0xd0, 0x8d, 0x82, 0xa9, 0x62, 0x84, 0xab, 0x12,
	0x12, 0xe2, 0x62, 0x01, 0x19, 0x02, 0xd5, 0x08, 0x66, 0x0b, 0x49, 0x70, 0xb1, 0x97, 0xa5, 0x16,
	0x15, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x83, 0x85, 0x61, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0xdc, 0xc4,
	0xbc, 0xcc, 0xb4, 0xd4, 0xe2, 0x12, 0x09, 0x16, 0xb0, 0x14, 0x9c, 0x9f, 0xc4, 0x06, 0x0e, 0x4c,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x71, 0x69, 0x39, 0x67, 0x01, 0x00, 0x00,
}
