package main

import (
	"github.com/houqingying/douyin-lite/pkg/config"
)

func main() {
	// 1. init config
	// 2. init db
	config.Init("mysql")
}
