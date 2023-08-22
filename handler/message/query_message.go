package message

import (
	"douyin-lite/service/message_service"
	"net/http"
	"strconv"

	"k8s.io/klog"

	"github.com/gin-gonic/gin"
)

type QueryMessageResp struct {
	Code        int32                          `json:"status_code"`
	Msg         string                         `json:"status_msg"`
	MessageList []*message_service.MessageInfo `json:"message_list"`
}

func QueryMessageHandler(c *gin.Context) {
	fromUserId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		klog.Errorf("user_id strconv.ParseInt error: %v", err)
		c.JSON(http.StatusOK, QueryMessageResp{
			Code: 403,
			Msg:  "user_id is invalid",
		})
		return
	}

	// 读出其他request参数并检查合法性
	toUserIdStr := c.Query("to_user_id")

	// 转换toUserId为int64，判断是否小于0；若不小于0，才可在之后流程转换为uint
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil || toUserId < 0 {
		c.JSON(http.StatusOK, QueryMessageResp{400, "to_user_id参数错误", nil})
		return
	}

	// 调用Service层，完成查找
	messageInfoList, err := message_service.QueryMessage(uint(fromUserId), uint(toUserId))
	if err != nil {
		c.JSON(http.StatusOK, QueryMessageResp{500, err.Error(), nil})
		return
	}

	c.JSON(http.StatusOK, QueryMessageResp{0, "查找消息成功", messageInfoList})
}
