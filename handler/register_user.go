package handler

import "douyin-lite/service"

type RegisterResp struct {
	Code   string `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

func RegisterUserHandler(userName string, password string) (*RegisterResp, error) {
	registerInfo, err := service.RegisterUser(userName, password)
	if err != nil {
		return &RegisterResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	return &RegisterResp{
		Code:   "0",
		Msg:    "success",
		UserId: registerInfo.UserId,
		Token:  registerInfo.Token,
	}, nil
}
