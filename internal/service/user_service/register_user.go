package user_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/middleware"
	"errors"
)

type RegisterInfo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func RegisterUser(userName string, userPassword string) (*RegisterInfo, error) {
	return NewRegisterUserFlow(userName, userPassword).Do()
}

type RegisterUserFlow struct {
	userName     string
	password     string
	userId       int64
	token        string
	registerInfo *RegisterInfo
}

func NewRegisterUserFlow(userName string, userPassword string) *RegisterUserFlow {
	return &RegisterUserFlow{
		userName: userName,
		password: userPassword,
	}
}

func (f *RegisterUserFlow) Do() (*RegisterInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.updateRegisterInfo()
	if err != nil {
		return nil, err
	}
	err = f.packRegisterInfo()
	if err != nil {
		return nil, err
	}
	return f.registerInfo, nil
}

func (f *RegisterUserFlow) checkParam() error {
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

func (f *RegisterUserFlow) updateRegisterInfo() error {
	isExist, _ := entity.NewUserDaoInstance().QueryIsUserExistByName(f.userName)
	if isExist {
		return errors.New("当前用户名已存在")
	}
	regInfo, err := entity.NewUserDaoInstance().CreateRegisterUser(f.userName, f.password)
	if err != nil {
		return err
	}
	f.userId = regInfo.ID
	token, err := middleware.ReleaseToken(regInfo.ID)
	if err != nil {
		return err
	}
	f.token = token
	return nil
}

func (f *RegisterUserFlow) packRegisterInfo() error {
	f.registerInfo = &RegisterInfo{
		UserId: f.userId,
		Token:  f.token,
	}
	return nil
}
