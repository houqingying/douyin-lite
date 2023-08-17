package test

import (
	"douyin-lite/repository"
	"douyin-lite/service/message_service"
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	repository.Init()
	err := message_service.SendMessage(1, 1, "你好")
	if err != nil {
		panic(err)
	}
}

func TestQueryMessage(t *testing.T) {
	repository.Init()
	messageInfoList, err := message_service.QueryMessage(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(messageInfoList))
}
