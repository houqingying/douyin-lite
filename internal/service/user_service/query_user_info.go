package user_service

import (
	"douyin-lite/internal/entity"
	"errors"
)

func QueryUserInfo(userId int64) (*entity.UserInfo, error) {
	return NewQueryUserInfoFlow(userId).Do()
}

type QueryUserInfoFlow struct {
	userId   int64
	userInfo *entity.UserInfo
}

func NewQueryUserInfoFlow(userId int64) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userId: userId,
	}
}

func (f *QueryUserInfoFlow) Do() (*entity.UserInfo, error) {
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
	userInfo, err := QueryAUserInfo2(f.userId)
	if err != nil {
		return err
	}
	f.userInfo = userInfo
	return nil
}
