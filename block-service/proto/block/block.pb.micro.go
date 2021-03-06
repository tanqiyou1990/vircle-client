// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/block/block.proto

package go_micro_srv_block

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

// Client API for Block service

type BlockService interface {
	//IPFS
	UploadIpfsContent(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UploadIpfsURL(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	//BlockChain
	UploadBlockData(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetBlockInfoByTxid(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type blockService struct {
	c    client.Client
	name string
}

func NewBlockService(name string, c client.Client) BlockService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.block"
	}
	return &blockService{
		c:    c,
		name: name,
	}
}

func (c *blockService) UploadIpfsContent(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Block.UploadIpfsContent", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockService) UploadIpfsURL(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Block.UploadIpfsURL", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockService) UploadBlockData(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Block.UploadBlockData", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockService) GetBlockInfoByTxid(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Block.GetBlockInfoByTxid", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Block service

type BlockHandler interface {
	//IPFS
	UploadIpfsContent(context.Context, *Request, *Response) error
	UploadIpfsURL(context.Context, *Request, *Response) error
	//BlockChain
	UploadBlockData(context.Context, *Request, *Response) error
	GetBlockInfoByTxid(context.Context, *Request, *Response) error
}

func RegisterBlockHandler(s server.Server, hdlr BlockHandler, opts ...server.HandlerOption) error {
	type block interface {
		UploadIpfsContent(ctx context.Context, in *Request, out *Response) error
		UploadIpfsURL(ctx context.Context, in *Request, out *Response) error
		UploadBlockData(ctx context.Context, in *Request, out *Response) error
		GetBlockInfoByTxid(ctx context.Context, in *Request, out *Response) error
	}
	type Block struct {
		block
	}
	h := &blockHandler{hdlr}
	return s.Handle(s.NewHandler(&Block{h}, opts...))
}

type blockHandler struct {
	BlockHandler
}

func (h *blockHandler) UploadIpfsContent(ctx context.Context, in *Request, out *Response) error {
	return h.BlockHandler.UploadIpfsContent(ctx, in, out)
}

func (h *blockHandler) UploadIpfsURL(ctx context.Context, in *Request, out *Response) error {
	return h.BlockHandler.UploadIpfsURL(ctx, in, out)
}

func (h *blockHandler) UploadBlockData(ctx context.Context, in *Request, out *Response) error {
	return h.BlockHandler.UploadBlockData(ctx, in, out)
}

func (h *blockHandler) GetBlockInfoByTxid(ctx context.Context, in *Request, out *Response) error {
	return h.BlockHandler.GetBlockInfoByTxid(ctx, in, out)
}
