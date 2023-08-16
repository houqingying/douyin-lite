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
		return false, errors.New("数据库异常")
	}
	return true, nil
}

func (*UserDao) QueryLoginUser(name string, password string) (*User, error) {
	qUser := User{}
	err := db.Where("name = ? and password = ?", name, password).First(&qUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("账号或者密码不对")
		}
		return nil, errors.New("数据库异常")
	}
	return &qUser, nil
}

func (*UserDao) QueryUserById(userId uint) (*User, error) {
	qUser := User{}
	err := db.Where("id = ?", userId).First(&qUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("该用户不存在")
		}
		return nil, errors.New("数据库异常")
	}
	return &qUser, nil
}
