package user_service

import (
	"douyin-lite/repository"
	"errors"
)

type RegisterInfo struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

func RegisterUser(userName string, userPassword string) (*RegisterInfo, error) {
	return NewRegisterUserFlow(userName, userPassword).Do()
}

type RegisterUserFlow struct {
	userName     string
	password     string
	userId       uint
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

func (f *RegisterUserFlow) updateRegisterInfo() error {
	isExist, _ := repository.NewUserDaoInstance().QueryIsUserExist(f.userName)
	if isExist {
		return errors.New("用户已经存在, 不需要再注册")
	}
	regInfo, err := repository.NewUserDaoInstance().CreateRegisterUser(f.userName, f.password)
	if err != nil {
		return err
	}
	f.userId = regInfo.ID
	f.token = "whatToken"
	return nil
}

func (f *RegisterUserFlow) packRegisterInfo() error {
	f.registerInfo = &RegisterInfo{
		UserId: f.userId,
		Token:  f.token,
	}
	return nil
}
