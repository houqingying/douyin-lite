package message

import (
	"douyin-lite/repository"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMessageHandler(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/douyin/message/action/", SendMessageHandler)
	r.GET("/douyin/message/chat/", QueryMessageHandler)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
