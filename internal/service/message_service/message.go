package message_service

import (
	"douyin-lite/internal/entity"
	"errors"
)

// MessageInfo 封装消息记录
type MessageInfo struct {
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	Content    string `json:"content"`      // 消息内容
	CreateTime string `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
}

var userDao = entity.NewUserDaoInstance()
var messageDao = entity.GetMessageDaoInstance()

// SendMessage
// @description 消息发送服务，将接收到的消息保存到MySQL数据库 TODO: 考虑要不要加入MQ
// @auth	hqy			2023/08/17
// @param	fromUserId	uint	发送方用户Id
// @param	toUserId	uint	接收方用户Id
// @param	content		string	消息内容
// @return	err			error	当执行出现错误时返回error，否则返回nil
func SendMessage(fromUserId uint, toUserId uint, content string) error {
	// 验证toUserId的合法性
	_, err := userDao.QueryUserById(int64(toUserId))
	if err != nil {
		return errors.New("消息接收方id不存在")
	}
	// 不能发消息给自身
	if fromUserId == toUserId {
		return errors.New("不可向自身发送消息")
	}
	err = messageDao.CreateMessage(fromUserId, toUserId, content)
	return err
}

// QueryMessage
// @description 查询某两个用户消息列表，转换为目标格式输出
// @auth	hqy				2023/08/17
// @param	fromUserId		uint			发送方用户Id
// @param	toUserId		uint			接收方用户Id
// @return	messageInfoList	[]*MessageInfo	需要将Id转换为int64类型，并转换时间格式
// @return	err				error			当执行出现错误时返回error，否则返回nil
func QueryMessage(fromUserId uint, toUserId uint) ([]*MessageInfo, error) {
	// 从DAO层查询消息列表
	messageList, err := messageDao.QueryMessage(fromUserId, toUserId)
	if err != nil {
		return nil, err
	}
	// 转换为目标格式输出
	var messageInfoList = make([]*MessageInfo, len(messageList))
	for i, message := range messageList {
		messageInfo := MessageInfo{
			ID:         int64(message.ID),
			FromUserID: int64(message.FromUserId),
			ToUserID:   int64(message.ToUserId),
			CreateTime: message.CreatedAt.Format("2006-01-02 15:04:05"),
			Content:    message.Content,
		}
		messageInfoList[i] = &messageInfo
	}
	return messageInfoList, nil
}

// QueryLastMessage
// @description 查询某两个用户最新一条消息，转换为目标格式输出
// @param	fromUserId		uint			发送方用户Id
// @param	toUserId		uint			接收方用户Id
// @return	messageInfo	*MessageInfo		需要将Id转换为int64类型，并转换时间格式
// @return	err				error			当执行出现错误时返回error，否则返回nil
func QueryLastMessage(fromUserId uint, toUserId uint) (*MessageInfo, error) {
	message, err := messageDao.QueryLastMessage(fromUserId, toUserId)
	if err != nil {
		return nil, err
	}
	messageInfo := MessageInfo{
		ID:         int64(message.ID),
		FromUserID: int64(message.FromUserId),
		ToUserID:   int64(message.ToUserId),
		CreateTime: message.CreatedAt.Format("2006-01-02 15:04:05"),
		Content:    message.Content,
	}
	return &messageInfo, nil
}
