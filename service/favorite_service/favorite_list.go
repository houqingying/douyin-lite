package favorite_service

import (
	"douyin-lite/repository"
)

// Favorite_List 获取点赞列表
func Favorite_List(userId uint) ([]Video, error) {
	videoList, err := repository.Query_Favorite_List(userId)
	return videoList, err
}
