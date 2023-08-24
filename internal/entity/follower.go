package entity

import (
	"douyin-lite/pkg/storage"
	"gorm.io/gorm"
)

func (*FollowingDao) QueryFollowerListByHostId(hostId uint) ([]*User, error) {
	var FollowerList []*Following
	err := storage.DB.Where("guest_id = ?", hostId).Find(&FollowerList).Error
	if err != nil {
		return nil, err
	}
	var UserList []*User
	var tempUser *User
	for _, follow := range FollowerList {
		tempUser = nil
		err := storage.DB.Where("id = ?", follow.HostId).Find(&tempUser).Error
		if err != nil {
			return nil, err
		}
		UserList = append(UserList, tempUser)
	}
	return UserList, nil
}

func (*FollowingDao) IncFollowerCnt(guestId uint) error {
	err := storage.DB.Model(&User{}).Where("id = ?", guestId).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) DecFollowerCnt(guestId uint) error {
	err := storage.DB.Model(&User{}).Where("id = ?", guestId).
		UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FollowingDao) QueryFriendById(hostId uint) ([]*User, error) {
	var FriendList []*Following
	err := storage.DB.Raw("SELECT * FROM following WHERE host_id = ? AND guest_id IN (SELECT host_id FROM following f WHERE f.guest_id = following.host_id)", hostId).Scan(&FriendList).Error
	if err != nil {
		return nil, err
	}
	var UserList []*User
	var tempUser *User
	for _, follow := range FriendList {
		tempUser = nil
		err := storage.DB.Where("id = ?", follow.GuestId).Find(&tempUser).Error
		if err != nil {
			return nil, err
		}
		UserList = append(UserList, tempUser)
	}
	return UserList, nil
}
