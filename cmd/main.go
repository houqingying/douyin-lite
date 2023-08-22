package main

import (
	"douyin-lite/configs"
	"douyin-lite/router"
	"fmt"
)

func main() {

	if err := configs.Init(); err != nil {
		fmt.Println("初始化失败")
		return
	}

	r := router.Init()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("启动服务失败")
		return
	}
}
