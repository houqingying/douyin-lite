package message

import (
	"douyin-lite/service/message_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SendMessageResp struct {
	Code int32  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

// ValidToken 验证请求携带token的合法性，如果合法，返回用户ID；如果非法，设置error
func ValidToken(token string) (uint, error) {
	// TODO: 负责登录功能的同学在某个文件中提供验证token合法性的方法，注意用户id的类型是int64
	return 1, nil
}

func SendMessageHandler(c *gin.Context) {
	token := c.Query("token")
	fromUserId, err := ValidToken(token)

	// token验证失败
	if err != nil {
		sendMessageResp := SendMessageResp{403, "用户token无效，拒绝用户请求"}
		c.JSON(http.StatusOK, sendMessageResp)
		return
	}

	// 读出其他request参数并检查合法性
	toUserIdStr := c.Query("to_user_id")

	// 转换toUserId为int64，判断是否小于0；若不小于0，才可在之后流程转换为uint
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil || toUserId < 0 {
		c.JSON(http.StatusOK, SendMessageResp{400, "to_user_id参数错误"})
		return
	}

	actionType := c.Query("action_type")
	if actionType != "1" {
		c.JSON(http.StatusOK, SendMessageResp{400, "不支持除发送消息外的其他请求"})
		return
	}

	content := c.Query("content")
	if content == "" {
		c.JSON(http.StatusOK, SendMessageResp{400, "发送内容不能为空"})
		return
	}

	// 调用Service层，完成发送服务
	err = message_service.SendMessage(fromUserId, uint(toUserId), content)
	if err != nil {
		c.JSON(http.StatusOK, SendMessageResp{500, err.Error()})
		return
	}

	c.JSON(http.StatusOK, SendMessageResp{0, "发送消息成功"})
}
