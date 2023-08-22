package user_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/middleware"
	"errors"
)

type LoginInfo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func LoginUser(userName string, password string) (*LoginInfo, error) {
	return NewLoginUserFlow(userName, password).Do()
}

type LoginUserFlow struct {
	userName      string
	password      string
	userId        int64
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
		return errors.New("用户名不能为空")
	}
	if len(f.userName) > 20 {
		return errors.New("用户名太长，请不要超过二十位字符")
	}
	if f.password == "" || len(f.password) <= 0 {
		return errors.New("密码不能为空")
	}
	return nil
}

func (f *LoginUserFlow) queryLoginInfo() error {
	loginUser, err := entity.NewUserDaoInstance().QueryLoginUser(f.userName, f.password)
	if err != nil {
		return err
	}
	f.userId = loginUser.ID
	token, err := middleware.ReleaseToken(loginUser.ID)
	if err != nil {
		return err
	}
	f.token = token
	return nil
}

func (f *LoginUserFlow) packLoginInfo() error {
	f.userLoginInfo = &LoginInfo{
		UserId: f.userId,
		Token:  f.token,
	}
	return nil
}
