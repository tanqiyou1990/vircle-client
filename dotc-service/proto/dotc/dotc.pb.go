// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dotc/dotc.proto

package go_micro_srv_dotc

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

type DataModel struct {
	DataName             string   `protobuf:"bytes,1,opt,name=dataName,proto3" json:"dataName,omitempty"`
	CreateTime           int64    `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,3,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataModel) Reset()         { *m = DataModel{} }
func (m *DataModel) String() string { return proto.CompactTextString(m) }
func (*DataModel) ProtoMessage()    {}
func (*DataModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebe87db631c5431, []int{0}
}

func (m *DataModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataModel.Unmarshal(m, b)
}
func (m *DataModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataModel.Marshal(b, m, deterministic)
}
func (m *DataModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataModel.Merge(m, src)
}
func (m *DataModel) XXX_Size() int {
	return xxx_messageInfo_DataModel.Size(m)
}
func (m *DataModel) XXX_DiscardUnknown() {
	xxx_messageInfo_DataModel.DiscardUnknown(m)
}

var xxx_messageInfo_DataModel proto.InternalMessageInfo

func (m *DataModel) GetDataName() string {
	if m != nil {
		return m.DataName
	}
	return ""
}

func (m *DataModel) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *DataModel) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type BlockData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ModelName            string   `protobuf:"bytes,2,opt,name=modelName,proto3" json:"modelName,omitempty"`
	DataHash             string   `protobuf:"bytes,3,opt,name=dataHash,proto3" json:"dataHash,omitempty"`
	TransHash            string   `protobuf:"bytes,4,opt,name=transHash,proto3" json:"transHash,omitempty"`
	TransTime            int64    `protobuf:"varint,5,opt,name=transTime,proto3" json:"transTime,omitempty"`
	BlockHeight          int64    `protobuf:"varint,6,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	BlockHash            string   `protobuf:"bytes,7,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockTime            int64    `protobuf:"varint,8,opt,name=blockTime,proto3" json:"blockTime,omitempty"`
	CreateTime           int64    `protobuf:"varint,9,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,10,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebe87db631c5431, []int{1}
}

func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockData.Unmarshal(m, b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
}
func (m *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(m, src)
}
func (m *BlockData) XXX_Size() int {
	return xxx_messageInfo_BlockData.Size(m)
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

func (m *BlockData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BlockData) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *BlockData) GetDataHash() string {
	if m != nil {
		return m.DataHash
	}
	return ""
}

func (m *BlockData) GetTransHash() string {
	if m != nil {
		return m.TransHash
	}
	return ""
}

func (m *BlockData) GetTransTime() int64 {
	if m != nil {
		return m.TransTime
	}
	return 0
}

func (m *BlockData) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *BlockData) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *BlockData) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *BlockData) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *BlockData) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type Request struct {
	DataName             string   `protobuf:"bytes,1,opt,name=dataName,proto3" json:"dataName,omitempty"`
	BlockContent         string   `protobuf:"bytes,2,opt,name=blockContent,proto3" json:"blockContent,omitempty"`
	FileUrl              string   `protobuf:"bytes,3,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
	DataHash             string   `protobuf:"bytes,4,opt,name=dataHash,proto3" json:"dataHash,omitempty"`
	Limit                string   `protobuf:"bytes,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebe87db631c5431, []int{2}
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

func (m *Request) GetDataName() string {
	if m != nil {
		return m.DataName
	}
	return ""
}

func (m *Request) GetBlockContent() string {
	if m != nil {
		return m.BlockContent
	}
	return ""
}

func (m *Request) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *Request) GetDataHash() string {
	if m != nil {
		return m.DataHash
	}
	return ""
}

func (m *Request) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebe87db631c5431, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Response struct {
	Datamodel            *DataModel        `protobuf:"bytes,1,opt,name=datamodel,proto3" json:"datamodel,omitempty"`
	Datamodels           []*DataModel      `protobuf:"bytes,2,rep,name=datamodels,proto3" json:"datamodels,omitempty"`
	BlockData            *BlockData        `protobuf:"bytes,3,opt,name=blockData,proto3" json:"blockData,omitempty"`
	BlockDatas           []*BlockData      `protobuf:"bytes,4,rep,name=blockDatas,proto3" json:"blockDatas,omitempty"`
	Errors               []*Error          `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	Code                 int32             `protobuf:"varint,6,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string            `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 string            `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	ResultMap            map[string]string `protobuf:"bytes,9,rep,name=resultMap,proto3" json:"resultMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_eebe87db631c5431, []int{4}
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

func (m *Response) GetDatamodel() *DataModel {
	if m != nil {
		return m.Datamodel
	}
	return nil
}

func (m *Response) GetDatamodels() []*DataModel {
	if m != nil {
		return m.Datamodels
	}
	return nil
}

func (m *Response) GetBlockData() *BlockData {
	if m != nil {
		return m.BlockData
	}
	return nil
}

func (m *Response) GetBlockDatas() []*BlockData {
	if m != nil {
		return m.BlockDatas
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

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

func (m *Response) GetResultMap() map[string]string {
	if m != nil {
		return m.ResultMap
	}
	return nil
}

func init() {
	proto.RegisterType((*DataModel)(nil), "go.micro.srv.dotc.DataModel")
	proto.RegisterType((*BlockData)(nil), "go.micro.srv.dotc.BlockData")
	proto.RegisterType((*Request)(nil), "go.micro.srv.dotc.Request")
	proto.RegisterType((*Error)(nil), "go.micro.srv.dotc.Error")
	proto.RegisterType((*Response)(nil), "go.micro.srv.dotc.Response")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.srv.dotc.Response.ResultMapEntry")
}

func init() { proto.RegisterFile("proto/dotc/dotc.proto", fileDescriptor_eebe87db631c5431) }

var fileDescriptor_eebe87db631c5431 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdf, 0x4e, 0xdb, 0x3e,
	0x14, 0xfe, 0x35, 0x4d, 0x0b, 0x39, 0xfc, 0xc4, 0xc0, 0xda, 0xa4, 0xa8, 0x43, 0x53, 0x95, 0x2b,
	0xb4, 0x8b, 0x6c, 0x62, 0x37, 0x13, 0x62, 0x17, 0x83, 0x32, 0xb1, 0x09, 0xb8, 0xf0, 0xc6, 0x03,
	0x98, 0xc4, 0x94, 0x08, 0x27, 0xce, 0x6c, 0x17, 0x89, 0x07, 0xd9, 0x73, 0xec, 0xdd, 0x78, 0x82,
	0xc9, 0xc7, 0xf9, 0x57, 0xe8, 0xda, 0x8b, 0xee, 0xa6, 0xf2, 0xf9, 0xf7, 0x9d, 0x2f, 0xdf, 0x39,
	0x76, 0xe1, 0x55, 0xa9, 0xa4, 0x91, 0xef, 0x52, 0x69, 0x12, 0xfc, 0x89, 0xd1, 0x26, 0xbb, 0x53,
	0x19, 0xe7, 0x59, 0xa2, 0x64, 0xac, 0xd5, 0x7d, 0x6c, 0x03, 0xd1, 0x14, 0x82, 0x09, 0x33, 0xec,
	0x42, 0xa6, 0x5c, 0x90, 0x11, 0x6c, 0xa6, 0xcc, 0xb0, 0x4b, 0x96, 0xf3, 0xb0, 0x37, 0xee, 0xed,
	0x07, 0xb4, 0xb1, 0xc9, 0x1b, 0x80, 0x44, 0x71, 0x66, 0xf8, 0x8f, 0x2c, 0xe7, 0xa1, 0x37, 0xee,
	0xed, 0xf7, 0x69, 0xc7, 0x63, 0xe3, 0xb3, 0x32, 0xad, 0xe3, 0x7d, 0x17, 0x6f, 0x3d, 0xd1, 0x6f,
	0x0f, 0x82, 0x63, 0x21, 0x93, 0x3b, 0xdb, 0x8e, 0x6c, 0x83, 0x97, 0xa5, 0x55, 0x0f, 0x2f, 0x4b,
	0xc9, 0x1e, 0x04, 0xb9, 0xa5, 0x80, 0xad, 0x3d, 0x74, 0xb7, 0x8e, 0x9a, 0xd7, 0x19, 0xd3, 0xb7,
	0x88, 0x5c, 0xf1, 0xb2, 0xb6, 0xad, 0x34, 0x8a, 0x15, 0x1a, 0x83, 0xbe, 0xab, 0x6c, 0x1c, 0x4d,
	0x14, 0x49, 0x0d, 0x90, 0x54, 0xeb, 0x20, 0x63, 0xd8, 0xba, 0xb6, 0x94, 0xce, 0x78, 0x36, 0xbd,
	0x35, 0xe1, 0x10, 0xe3, 0x5d, 0x97, 0xad, 0x77, 0xa6, 0x45, 0xdf, 0x70, 0xe8, 0x8d, 0xa3, 0x89,
	0x22, 0xfa, 0xa6, 0x43, 0x6f, 0x1c, 0x4f, 0x14, 0x0b, 0x56, 0x28, 0x06, 0xcf, 0x14, 0xfb, 0xd5,
	0x83, 0x0d, 0xca, 0x7f, 0xce, 0xb8, 0x36, 0x4b, 0x27, 0x13, 0xc1, 0xff, 0xd8, 0xf4, 0x44, 0x16,
	0x86, 0x17, 0xa6, 0x92, 0x6f, 0xce, 0x47, 0x42, 0xd8, 0xb8, 0xc9, 0x04, 0xbf, 0x52, 0xa2, 0x12,
	0xb0, 0x36, 0xe7, 0xb4, 0xf5, 0x9f, 0x68, 0xfb, 0x12, 0x06, 0x22, 0xcb, 0x33, 0x83, 0xca, 0x05,
	0xd4, 0x19, 0xd1, 0x27, 0x18, 0x9c, 0x2a, 0x25, 0x15, 0x21, 0xe0, 0x27, 0x32, 0x75, 0x84, 0x06,
	0x14, 0xcf, 0x56, 0xd2, 0x94, 0xeb, 0x44, 0x65, 0xa5, 0xc9, 0x64, 0x51, 0x71, 0xe9, 0xba, 0xa2,
	0xc7, 0x3e, 0x6c, 0x52, 0xae, 0x4b, 0x59, 0x68, 0x4e, 0x0e, 0x21, 0xb0, 0xdd, 0x70, 0xd4, 0x88,
	0xb3, 0x75, 0xb0, 0x17, 0x3f, 0xdb, 0xd2, 0xb8, 0x59, 0x51, 0xda, 0xa6, 0x93, 0x23, 0x80, 0xc6,
	0xd0, 0xa1, 0x37, 0xee, 0xaf, 0x2c, 0xee, 0xe4, 0xdb, 0xce, 0xd7, 0xf5, 0x3a, 0xa2, 0x26, 0x8b,
	0x8b, 0x9b, 0x95, 0xa5, 0x6d, 0xba, 0xed, 0xdc, 0x18, 0x3a, 0xf4, 0xff, 0xda, 0xb9, 0x2d, 0xee,
	0xe4, 0x93, 0xf7, 0x30, 0xe4, 0x56, 0x3f, 0x1d, 0x0e, 0xb0, 0x32, 0x5c, 0x50, 0x89, 0x02, 0xd3,
	0x2a, 0xaf, 0x11, 0x7a, 0xd8, 0x11, 0x7a, 0x07, 0xfa, 0xb9, 0x9e, 0x56, 0x3b, 0x69, 0x8f, 0x36,
	0xcb, 0x7e, 0x1f, 0x2e, 0x62, 0x40, 0xf1, 0x4c, 0xce, 0x20, 0x50, 0x5c, 0xcf, 0x84, 0xb9, 0x60,
	0x65, 0x18, 0x60, 0xbb, 0xb7, 0x0b, 0xda, 0xd5, 0xf3, 0xb0, 0x07, 0x97, 0x7c, 0x5a, 0x18, 0xf5,
	0x40, 0xdb, 0xe2, 0xd1, 0x11, 0x6c, 0xcf, 0x07, 0x2d, 0x83, 0x3b, 0xfe, 0x50, 0xad, 0xa3, 0x3d,
	0xda, 0x7d, 0xb9, 0x67, 0x62, 0x56, 0xdf, 0x60, 0x67, 0x1c, 0x7a, 0x1f, 0x7b, 0x07, 0x8f, 0x3e,
	0xf8, 0x13, 0x69, 0x12, 0xf2, 0x0d, 0xb6, 0x4e, 0xf0, 0x0a, 0xb8, 0x17, 0x67, 0xe9, 0xbc, 0x46,
	0xaf, 0x97, 0x50, 0x8d, 0xfe, 0xb3, 0x58, 0x13, 0x2e, 0xf8, 0x3f, 0xc1, 0x3a, 0x87, 0x17, 0x57,
	0xa5, 0x90, 0x2c, 0x6d, 0xdf, 0xa8, 0xd1, 0xc2, 0x0a, 0xbc, 0x8f, 0xab, 0xd0, 0xbe, 0x40, 0xe0,
	0xd0, 0xae, 0xe8, 0xf9, 0x3a, 0x38, 0x97, 0xb0, 0xf3, 0x9d, 0x0b, 0x9e, 0x98, 0xe3, 0x87, 0x49,
	0x7d, 0x29, 0x97, 0xc1, 0x2d, 0x5d, 0x42, 0xc4, 0xdb, 0x3d, 0x97, 0x2c, 0xfd, 0x2c, 0xc4, 0x71,
	0xbb, 0x8f, 0x6b, 0xf1, 0xdb, 0x75, 0xdf, 0x79, 0x80, 0x78, 0x27, 0xb7, 0x2c, 0x2b, 0xd6, 0xc1,
	0xc3, 0x29, 0xd8, 0x07, 0x10, 0xe1, 0xbe, 0x16, 0x37, 0x72, 0x0d, 0xb4, 0xeb, 0x21, 0xfe, 0xeb,
	0x7d, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x27, 0x4d, 0x74, 0x0e, 0x07, 0x00, 0x00,
}
