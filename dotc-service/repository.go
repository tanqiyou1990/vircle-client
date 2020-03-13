package main

import (
	pb "vircle/dotc-service/proto/dotc"

	"github.com/jinzhu/gorm"
)

// Repository 数据库操作
type Repository interface {
	InsertModel(*pb.DataModel) error
	GetAllModels() ([]*pb.DataModel, error)
	GetOneModel(name string) (*pb.DataModel, error)
	DeleteModel(name string) error
	InsertBlockData(*pb.BlockData) error
	GetAllDatas() ([]*pb.BlockData, error)
	GetOneData(id string) (*pb.BlockData, error)
	GetByDataHash(hash string) (*pb.BlockData, error)
	DeleteBlockData(id string) error
	GetUnUploadDatas(limit int) ([]*pb.BlockData, error)
	GetUnUpdateDatas(limit int) ([]*pb.BlockData, error)
	UpdateBlockDataByID(*pb.BlockData) error
}

// DotcRepository 返回gorm
type DotcRepository struct {
	db *gorm.DB
}

// InsertModel 创建一个新的数据类型
func (repo *DotcRepository) InsertModel(d *pb.DataModel) error {
	if err := repo.db.Create(d).Error; err != nil {
		return err
	}
	return nil
}

// GetAllModels 获取所有数据类型
func (repo *DotcRepository) GetAllModels() ([]*pb.DataModel, error) {
	var mds []*pb.DataModel
	if err := repo.db.Find(&mds).Error; err != nil {
		return nil, err
	}
	return mds, nil
}

// GetOneModel 通过数据类型名称查找
func (repo *DotcRepository) GetOneModel(name string) (*pb.DataModel, error) {
	md := new(pb.DataModel)
	if err := repo.db.Where(&pb.DataModel{DataName: name}).First(&md).Error; err != nil {
		return nil, err
	}
	return md, nil
}

// DeleteModel 通过数据类型名称删除
func (repo *DotcRepository) DeleteModel(name string) error {
	if err := repo.db.Where(&pb.DataModel{DataName: name}).Delete(pb.DataModel{}).Error; err != nil {
		return err
	}
	return nil
}

// InsertBlockData 插入一条上链数据
func (repo *DotcRepository) InsertBlockData(d *pb.BlockData) error {
	if err := repo.db.Create(d).Error; err != nil {
		return err
	}
	return nil
}

// GetAllDatas 获取所有上链数据
func (repo *DotcRepository) GetAllDatas() ([]*pb.BlockData, error) {
	var bds []*pb.BlockData
	if err := repo.db.Find(&bds).Error; err != nil {
		return nil, err
	}
	return bds, nil
}

// GetOneData 根据ID查找
func (repo *DotcRepository) GetOneData(id string) (*pb.BlockData, error) {
	bd := new(pb.BlockData)
	if err := repo.db.Model(&pb.BlockData{}).Where("id = ?", id).First(&bd).Error; err != nil {
		return nil, err
	}
	return bd, nil
}

// GetByDataHash 根据数据哈希查找
func (repo *DotcRepository) GetByDataHash(hash string) (*pb.BlockData, error) {
	md := new(pb.BlockData)
	if err := repo.db.Where(&pb.BlockData{DataHash: hash}).First(&md).Error; err != nil {
		return nil, err
	}
	return md, nil
}

// DeleteBlockData 根据id删除
func (repo *DotcRepository) DeleteBlockData(id string) error {
	if err := repo.db.Where(&pb.BlockData{Id: id}).Delete(pb.BlockData{}).Error; err != nil {
		return err
	}
	return nil
}

// GetUnUploadDatas 获取固定条数的待上链数据
func (repo *DotcRepository) GetUnUploadDatas(limit int) ([]*pb.BlockData, error) {
	var bds []*pb.BlockData
	if limit < 0 {
		err := repo.db.Model(&pb.BlockData{}).Where("data_hash <> ? AND trans_hash = ?", "-1", "-1").Find(&bds).Error
		if err != nil {
			return nil, err
		}

		return bds, nil
	}
	err := repo.db.Model(&pb.BlockData{}).Where("data_hash <> ? AND trans_hash = ?", "-1", "-1").Limit(limit).Find(&bds).Error
	if err != nil {
		return nil, err
	}
	return bds, nil
}

// GetUnUpdateDatas 获取固定条数的待更新区块信息的数据
func (repo *DotcRepository) GetUnUpdateDatas(limit int) ([]*pb.BlockData, error) {
	var bds []*pb.BlockData
	if limit < 0 {
		err := repo.db.Model(&pb.BlockData{}).Where("data_hash <> ? AND trans_hash <> ? AND block_hash = ?", "-1", "-1", "-1").Find(&bds).Error
		if err != nil {
			return nil, err
		}
		return bds, nil
	}
	err := repo.db.Model(&pb.BlockData{}).Where("data_hash <> ? AND trans_hash <> ? AND block_hash = ?", "-1", "-1", "-1").Limit(limit).Find(&bds).Error
	if err != nil {
		return nil, err
	}

	return bds, nil
}

// UpdateBlockDataByID 根据ID 更新
func (repo *DotcRepository) UpdateBlockDataByID(d *pb.BlockData) error {
	err := repo.db.Model(&pb.BlockData{}).Updates(&d).Error
	if err != nil {
		return err
	}
	return nil
}
