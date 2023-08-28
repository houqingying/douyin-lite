package relation_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/repository"
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

func (f *QueryFollowerInfoFlow) packFollowerInfo() error {
	f.followerListInfo = &FollowerListInfo{
		UserInfoList: f.userinfoList,
	}
	return nil
}
