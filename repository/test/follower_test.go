package test

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestFollowingDao_QueryFollowerListByHostId(t *testing.T) {
	repository.Init()
	userList, err := repository.NewFollowingDaoInstance().QueryFollowerListByHostId(1)
	if err != nil {
		panic(err)
	}
	for _, user := range userList {
		fmt.Println(user.Name, " ", user.ID)
	}
}

func TestFollowingDao_IncFollowerCnt(t *testing.T) {
	repository.Init()
	err := repository.NewFollowingDaoInstance().IncFollowerCnt(4)
	if err != nil {
		panic(err)
	}
}

func TestFollowingDao_DecFollowerCnt(t *testing.T) {
	repository.Init()
	err := repository.NewFollowingDaoInstance().DecFollowerCnt(1)
	if err != nil {
		panic(err)
	}
}
