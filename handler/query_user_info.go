package handler

import (
	"douyin-lite/service/user_service"
	"strconv"
)

type QueryUserInfoResp struct {
	Code string                 `json:"status_code"`
	Msg  string                 `json:"status_msg"`
	User *user_service.UserInfo `json:"user"`
}

func QueryUserInfoHandler(userIdStr string) (*QueryUserInfoResp, error) {
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return &QueryUserInfoResp{
			Code: "-1",
			Msg:  err.Error(),
			User: nil,
		}, err
	}
	userInfo, err := user_service.QueryUserInfo(uint(userId))
	if err != nil {
		return &QueryUserInfoResp{
			Code: "-1",
			Msg:  err.Error(),
			User: nil,
		}, err
	}
	return &QueryUserInfoResp{
		Code: "0",
		Msg:  "success",
		User: userInfo,
	}, nil
}
