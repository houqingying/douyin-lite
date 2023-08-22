package user_service

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestNewRegisterUserFlow(t *testing.T) {
	repository.Init()
	registerInfo, err := RegisterUser("maoj", "12334")
	if err != nil {
		panic(err)
	}
	fmt.Println(registerInfo)
}
