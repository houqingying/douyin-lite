package favorite_service

import (
	"douyin-lite/internal/entity"
	"fmt"
)

// FavoriteAction 点赞操作
func FavoriteAction(userId int64, videoId int64, actionType int64) (err error) {
	//1-点赞
	if actionType == 1 {
		fmt.Println("点赞")
		favoriteAction := entity.NewFavoriteDaoInstance()
		//找不到时会返回错误
		//如果没有记录-Create，如果有了记录-修改State
		result, favoriteExist := favoriteAction.IsFavoriteExist(userId, videoId)
		if !result { //不存在
			fmt.Println("记录不存在")
			err := favoriteAction.CreateFavorite(userId, videoId)
			if err != nil {
				return err
			}
			favoriteAction.UpdateFavoriteCount(videoId, 1)
			//if err != nil {
			//	fmt.Println("更新出错，err=?", err)
			//	return err
			//}
			//userId的favorite_count增加
			favoriteAction.AddFavoriteCount(userId)
			//videoId对应的userId的total_favorite增加
			GuestId, err := favoriteAction.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			favoriteAction.AddTotalFavorited(GuestId)
		} else { //存在
			fmt.Println("记录存在")
			if favoriteExist.State == 0 { //state为0-video的favorite_count加1
				favoriteAction.UpdateFavoriteCount(videoId, 1)
				favoriteAction.UpdateFavoriteState(videoId, 1)
				//userId的favorite_count增加
				favoriteAction.AddFavoriteCount(userId)
				//videoId对应的userId的total_favorite增加
				GuestId, err := favoriteAction.GetVideoAuthor(videoId)
				if err != nil {
					return err
				}
				favoriteAction.AddTotalFavorited(GuestId)
			}
			//state为1-video的favorite_count不变
			return nil
		}
	} else { //2-取消点赞
		fmt.Printf("取消点赞")
		favoriteCancel := entity.NewFavoriteDaoInstance()
		result, favoriteExist := favoriteCancel.IsFavoriteExist(userId, videoId)
		if !result { //找不到这条记录，取消点赞失败，创建记录
			err := favoriteCancel.CreateFavorite(userId, videoId)
			if err != nil {
				return err
			}
			//userId的favorite_count增加
			favoriteCancel.ReduceFavoriteCount(userId)
			//videoId对应的userId的total_favorite增加
			GuestId, err := favoriteCancel.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			favoriteCancel.ReduceTotalFavorited(GuestId)
			return err
		}
		//存在
		if favoriteExist.State == 1 { //state为1-video的favorite_count减1
			favoriteCancel.UpdateFavoriteCount(videoId, -1)
			//更新State
			favoriteCancel.UpdateFavoriteState(videoId, 0)
			favoriteCancel.ReduceFavoriteCount(userId)
			//videoId对应的userId的total_favorite增加
			GuestId, err := favoriteCancel.GetVideoAuthor(videoId)
			if err != nil {
				return err
			}
			favoriteCancel.ReduceTotalFavorited(GuestId)
			return err
		}
		//state为0-video的favorite_count不变
		return nil
	}
	return nil
}
