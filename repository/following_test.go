package repository

import (
	"fmt"
	"testing"
)

func TestFollowingDao_CreateFollowing(t *testing.T) {
	Init()
	err := followingDao.CreateFollowing(1, 2)
	if err != nil {
		panic(err)
	}
	err = followingDao.CreateFollowing(1, 3)
	if err != nil {
		panic(err)
	}
}

func TestFollowingDao_DeleteFollowing(t *testing.T) {
	Init()
	err := followingDao.DeleteFollowing(1, 2)
	if err != nil {
		panic(err)
	}
	err = followingDao.DeleteFollowing(1, 3)
	if err != nil {
		panic(err)
	}
}

func TestFollowingDao_QueryFollowingListByHostId(t *testing.T) {
	Init()
	UserList, err := followingDao.QueryFollowingListByHostId(1)
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

func TestFollowingDao_DecFollowingCnt(t *testing.T) {
	Init()
	err := followingDao.DecFollowingCnt(1)
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

func TestFollowingDao_UnfollowAction(t *testing.T) {
	Init()
	err := followingDao.UnfollowAction(1, 2)
	if err != nil {
		panic(err)
	}
}
