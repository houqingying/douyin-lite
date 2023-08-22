package entity

import (
	"douyin-lite/configs"
	"fmt"
	"testing"
	"time"
)

func Test_Message_Init(t *testing.T) {
	err := configs.Init()
	if err != nil {
		panic(err)
	}
}

func TestMessageDao_Singleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			_ = GetMessageDaoInstance()
		}()
	}
	time.Sleep(time.Second)
}

func TestMessageDao_CreateMessage(t *testing.T) {
	err := configs.Init()
	if err != nil {
		panic(err)
	}
	err = GetMessageDaoInstance().CreateMessage(1, 2, "hello")
	if err != nil {
		panic(err)
	}
}

func TestMessageDao_QueryMessage(t *testing.T) {
	err := configs.Init()
	if err != nil {
		panic(err)
	}
	var messageList []*Message
	messageList, err = GetMessageDaoInstance().QueryMessage(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(messageList))
}
