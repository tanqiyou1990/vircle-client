package main

import (
	"context"
	"errors"
	"time"
	pu "vircle/uc-service/proto/uc"

	"github.com/jinzhu/gorm"
)

// Uc 接收链上数据查询结果
type Uc struct {
	repo Repository
}

// CreateUser 创建一个用户
func (e *Uc) CreateUser(ctx context.Context, req *pu.User, rsp *pu.Response) error {

	req.CreateTime = time.Now().Unix()
	req.UpdateTime = time.Now().Unix()
	if req.Name == "" || req.Phone == "" {
		return errors.New("缺少必要参数")
	}
	req.UseFlag = 1
	if err := e.repo.InsertUser(req); err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.User = req
	return nil
}

// DeleteUser 删除一个用户
func (e *Uc) DeleteUser(ctx context.Context, req *pu.Request, rsp *pu.Response) error {

	if req.Id == "" {
		return errors.New("缺少必要参数")
	}

	err := e.repo.DeleteUser(req.Id)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	return nil
}

// CreateBlockAccount 记录一个钱包账户
func (e *Uc) CreateBlockAccount(ctx context.Context, req *pu.BlockAccount, rsp *pu.Response) error {

	req.CreateTime = time.Now().Unix()
	req.UpdateTime = time.Now().Unix()
	if req.UserId == "" || req.Mnemonic == "" || req.Name == "" {
		return errors.New("缺少必要参数")
	}
	if err := e.repo.InsertBlockAccount(req); err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Account = req
	return nil
}

// FindAccountsByUID 查找用户的所有钱包账户
func (e *Uc) FindAccountsByUID(ctx context.Context, req *pu.Request, rsp *pu.Response) error {

	if req.UserId == "" {
		return errors.New("缺少必要参数")
	}

	a := &pu.BlockAccount{
		UserId: req.UserId,
	}

	list, err := e.repo.FindAccounts(a)
	//对于返回空的非异常情况
	if err == gorm.ErrRecordNotFound {
		rsp.Msg = "操作成功"
		rsp.Accounts = nil
		return nil
	}
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Accounts = list
	return nil
}
