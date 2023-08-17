package follow

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestQueryFollowListHandler(t *testing.T) {
	repository.Init()
	followListResp, err := QueryFollowListHandler("2")
	if err != nil {
		panic(err)
	}
	for _, userinfo := range followListResp.UserInfoList {
		fmt.Println(userinfo)
	}
}
