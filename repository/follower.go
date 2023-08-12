package repository

import "gorm.io/gorm"

func (*FollowingDao) QueryFollowerListByHostId(hostId uint) ([]*Following, error) {
	var FollowerList []*Following
	err := db.Where("guest_id = ?", hostId).Find(&FollowerList).Error
	if err != nil {
		return nil, err
	}
	return FollowerList, nil
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