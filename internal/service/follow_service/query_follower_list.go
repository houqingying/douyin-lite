package follow_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
	"errors"
)

type FollowerListInfo struct {
	UserInfoList []*user_service.UserInfo `json:"user_list"`
}

func QueryFollowerListInfo(hostId uint) (*FollowerListInfo, error) {
	return NewQueryFollowerListInfoFlow(hostId).Do()
}

type QueryFollowerInfoFlow struct {
	hostId           uint
	followerListInfo *FollowerListInfo
	userinfoList     []*user_service.UserInfo
}

func NewQueryFollowerListInfoFlow(hostId uint) *QueryFollowerInfoFlow {
	return &QueryFollowerInfoFlow{
		hostId: hostId,
	}
}

func (f *QueryFollowerInfoFlow) Do() (*FollowerListInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareFollowerInfo()
	if err != nil {
		return nil, err
	}
	err = f.packFollowerInfo()
	if err != nil {
		return nil, err
	}
	return f.followerListInfo, nil
}

func (f *QueryFollowerInfoFlow) checkParam() error {
	if f.hostId <= 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFollowerInfoFlow) prepareFollowerInfo() error {
	userList, err := entity.NewFollowingDaoInstance().QueryFollowerListByHostId(f.hostId)
	if err != nil {
		return errors.New("DB Find Follower Error")
	}
	var userInfoList = make([]*user_service.UserInfo, len(userList))
	for i := 0; i < len(userList); i++ {
		userInfoList[i] = &user_service.UserInfo{}
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

func (f *QueryFollowerInfoFlow) packFollowerInfo() error {
	f.followerListInfo = &FollowerListInfo{
		UserInfoList: f.userinfoList,
	}
	return nil
}
