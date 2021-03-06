// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/dotc/dotc.proto

package go_micro_srv_dotc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Dotc service

type DotcService interface {
	// rpc Call(Request) returns (Response) {}
	//Model
	CreateModel(ctx context.Context, in *DataModel, opts ...client.CallOption) (*Response, error)
	// rpc GetModelByName(Request) returns (Response) {}
	DeleteModel(ctx context.Context, in *DataModel, opts ...client.CallOption) (*Response, error)
	//BlockData
	UploadBlockData(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UploadURL(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	SelectByDataHash(ctx context.Context, in *Request, opts ...client.CallOption) (*BlockData, error)
	LoadAllBlockDatas(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Upload2BlockChain(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateBlockInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type dotcService struct {
	c    client.Client
	name string
}

func NewDotcService(name string, c client.Client) DotcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.dotc"
	}
	return &dotcService{
		c:    c,
		name: name,
	}
}

func (c *dotcService) CreateModel(ctx context.Context, in *DataModel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.CreateModel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) DeleteModel(ctx context.Context, in *DataModel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.DeleteModel", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) UploadBlockData(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.UploadBlockData", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) UploadURL(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.UploadURL", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) SelectByDataHash(ctx context.Context, in *Request, opts ...client.CallOption) (*BlockData, error) {
	req := c.c.NewRequest(c.name, "Dotc.SelectByDataHash", in)
	out := new(BlockData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) LoadAllBlockDatas(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.LoadAllBlockDatas", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) Upload2BlockChain(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.Upload2BlockChain", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dotcService) UpdateBlockInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Dotc.UpdateBlockInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dotc service

type DotcHandler interface {
	// rpc Call(Request) returns (Response) {}
	//Model
	CreateModel(context.Context, *DataModel, *Response) error
	// rpc GetModelByName(Request) returns (Response) {}
	DeleteModel(context.Context, *DataModel, *Response) error
	//BlockData
	UploadBlockData(context.Context, *Request, *Response) error
	UploadURL(context.Context, *Request, *Response) error
	SelectByDataHash(context.Context, *Request, *BlockData) error
	LoadAllBlockDatas(context.Context, *Request, *Response) error
	Upload2BlockChain(context.Context, *Request, *Response) error
	UpdateBlockInfo(context.Context, *Request, *Response) error
}

func RegisterDotcHandler(s server.Server, hdlr DotcHandler, opts ...server.HandlerOption) error {
	type dotc interface {
		CreateModel(ctx context.Context, in *DataModel, out *Response) error
		DeleteModel(ctx context.Context, in *DataModel, out *Response) error
		UploadBlockData(ctx context.Context, in *Request, out *Response) error
		UploadURL(ctx context.Context, in *Request, out *Response) error
		SelectByDataHash(ctx context.Context, in *Request, out *BlockData) error
		LoadAllBlockDatas(ctx context.Context, in *Request, out *Response) error
		Upload2BlockChain(ctx context.Context, in *Request, out *Response) error
		UpdateBlockInfo(ctx context.Context, in *Request, out *Response) error
	}
	type Dotc struct {
		dotc
	}
	h := &dotcHandler{hdlr}
	return s.Handle(s.NewHandler(&Dotc{h}, opts...))
}

type dotcHandler struct {
	DotcHandler
}

func (h *dotcHandler) CreateModel(ctx context.Context, in *DataModel, out *Response) error {
	return h.DotcHandler.CreateModel(ctx, in, out)
}

func (h *dotcHandler) DeleteModel(ctx context.Context, in *DataModel, out *Response) error {
	return h.DotcHandler.DeleteModel(ctx, in, out)
}

func (h *dotcHandler) UploadBlockData(ctx context.Context, in *Request, out *Response) error {
	return h.DotcHandler.UploadBlockData(ctx, in, out)
}

func (h *dotcHandler) UploadURL(ctx context.Context, in *Request, out *Response) error {
	return h.DotcHandler.UploadURL(ctx, in, out)
}

func (h *dotcHandler) SelectByDataHash(ctx context.Context, in *Request, out *BlockData) error {
	return h.DotcHandler.SelectByDataHash(ctx, in, out)
}

func (h *dotcHandler) LoadAllBlockDatas(ctx context.Context, in *Request, out *Response) error {
	return h.DotcHandler.LoadAllBlockDatas(ctx, in, out)
}

func (h *dotcHandler) Upload2BlockChain(ctx context.Context, in *Request, out *Response) error {
	return h.DotcHandler.Upload2BlockChain(ctx, in, out)
}

func (h *dotcHandler) UpdateBlockInfo(ctx context.Context, in *Request, out *Response) error {
	return h.DotcHandler.UpdateBlockInfo(ctx, in, out)
}
