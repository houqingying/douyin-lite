package user

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryUserInfoHandler(t *testing.T) {
	repository.Init()
	userInfoResp, err := QueryUserInfo("21")
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp.User)
}
