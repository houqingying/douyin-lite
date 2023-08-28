package entity

import (
	"douyin-lite/pkg/storage"
	"gorm.io/gorm"
	"sync"
)

type Count struct {
	gorm.Model
	ID            int64 `json:"id" gorm:"id,omitempty"`
	FollowCount   int64 `json:"follow_cnt" gorm:"comment: 关注总数"`
	FollowerCount int64 `json:"follower_cnt" gorm:"comment:粉丝总数"`
}

type CountDao struct {
}

var countDao *CountDao
var CountOnce sync.Once

func NewCountDaoInstance() *CountDao {
	CountOnce.Do(func() {
		countDao = &CountDao{}
	})
	return countDao
}

func (*CountDao) CreateCount(id int64) (*Count, error) {
	newCount := Count{
		Model:         gorm.Model{},
		ID:            id,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := storage.DB.Create(&newCount).Error
	if err != nil {
		return nil, err
	}
	return &newCount, nil
}

func (*CountDao) QueryFollowingCount(id int64) (*int64, error) {
	var followCount int64
	err := storage.DB.Model(&Count{}).Select("follow_cnt").Where("id = ?", id).First(&followCount).Error
	if err != nil {
		return nil, err
	}
	return &followCount, nil
}

func (*CountDao) QueryFollowerCount(id int64) (*int64, error) {
	var followerCount int64
	err := storage.DB.Model(&Count{}).Select("follower_cnt").Where("id = ?", id).First(&followerCount).Error
	if err != nil {
		return nil, err
	}
	return &followerCount, nil
}
