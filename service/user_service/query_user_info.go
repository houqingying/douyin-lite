package user_service

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

func QueryUserInfo(userId uint) (*UserInfo, error) {
	return NewQueryUserInfoFlow(userId).Do()
}

type QueryUserInfoFlow struct {
	userId   uint
	userInfo *UserInfo
}

func NewQueryUserInfoFlow(userId uint) *QueryUserInfoFlow {
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
	if f.userId <= 0 {
		return errors.New("user id should be larger than 0")
	}
	return nil
}

func (f *QueryUserInfoFlow) prepareUserInfo() error {
	qUser, err := repository.NewUserDaoInstance().QueryUserById(f.userId)
	if err != nil {
		return err
	}
	newUserInfo := UserInfo{}
	newUserInfo.ID = qUser.ID
	newUserInfo.Name = qUser.Name
	newUserInfo.FollowingCount = qUser.FollowingCount
	newUserInfo.FollowerCount = qUser.FollowerCount
	newUserInfo.IsFollow = true
	newUserInfo.Avatar = qUser.Avatar
	newUserInfo.BackgroundImage = qUser.BackgroundImage
	newUserInfo.Signature = qUser.Signature
	newUserInfo.TotalFavorited = qUser.TotalFavorited
	newUserInfo.WorkCount = qUser.WorkCount
	newUserInfo.FavoriteCount = qUser.FavoriteCount
	f.userInfo = &newUserInfo
	return nil
}