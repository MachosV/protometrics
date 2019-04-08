// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protoCar.proto

package car

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

type CarMessage struct {
	Make                 *string  `protobuf:"bytes,1,opt,name=make" json:"make,omitempty"`
	Model                *string  `protobuf:"bytes,2,opt,name=model" json:"model,omitempty"`
	Year                 *int32   `protobuf:"varint,3,opt,name=year" json:"year,omitempty"`
	Color                *string  `protobuf:"bytes,4,opt,name=color" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarMessage) Reset()         { *m = CarMessage{} }
func (m *CarMessage) String() string { return proto.CompactTextString(m) }
func (*CarMessage) ProtoMessage()    {}
func (*CarMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_65fa50b535b369cf, []int{0}
}

func (m *CarMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarMessage.Unmarshal(m, b)
}
func (m *CarMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarMessage.Marshal(b, m, deterministic)
}
func (m *CarMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarMessage.Merge(m, src)
}
func (m *CarMessage) XXX_Size() int {
	return xxx_messageInfo_CarMessage.Size(m)
}
func (m *CarMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CarMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CarMessage proto.InternalMessageInfo

func (m *CarMessage) GetMake() string {
	if m != nil && m.Make != nil {
		return *m.Make
	}
	return ""
}

func (m *CarMessage) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *CarMessage) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *CarMessage) GetColor() string {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return ""
}

func init() {
	proto.RegisterType((*CarMessage)(nil), "car.CarMessage")
}

func init() { proto.RegisterFile("protoCar.proto", fileDescriptor_65fa50b535b369cf) }

var fileDescriptor_65fa50b535b369cf = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x77, 0x4e, 0x2c, 0xd2, 0x03, 0x33, 0x84, 0x98, 0x93, 0x13, 0x8b, 0x94, 0x12, 0xb8, 0xb8,
	0x9c, 0x13, 0x8b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x72, 0x13,
	0xb3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xdc,
	0xfc, 0x94, 0xd4, 0x1c, 0x09, 0x26, 0xb0, 0x20, 0x84, 0x03, 0x52, 0x59, 0x99, 0x9a, 0x58, 0x24,
	0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x83, 0x54, 0x26, 0xe7, 0xe7, 0xe4, 0x17, 0x49,
	0xb0, 0x40, 0x54, 0x82, 0x39, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x27, 0x67, 0x82, 0x77,
	0x00, 0x00, 0x00,
}
