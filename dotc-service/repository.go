package main

import (
    "github.com/jinzhu/gorm"
    pb "github.com/tanqiyou1990/vircle-client/dotc-service/proto/dotc"
)

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
}

type DotcRepository struct {
    db *gorm.DB
}

func (repo *DotcRepository) InsertModel(d *pb.DataModel) error  {
	if err := repo.db.Create(d).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DotcRepository) GetAllModels() ([]*pb.DataModel, error)  {
	var mds []*pb.DataModel
	if err := repo.db.Find(&mds).Error; err != nil {
		return nil, err
	}
	return mds, nil
}

func (repo *DotcRepository) GetOneModel(name string) (*pb.DataModel, error)  {
	var md *pb.DataModel
	md.DataName = name
	if err := repo.db.First(&md).Error; err != nil {
		return nil, err
	}
	return md, nil
}

func (repo *DotcRepository) DeleteModel(name string) error {
	if err := repo.db.Where(&pb.DataModel{DataName: name}).Delete(pb.DataModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DotcRepository) InsertBlockData(d *pb.BlockData) error  {
	if err := repo.db.Create(d).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DotcRepository) GetAllDatas() ([]*pb.BlockData, error)  {
	var bds []*pb.BlockData
	if err := repo.db.Find(&bds).Error; err != nil {
		return nil, err
	}
	return bds, nil
}

func (repo *DotcRepository) GetOneData(id string) (*pb.BlockData, error)  {
	var bd *pb.BlockData
	bd.Id = id
	if err := repo.db.First(&bd).Error; err != nil {
		return nil, err
	}
	return bd, nil
}

func (repo *DotcRepository) GetByDataHash(hash string) (*pb.BlockData, error) {
	var bd *pb.BlockData
	bd.DataHash = hash
	if err := repo.db.First(&bd).Error; err != nil {
		return nil, err
	}
	return bd, nil
}

func (repo *DotcRepository) DeleteBlockData(id string) error {
	if err := repo.db.Where(&pb.BlockData{Id: id}).Delete(pb.BlockData{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DotcRepository) GetUnUploadDatas(limit int) ([]*pb.BlockData, error)  {
	var bds []*pb.BlockData
	if limit < 0 {
		err := repo.db.Model(&pb.BlockData{}).Where("DataHash <> ? AND TransHash = ?","-1","-1").Find(&bds).Error
		if err != nil {
			return nil, err
		}

		return bds, nil
	} else {
		err := repo.db.Model(&pb.BlockData{}).Where("DataHash <> ? AND TransHash = ?","-1","-1").Limit(limit).Find(&bds).Error
		if err != nil {
			return nil, err
		}

		return bds, nil
	}

}

func (repo *DotcRepository) GetUnUpdateDatas(limit int) ([]*pb.BlockData, error)  {
	var bds []*pb.BlockData
	if limit < 0 {
		err := repo.db.Model(&pb.BlockData{}).Where("DataHash <> ? AND TransHash <> ? AND BlockHash = ?","-1","-1","-1").Find(&bds).Error
		if err != nil {
			return nil, err
		}

		return bds, nil
	} else {
		err := repo.db.Model(&pb.BlockData{}).Where("DataHash <> ? AND TransHash <> ? AND BlockHash = ?","-1","-1","-1").Limit(limit).Find(&bds).Error
		if err != nil {
			return nil, err
		}

		return bds, nil
	}
}