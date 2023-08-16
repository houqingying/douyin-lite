package db

import (
	"fmt"

	"github.com/houqingying/douyin-lite/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	dbConfig = config.Init("mysql")
)

func GetDsn() string {
	Conf := config.GetConfig(dbConfig)
	dbConf := Conf.Database
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname, dbConf.Timeout)
}
func Init() error {
	dsn := GetDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate()
	return err
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return db
}
