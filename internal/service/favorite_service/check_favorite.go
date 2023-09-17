package favorite_service

import "douyin-lite/internal/entity"

// CheckFavorite 查询某用户是否点赞某视频
func CheckFavorite(uid int64, vid int64) bool {
	isFavorite, err := entity.NewFavoriteDaoInstance().QueryCheckFavorite(uid, vid)
	if err != nil {
		return false
	}
	return isFavorite
}
