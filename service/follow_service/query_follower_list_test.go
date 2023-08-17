package follow_service

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryFollowerListInfo(t *testing.T) {
	repository.Init()
	followerListInfo, err := QueryFollowerListInfo(1)
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followerListInfo.UserInfoList {
		fmt.Println(userinfo.Name)
	}
}
