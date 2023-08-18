package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"douyin-lite/service/follow_service"
)

type FriendList struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`

	FriendInfoList *follow_service.FriendListInfo `json:"friend_info_list,omitempty"`
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
	friendListData, err := follow_service.QueryFriendListInfo(uint(hostId))
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
