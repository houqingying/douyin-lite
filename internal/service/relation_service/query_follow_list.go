package relation_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/repository"
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
	var userInfoList = make([]*user_service.UserInfo, len(idList))
	for i, id := range idList {
		user, err := entity.NewUserDaoInstance().QueryUserById(id)
		if err != nil {
			return err
		}
		followCnt, err := repository.QueryFollowCnt(id)
		if err != nil {
			return err
		}
		followerCnt, err := repository.QueryFollowerCnt(id)
		if err != nil {
			return err
		}
		isFollow, err := entity.NewFollowingDaoInstance().QueryisFollow(f.hostId, id)
		if err != nil {
			return err
		}
		userInfoList[i] = &user_service.UserInfo{
			ID:              user.ID,
			Name:            user.Name,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			FollowingCount:  *followCnt,
			FollowerCount:   *followerCnt,
			IsFollow:        isFollow,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}
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
