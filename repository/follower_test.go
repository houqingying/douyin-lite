package repository

import (
	"fmt"
	"testing"
)

func TestFollowingDao_QueryFollowerListByHostId(t *testing.T) {
	Init()
	userList, err := NewFollowingDaoInstance().QueryFollowerListByHostId(1)
	if err != nil {
		panic(err)
	}
	for _, user := range userList {
		fmt.Println(user.Name, " ", user.ID)
	}
}

func TestFollowingDao_IncFollowerCnt(t *testing.T) {
	Init()
	err := NewFollowingDaoInstance().IncFollowerCnt(4)
	if err != nil {
		panic(err)
	}
}

func TestFollowingDao_DecFollowerCnt(t *testing.T) {
	Init()
	err := NewFollowingDaoInstance().DecFollowerCnt(1)
	if err != nil {
		panic(err)
	}
}
