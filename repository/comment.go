package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoId   string `json:"video_id"`
	UserId    string `json:"user_id"`
	CommentId string `json:"comment_id"`
}

type CommentDao struct {
}

var commentDao *UserDao
var CommentOnce sync.Once

func NewCommentDaoInstance() *UserDao {
	UserOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (*UserDao) CreateComment(name string, followingCnt uint, followerCnt uint) error {
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
