package test

import (
	"douyin-lite/repository"
	"douyin-lite/service/user_service"
	"fmt"
	"testing"
)

func TestNewRegisterUserFlow(t *testing.T) {
	repository.Init()
	registerInfo, err := user_service.RegisterUser("maoj", "12334")
	if err != nil {
		panic(err)
	}
	fmt.Println(registerInfo)
}
