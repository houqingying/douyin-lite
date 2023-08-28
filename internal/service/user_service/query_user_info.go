package user_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/repository"
	"errors"
)

type UserInfo struct {
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

func QueryUserInfo(userId int64) (*UserInfo, error) {
	return NewQueryUserInfoFlow(userId).Do()
}

type QueryUserInfoFlow struct {
	userId   int64
	userInfo *UserInfo
}

func NewQueryUserInfoFlow(userId int64) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userId: userId,
	}
}

func (f *QueryUserInfoFlow) Do() (*UserInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareUserInfo()
	if err != nil {
		return nil, err
	}
	return f.userInfo, nil
}

func (f *QueryUserInfoFlow) checkParam() error {
	if f.userId < 0 {
		return errors.New("user id should be larger than 0")
	}
	return nil
}

func (f *QueryUserInfoFlow) prepareUserInfo() error {
	qUser, err := entity.NewUserDaoInstance().QueryUserById(f.userId)
	if err != nil {
		return err
	}
	followCnt, err := repository.QueryFollowCnt(f.userId)
	if err != nil {
		return err
	}
	followerCnt, err := repository.QueryFollowerCnt(f.userId)
	if err != nil {
		return err
	}
	newUserInfo := UserInfo{}
	newUserInfo.ID = qUser.ID
	newUserInfo.Name = qUser.Name
	newUserInfo.IsFollow = true
	newUserInfo.Avatar = qUser.Avatar
	newUserInfo.BackgroundImage = qUser.BackgroundImage
	newUserInfo.Signature = qUser.Signature
	newUserInfo.TotalFavorited = qUser.TotalFavorited
	newUserInfo.WorkCount = qUser.WorkCount
	newUserInfo.FavoriteCount = qUser.FavoriteCount
	newUserInfo.FollowingCount = *followCnt
	newUserInfo.FollowerCount = *followerCnt
	f.userInfo = &newUserInfo
	return nil
}

func QueryUserList(hostId int64, idList *[]int64) ([]*UserInfo, error) {
	var userInfoList = make([]*UserInfo, len(*idList))
	for i, id := range *idList {
		user, err := entity.NewUserDaoInstance().QueryUserById(id)
		if err != nil {
			return nil, err
		}
		followCnt, err := repository.QueryFollowCnt(id)
		if err != nil {
			return nil, err
		}
		followerCnt, err := repository.QueryFollowerCnt(id)
		if err != nil {
			return nil, err
		}
		isFollow, err := entity.NewFollowingDaoInstance().QueryisFollow(hostId, id)
		if err != nil {
			return nil, err
		}
		userInfoList[i] = &UserInfo{
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
	return userInfoList, nil
}
