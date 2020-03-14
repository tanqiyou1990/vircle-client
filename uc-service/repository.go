package main

import (
	pu "vircle/uc-service/proto/uc"

	"github.com/jinzhu/gorm"
)

// Repository 数据库操作
type Repository interface {
	// user
	InsertUser(*pu.User) error
	DeleteUser(string) error
	UpdateUser(*pu.User) error
	FindUsers(*pu.User) ([]*pu.User, error)
	FindOneUser(string) (*pu.User, error)

	// account
	InsertBlockAccount(*pu.BlockAccount) error
	FindAccounts(*pu.BlockAccount) ([]*pu.BlockAccount, error)
	FindAccountByID(string) (*pu.BlockAccount, error)
}

// UcRepository 返回gorm
type UcRepository struct {
	db *gorm.DB
}

// InsertUser 创建一个用户
func (repo *UcRepository) InsertUser(d *pu.User) error {
	if err := repo.db.Create(d).Error; err != nil {
		return err
	}
	return nil
}

// InsertBlockAccount 创建一个用户钱包
func (repo *UcRepository) InsertBlockAccount(d *pu.BlockAccount) error {
	if err := repo.db.Create(d).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func (repo *UcRepository) DeleteUser(id string) error {
	if err := repo.db.Where(&pu.User{Id: id}).Delete(pu.User{}).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户
func (repo *UcRepository) UpdateUser(d *pu.User) error {
	err := repo.db.Model(&pu.User{}).Updates(&d).Error
	if err != nil {
		return err
	}
	return nil
}

// FindUsers 查找用户
func (repo *UcRepository) FindUsers(d *pu.User) ([]*pu.User, error) {
	var users []*pu.User
	err := repo.db.Model(&pu.User{}).Where(&d).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindAccounts 查找账户
func (repo *UcRepository) FindAccounts(d *pu.BlockAccount) ([]*pu.BlockAccount, error) {
	var accounts []*pu.BlockAccount
	err := repo.db.Model(&pu.BlockAccount{}).Where(&d).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// FindOneUser 查找用户
func (repo *UcRepository) FindOneUser(id string) (*pu.User, error) {
	u := new(pu.User)
	err := repo.db.Model(&pu.User{}).Where(&pu.User{Id: id}).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindAccountByID 查找钱包账户
func (repo *UcRepository) FindAccountByID(id string) (*pu.BlockAccount, error) {
	a := new(pu.BlockAccount)
	err := repo.db.Model(&pu.BlockAccount{}).Where(&pu.BlockAccount{Id: id}).First(&a).Error
	if err != nil {
		return nil, err
	}
	return a, nil
}
