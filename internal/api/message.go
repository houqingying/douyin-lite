package api

import (
	"douyin-lite/internal/service/message_service"
	"net/http"
	"strconv"

	"k8s.io/klog"

	"github.com/gin-gonic/gin"
)

type SendMessageResp struct {
	Code int32  `json:"status_code"`
	Msg  string `json:"status_msg,omitempty"`
}

func Message(c *gin.Context) {
	userId, got := c.Get("user_id")
	if !got {
		klog.Errorf("user_id didn't set properly, something may be wrong with the jwt")
		c.JSON(http.StatusOK, SendMessageResp{
			Code: 403,
			Msg:  "user_id is invalid",
		})
		return
	}
	fromUserId, _ := userId.(int64)

	// 读出其他request参数并检查合法性
	toUserIdStr := c.Query("to_user_id")

	// 转换toUserId为int64，判断是否小于0；若不小于0，才可在之后流程转换为int64
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
	err = message_service.SendMessage(fromUserId, toUserId, content)
	if err != nil {
		c.JSON(http.StatusOK, SendMessageResp{500, err.Error()})
		return
	}

	c.JSON(http.StatusOK, SendMessageResp{0, "发送消息成功"})
}

type QueryMessageResp struct {
	Code        int32                          `json:"status_code"`
	Msg         string                         `json:"status_msg,omitempty"`
	MessageList []*message_service.MessageInfo `json:"message_list"`
}

func MessageList(c *gin.Context) {
	userId, got := c.Get("user_id")
	if !got {
		klog.Errorf("user_id didn't set properly, something may be wrong with the jwt")
		c.JSON(http.StatusOK, SendMessageResp{
			Code: 403,
			Msg:  "user_id is invalid",
		})
		return
	}
	fromUserId, _ := userId.(int64)

	// 读出其他request参数并检查合法性
	toUserIdStr := c.Query("to_user_id")

	// 转换toUserId为int64，判断是否小于0；若不小于0，才可在之后流程转换为int64
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil || toUserId < 0 {
		c.JSON(http.StatusOK, QueryMessageResp{400, "to_user_id参数错误", nil})
		return
	}

	// 获取pre_msg_time 并获取最后一条消息
	preMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	lastMsg, err := message_service.QueryLastMessage(fromUserId, toUserId)

	if err != nil {
		// 获取最后一条消息失败
		c.JSON(http.StatusOK, QueryMessageResp{500, err.Error(), nil})
		return
	}

	// 判断最后一条消息时间与pre_msg_time的关系，如果lastTime <= preMsgTime，则无需更新消息，返回空列表
	lastTime, _ := strconv.ParseInt(lastMsg.CreateTime, 10, 64)
	if lastTime <= preMsgTime {
		c.JSON(http.StatusOK, QueryMessageResp{0, "暂时没有新消息", nil})
		return
	}

	// 调用Service层，完成查找
	messageInfoList, err := message_service.QueryMessage(fromUserId, toUserId)
	if err != nil {
		c.JSON(http.StatusOK, QueryMessageResp{500, err.Error(), nil})
		return
	}

	c.JSON(http.StatusOK, QueryMessageResp{0, "查找消息成功", messageInfoList})
}
