package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"douyin-lite/service/follow_service"
)

type FollowerListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*follow_service.FollowerListInfo
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
	followListData, err := follow_service.QueryFollowerListInfo(uint(hostId))
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
