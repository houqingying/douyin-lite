package config

import (
	"os"

	"github.com/fsnotify/fsnotify"
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
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// Setup initialize configuration
func Setup() error {
	var configuration *Configuration

	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)
	viper.AddConfigPath(FilePath)

	if err := viper.ReadInConfig(); err != nil {
		klog.Errorf("Error reading config file, %s", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("Config file changed: %s\n", e.Name)
		Config = GetConfig(viper.GetViper())
	})
	viper.AllSettings()
	Config = GetConfig(viper.GetViper())

	Config = configuration

	return nil
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
