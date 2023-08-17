package test

import (
	"douyin-lite/handler/follow"
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryFollowerListHandler(t *testing.T) {
	repository.Init()
	followerListResp, err := follow.QueryFollowerListHandler("1")
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followerListResp.UserInfoList {
		fmt.Println(userinfo)
	}
}
