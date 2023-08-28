package configs

import (
	"douyin-lite/internal/entity"
	"douyin-lite/pkg/ormx"
	"douyin-lite/pkg/storage"
)

const (
	USERNAME       = "douyin"                                                                         //账号
	PASSWORD       = "douyin@2023"                                                                    //密码
	HOST           = "39.105.199.147"                                                                 //数据库地址，可以是Ip或者域名
	PORT           = "3306"                                                                           //数据库端口
	DBNAME         = "douyin"                                                                         //数据库名
	PARAMETERS     = "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true&timeout=10s" //连接超时，10秒
	DEBUG          = false                                                                            // 是否开启debug模式
	DB_TYPE        = "mysql"                                                                          // 数据库类型
	MAX_LIFE_TIME  = 7200                                                                             // 设置连接可以重复使用的最长时间(s)
	MAX_OPEN_CONNS = 150                                                                              // 最大开启链接
	MAX_IDLE_CONNS = 50                                                                               // 设置空闲状态下的最大连接数
	TABLE_PREFIX   = ""
)

func DbInit() error {
	// 数据库初始化
	err := storage.InitDB(ormx.DBConfig{
		Debug:        DEBUG,
		DBType:       DB_TYPE,
		MaxLifetime:  MAX_LIFE_TIME,
		MaxOpenConns: MAX_OPEN_CONNS,
		MaxIdleConns: MAX_IDLE_CONNS,
		TablePrefix:  TABLE_PREFIX,
		Host:         HOST,
		Port:         PORT,
		User:         USERNAME,
		Password:     PASSWORD,
		DBName:       DBNAME,
		Parameters:   PARAMETERS,
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
	)
	return err
}
