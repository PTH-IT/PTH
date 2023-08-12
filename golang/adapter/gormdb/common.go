package gormdb

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type GormDb interface {
	Start(gormdb *gorm.DB)
	Begin() *gorm.DB
	Commit(gormdb *gorm.DB) error
	Rollback() error
}

func NewGormDb() GormDb {
	return &gormDb{}
}

type gormDb struct {
}

func (db *gormDb) Start(gormdb *gorm.DB) {
	DB = gormdb
}

func (db *gormDb) Begin() *gorm.DB {
	DB = DB.Begin()
	return DB
}

func (db *gormDb) Commit(gormdb *gorm.DB) error {
	if DB == nil {
		return nil
	}
	DB = gormdb.Statement.Commit()
	return DB.Error
}

func (db *gormDb) Rollback() error {
	if DB == nil {
		return nil
	}
	err := DB.Rollback().Error

	return err
}
