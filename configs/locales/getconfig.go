package locales

import (
	"github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	System *System           `yaml:"system"`
	MySql  map[string]*MySql `yaml:"mysql"`
	Redis  map[string]*Redis `yaml:"redis"`
}

type System struct {
	StartTime string `yaml:"startTime"`
	MachineID int    `yaml:"machineID"`
}

type MySql struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Dbname       string `yaml:"dbname"`
	Parameters   string `yaml:"parameters"`
	Debug        bool   `yaml:"debug"`
	DbType       string `yaml:"dbType"`
	MaxLifeTime  int    `yaml:"maxLifeTime"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxLdleConns int    `yaml:"maxLdleConns"`
	TablePrefix  string `yaml:"tablePrefix"`
}

type Redis struct {
	RedisAddr     string `yaml:"redisAddr"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/configs/locales")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
