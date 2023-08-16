package user_service

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestLoginUser(t *testing.T) {
	repository.Init()
	loginInfo, err := LoginUser("feyman", "ywuiqjqiq")
	if err != nil {
		panic(err)
	}
	fmt.Println(loginInfo)
}
