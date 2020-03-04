package main

import (
	"fmt"
	"time"
	"errors"
	"context"
	"github.com/micro/go-micro/util/log"
	block "github.com/tanqiyou1990/vircle-client/block-service/proto/block"
	dotc "dotc-service/proto/dotc"
)

type Dotc struct{
	repo Repository
	blockClient block.BlockService
}

func (e *Dotc) CreateModel(ctx context.Context, req *dotc.DataModel, rsp *dotc.Response) error {
	log.Log("--CreateModel--")
	if req.DataName == "" {
		return errors.New("数据名称不能为空")
	}
	req.CreateTime = time.Now().Unix()
	req.UpdateTime = time.Now().Unix()
	if err := e.repo.InsertModel(req); err != nil {
		return errors.New(fmt.Sprintf("error creating DataModel: %v", err))
	}
	rsp.Msg = "操作成功"
	return nil
}

// func (e *Dotc) GetModelByName(ctx context.Context, req *dotc.DataModel, rsp *dotc.Response) error {
// 	log.Log("--GetModelByName--")
// 	if req.DataName == "" {
// 		return errors.New("参数不能为空")
// 	}
// 	m, err := e.repo.GetOneModel(req.DataName)
// 	if err != nil {
// 		return err
// 	}
// 	rsp.DataModel = m
// 	return nil
// }

func (e *Dotc) DeleteModel(ctx context.Context, req *dotc.DataModel, rsp *dotc.Response) error {
	log.Log("--DeleteModel--")
	if req.DataName == "" {
		return errors.New("参数不能为空")
	}
	err := e.repo.DeleteModel(req.DataName)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	return nil
}

func (e *Dotc) UploadBlockData(ctx context.Context, req *dotc.Request, rsp *dotc.Response) error {
	log.Log("--UploadBlockData--")
	if req.DataName == "" || req.BlockContent == "" {
		return errors.New("参数不能为空")
	}

	//检查数据类型是否存在
	m, err := e.repo.GetOneModel(req.DataName)
	if err != nil {
		return err
	}
	if m == nil {
		return errors.New("数据类型不存在")
	}

	//上传IPFS
	ipfsResp, err := e.vesselClient.UploadIpfsContent(ctx.Background(),&block.Request{
		content: req.BlockContent
	})
	if err !=nil {
		return err
	}

	var blockdt *dotc.BlockData
	blockdt.ModelName = req.DataName
	blockdt.DataHash = ipfsResp.Data
	blockdt.transHash = "-1"
	blockdt.blockHash = "-1"
	blockdt.createTime = time.Now().Unix()
	blockdt.updateTime = time.Now().Unix()

	err = e.repo.InsertBlockData(blockdt)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Data = ipfsResp.Data
	return nil

}