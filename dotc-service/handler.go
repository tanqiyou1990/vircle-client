package main

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	block "vircle/block-service/proto/block"
	dotc "vircle/dotc-service/proto/dotc"

	"github.com/jinzhu/gorm"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
)

// Dotc 接收链上数据查询结果
type Dotc struct {
	repo        Repository
	blockClient block.BlockService
}

// BlockTransaction 接收链上查询数据
type BlockTransaction struct {
	Blockhash string `json:"blockhash"`
	Time      int64  `json:"time"`
}

// BlockBody 接收链上查询数据
type BlockBody struct {
	Height int64 `json:"height"`
	Time   int64 `json:"time"`
}

// CreateModel 创建一个数据类型
func (e *Dotc) CreateModel(ctx context.Context, req *dotc.DataModel, rsp *dotc.Response) error {
	if req.DataName == "" {
		return errors.New("数据名称不能为空")
	}
	req.CreateTime = time.Now().Unix()
	req.UpdateTime = time.Now().Unix()
	if err := e.repo.InsertModel(req); err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	return nil
}

// DeleteModel 删除一个数据类型
func (e *Dotc) DeleteModel(ctx context.Context, req *dotc.DataModel, rsp *dotc.Response) error {
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

// SelectByDataHash 根据数据哈希查询数据
func (e *Dotc) SelectByDataHash(ctx context.Context, req *dotc.Request, rsp *dotc.BlockData) error {
	if req.DataHash == "" {
		return errors.New("参数不能为空")
	}
	bd, err := e.repo.GetByDataHash(req.DataHash)
	//对于返回空的非异常情况
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	*rsp = *bd
	return nil
}

// LoadAllBlockDatas 查询所有数据
func (e *Dotc) LoadAllBlockDatas(ctx context.Context, req *dotc.Request, rsp *dotc.Response) error {
	datas, err := e.repo.GetAllDatas()
	//对于返回空的非异常情况
	if err == gorm.ErrRecordNotFound {
		rsp.BlockData = nil
		return nil
	}
	if err != nil {
		return err
	}
	rsp.BlockDatas = datas
	return nil
}

// UploadBlockData 数据上链
func (e *Dotc) UploadBlockData(ctx context.Context, req *dotc.Request, rsp *dotc.Response) error {
	if req.DataName == "" || req.BlockContent == "" {
		return errors.New("参数不能为空")
	}

	//检查数据类型是否存在
	m := new(dotc.DataModel)
	m, err := e.repo.GetOneModel(req.DataName)
	//对于返回空的非异常情况
	if err == gorm.ErrRecordNotFound {
		return errors.New("数据类型不存在")
	}
	if err != nil {
		return err
	}
	if m == nil {
		return errors.New("数据类型不存在")
	}

	//上传IPFS
	ctxx, cancelFn := context.WithTimeout(context.Background(), 50*time.Second)
	ipfsResp, err := e.blockClient.UploadIpfsContent(ctxx, &block.Request{Content: req.BlockContent}, func(o *microclient.CallOptions) {
		o.RequestTimeout = time.Second * 30
		o.DialTimeout = time.Second * 30
	})
	defer cancelFn()
	if err != nil {
		return err
	}
	bd := new(dotc.BlockData)
	bd, err = e.repo.GetByDataHash(ipfsResp.Data)
	if err == nil && bd != nil {
		rsp.Msg = "操作成功"
		rsp.Data = ipfsResp.Data
		rsp.BlockData = bd
		return nil
	}
	if !strings.Contains(err.Error(), "record not found") {
		return err
	}

	blockdt := new(dotc.BlockData)
	blockdt.ModelName = req.DataName
	blockdt.DataHash = ipfsResp.Data
	blockdt.TransHash = "-1"
	blockdt.BlockHash = "-1"
	blockdt.CreateTime = time.Now().Unix()
	blockdt.UpdateTime = time.Now().Unix()

	err = e.repo.InsertBlockData(blockdt)
	if err != nil {
		return err
	}
	rsp.Msg = "操作成功"
	rsp.Data = ipfsResp.Data
	rsp.BlockData = blockdt
	return nil

}

// UploadURL 根据URL上传文件
func (e *Dotc) UploadURL(ctx context.Context, req *dotc.Request, rsp *dotc.Response) error {
	if req.FileUrl == "" {
		return errors.New("参数不能为空")
	}

	//上传IPFS
	ctxx, cancelFn := context.WithTimeout(context.Background(), 50*time.Second)
	ipfsResp, err := e.blockClient.UploadIpfsURL(ctxx, &block.Request{IpfsUrl: req.FileUrl}, func(o *microclient.CallOptions) {
		o.RequestTimeout = time.Second * 30
		o.DialTimeout = time.Second * 30
	})
	defer cancelFn()
	if err != nil {
		return err
	}

	rsp.Msg = "操作成功"
	rsp.Data = ipfsResp.Data
	return nil

}

// Upload2BlockChain 数据上传到区块链
func (e *Dotc) Upload2BlockChain(ctx context.Context, req *dotc.Request, rsp *dotc.Response) error {
	var limt int
	limt = -1
	if req.Limit != "" {
		var err error
		limt, err = strconv.Atoi(req.Limit)
		if err != nil {
			return errors.New("参数错误")
		}
	}

	dataList, err := e.repo.GetUnUploadDatas(limt)
	if err != nil {
		return err
	}

	if len(dataList) < 1 {
		rsp.Msg = "无待上链数据"
		return nil
	}

	rss := make(map[string]string)

	for _, item := range dataList {

		ctxx, cancelFn := context.WithTimeout(context.Background(), 50*time.Second)
		blockRsp, err := e.blockClient.UploadBlockData(ctxx, &block.Request{Content: item.DataHash}, func(o *microclient.CallOptions) {
			o.RequestTimeout = time.Second * 30
			o.DialTimeout = time.Second * 30
		})
		defer cancelFn()
		if err != nil {
			log.Log(err)
			rss[item.DataHash] = "上传失败:" + err.Error()
			continue
		}
		if len(blockRsp.Txid) < 5 {
			log.Log("txid异常:", blockRsp.Txid)
			rss[item.DataHash] = "txid异常:" + blockRsp.Txid
			continue
		}
		rss[item.DataHash] = "成功:" + blockRsp.Txid
		item.TransHash = blockRsp.Txid
		item.UpdateTime = time.Now().Unix()
		err = e.repo.UpdateBlockDataByID(item)
		if err != nil {
			log.Log("更新区块数据异常:", blockRsp.Txid)
			rss[item.DataHash] = "更新区块数据异常:" + err.Error()
			continue
		}
	}

	rsp.Msg = "执行完毕"
	rsp.ResultMap = rss
	return nil

}

// UpdateBlockInfo 更新区块信息
func (e *Dotc) UpdateBlockInfo(ctx context.Context, req *dotc.Request, rsp *dotc.Response) error {
	var limt int
	limt = -1
	if req.Limit != "" {
		var err error
		limt, err = strconv.Atoi(req.Limit)
		if err != nil {
			return errors.New("参数错误")
		}
	}

	dataList, err := e.repo.GetUnUpdateDatas(limt)
	if err != nil {
		return err
	}

	if len(dataList) < 1 {
		rsp.Msg = "无待更新数据"
		return nil
	}

	rss := make(map[string]string)

	for _, item := range dataList {

		ctxx, cancelFn := context.WithTimeout(context.Background(), 50*time.Second)
		blockRsp, err := e.blockClient.GetBlockInfoByTxid(ctxx, &block.Request{Txid: item.TransHash}, func(o *microclient.CallOptions) {
			o.RequestTimeout = time.Second * 30
			o.DialTimeout = time.Second * 30
		})
		defer cancelFn()

		if err != nil {
			log.Log(err)
			rss[item.TransHash] = "查询失败:" + err.Error()
			continue
		}
		if len(blockRsp.Data) < 5 {
			log.Log("返回数据异常:", blockRsp.Data)
			rss[item.TransHash] = "返回数据异常:" + blockRsp.Data
			continue
		}
		rss[item.TransHash] = "查询成功"
		var resultMap map[string]string
		if err := json.Unmarshal([]byte(blockRsp.Data), &resultMap); err != nil {
			log.Log("解析数据异常:", err)
			rss[item.TransHash] = "解析数据异常:" + err.Error()
			continue
		}

		transactionJSON := new(BlockTransaction)
		if err := json.Unmarshal([]byte(resultMap["transaction"]), &transactionJSON); err != nil {
			log.Log("解析数据异常:", err)
			rss[item.TransHash] = "解析数据异常:" + err.Error()
			continue
		}
		blockJSON := new(BlockBody)
		if err := json.Unmarshal([]byte(resultMap["blockData"]), &blockJSON); err != nil {
			log.Log("解析数据异常:", err)
			rss[item.TransHash] = "解析数据异常:" + err.Error()
			continue
		}
		item.TransTime = transactionJSON.Time
		item.BlockHash = transactionJSON.Blockhash
		item.BlockHeight = blockJSON.Height
		item.BlockTime = blockJSON.Time
		item.UpdateTime = time.Now().Unix()
		err = e.repo.UpdateBlockDataByID(item)
		if err != nil {
			log.Log("更新数据异常:", blockRsp.Txid)
			rss[item.DataHash] = "更新数据异常:" + err.Error()
			continue
		}
	}

	rsp.Msg = "执行完毕"
	rsp.ResultMap = rss
	return nil

}
