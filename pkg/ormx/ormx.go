package ormx

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DBConfig GORM DBConfig
type DBConfig struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	Parameters   string
}

// New Create gorm.DB instance
func New(c DBConfig) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch strings.ToLower(c.DBType) {
	case "mysql":
		c.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			c.User, c.Password, c.Host, c.Port, c.DBName, c.Parameters)
		dialector = mysql.Open(c.DSN)
	//case "postgres":
	//	dialector = postgres.Open(c.DSN)
	default:
		return nil, fmt.Errorf("dialector(%s) not supported", c.DBType)
	}

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	return db, nil
}
