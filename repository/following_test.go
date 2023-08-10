package repository

import (
	"fmt"
	"testing"
)

func TestFollowingDao_CreateFollowing(t *testing.T) {
	Init()
	followingDao.CreateFollowing(1, 2)
	followingDao.CreateFollowing(1, 3)
}

func TestFollowingDao_QueryFollowingListByHostId(t *testing.T) {
	Init()
	UserList, err := followingDao.QueryFollowingListByHostId(4)
	if err != nil {
		panic(err)
	}
	for _, user := range UserList {
		fmt.Println(user.Name)
	}
}

func TestFollowingDao_IncFollowingCnt(t *testing.T) {
	Init()
	err := followingDao.IncFollowingCnt(4)
	if err != nil {
		panic(err)
	}
}

func TestFollowingDao_FollowAction(t *testing.T) {
	Init()
	err := followingDao.FollowAction(1, 2)
	if err != nil {
		panic(err)
	}
	err = followingDao.FollowAction(1, 3)
	if err != nil {
		panic(err)
	}
	err = followingDao.FollowAction(2, 3)
	if err != nil {
		panic(err)
	}
	err = followingDao.FollowAction(2, 1)
	if err != nil {
		panic(err)
	}
}
