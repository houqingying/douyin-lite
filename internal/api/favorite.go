package api

import (
	favorite_service2 "douyin-lite/internal/service/favorite_service"
	"github.com/gin-gonic/gin"
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
	user_Id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: 403,
			StatusMsg:  "用户不存在",
		})
		return
	}
	// get action_type
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 0) {
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

	err = favorite_service2.Favorite_Action(uint(user_Id), uint(video_id), uint(actionType))
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
	userIdStr := c.Query("user_id")
	followListResp, err := QueryFavoriteList(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, followListResp)
		return
	}
	c.JSON(http.StatusOK, followListResp)
}

func QueryFavoriteList(hostIdStr string) (*FavoriteListResp, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FavoriteListResp{
			Code: "403",
			Msg:  "找不到用户",
		}, err
	}
	favoriteListData, err := favorite_service2.QueryFavoriteListInfo(uint(hostId))
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
