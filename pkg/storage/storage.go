package storage

import (
	"douyin-lite/pkg/ormx"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RdbUserCount *redis.Client

var RdbFollower *redis.Client

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
