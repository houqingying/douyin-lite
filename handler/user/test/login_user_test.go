package test

import (
	"douyin-lite/handler/user"
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestLoginUserHandler(t *testing.T) {
	repository.Init()
	loginResp, err := user.LoginUserHandler("mao122", "123")
	if err != nil {
		panic(err)
	}
	fmt.Println(loginResp)
}
