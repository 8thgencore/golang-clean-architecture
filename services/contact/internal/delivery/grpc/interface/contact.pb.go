// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

package contact

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateGroupRequest struct {
	CreatedBy            string   `protobuf:"bytes,1,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (m *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(m, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *CreateGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateGroupRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GroupResponse struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	ContactCount         uint64               `protobuf:"varint,6,opt,name=contactCount,proto3" json:"contactCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GroupResponse) Reset()         { *m = GroupResponse{} }
func (m *GroupResponse) String() string { return proto.CompactTextString(m) }
func (*GroupResponse) ProtoMessage()    {}
func (*GroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

func (m *GroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupResponse.Unmarshal(m, b)
}
func (m *GroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupResponse.Marshal(b, m, deterministic)
}
func (m *GroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupResponse.Merge(m, src)
}
func (m *GroupResponse) XXX_Size() int {
	return xxx_messageInfo_GroupResponse.Size(m)
}
func (m *GroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GroupResponse proto.InternalMessageInfo

func (m *GroupResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GroupResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GroupResponse) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

func (m *GroupResponse) GetContactCount() uint64 {
	if m != nil {
		return m.ContactCount
	}
	return 0
}

type CreateGroupResponse struct {
	Response             *GroupResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateGroupResponse) Reset()         { *m = CreateGroupResponse{} }
func (m *CreateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGroupResponse) ProtoMessage()    {}
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}

func (m *CreateGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupResponse.Unmarshal(m, b)
}
func (m *CreateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupResponse.Marshal(b, m, deterministic)
}
func (m *CreateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupResponse.Merge(m, src)
}
func (m *CreateGroupResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGroupResponse.Size(m)
}
func (m *CreateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupResponse proto.InternalMessageInfo

func (m *CreateGroupResponse) GetResponse() *GroupResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type UpdateGroupRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedBy            string   `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}

func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(m, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateGroupRequest) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateGroupRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateGroupResponse struct {
	Response             *GroupResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateGroupResponse) Reset()         { *m = UpdateGroupResponse{} }
func (m *UpdateGroupResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupResponse) ProtoMessage()    {}
func (*UpdateGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}

func (m *UpdateGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupResponse.Unmarshal(m, b)
}
func (m *UpdateGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupResponse.Marshal(b, m, deterministic)
}
func (m *UpdateGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupResponse.Merge(m, src)
}
func (m *UpdateGroupResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupResponse.Size(m)
}
func (m *UpdateGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupResponse proto.InternalMessageInfo

func (m *UpdateGroupResponse) GetResponse() *GroupResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type DeleteGroupRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedBy            string   `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGroupRequest) Reset()         { *m = DeleteGroupRequest{} }
func (m *DeleteGroupRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupRequest) ProtoMessage()    {}
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{5}
}

func (m *DeleteGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupRequest.Unmarshal(m, b)
}
func (m *DeleteGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupRequest.Merge(m, src)
}
func (m *DeleteGroupRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupRequest.Size(m)
}
func (m *DeleteGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupRequest proto.InternalMessageInfo

func (m *DeleteGroupRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteGroupRequest) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

type DeleteGroupResponse struct {
	Response             *GroupResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteGroupResponse) Reset()         { *m = DeleteGroupResponse{} }
func (m *DeleteGroupResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteGroupResponse) ProtoMessage()    {}
func (*DeleteGroupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{6}
}

func (m *DeleteGroupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGroupResponse.Unmarshal(m, b)
}
func (m *DeleteGroupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGroupResponse.Marshal(b, m, deterministic)
}
func (m *DeleteGroupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGroupResponse.Merge(m, src)
}
func (m *DeleteGroupResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteGroupResponse.Size(m)
}
func (m *DeleteGroupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGroupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGroupResponse proto.InternalMessageInfo

func (m *DeleteGroupResponse) GetResponse() *GroupResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateGroupRequest)(nil), "contact.CreateGroupRequest")
	proto.RegisterType((*GroupResponse)(nil), "contact.GroupResponse")
	proto.RegisterType((*CreateGroupResponse)(nil), "contact.CreateGroupResponse")
	proto.RegisterType((*UpdateGroupRequest)(nil), "contact.UpdateGroupRequest")
	proto.RegisterType((*UpdateGroupResponse)(nil), "contact.UpdateGroupResponse")
	proto.RegisterType((*DeleteGroupRequest)(nil), "contact.DeleteGroupRequest")
	proto.RegisterType((*DeleteGroupResponse)(nil), "contact.DeleteGroupResponse")
}

func init() {
	proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15)
}

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6f, 0xa2, 0x40,
	0x18, 0x15, 0x64, 0xdd, 0xf5, 0x43, 0x3d, 0x7c, 0x26, 0x1b, 0xe2, 0x6e, 0x53, 0xc2, 0xc9, 0x13,
	0x26, 0xf4, 0xd4, 0x78, 0x52, 0x9a, 0x34, 0xf6, 0x48, 0xdb, 0x4b, 0x2f, 0x0d, 0xc2, 0x68, 0x26,
	0x11, 0x86, 0xc2, 0xd0, 0xc4, 0x9f, 0xdd, 0xa4, 0x3f, 0xa0, 0x11, 0x06, 0x3b, 0x88, 0xa6, 0x89,
	0xde, 0xc6, 0xf7, 0xbd, 0x79, 0xf3, 0xbd, 0xf7, 0x04, 0xfa, 0x01, 0x8b, 0xb9, 0x1f, 0x70, 0x3b,
	0x49, 0x19, 0x67, 0xf8, 0x5b, 0xfc, 0x1c, 0x5d, 0xaf, 0x19, 0x5b, 0x6f, 0xc8, 0xa4, 0x80, 0x97,
	0xf9, 0x6a, 0xc2, 0x69, 0x44, 0x32, 0xee, 0x47, 0x49, 0xc9, 0xb4, 0x28, 0xa0, 0x9b, 0x12, 0x9f,
	0x93, 0xfb, 0x94, 0xe5, 0x89, 0x47, 0xde, 0x72, 0x92, 0x71, 0xbc, 0x02, 0x08, 0x0a, 0x34, 0x7c,
	0x5d, 0x6e, 0x0d, 0xc5, 0x54, 0xc6, 0x5d, 0xaf, 0x2b, 0x90, 0xf9, 0x16, 0x11, 0xb4, 0xd8, 0x8f,
	0x88, 0xa1, 0x16, 0x83, 0xe2, 0x8c, 0x26, 0xe8, 0x21, 0xc9, 0x82, 0x94, 0x26, 0x9c, 0xb2, 0xd8,
	0x68, 0x17, 0x23, 0x19, 0xb2, 0x3e, 0x14, 0xe8, 0x8b, 0x57, 0xb2, 0x84, 0xc5, 0x19, 0xc1, 0x01,
	0xa8, 0x34, 0x14, 0xf2, 0x2a, 0x0d, 0xcf, 0xd3, 0xc5, 0xdb, 0xef, 0x65, 0x7d, 0x6e, 0x68, 0xa6,
	0x32, 0xd6, 0x9d, 0x91, 0x5d, 0x1a, 0xb7, 0x2b, 0xe3, 0xf6, 0x53, 0x65, 0x7c, 0x6f, 0x64, 0xc6,
	0x71, 0x0a, 0x7a, 0xc4, 0x42, 0xba, 0xa2, 0xe5, 0xdd, 0x5f, 0x3f, 0xde, 0x85, 0x8a, 0x3e, 0xe3,
	0x68, 0x41, 0x4f, 0xc4, 0xec, 0xb2, 0x3c, 0xe6, 0x46, 0xc7, 0x54, 0xc6, 0x9a, 0x57, 0xc3, 0xac,
	0x05, 0x0c, 0x6b, 0xf1, 0x0a, 0xe3, 0x0e, 0xfc, 0x49, 0xc5, 0xb9, 0xb0, 0xaf, 0x3b, 0x7f, 0xed,
	0xaa, 0xc1, 0x1a, 0xd3, 0xdb, 0xf3, 0xac, 0x2d, 0xe0, 0x73, 0x12, 0x1e, 0x36, 0x75, 0x18, 0x61,
	0xbd, 0x39, 0xf5, 0x54, 0x73, 0xed, 0xd3, 0x09, 0x6b, 0xcd, 0xe6, 0x16, 0x30, 0xac, 0x3d, 0x7d,
	0x81, 0x0b, 0x17, 0xf0, 0x8e, 0x6c, 0xc8, 0x45, 0x2e, 0x76, 0xfb, 0xd4, 0x44, 0xce, 0xdf, 0xc7,
	0xf9, 0x54, 0x60, 0xe0, 0x96, 0x9c, 0x47, 0x92, 0xbe, 0xd3, 0x80, 0xe0, 0x03, 0xe8, 0x52, 0x67,
	0xf8, 0x6f, 0xaf, 0xd1, 0xfc, 0x50, 0x46, 0xff, 0x8f, 0x0f, 0x85, 0xd9, 0xd6, 0x4e, 0x4b, 0x4a,
	0x4e, 0xd2, 0x6a, 0x56, 0x29, 0x69, 0x1d, 0x09, 0xbb, 0xd4, 0x92, 0x5c, 0x4b, 0x5a, 0xcd, 0x40,
	0x25, 0xad, 0x23, 0x41, 0x59, 0xad, 0x79, 0xef, 0x05, 0xec, 0xc9, 0x54, 0x70, 0x96, 0x9d, 0xe2,
	0x9f, 0x7e, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x24, 0x92, 0x42, 0xbe, 0x46, 0x04, 0x00, 0x00,
}
