package repository

import (
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name            string
	Avatar          string
	BackgroundImage string
	Signature       string
	FollowingCount  int64
	FollowerCount   int64
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
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

func (*UserDao) CreateUser(name string, followingCnt int64, followerCnt int64) error {
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
