package entity

import (
	"douyin-lite/pkg/storage"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

type Favorite struct {
	gorm.Model
	ID      int64 `json:"id" gorm:"id,omitempty"`
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	State   int32
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
func (*FavoriteDao) Query_Favorite_List(userId int64) ([]Video, error) {
	//查询当前用户点赞视频
	var favoriteList []Favorite
	videoList := make([]Video, 0)
	if err := storage.DB.Table("favorites").Where("user_id=? AND state=?", userId, 1).Find(&favoriteList).Error; err != nil { //找不到记录
		return videoList, nil
	}
	for _, m := range favoriteList {
		var video = Video{}
		if err := storage.DB.Table("video").Where("id=?", m.VideoId).Find(&video).Error; err != nil {
			return nil, err
		}
		videoList = append(videoList, video)
	}
	return videoList, nil
}

func (f *Favorite) GetFavoriteListResp(userId int64) (videos []*Favorite, err error) {
	err = storage.DB.Model(&f).
		Where("user_id = ?", userId).
		Where("state = ?", 1).
		Find(&videos).Error
	return
}

// 查看当前用户对已知视频是否点赞
func (*FavoriteDao) Query_Check_Favorite(userId int64, videoId int64) (bool, error) {
	var total int64
	if err := storage.DB.Table("favorite").Where("user_id = ? AND video_id = ? AND state = 1", userId, videoId).Count(&total).Error; gorm.IsRecordNotFoundError(err) { //没有该条记录
		return false, err
	}
	if total == 0 {
		return false, nil
	}
	return true, nil
}

func (*FavoriteDao) AddFavoriteCount(HostId int64) {
	/*fmt.Println("hostid=?", HostId)
	if err := storage.DB.Model(&User{}).Where("id=?", HostId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil*/
	result := storage.DB.Exec("UPDATE user SET favorite_count = favorite_count + 1 WHERE id = ?", HostId)

	if result.Error != nil {
		// 处理更新操作失败的情况
		fmt.Println("更新失败err", result.Error)
	} else {
		// 更新操作成功
		fmt.Println("更新ing")
	}
}

func (*FavoriteDao) AddTotalFavorited(HostId int64) {
	//if err := storage.DB.Model(&User{}).Where("id=?", HostId).
	//	Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error; err != nil {
	//	return err
	//}
	//return nil
	result := storage.DB.Exec("UPDATE user SET total_favorited = total_favorited + 1 WHERE id = ?", HostId)

	if result.Error != nil {
		// 处理更新操作失败的情况
		fmt.Println("更新失败err", result.Error)
	} else {
		// 更新操作成功
		fmt.Println("更新ing")
	}
}

func (*FavoriteDao) GetVideoAuthor(videoId int64) (int64, error) {
	var video Video
	if err := storage.DB.Table("video").Where("id = ?", videoId).Find(&video).Error; err != nil {
		return int64(video.ID), err
	}
	return video.AuthorId, nil
}

func (*FavoriteDao) ReduceFavoriteCount(HostId int64) {
	//if err := storage.DB.Model(&User{}).
	//	Where("id=?", HostId).
	//	Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
	//	return err
	//}
	//return nil

	query := "UPDATE user SET favorite_count = favorite_count - 1 WHERE id = ?"

	// 使用 Exec 方法执行更新操作
	result := storage.DB.Exec(query, HostId)

	if result.Error != nil {
		// 处理更新操作失败的情况
		fmt.Println("更新失败err", result.Error)
	} else {
		// 更新操作成功
		fmt.Println("更新ing")
	}
}

func (*FavoriteDao) ReduceTotalFavorited(HostId int64) {
	//if err := storage.DB.Table("user").
	//	Where("id = ?", HostId).
	//	Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error; err != nil {
	//	return err
	//}
	//return nil

	query := "UPDATE user SET total_favorited = total_favorited - 1 WHERE id = ?"

	// 使用 Exec 方法执行更新操作
	result := storage.DB.Exec(query, HostId)

	if result.Error != nil {
		// 处理更新操作失败的情况
		fmt.Println("更新失败err", result.Error)
	} else {
		// 更新操作成功
		fmt.Println("更新ing")
	}
}

func (*FavoriteDao) IsFavoriteExist(userId int64, videoId int64) (bool, Favorite) {
	var favoriteExist = Favorite{} //找不到时会返回错误
	result := storage.DB.Table("favorite").Where("user_id = ? AND video_id = ?", userId, videoId).First(&favoriteExist).Error
	if result != nil {
		fmt.Println("result=?", result)
		return false, favoriteExist
	}
	return true, favoriteExist
}

func (*FavoriteDao) CreateFavorite(userId int64, videoId int64) error {
	newFavorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
		State:   1,
	}
	err := storage.DB.Create(&newFavorite).Error
	fmt.Printf("创建成功")
	if err != nil {
		fmt.Println("创建失败")
		return err
	}
	return nil
}

func (*FavoriteDao) UpdateFavoriteCount(VideoId int64, count int64) {
	//result := storage.DB.Table("video").
	//	Where("id = ?", VideoId).
	//	Update("favorite_count", gorm.Expr("favorite_count + ?", count))
	result := storage.DB.Exec("UPDATE video SET favorite_count = favorite_count + ? WHERE id = ?", count, VideoId)

	if result.Error != nil {
		// 处理更新操作失败的情况
		fmt.Println("更新失败err", result.Error)
	} else {
		// 更新操作成功
		fmt.Println("更新ing")
	}

}

func (*FavoriteDao) UpdateFavoriteState(VideoId int64, state int32) {
	//return storage.DB.Table("favorite").Where("video_id = ?", VideoId).Update("state", state).Error
	result := storage.DB.Exec("UPDATE favorite SET state = ? WHERE video_id = ?", state, VideoId)
	if result.Error != nil {
		// 处理更新操作失败的情况
		fmt.Println("更新失败err", result.Error)
	} else {
		// 更新操作成功
		fmt.Println("更新ing")
	}
}
