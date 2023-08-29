package favorite_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
	"errors"
)

type FavoriteListInfo struct {
	UserInfoList []*user_service.UserInfo `json:"user_list"`
}

func QueryFavoriteListInfo(hostId int64) (*FavoriteListInfo, error) {
	return NewQueryFavoriteListInfoFlow(hostId).Do()
}

type QueryFavoriteInfoFlow struct {
	hostId           int64
	favoriteListInfo *FavoriteListInfo
	userinfoList     []*user_service.UserInfo
}

func NewQueryFavoriteListInfoFlow(hostId int64) *QueryFavoriteInfoFlow {
	return &QueryFavoriteInfoFlow{
		hostId: hostId,
	}
}

func (f *QueryFavoriteInfoFlow) Do() (*FavoriteListInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareFavoriteInfo()
	if err != nil {
		return nil, err
	}
	err = f.packFavoriteInfo()
	if err != nil {
		return nil, err
	}
	return f.favoriteListInfo, nil
}

func (f *QueryFavoriteInfoFlow) checkParam() error {
	if f.hostId <= 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFavoriteInfoFlow) prepareFavoriteInfo() error {
	userList, err := entity.NewFavoriteDaoInstance().Query_Favorite_List(f.hostId)
	if err != nil {
		return errors.New("DB Find Favorite Error")
	}
	var userInfoList = make([]*user_service.UserInfo, len(userList))
	for i := 0; i < len(userList); i++ {
		userInfoList[i] = &user_service.UserInfo{}
		userInfoList[i].ID = int64(userList[i].ID)
		userInfoList[i].Name = userList[i].Title
		userInfoList[i].FavoriteCount = int64(userList[i].FavoriteCount)
	}
	f.userinfoList = userInfoList
	return nil
}

func (f *QueryFavoriteInfoFlow) packFavoriteInfo() error {
	f.favoriteListInfo = &FavoriteListInfo{
		UserInfoList: f.userinfoList,
	}
	return nil
}

// Favorite_List 获取点赞列表
//func Favorite_List(userId int64) ([]repository.Video, error) {
//	videoList, err := repository.NewFavoriteDaoInstance().Query_Favorite_List(userId)
//	return videoList, err
//}
