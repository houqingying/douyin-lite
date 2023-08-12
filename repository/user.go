package repository

import (
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  uint   `json:"follow_count"`
	FollowerCount   uint   `json:"follower_count"`
	TotalFavorited  uint   `json:"total_favorited"`
	WorkCount       uint   `json:"work_count"`
	FavoriteCount   uint   `json:"favorite_count"`
}

type UserDao struct {
}

var userDao *UserDao
var UserOnce sync.Once

func NewUserDaoInstance() *UserDao {
	UserOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) CreateUser(name string, followingCnt uint, followerCnt uint) error {
	newUser := User{
		Name:           name,
		FollowingCount: followingCnt,
		FollowerCount:  followerCnt,
	}
	err := db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}
