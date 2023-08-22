package storage

import (
	"douyin-lite/pkg/ormx"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg ormx.DBConfig) error {
	db, err := newGormDB(cfg)
	if err == nil {
		DB = db
	}
	return err
}

func newGormDB(cfg ormx.DBConfig) (*gorm.DB, error) {
	return ormx.New(cfg)
}
