package main

import (
	"context"
	"errors"
	"vircle/block-service/chain"
	"vircle/block-service/ipfs"
	block "vircle/block-service/proto/block"
)

// Block IPFS及链上相关功能
type Block struct{}

// UploadIpfsContent 将数据上传至IPFS
func (e *Block) UploadIpfsContent(ctx context.Context, req *block.Request, rsp *block.Response) error {
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

// UploadIpfsURL 根据文件的URL上传到IPFS
func (e *Block) UploadIpfsURL(ctx context.Context, req *block.Request, rsp *block.Response) error {
	if req.IpfsUrl == "" {
		return errors.New("参数不能为空")
	}

	hash, err := ipfs.UploadURLFile(req.IpfsUrl)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Data = hash
	return nil
}

// UploadBlockData 将字符串通过创建一笔交易发送到链上，并返回交易哈希
func (e *Block) UploadBlockData(ctx context.Context, req *block.Request, rsp *block.Response) error {
	if req.Content == "" {
		return errors.New("参数不能为空")
	}

	txid, err := chain.CreateTxWithContent(req.Content)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Txid = txid
	return nil
}

// GetBlockInfoByTxid 根据交易哈希查找区块信息
func (e *Block) GetBlockInfoByTxid(ctx context.Context, req *block.Request, rsp *block.Response) error {
	if req.GetTxid() == "" {
		return errors.New("参数不能为空")
	}

	blockJSON, err := chain.CheckBlockData(req.GetTxid())
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Data = blockJSON
	return nil
}
