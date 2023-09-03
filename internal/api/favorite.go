package api

import (
	favorite_service2 "douyin-lite/internal/service/favorite_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"net/http"
	"strconv"
)

type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Favorite(c *gin.Context) {
	// klog.Info("post relation action")
	// get guest_id
	user_Id, got := c.Get("user_id")
	if !got {
		klog.Errorf("user_id didn't set properly, something may be wrong with the jwt")
		c.JSON(http.StatusOK, SendMessageResp{
			Code: 403,
			Msg:  "user_id is invalid",
		})
		return
	}
	user_IdInt64, _ := user_Id.(int64)
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		fmt.Println("actionType=?", actionType)
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: -1,
			StatusMsg:  "操作无效",
		})
		return
	}
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: -1,
			StatusMsg:  "video_id is invalid",
		})
		return
	}

	err = favorite_service2.Favorite_Action(user_IdInt64, video_id, actionType)
	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteActionResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		})
	}
}

type FavoriteListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*favorite_service2.FavoriteListInfo
}

func FavoriteList(c *gin.Context) {
	user_Id, got := c.Get("user_id")
	if !got {
		klog.Errorf("user_id didn't set properly, something may be wrong with the jwt")
		c.JSON(http.StatusOK, SendMessageResp{
			Code: 403,
			Msg:  "user_id is invalid",
		})
		return
	}
	user_IdInt64, _ := user_Id.(int64)
	followListResp, err := QueryFavoriteList(user_IdInt64)
	if err != nil {
		c.JSON(http.StatusOK, followListResp)
		return
	}
	c.JSON(http.StatusOK, followListResp)
}

func QueryFavoriteList(hostInt64 int64) (*FavoriteListResp, error) {
	//hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	//if err != nil {
	//	return &FavoriteListResp{
	//		Code: "403",
	//		Msg:  "找不到用户",
	//	}, err
	//}
	favoriteListData, err := favorite_service2.QueryFavoriteListInfo(hostInt64)
	if err != nil {
		return &FavoriteListResp{
			Code: "-1",
			Msg:  "查找失败",
		}, err
	}
	return &FavoriteListResp{
		Code:             "200",
		Msg:              "查找成功",
		FavoriteListInfo: favoriteListData,
	}, nil
}
