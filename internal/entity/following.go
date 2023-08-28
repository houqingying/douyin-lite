package entity

import (
	"douyin-lite/pkg/storage"
	"errors"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type Following struct {
	gorm.Model
	HostId  int64 `gorm:"uniqueIndex:host_guest"`
	GuestId int64 `gorm:"uniqueIndex:host_guest"`
}

func (Following) TableName() string {
	return "following"
}

type FollowingDao struct {
}

var followingDao *FollowingDao
var followingOnce sync.Once

func NewFollowingDaoInstance() *FollowingDao {
	followingOnce.Do(func() {
		followingDao = &FollowingDao{}
	})
	return followingDao
}

func (*FollowingDao) FollowAction(hostId int64, guestId int64) error {
	errTran := storage.DB.Transaction(func(tx *gorm.DB) error {
		err := followingDao.CreateFollowing(hostId, guestId)
		if err != nil {
			return err
		}
		//TODO
		//err = followingDao.IncFollowingCnt(hostId)
		//if err != nil {
		//	return err
		//}
		//err = followingDao.IncFollowerCnt(guestId)
		//if err != nil {
		//	return err
		//}
		return nil
	})
	if errTran != nil {
		return errTran
	}
	return nil
}

func (*FollowingDao) UnfollowAction(hostId int64, guestId int64) error {
	errTran := storage.DB.Transaction(func(tx *gorm.DB) error {
		err := followingDao.DeleteFollowing(hostId, guestId)
		if err != nil {
			return err
		}
		//TODO
		//err = followingDao.DecFollowingCnt(hostId)
		//if err != nil {
		//	return err
		//}
		//err = followingDao.DecFollowerCnt(guestId)
		//if err != nil {
		//	return err
		//}
		return nil
	})
	if errTran != nil {
		return errTran
	}
	return nil
}

func (*FollowingDao) QueryFollowingListByHostId(hostId int64) ([]*User, error) {
	var FollowingList []*Following
	err := storage.DB.Where("host_id = ?", hostId).Find(&FollowingList).Error
	if err != nil {
		return nil, err
	}
	var UserList []*User
	var tempUser *User
	for _, follow := range FollowingList {
		tempUser = nil
		err := storage.DB.Where("id = ?", follow.GuestId).Find(&tempUser).Error
		if err != nil {
			return nil, err
		}
		UserList = append(UserList, tempUser)
	}
	return UserList, nil
}

func (*FollowingDao) QueryFollowingIdList(hostId int64) ([]int64, error) {
	var idList []int64
	err := storage.DB.Model(&Following{}).Select("guest_id").Where("host_id = ?", hostId).Find(&idList).Error
	if err != nil {
		return nil, err
	}
	return idList, nil
}

func (*FollowingDao) QueryFollowerIdList(hostId int64) ([]int64, error) {
	var idList []int64
	err := storage.DB.Model(&Following{}).Select("host_id").Where("guest_id = ?", hostId).Find(&idList).Error
	if err != nil {
		return nil, err
	}
	return idList, nil
}

func (*FollowingDao) CreateFollowing(hostId int64, guestId int64) error {
	newFollowing := Following{
		HostId:  hostId,
		GuestId: guestId,
	}
	err := storage.DB.Create(&newFollowing).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) DeleteFollowing(hostId int64, guestId int64) error {
	err := storage.DB.Where("host_id = ? AND guest_id = ?", hostId, guestId).
		Delete(&Following{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) QueryisFollow(hostId int64, guestId int64) (bool, error) {
	followItem := &Following{}
	err := storage.DB.Where("host_id = ? AND guest_id = ?", hostId, guestId).First(&followItem).Error
	fmt.Println(err)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	}

	return true, nil
}
