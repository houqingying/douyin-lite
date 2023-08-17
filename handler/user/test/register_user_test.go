package test

import (
	"douyin-lite/handler/user"
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	repository.Init()
	registerResp, err := user.RegisterUserHandler("mao123", "123")
	if err != nil {
		panic(err)
	}
	fmt.Println(registerResp)
}
