package test

import (
	"douyin-lite/repository"
	"douyin-lite/service/follow_service"
	"fmt"
	"testing"
)

func TestQueryFollowListInfo(t *testing.T) {
	repository.Init()
	followListInfo, err := follow_service.QueryFollowListInfo(2)
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followListInfo.UserInfoList {
		fmt.Println(userinfo)
	}
}
