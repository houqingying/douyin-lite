package entity

import (
	"douyin-lite/pkg/storage"
	"sync"

	"gorm.io/gorm"
)

// Message 用户消息对象
type Message struct {
	gorm.Model        // 自动生成Unsigned BIGINT作为ID
	FromUserId uint   `gorm:"column:from_user_id;"`               // 接收方ID，BIGINT
	ToUserId   uint   `gorm:"column:to_user_id;"`                 // 发送方ID，BIGINT
	Content    string `gorm:"column:content;type:varchar(1024);"` // 消息内容，VARCHAR(1024)
}

// TableName 表名映射函数，映射至MySQL message表
func (Message) TableName() string {
	return "message"
}

// MessageDao 使用单例模式
type MessageDao struct {
}

var messageDao *MessageDao
var messageOnce sync.Once

func GetMessageDaoInstance() *MessageDao {
	messageOnce.Do(func() {
		messageDao = &MessageDao{}
	})
	return messageDao
}

// CreateMessage 定义MessageDao类型的发送消息方法，将接收到的消息保存到MySQL数据库
// @auth	hqy			2023/08/17
// @param	fromUserId	uint	发送方用户Id
// @param	toUserId	uint	接收方用户Id
// @param	content		string	消息内容
// @return	err			error	当执行出现错误时返回error，否则返回nil
func (*MessageDao) CreateMessage(fromUserId uint, toUserId uint, content string) error {
	message := Message{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Content:    content,
	}
	err := storage.DB.Create(&message).Error
	return err
}

// QueryMessage 查询数据库中fromUserId和toUserId间的所有聊天记录
// @auth	hqy				2023/08/17
// @param	fromUserId		uint		发送方用户Id
// @param	toUserId		uint		接收方用户Id
// @return	messageList 	[]*Message	消息记录列表
// @return	err				error		当执行出现错误时返回error，否则返回nil
func (*MessageDao) QueryMessage(fromUserId uint, toUserId uint) ([]*Message, error) {
	var messageList []*Message

	// 查询表中发送方和接收方参数均在 {fromUserId, toUserId}中的记录
	err := storage.DB.Where("from_user_id in ? AND to_user_id in ?",
		[]uint{fromUserId, toUserId}, []uint{fromUserId, toUserId}).Find(&messageList).Error

	if err != nil {
		return nil, err
	}

	return messageList, nil
}

// QueryLastMessage QueryMessageByDate 查询数据库中fromUserId和toUserId间的最后一条聊天记录
// @param 	fromUserId		uint		发送方用户Id
// @param 	toUserId		uint		接收方用户Id
// @return message *Message	消息记录
// @return err		error					当执行出现错误时返回error，否则返回nil
func (*MessageDao) QueryLastMessage(fromUserId uint, toUserId uint) (*Message, error) {
	message := Message{}
	err := storage.DB.Model(&Message{}).Where("(to_user_id=? and from_user_id=?) or (to_user_id=? and from_user_id=?)",
		toUserId, fromUserId, fromUserId, toUserId).Order("create_at desc").First(&message).Error

	if err != nil {
		return nil, err
	}

	return &message, nil
}
