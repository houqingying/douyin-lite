package test

import (
	"douyin-lite/handler/user"
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryUserInfoHandler(t *testing.T) {
	repository.Init()
	userInfoResp, err := user.QueryUserInfoHandler("21")
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp.User)
}
