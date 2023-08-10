package repository

import (
	"fmt"
	"testing"
)

func TestFollowerDao_QueryFollowerListByHostId(t *testing.T) {
	Init()
	res, err := followingDao.QueryFollowerListByHostId(3)
	if err != nil {
		panic(err)
	}
	for _, id := range res {
		fmt.Println(id.HostId, " ", id.GuestId)
	}
}

func TestFollowerDao_IncFollowerCnt(t *testing.T) {
	Init()
	err := followingDao.IncFollowerCnt(4)
	if err != nil {
		panic(err)
	}
}
