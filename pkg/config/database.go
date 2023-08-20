package config

type DatabaseConfiguration struct {
	Driver   string `yaml:"driver"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	LogMode  bool   `yaml:"log_mode"`
	Timeout  string `yaml:"timeout"`
}
