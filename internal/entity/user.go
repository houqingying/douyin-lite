package entity

import (
	"douyin-lite/pkg/storage"
	"errors"
	"sync"

	conf "douyin-lite/configs/locales"
	"douyin-lite/pkg/snowflake"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              int64  `json:"id" gorm:"id,omitempty"`
	Name            string `json:"name" gorm:"comment:用户名"`
	Password        string `json:"password" gorm:"comment:用户密码"`
	Avatar          string `json:"avatar" gorm:"comment:用户头像"`
	BackgroundImage string `json:"background_image" gorm:"comment:用户背景主图"`
	Signature       string `json:"signature" gorm:"comment:用户签名"`
	TotalFavorited  int64  `json:"total_favorited" gorm:"comment:获赞总数"`
	WorkCount       int64  `json:"work_count" gorm:"comment:作品总数"`
	FavoriteCount   int64  `json:"favorite_count" gorm:"comment:点赞总数"`
}
type UserInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
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

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Avatar = "https://douyin-lite.oss-cn-hangzhou.aliyuncs.com/avater/default.jpg"
	user.BackgroundImage = "https://douyin-lite.oss-cn-hangzhou.aliyuncs.com/background_image/default.jpg"
	user.Signature = "你所热爱的，就是你的生活"
	return nil
}

func (*UserDao) CreateUser(name string) error {
	err := snowflake.InitSnowflakeNode(conf.Config.System.StartTime, int64(conf.Config.System.MachineID))
	if err != nil {
		return err
	}
	id := snowflake.GenerateID()
	newUser := User{
		Name: name,
		ID:   id,
	}
	err = storage.DB.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (*UserDao) CreateRegisterUser(name string, password string) (*User, error) {
	err := snowflake.InitSnowflakeNode(conf.Config.System.StartTime, int64(conf.Config.System.MachineID))
	if err != nil {
		return nil, err
	}

	newUser := User{
		Name:     name,
		Password: password,
		ID:       snowflake.GenerateID(),
	}

	err = storage.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (*UserDao) QueryIsUserExistByName(name string) (bool, error) {
	err := storage.DB.Where("name = ?", name).First(&User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, errors.New("数据库异常")
	}
	return true, nil
}

func (*UserDao) QueryIsUserExistById(userId int64) (bool, error) {
	err := storage.DB.Where("id = ?", userId).First(&User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.New("数据库异常")
	}
	return true, nil
}

func (*UserDao) QueryLoginUser(name string, password string) (*User, error) {
	qUser := User{}
	err := storage.DB.Where("name = ? and password = ?", name, password).First(&qUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("账号或者密码不对")
		}
		return nil, errors.New("数据库异常")
	}
	return &qUser, nil
}

func (*UserDao) QueryUserById(userId int64) (*User, error) {
	qUser := User{}
	err := storage.DB.Where("id = ?", userId).First(&qUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("该用户不存在")
		}
		return nil, errors.New("数据库异常")
	}
	return &qUser, nil
}

func (u *User) UpdateUserWorkCount(count int64) error {
	return storage.DB.Model(&User{}).Where("id = ?", u.ID).Update("work_count", count).Error
}
