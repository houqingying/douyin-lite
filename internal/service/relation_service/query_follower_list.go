package relation_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
	"errors"
)

type FollowerListInfo struct {
	UserInfoList []*user_service.UserInfo `json:"user_list"`
}

func QueryFollowerListInfo(hostId int64) (*FollowerListInfo, error) {
	return NewQueryFollowerListInfoFlow(hostId).Do()
}

type QueryFollowerInfoFlow struct {
	hostId           int64
	followerListInfo *FollowerListInfo
	userinfoList     []*user_service.UserInfo
}

func NewQueryFollowerListInfoFlow(hostId int64) *QueryFollowerInfoFlow {
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
	if f.hostId < 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFollowerInfoFlow) prepareFollowerInfo() error {
	idList, err := entity.NewFollowingDaoInstance().QueryFollowerIdList(f.hostId)
	if err != nil {
		return err
	}
	userInfoList, err := user_service.QueryUserInfoList(f.hostId, &idList)
	if err != nil {
		return errors.New("DB Find Follower Error")
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
