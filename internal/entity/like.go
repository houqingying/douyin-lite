package entity

import (
	"douyin-lite/pkg/storage"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

type Favorite struct {
	gorm.Model
	UserId  uint `json:"user_id"`
	VideoId uint `json:"video_id"`
	State   uint
}

func (Favorite) TableName() string {
	return "favorite"
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(func() {
		favoriteDao = &FavoriteDao{}
	})
	return favoriteDao
}

// 查询当前用户点赞视频
func (*FavoriteDao) Query_Favorite_List(userId uint) ([]Video, error) {
	//查询当前用户点赞视频
	var favoriteList []Favorite
	videoList := make([]Video, 0)
	if err := storage.DB.Table("favorites").Where("user_id=? AND state=?", userId, 1).Find(&favoriteList).Error; err != nil { //找不到记录
		return videoList, nil
	}
	for _, m := range favoriteList {
		var video = Video{}
		if err := storage.DB.Table("videos").Where("id=?", m.VideoId).Find(&video).Error; err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}
	return videoList, nil
}

// 查看当前用户对已知视频是否点赞
func (*FavoriteDao) Query_Check_Favorite(userId uint, videoId uint) (bool, error) {
	var total int64
	if err := storage.DB.Table("favorites").Where("user_id = ? AND video_id = ? AND state = 1", userId, videoId).Count(&total).Error; gorm.IsRecordNotFoundError(err) { //没有该条记录
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func (*FavoriteDao) AddFavoriteCount(HostId uint) error {
	if err := storage.DB.Model(&User{}).Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) AddTotalFavorited(HostId uint) error {
	if err := storage.DB.Model(&User{}).Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
func (*FavoriteDao) GetVideoAuthor(videoId uint) (uint, error) {
	var video Video
	if err := storage.DB.Table("videos").Where("id = ?", videoId).Find(&video).Error; err != nil {
		return video.ID, err
	}
	return video.AuthorId, nil
}

func (*FavoriteDao) ReduceFavoriteCount(HostId uint) error {
	if err := storage.DB.Model(&User{}).
		Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) ReduceTotalFavorited(HostId uint) error {
	if err := storage.DB.Model(&User{}).
		Where("id=?", HostId).
		Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) IsFavoriteExist(userId uint, videoId uint) (bool, Favorite) {
	var favoriteExist = Favorite{} //找不到时会返回错误
	result := storage.DB.Table("favorites").Where("user_id = ? AND video_id = ?", userId, videoId).First(&favoriteExist)
	if result != nil {
		return false, favoriteExist
	}
	return true, favoriteExist
}

func (*FavoriteDao) CreateFavorite(userId uint, videoId uint) error {
	newFavorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
		State:   1,
	}
	err := storage.DB.Create(&newFavorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (*FavoriteDao) UpdateFavoriteCount(VideoId uint, count int8) {
	storage.DB.Table("videos").Where("id = ?", VideoId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", count))
}

func (*FavoriteDao) UpdateFavoriteState(VideoId uint, state int8) {
	storage.DB.Table("favorites").Where("video_id = ?", VideoId).Update("state", state)
}

// UpdateFavorite 根据userId，videoId,actionType点赞或者取消赞
func UpdateFavorite(userId uint, videoId uint, state uint) error {
	//更新当前用户观看视频的点赞状态“cancel”，返回错误结果
	err := storage.DB.Model(&Favorite{}).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		Update("State", state).Error
	//如果出现错误，返回更新数据库失败
	if err != nil {
		log.Println(err.Error())
		return errors.New("update data fail")
	}
	//更新操作成功
	return nil
}
