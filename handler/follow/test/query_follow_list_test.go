package test

import (
	"douyin-lite/handler/follow"
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryFollowListHandler(t *testing.T) {
	repository.Init()
	followListResp, err := follow.QueryFollowListHandler("2")
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followListResp.UserInfoList {
		fmt.Println(userinfo)
	}
}
