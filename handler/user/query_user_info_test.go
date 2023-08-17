package user

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryUserInfoHandler(t *testing.T) {
	repository.Init()
	userInfoResp, err := QueryUserInfoHandler("21")
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfoResp.User)
}
