// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/block/block.proto

package go_micro_srv_block

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

type Request struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	IpfsUrl              string   `protobuf:"bytes,2,opt,name=ipfsUrl,proto3" json:"ipfsUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2be3018ad46957, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Request) GetIpfsUrl() string {
	if m != nil {
		return m.IpfsUrl
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2be3018ad46957, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.block.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.block.Response")
}

func init() { proto.RegisterFile("proto/block/block.proto", fileDescriptor_2a2be3018ad46957) }

var fileDescriptor_2a2be3018ad46957 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x86, 0x90, 0x7a, 0x60, 0x11, 0x21, 0xa1, 0xf4, 0x7c,
	0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0x3d, 0xb0, 0x8c, 0x92, 0x2d, 0x17,
	0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49,
	0x6a, 0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x0b, 0x92, 0xc9, 0x2c, 0x48,
	0x2b, 0x0e, 0x2d, 0xca, 0x91, 0x60, 0x82, 0xc8, 0x40, 0xb9, 0x4a, 0x06, 0x5c, 0x1c, 0x41, 0xa9,
	0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x50, 0xbd, 0x20,
	0xa6, 0x90, 0x10, 0x17, 0x4b, 0x4a, 0x62, 0x49, 0x22, 0x54, 0x13, 0x98, 0x6d, 0xb4, 0x9c, 0x91,
	0x8b, 0xd5, 0x09, 0x64, 0xb5, 0x50, 0x00, 0x97, 0x60, 0x68, 0x41, 0x4e, 0x7e, 0x62, 0x8a, 0x67,
	0x41, 0x5a, 0xb1, 0x33, 0xd4, 0x2a, 0x69, 0x3d, 0x4c, 0x47, 0xea, 0x41, 0x5d, 0x28, 0x25, 0x83,
	0x5d, 0x12, 0x62, 0xbf, 0x12, 0x83, 0x90, 0x0f, 0x17, 0x2f, 0xc2, 0xc4, 0xd0, 0xa2, 0x1c, 0x8a,
	0x4c, 0x4b, 0x62, 0x03, 0x87, 0x9a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x72, 0xa1, 0x48,
	0x50, 0x01, 0x00, 0x00,
}
