package favorite_service

import (
	"douyin-lite/internal/entity"
)

// Favorite_Action 点赞操作
func Favorite_Action(userId uint, videoId uint, actionType uint) (err error) {
	//1-点赞
	if actionType == 1 {
		favoriteAction := entity.NewFavoriteDaoInstance()
		//找不到时会返回错误
		//如果没有记录-Create，如果有了记录-修改State
		result, favoriteExist := favoriteAction.IsFavoriteExist(userId, videoId)
		if !result { //不存在
			favoriteAction.CreateFavorite(userId, videoId)
			favoriteAction.UpdateFavoriteCount(videoId, 1)
			//userId的favorite_count增加
			if err := favoriteAction.AddFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := favoriteAction.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := favoriteAction.AddTotalFavorited(GuestId); err != nil {
				return err
			}
		} else { //存在
			if favoriteExist.State == 0 { //state为0-video的favorite_count加1
				favoriteAction.UpdateFavoriteCount(videoId, 1)
				favoriteAction.UpdateFavoriteState(videoId, 1)
				//userId的favorite_count增加
				if err := favoriteAction.AddFavoriteCount(userId); err != nil {
					return err
				}
				//videoId对应的userId的total_favorite增加
				GuestId, err := favoriteAction.GetVideoAuthor(videoId)
				if err != nil {
					return err
				}
				if err := favoriteAction.AddTotalFavorited(GuestId); err != nil {
					return err
				}
			}
			//state为1-video的favorite_count不变
			return nil
		}
	} else { //2-取消点赞
		favoriteCancel := entity.NewFavoriteDaoInstance()
		result, favoriteExist := favoriteCancel.IsFavoriteExist(userId, videoId)
		if !result { //找不到这条记录，取消点赞失败，创建记录
			favoriteCancel.CreateFavorite(userId, videoId)
			//userId的favorite_count增加
			if err := favoriteCancel.ReduceFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := favoriteCancel.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := favoriteCancel.ReduceTotalFavorited(GuestId); err != nil {
				return err
			}
			return err
		}
		//存在
		if favoriteExist.State == 1 { //state为1-video的favorite_count减1
			favoriteCancel.UpdateFavoriteCount(videoId, -1)
			//更新State
			favoriteCancel.UpdateFavoriteState(videoId, 0)
			if err := favoriteCancel.ReduceFavoriteCount(userId); err != nil {
				return err
			}
			//videoId对应的userId的total_favorite增加
			GuestId, err := favoriteCancel.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			if err := favoriteCancel.ReduceTotalFavorited(GuestId); err != nil {
				return err
			}
			return err
		}
		//state为0-video的favorite_count不变
		return nil
	}
	return nil
}
