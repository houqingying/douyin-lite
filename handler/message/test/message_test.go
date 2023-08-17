package test

import (
	"douyin-lite/handler/message"
	"douyin-lite/repository"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	repository.Init()
	r := gin.Default()
	r.POST("/douyin/message/action/", message.SendMessageHandler)
	r.GET("/douyin/message/chat/", message.QueryMessageHandler)
	r.Run()
}
