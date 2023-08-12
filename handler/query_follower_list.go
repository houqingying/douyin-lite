package handler

import (
	"douyin-lite/service"
	"strconv"
)

type FollowerListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*service.FollowerListInfo
}

func QueryFollowerListHandler(hostIdStr string) (*FollowerListResp, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FollowerListResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	followListData, err := service.QueryFollowerListInfo(uint(hostId))
	if err != nil {
		return &FollowerListResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	return &FollowerListResp{
		Code:             "0",
		Msg:              "success",
		FollowerListInfo: followListData,
	}, nil
}
