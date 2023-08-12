package service

import (
	"douyin-lite/repository"
	"errors"
)

type UserInfo struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  uint   `json:"follow_count"`
	FollowerCount   uint   `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	TotalFavorited  uint   `json:"total_favorited"`
	WorkCount       uint   `json:"work_count"`
	FavoriteCount   uint   `json:"favorite_count"`
}

type FollowListInfo struct {
	UserInfoList []*UserInfo `json:"user_list"`
}

func QueryFollowListInfo(hostId uint) (*FollowListInfo, error) {
	return NewQueryFollowListInfoFlow(hostId).Do()
}

type QueryFollowInfoFlow struct {
	hostId         uint
	followListInfo *FollowListInfo

	userinfoList []*UserInfo
}

func NewQueryFollowListInfoFlow(hostId uint) *QueryFollowInfoFlow {
	return &QueryFollowInfoFlow{
		hostId: hostId,
	}
}

func (f *QueryFollowInfoFlow) Do() (*FollowListInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareFollowInfo()
	if err != nil {
		return nil, err
	}
	err = f.packFollowInfo()
	if err != nil {
		return nil, err
	}
	return f.followListInfo, nil
}

func (f *QueryFollowInfoFlow) checkParam() error {
	if f.hostId <= 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFollowInfoFlow) prepareFollowInfo() error {
	userList, err := repository.NewFollowingDaoInstance().QueryFollowingListByHostId(f.hostId)
	if err != nil {
		return errors.New("DB Find Following Error")
	}
	var userInfoList = make([]*UserInfo, len(userList))
	for i := 0; i < len(userList); i++ {
		userInfoList[i] = &UserInfo{}
		userInfoList[i].ID = userList[i].ID
		userInfoList[i].Name = userList[i].Name
		userInfoList[i].FollowingCount = userList[i].FollowingCount
		userInfoList[i].FollowerCount = userList[i].FollowerCount
		userInfoList[i].IsFollow = true
		userInfoList[i].Avatar = userList[i].Avatar
		userInfoList[i].BackgroundImage = userList[i].BackgroundImage
		userInfoList[i].Signature = userList[i].Signature
		userInfoList[i].TotalFavorited = userList[i].TotalFavorited
		userInfoList[i].WorkCount = userList[i].WorkCount
		userInfoList[i].FavoriteCount = userList[i].FavoriteCount
	}
	f.userinfoList = userInfoList
	return nil
}

func (f *QueryFollowInfoFlow) packFollowInfo() error {
	f.followListInfo = &FollowListInfo{
		UserInfoList: f.userinfoList,
	}
	return nil
}
