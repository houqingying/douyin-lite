package test

import (
	"douyin-lite/repository"
	"douyin-lite/service/user_service"
	"fmt"
	"testing"
)

func TestLoginUser(t *testing.T) {
	repository.Init()
	loginInfo, err := user_service.LoginUser("feyman", "ywuiqjqiq")
	if err != nil {
		panic(err)
	}
	fmt.Println(loginInfo)
}
