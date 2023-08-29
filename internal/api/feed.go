package api

import (
	"douyin-lite/internal/service/video_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResp struct {
	Code int32  `json:"status_code"`
	Msg  string `json:"status_msg"`
	*video_service.FeedVideoListVO
}

func Feed(c *gin.Context) {
	rawTimestamp := c.Query("latest_time")
	var latestTime time.Time
	if rawTimestamp != "" {
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err == nil {
			latestTime = time.Unix(intTime/1000, 0)
		}
	} else {
		latestTime = time.Now()
	}
	//目前没有考虑用户登录的情况
	videoList, err := video_service.QueryFeedVideoList(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, FeedResp{
			Code: -1,
			Msg:  "获取视频流失败",
		})
		return
	}
	c.JSON(http.StatusOK, FeedResp{
		Code:            0,
		Msg:             "获取视频流成功",
		FeedVideoListVO: videoList,
	})

}
