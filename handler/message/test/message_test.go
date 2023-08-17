package test

import (
	"douyin-lite/handler/message"
	"douyin-lite/repository"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/douyin/message/action/", message.SendMessageHandler)
	r.GET("/douyin/message/chat/", message.QueryMessageHandler)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
