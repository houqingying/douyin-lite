package api

import (
	"douyin-lite/internal/service/video_service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
	id, _ := c.Get("user_id")
	userId, _ := id.(int64)
	videoList, err := video_service.QueryFeedVideoList(userId, latestTime)
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
