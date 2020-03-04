package main

import (
	"errors"
	"context"
	"github.com/micro/go-micro/util/log"
	"vircle-client/block-service/ipfs"
	block "vircle-client/block-service/proto/block"
	
)

type Block struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Block) UploadIpfsContent(ctx context.Context, req *block.Request, rsp *block.Response) error {
	log.Log("-UploadIpfsContent-")
	if req.Content == "" {
		return errors.New("参数不能为空")
	}

	hash, err := ipfs.UploadIPFS(req.Content)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Data = hash
	return nil
}

func (e *Block) UploadIpfsUrl(ctx context.Context, req *block.Request, rsp *block.Response) error {
	log.Log("-UploadIpfsUrl-")
	if req.IpfsUrl == "" {
		return errors.New("参数不能为空")
	}

	hash, err := ipfs.UploadUrlFile(req.IpfsUrl)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Data = hash
	return nil
}