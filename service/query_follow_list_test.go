package service

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryFollowListInfo(t *testing.T) {
	repository.Init()
	followListInfo, err := QueryFollowListInfo(2)
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followListInfo.UserInfoList {
		fmt.Println(userinfo)
	}
}
