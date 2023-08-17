package test

import (
	"douyin-lite/repository"
	"douyin-lite/service/follow_service"
	"fmt"
	"testing"
)

func TestQueryFollowerListInfo(t *testing.T) {
	repository.Init()
	followerListInfo, err := follow_service.QueryFollowerListInfo(1)
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followerListInfo.UserInfoList {
		fmt.Println(userinfo.Name)
	}
}
