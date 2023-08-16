package config

import (
	"os"

	"github.com/spf13/viper"
	"k8s.io/klog"
)

const (
	FileName = "config"
	FileType = "yaml"
	FilePath = "./conf"
)

var Config *Configuration

type Configuration struct {
	Database Database
}

func Init(configName string) *viper.Viper {
	v := viper.New()
	v.SetConfigType(FileType)
	v.SetConfigName(configName)
	v.AddConfigPath("./conf")
	v.AddConfigPath("../conf")
	v.AddConfigPath("../../conf")
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("errno is %+v", err)
	}
	return v
}

// GetConfig helps you to get configuration data
func GetConfig(vip *viper.Viper) *Configuration {
	setting := new(Configuration)
	// unmarshal config
	if err := vip.Unmarshal(setting); err != nil {
		klog.Errorf("Unmarshal yaml failed: %v", err)
		os.Exit(-1)
	}
	return setting
}
