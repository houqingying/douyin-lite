package user_service

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestNewQueryUserInfoFlow(t *testing.T) {
	repository.Init()
	userInfo, err := QueryUserInfo(21)
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfo)
}
