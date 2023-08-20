package favorite_service

import (
	"douyin-lite/repository"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `json:"name"`
	Password       string `json:"password"`
	FollowCount    uint   `json:"follow_count"`
	FollowerCount  uint   `json:"follower_count"`
	TotalFavorited uint   `json:"total_favorited"`
	FavoriteCount  uint   `json:"favorite_count"`
}

type Video struct {
	gorm.Model
	AuthorId      uint   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	Title         string `json:"title"`
}

type Favorite struct {
	gorm.Model
	UserId  uint `json:"user_id"`
	VideoId uint `json:"video_id"`
	State   uint
}

// Favorite_Action 点赞操作
func Favorite_Action(userId uint, videoId uint, actionType uint) (err error) {

	//1-点赞
	if actionType == 1 {
		favoriteAction := Favorite{
			UserId:  userId,
			VideoId: videoId,
			State:   1, //1-已点赞
		}
		//找不到时会返回错误
		//如果没有记录-Create，如果有了记录-修改State
		result := repository.IsFavoriteExist(userId, videoId)
		if result != nil { //不存在
			if err := repository.Db.Table("favorites").Create(&favoriteAction).Error; err != nil { //创建记录
				return err
			}
			repository.Db.Table("videos").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1"))
			//userId的favorite_count增加
			if err := repository.AddFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := repository.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := repository.AddTotalFavorited(GuestId); err != nil {
				return err
			}
		} else { //存在
			if result.State == 0 { //state为0-video的favorite_count加1
				repository.Db.Table("videos").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1"))
				repository.Db.Table("favorites").Where("video_id = ?", videoId).Update("state", 1)
				//userId的favorite_count增加
				if err := repository.AddFavoriteCount(userId); err != nil {
					return err
				}
				//videoId对应的userId的total_favorite增加
				GuestId, err := repository.GetVideoAuthor(videoId)
				if err != nil {
					return err
				}
				if err := repository.AddTotalFavorited(GuestId); err != nil {
					return err
				}
			}
			//state为1-video的favorite_count不变
			return nil
		}
	} else { //2-取消点赞
		var favoriteCancel = Favorite{}
		favoriteActionCancel := Favorite{
			UserId:  userId,
			VideoId: videoId,
			State:   0, //0-未点赞
		}
		if err := repository.Db.Table("favorites").Where("user_id = ? AND video_id = ?", userId, videoId).First(&favoriteCancel).Error; err != nil { //找不到这条记录，取消点赞失败，创建记录
			repository.Db.Table("favorites").Create(&favoriteActionCancel)
			//userId的favorite_count增加
			if err := repository.ReduceFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := repository.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := repository.ReduceTotalFavorited(GuestId); err != nil {
				return err
			}
			return err
		}
		//存在
		if favoriteCancel.State == 1 { //state为1-video的favorite_count减1
			repository.Db.Table("videos").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1"))
			//更新State
			repository.Db.Table("favorites").Where("video_id = ?", videoId).Update("state", 0)
			if err := repository.ReduceFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := repository.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := repository.ReduceTotalFavorited(GuestId); err != nil {
				return err
			}
			return err
		}
		//state为0-video的favorite_count不变
		return nil
	}
	return nil
}
