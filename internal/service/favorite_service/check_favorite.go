package favorite_service

import "douyin-lite/internal/entity"

// Check_Favorite 查询某用户是否点赞某视频
func Check_Favorite(uid uint, vid uint) bool {
	isFavorite, err := entity.NewFavoriteDaoInstance().Query_Check_Favorite(uid, vid)
	if err != nil {
		return false
	}
	return isFavorite
}
