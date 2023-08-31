package favorite_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
	"errors"
)

type FavoriteListInfo struct {
	UserInfoList []*QueryFavoriteInfo `json:"user_list"`
}

func QueryFavoriteListInfo(hostId int64) (*FavoriteListInfo, error) {
	return NewQueryFavoriteListInfoFlow(hostId).Do()
}

type QueryFavoriteInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type QueryFavoriteListInfoFlow struct {
	hostId           int64
	favoriteListInfo *FavoriteListInfo
	favoriteinfoList []*QueryFavoriteInfo
}

func NewQueryFavoriteListInfoFlow(hostId int64) *QueryFavoriteListInfoFlow {
	return &QueryFavoriteListInfoFlow{
		hostId: hostId,
	}
}

func (f *QueryFavoriteListInfoFlow) Do() (*FavoriteListInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareFavoriteInfo()
	if err != nil {
		return nil, err
	}
	// err = f.packFavoriteInfo()
	// if err != nil {
	// 	return nil, err
	// }
	return f.favoriteListInfo, nil
}

func (f *QueryFavoriteListInfoFlow) checkParam() error {
	if f.hostId <= 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFavoriteListInfoFlow) prepareFavoriteInfo() error {
	videoList, err := entity.NewFavoriteDaoInstance().Query_Favorite_List(f.hostId)
	if err != nil {
		return errors.New("DB Find Favorite Error")
	}
	var AuthorInfoList = make([]*QueryFavoriteInfo, len(videoList))
	for i := 0; i < len(videoList); i++ {
		AuthorId := videoList[i].AuthorId
		authorInfo, _ := user_service.QueryAUserInfo2(AuthorId)
		AuthorInfoList = append(AuthorInfoList, (*QueryFavoriteInfo)(authorInfo))
	}
	f.favoriteinfoList = AuthorInfoList
	return nil
}

// func (f *QueryFavoriteListInfoFlow) packFavoriteInfo() error {
// 	f.favoriteListInfo = &FavoriteListInfo{
// 		UserInfoList: f.QueryFavoriteInfo,
// 	}
// 	return nil
// }
