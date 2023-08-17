package test

import (
	"douyin-lite/repository"
	"douyin-lite/service/user_service"
	"fmt"
	"testing"
)

func TestNewQueryUserInfoFlow(t *testing.T) {
	repository.Init()
	userInfo, err := user_service.QueryUserInfo(21)
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfo)
}
