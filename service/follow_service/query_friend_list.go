package follow_service

import (
	"douyin-lite/repository"
	"douyin-lite/service/user_service"
	"errors"
)

type FriendListInfo struct {
	UserInfoList []*user_service.UserInfo `json:"user_list"`
}

func QueryFriendListInfo(hostId uint) (*FriendListInfo, error) {
	return NewQueryFriendListInfoFlow(hostId).Do()
}

type QueryFriendListInfoFlow struct {
	hostId         uint
	friendListInfo *FriendListInfo

	userinfoList []*user_service.UserInfo
}

func NewQueryFriendListInfoFlow(hostId uint) *QueryFriendListInfoFlow {
	return &QueryFriendListInfoFlow{
		hostId: hostId,
	}
}

var FollowingDao = repository.NewFollowingDaoInstance()

func (f *QueryFriendListInfoFlow) Do() (*FriendListInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareFriendInfo()
	if err != nil {
		return nil, err
	}
	err = f.packFriendInfo()
	if err != nil {
		return nil, err
	}
	return f.friendListInfo, nil
}

func (f *QueryFriendListInfoFlow) checkParam() error {
	if f.hostId <= 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFriendListInfoFlow) prepareFriendInfo() error {
	friendList, err := FollowingDao.QueryFriendById(f.hostId)
	if err != nil {
		return errors.New("DB Find Friend Error")
	}
	var friendInfoList = make([]*user_service.UserInfo, len(friendList))
	for i, friend := range friendList {
		friendInfoList[i] = &user_service.UserInfo{
			ID:              friend.ID,
			Name:            friend.Name,
			Avatar:          friend.Avatar,
			BackgroundImage: friend.BackgroundImage,
			Signature:       friend.Signature,
			FollowingCount:  friend.FollowingCount,
			FollowerCount:   friend.FollowerCount,
			IsFollow:        true,
			TotalFavorited:  friend.TotalFavorited,
			WorkCount:       friend.WorkCount,
			FavoriteCount:   friend.FavoriteCount,
		}
	}
	f.userinfoList = friendInfoList
	//TODO: get recently message between host and guest
	return nil
}

func (f *QueryFriendListInfoFlow) packFriendInfo() error {
	f.friendListInfo = &FriendListInfo{
		UserInfoList: f.userinfoList,
	}
	return nil
}
