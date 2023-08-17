package user

import (
	"douyin-lite/service/user_service"
)

type LoginResp struct {
	Code   string `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

func LoginUserHandler(userName string, password string) (*LoginResp, error) {
	loginInfo, err := user_service.LoginUser(userName, password)
	if err != nil {
		return &LoginResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	return &LoginResp{
		Code:   "0",
		Msg:    "success",
		UserId: loginInfo.UserId,
		Token:  loginInfo.Token,
	}, nil
}
