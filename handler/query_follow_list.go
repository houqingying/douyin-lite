package handler

import (
	"douyin-lite/service/follow_service"
	"strconv"
)

type FollowListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*follow_service.FollowListInfo
}

func QueryFollowListHandler(hostIdStr string) (*FollowListResp, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FollowListResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	followListData, err := follow_service.QueryFollowListInfo(uint(hostId))
	if err != nil {
		return &FollowListResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	return &FollowListResp{
		Code:           "0",
		Msg:            "success",
		FollowListInfo: followListData,
	}, nil
}
