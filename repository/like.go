package repository

import (
	"errors"
	"log"
)

// IsVideoLikedByUser 获取视频点赞信息, 当前用户是否点赞
func IsVideoLikedByUser(userId int64, videoId int64) (int8, error) {
	var isLiked int8
	result := db.Model(Like{}).Select("liked").Where("user_id= ? and video_id= ?", userId, videoId).First(&isLiked)
	c := result.RowsAffected
	if c == 0 {
		return -1, errors.New("current user haven not liked current video")
	}
	if result.Error != nil {
		//如果查询数据库失败，返回获取likeInfo信息失败
		log.Println(result.Error)
	}
	return isLiked, nil
}
