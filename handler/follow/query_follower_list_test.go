package follow

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryFollowerListHandler(t *testing.T) {
	repository.Init()
	followerListResp, err := QueryFollowerListHandler("1")
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followerListResp.UserInfoList {
		fmt.Println(userinfo)
	}
}
