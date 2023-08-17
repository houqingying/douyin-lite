package test

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
	"time"
)

func Test_Message_Init(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}
}

func TestMessageDao_Singleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			_ = repository.GetMessageDaoInstance()
		}()
	}
	time.Sleep(time.Second)
}

func TestMessageDao_CreateMessage(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}
	err = repository.GetMessageDaoInstance().CreateMessage(1, 2, "hello")
	if err != nil {
		panic(err)
	}
}

func TestMessageDao_QueryMessage(t *testing.T) {
	err := repository.Init()
	if err != nil {
		panic(err)
	}
	var messageList []*repository.Message
	messageList, err = repository.GetMessageDaoInstance().QueryMessage(1, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(messageList))
}
