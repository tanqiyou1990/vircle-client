package main

import (
	"fmt"
	"time"
	"errors"
	"context"
	"github.com/micro/go-micro/util/log"
	dotc "dotc-service/proto/dotc"
)

type Dotc struct{
	repo Repository
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

func (e *Dotc) CreateBlockData(ctx context.Context, req *dotc.BlockData, rsp *dotc.Response) error {
	log.Log("--CreateBlockData--")
	if req.ModelName == "" || req.DataHash == "" {
		return errors.New("参数异常")
	}
	m, err := e.repo.GetOneModel(req.ModelName)
	if err != nil {
		return err
	}
	if m == nil {
		return errors.New("数据类型不存在")
	}
	req.TransHash = "-1"
	req.BlockHash = "-1"
	req.CreateTime = time.Now().Unix()
	req.UpdateTime = time.Now().Unix()
	err = e.repo.InsertBlockData(req)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	return nil
}
