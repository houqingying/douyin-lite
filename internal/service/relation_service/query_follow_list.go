package relation_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
	"errors"
)

type FollowListInfo struct {
	UserInfoList []*user_service.UserInfo `json:"user_list"`
}

func QueryFollowListInfo(hostId int64) (*FollowListInfo, error) {
	return NewQueryFollowListInfoFlow(hostId).Do()
}

type QueryFollowInfoFlow struct {
	hostId         int64
	followListInfo *FollowListInfo

	userinfoList []*user_service.UserInfo
}

func NewQueryFollowListInfoFlow(hostId int64) *QueryFollowInfoFlow {
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
	if f.hostId < 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFollowInfoFlow) prepareFollowInfo() error {
	idList, err := entity.NewFollowingDaoInstance().QueryFollowingIdList(f.hostId)
	if err != nil {
		return err
	}
	userInfoList, err := user_service.QueryUserInfoList(f.hostId, &idList)
	if err != nil {
		return errors.New("DB Find Following Error")
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
