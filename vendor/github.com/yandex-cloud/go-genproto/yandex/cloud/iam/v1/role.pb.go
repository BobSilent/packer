// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/role.proto

package iam // import "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Role resource. For more information, see [Roles](/docs/iam/concepts/access-control/roles).
type Role struct {
	// ID of the role.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Description of the role. 0-256 characters long.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_role_656bcd2fea39e7f2, []int{0}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Role)(nil), "yandex.cloud.iam.v1.Role")
}

func init() {
	proto.RegisterFile("yandex/cloud/iam/v1/role.proto", fileDescriptor_role_656bcd2fea39e7f2)
}

var fileDescriptor_role_656bcd2fea39e7f2 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xab, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0xd4,
	0x2f, 0xca, 0xcf, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0xc8, 0xeb, 0x81,
	0xe5, 0xf5, 0x32, 0x13, 0x73, 0xf5, 0xca, 0x0c, 0x95, 0x2c, 0xb8, 0x58, 0x82, 0xf2, 0x73, 0x52,
	0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53,
	0x84, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24,
	0x98, 0xc0, 0x12, 0xc8, 0x42, 0x4e, 0xb6, 0x51, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0xb3, 0x75, 0x21, 0x76, 0xa7, 0xe7, 0xeb, 0xa6, 0xa7, 0xe6, 0x81,
	0x6d, 0xd5, 0xc7, 0xe2, 0x28, 0xeb, 0xcc, 0xc4, 0xdc, 0x24, 0x36, 0xb0, 0xb4, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x48, 0x72, 0x28, 0xb9, 0xb6, 0x00, 0x00, 0x00,
}
