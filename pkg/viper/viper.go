package viper

import (
	v "github.com/spf13/viper"
	"k8s.io/klog"
)

const (
	FileType = "yaml"
	FilePath = "./conf"
)

type Config struct {
	Viper *v.Viper
}

func Init(configName string) Config {
	//config := Config{Viper: viper.New()}
	//v := config.Viper
	v.SetConfigType(FileType)
	v.SetConfigName(configName)
	v.AddConfigPath("./conf")
	v.AddConfigPath("../conf")
	v.AddConfigPath("../../conf")
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("errno is %+v", err)
	}
	return Config{Viper: v.GetViper()}
}
