package api

import (
	follow_service2 "douyin-lite/internal/service/relation_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

type RelationActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func RelationActionHandler(c *gin.Context) {
	klog.Info("post relation action")
	// get guest_id
	guestId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		klog.Errorf("to_user_id strconv.ParseInt error: %v", err)
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: -1,
			StatusMsg:  "to_user_id is invalid",
		})
		return
	}
	// get action_type
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		klog.Errorf("action_type strconv.ParseInt error: %v", err)
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: -1,
			StatusMsg:  "action_type is invalid",
		})
		return
	}

	hostId := c.GetInt64("user_id")
	//hostId, err := strconv.ParseInt(c.Get("user_id"), 10, 64)
	//if err != nil {
	//	klog.Errorf("to_user_id strconv.ParseInt error: %v", err)
	//	c.JSON(http.StatusOK, RelationActionResp{
	//		StatusCode: -1,
	//		StatusMsg:  "user_id is invalid",
	//	})
	//	return
	//}

	err = follow_service2.FollowAction(uint(hostId), uint(guestId), uint(actionType))
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: 0,
			StatusMsg:  "success",
		})
	}
}

type FollowListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*follow_service2.FollowListInfo
}

func QueryFollowListHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")
	//tokenStr := c.Param("token")
	followListResp, err := QueryFollowList(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, followListResp)
		return
	}
	c.JSON(http.StatusOK, followListResp)
}

func QueryFollowList(hostIdStr string) (*FollowListResp, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FollowListResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	followListData, err := follow_service2.QueryFollowListInfo(int64(uint(hostId)))
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

type FollowerListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*follow_service2.FollowerListInfo
}

func QueryFollowerListHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")
	//tokenStr := c.Param("token")
	followListResp, err := QueryFollowerList(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, followListResp)
		return
	}
	c.JSON(http.StatusOK, followListResp)
}

func QueryFollowerList(hostIdStr string) (*FollowerListResp, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FollowerListResp{
			Code: "-1",
			Msg:  err.Error(),
		}, err
	}
	followListData, err := follow_service2.QueryFollowerListInfo(int64(uint(hostId)))
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

type FriendList struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`

	FriendInfoList *follow_service2.FriendListInfo `json:"friend_info_list,omitempty"`
}

func QueryFriendListHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")
	friendListResp, err := QueryFriendList(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, friendListResp)
		return
	}
	c.JSON(http.StatusOK, friendListResp)
}

func QueryFriendList(hostIdStr string) (*FriendList, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FriendList{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}, err
	}
	friendListData, err := follow_service2.QueryFriendListInfo(uint(hostId))
	if err != nil {
		return &FriendList{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}, err
	}
	return &FriendList{
		StatusCode:     0,
		StatusMsg:      "success",
		FriendInfoList: friendListData,
	}, nil
}
