package user_service

import (
	"douyin-lite/repository"
	"errors"
)

type LoginInfo struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

func LoginUser(userName string, password string) (*LoginInfo, error) {
	return NewLoginUserFlow(userName, password).Do()
}

type LoginUserFlow struct {
	userName      string
	password      string
	userId        uint
	token         string
	userLoginInfo *LoginInfo
}

func NewLoginUserFlow(userName string, password string) *LoginUserFlow {
	return &LoginUserFlow{
		userName: userName,
		password: password,
	}
}

func (f *LoginUserFlow) Do() (*LoginInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.queryLoginInfo()
	if err != nil {
		return nil, err
	}
	err = f.packLoginInfo()
	if err != nil {
		return nil, err
	}
	return f.userLoginInfo, nil
}

func (f *LoginUserFlow) checkParam() error {
	if f.userName == "" || len(f.userName) <= 0 {
		return errors.New("用户名不合法")
	}
	if len(f.userName) > 20 {
		return errors.New("用户名不合法")
	}
	if f.password == "" || len(f.password) <= 0 {
		return errors.New("密码不合法")
	}
	if len(f.password) > 20 {
		return errors.New("密码不合法")
	}
	return nil
}

func (f *LoginUserFlow) queryLoginInfo() error {
	loginUser, err := repository.NewUserDaoInstance().QueryLoginUser(f.userName, f.password)
	if err != nil {
		return err
	}
	f.userId = loginUser.ID
	f.token = "whatToken"
	return nil
}

func (f *LoginUserFlow) packLoginInfo() error {
	f.userLoginInfo = &LoginInfo{
		UserId: f.userId,
		Token:  f.token,
	}
	return nil
}
