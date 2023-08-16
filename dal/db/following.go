package db

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

const TableNameFollowing = "following"

// Following mapped from table <following>
type Following struct {
	gorm.Model
	HostID  uint `gorm:"column:host_id;not null" json:"host_id"`
	GuestID uint `gorm:"column:guest_id;not null" json:"guest_id"`
}

// TableName Following's table name
func (*Following) TableName() string {
	return TableNameFollowing
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

func (*FollowingDao) FollowAction(hostId uint, guestId uint) error {
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

func (*FollowingDao) UnfollowAction(hostId uint, guestId uint) error {
	errTran := db.Transaction(func(tx *gorm.DB) error {
		err := followingDao.DeleteFollowing(hostId, guestId)
		if err != nil {
			return err
		}
		err = followingDao.DecFollowingCnt(hostId)
		if err != nil {
			return err
		}
		err = followingDao.DecFollowerCnt(guestId)
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

func (*FollowingDao) QueryFollowingListByHostId(hostId uint) ([]*User, error) {
	var FollowingList []*Following
	err := db.Where("host_id = ?", hostId).Find(&FollowingList).Error
	if err != nil {
		return nil, err
	}
	var UserList []*User
	var tempUser *User
	for _, follow := range FollowingList {
		tempUser = nil
		err := db.Where("id = ?", follow.GuestID).Find(&tempUser).Error
		if err != nil {
			return nil, err
		}
		UserList = append(UserList, tempUser)
	}
	return UserList, nil
}

func (*FollowingDao) CreateFollowing(hostId uint, guestId uint) error {
	newFollowing := Following{
		HostID:  hostId,
		GuestID: guestId,
	}
	err := db.Create(&newFollowing).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) DeleteFollowing(hostId uint, guestId uint) error {
	err := db.Where("host_id = ? AND guest_id = ?", hostId, guestId).
		Delete(&Following{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) IncFollowingCnt(hostId uint) error {
	err := db.Model(&User{}).Where("id = ?", hostId).
		UpdateColumn("following_count", gorm.Expr("following_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) DecFollowingCnt(hostId uint) error {
	err := db.Model(&User{}).Where("id = ?", hostId).
		UpdateColumn("following_count", gorm.Expr("following_count - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) QueryisFollow(hostId uint, guestId uint) (bool, error) {
	followItem := &Following{}
	err := db.Where("host_id = ? AND guest_id = ?", hostId, guestId).First(&followItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return true, err
	}

	return true, nil
}

func (*FollowingDao) QueryFollowerListByHostId(hostId uint) ([]*User, error) {
	var FollowerList []*Following
	err := db.Where("guest_id = ?", hostId).Find(&FollowerList).Error
	if err != nil {
		return nil, err
	}
	var UserList []*User
	var tempUser *User
	for _, follow := range FollowerList {
		tempUser = nil
		err := db.Where("id = ?", follow.HostID).Find(&tempUser).Error
		if err != nil {
			return nil, err
		}
		UserList = append(UserList, tempUser)
	}
	return UserList, nil
}

func (*FollowingDao) IncFollowerCnt(guestId uint) error {
	err := db.Model(&User{}).Where("id = ?", guestId).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) DecFollowerCnt(guestId uint) error {
	err := db.Model(&User{}).Where("id = ?", guestId).
		UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}
