package repository

import (
	"errors"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name            string `json:"name"`
	Password        string `json:"password"`
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

func (*UserDao) CreateRegisterUser(name string, password string) (*User, error) {
	newUser := User{
		Name:     name,
		Password: password,
	}
	err := db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (*UserDao) QueryIsUserExist(name string) (bool, error) {
	err := db.Where("name = ?", name).First(&User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return true, err
	}
	return true, nil
}
