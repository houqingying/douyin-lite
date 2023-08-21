package repository

import (
	"douyin-lite/service/favorite_service"
	"github.com/jinzhu/gorm"
)

// 查询当前用户点赞视频
func Query_Favorite_List(userId uint) ([]favorite_service.Video, error) {
	//查询当前用户点赞视频
	var favoriteList []favorite_service.Favorite
	videoList := make([]favorite_service.Video, 0)
	if err := db.Table("favorites").Where("user_id=? AND state=?", userId, 1).Find(&favoriteList).Error; err != nil { //找不到记录
		return videoList, nil
	}
	for _, m := range favoriteList {

		var video = favorite_service.Video{}
		if err := db.Table("videos").Where("id=?", m.VideoId).Find(&video).Error; err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}
	return videoList, nil
}

func Query_Check_Favorite(userId uint, videoId uint) (bool, error) {
	var total int64
	if err := db.Table("favorites").Where("user_id = ? AND video_id = ? AND state = 1", userId, videoId).Count(&total).Error; gorm.IsRecordNotFoundError(err) { //没有该条记录
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func AddFavoriteCount(HostId uint) error {
	if err := db.Model(User{}).
		Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count+?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func AddTotalFavorited(HostId uint) error {
	if err := db.Model(User{}).
		Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited+?", 1)).Error; err != nil {
		return err
	}
	return nil
}
func GetVideoAuthor(videoId uint) (uint, error) {
	var video favorite_service.Video
	if err := db.Table("videos").Where("id = ?", videoId).Find(&video).Error; err != nil {
		return video.ID, err
	}
	return video.AuthorId, nil
}

func ReduceFavoriteCount(HostId uint) error {
	if err := db.Model(User{}).
		Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count-?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func ReduceTotalFavorited(HostId uint) error {
	if err := db.Model(User{}).
		Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited-?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func IsFavoriteExist(userId uint, videoId uint) (bool, favorite_service.Favorite) {
	var favoriteExist = favorite_service.Favorite{} //找不到时会返回错误
	result := db.Table("favorites").Where("user_id = ? AND video_id = ?", userId, videoId).First(&favoriteExist)
	if result != nil {
		return false, favoriteExist
	}
	return true, favoriteExist
}

func CreatFavoriteAction(favoriteAction *favorite_service.Favorite) error {
	if err := db.Table("favorites").Create(&favoriteAction).Error; err != nil { //创建记录
		return err
	}
	return nil
}

func UpdateFavoriteCount(favoriteAction favorite_service.Favorite, count int8) {
	db.Table("videos").Where("id = ?", favoriteAction.VideoId).Update("favorite_count", gorm.Expr("favorite_count + {}", count))
}

func UpdateFavoriteState(favoriteAction favorite_service.Favorite, state int8) {
	db.Table("favorites").Where("video_id = ?", favoriteAction.VideoId).Update("state", state)
}
