package configs

import (
	conf "douyin-lite/configs/locales"
	"douyin-lite/internal/entity"
	"douyin-lite/pkg/ormx"
	"douyin-lite/pkg/storage"
)

func DbInit() error {
	mConfig := conf.Config.MySql["default"]
	// 数据库初始化
	err := storage.InitDB(ormx.DBConfig{
		Debug:        mConfig.Debug,
		DBType:       mConfig.DbType,
		MaxLifetime:  mConfig.MaxLifeTime,
		MaxOpenConns: mConfig.MaxOpenConns,
		MaxIdleConns: mConfig.MaxLdleConns,
		TablePrefix:  mConfig.TablePrefix,
		Host:         mConfig.Host,
		Port:         mConfig.Port,
		User:         mConfig.Username,
		Password:     mConfig.Password,
		DBName:       mConfig.Dbname,
		Parameters:   mConfig.Parameters,
	})
	if err != nil {
		return err
	}

	err = storage.DB.Debug().Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		&entity.Following{},
		&entity.User{},
		&entity.Message{},
		&entity.Comment{},
		&entity.Video{},
		&entity.Count{},
		&entity.Favorite{},
	)
	return err
}
