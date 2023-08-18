package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

// 定义likestructure
type Like struct {
	gorm.Model
	// Id    int64
	user_id  int64
	video_id int64
	status   int8 // 1-点赞，0-未点赞
	// CreatedAt time.Time
	// UpdatedAt time.Time
}

// IsVideoLikedByUser 获取视频点赞信息, 当前用户是否点赞
func IsVideoLikedByUser(userId int64, videoId int64) (int8, error) {
	var status int8
	result := db.Model(Like{}).Select("status").Where("user_id= ? and video_id= ?", userId, videoId).First(&status)
	c := result.RowsAffected
	if c == 0 {
		return -1, errors.New("current user haven not liked current video")
	}
	if result.Error != nil {
		//如果查询数据库失败，返回获取likeInfo信息失败
		log.Println(result.Error)
	}
	return status, nil
}
