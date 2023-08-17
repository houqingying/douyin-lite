package message

import (
	"douyin-lite/service/message_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueryMessageResp struct {
	Code        int32                          `json:"status_code"`
	Msg         string                         `json:"status_msg"`
	MessageList []*message_service.MessageInfo `json:"message_list"`
}

func QueryMessageHandler(c *gin.Context) {
	token := c.Query("token")
	fromUserId, err := ValidToken(token)

	// token验证失败
	if err != nil {
		sendMessageResp := QueryMessageResp{403, "用户token无效，拒绝用户请求", nil}
		c.JSON(http.StatusOK, sendMessageResp)
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
	messageInfoList, err := message_service.QueryMessage(fromUserId, uint(toUserId))
	if err != nil {
		c.JSON(http.StatusOK, QueryMessageResp{500, err.Error(), nil})
		return
	}

	c.JSON(http.StatusOK, QueryMessageResp{0, "查找消息成功", messageInfoList})
}
