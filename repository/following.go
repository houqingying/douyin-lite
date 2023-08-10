package repository

import (
	"gorm.io/gorm"
	"sync"
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

func (*FollowingDao) QueryFollowingListByHostId(hostId int64) ([]*User, error) {
	var FollowingList []*Following
	err := db.Where("host_id = ?", hostId).Find(&FollowingList).Error
	if err != nil {
		return nil, err
	}
	var UserList []*User
	var tempUser *User
	for _, follow := range FollowingList {
		tempUser = nil
		err := db.Where("id = ?", follow.GuestId).Find(&tempUser).Error
		if err != nil {
			return nil, err
		}
		UserList = append(UserList, tempUser)
	}
	return UserList, nil
}

func (*FollowingDao) CreateFollowing(hostId int64, guestId int64) error {
	newFollowing := Following{
		HostId:  hostId,
		GuestId: guestId,
	}
	err := db.Create(&newFollowing).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) IncFollowingCnt(hostId int64) error {
	err := db.Model(&User{}).Where("id = ?", hostId).
		UpdateColumn("following_count", gorm.Expr("following_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) FollowAction(hostId int64, guestId int64) error {
	errTran := db.Transaction(func(tx *gorm.DB) error {
		err := followingDao.CreateFollowing(hostId, guestId)
		if err != nil {
			return err
		}
		err = followingDao.IncFollowingCnt(hostId)
		if err != nil {
			return err
		}
		err = followingDao.IncFollowerCnt(guestId)
		if err != nil {
			return err
		}
		return nil
	})
	if errTran != nil {
		return errTran
	}
	return nil
}
