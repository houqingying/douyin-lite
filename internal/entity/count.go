package entity

import (
	"douyin-lite/pkg/storage"
	"gorm.io/gorm"
	"sync"
)

type Count struct {
	gorm.Model
	ID       int64  `gorm:"primary_key;auto_increment:false"`
	CountKey string `gorm:"size:30;primary_key;auto_increment:false"`
	CountVal int64  `json:"count_val"`
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

func (*CountDao) QueryFollowingCount(id int64) (*int64, error) {
	countKey := "follow_count"
	var followCount int64
	err := storage.DB.Model(&Count{}).Select("count_val").
		Where("id = ? and count_key = ?", id, countKey).First(&followCount).Error
	if err != nil {
		return nil, err
	}
	return &followCount, nil
}

func (*CountDao) QueryFollowerCount(id int64) (*int64, error) {
	countKey := "follower_count"
	var followerCount int64
	err := storage.DB.Model(&Count{}).Select("count_val").
		Where("id = ? and count_key = ?", id, countKey).First(&followerCount).Error
	if err != nil {
		return nil, err
	}
	return &followerCount, nil
}

func (*CountDao) SaveFollowingCount(id int64, val int64) error {
	countKey := "follow_count"
	err := storage.DB.First(&Count{
		ID:       id,
		CountKey: countKey,
	}).Error
	var err2 error = nil
	if err != nil {
		err2 = storage.DB.Create(&Count{
			ID:       id,
			CountKey: countKey,
			CountVal: val,
		}).Error
	} else {
		err2 = storage.DB.Model(&Count{}).
			Where("id = ? and count_key = ?", id, countKey).Update("count_val", val).Error
	}
	if err2 != nil {
		return err2
	}
	return nil
}

// 使用 Save 方法保存数据

func (*CountDao) SaveFollowerCount(id int64, val int64) error {
	countKey := "follower_count"
	err := storage.DB.First(&Count{
		ID:       id,
		CountKey: countKey,
	}).Error
	var err2 error = nil
	if err != nil {
		err2 = storage.DB.Create(&Count{
			ID:       id,
			CountKey: countKey,
			CountVal: val,
		}).Error
	} else {
		err2 = storage.DB.Model(&Count{}).
			Where("id = ? and count_key = ?", id, countKey).Update("count_val", val).Error
	}
	if err2 != nil {
		return err2
	}
	return nil
}

//func (m *Count) BeforeSave(tx *gorm.DB) error {
//	if m.CreatedAt.IsZero() {
//		m.CreatedAt = time.Now()
//	}
//	if m.UpdatedAt.IsZero() {
//		m.UpdatedAt = time.Now()
//	}
//	return nil
//}
