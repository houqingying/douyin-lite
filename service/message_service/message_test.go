package message_service

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}
	err = SendMessage(1, 1, "你好")
	if err != nil {
		panic(err)
	}
}

func TestQueryMessage(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}
	messageInfoList, err := QueryMessage(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(messageInfoList))
}
