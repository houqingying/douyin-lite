package user

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	repository.Init()
	registerResp, err := RegisterUserHandler("mao123", "123")
	if err != nil {
		panic(err)
	}
	fmt.Println(registerResp)
}
