package repository

import (
	"fmt"
	"testing"
)

func TestFollowingDao_QueryFollowerListByHostId(t *testing.T) {
	Init()
	res, err := followingDao.QueryFollowerListByHostId(3)
	if err != nil {
		panic(err)
	}
	for _, id := range res {
		fmt.Println(id.HostId, " ", id.GuestId)
	}
}

func TestFollowingDao_IncFollowerCnt(t *testing.T) {
	Init()
	err := followingDao.IncFollowerCnt(4)
	if err != nil {
		panic(err)
	}
}

func TestFollowingDao_DecFollowerCnt(t *testing.T) {
	Init()
	err := followingDao.DecFollowerCnt(1)
	if err != nil {
		panic(err)
	}
}
