package db

import (
	"gorm.io/gorm"
)

const TableNameMessage = "message"

// Message mapped from table <message>
type Message struct {
	gorm.Model
	HostID  uint   `gorm:"column:host_id;not null" json:"host_id"`
	GuestID uint   `gorm:"column:guest_id;not null" json:"guest_id"`
	Content string `gorm:"column:content;not null" json:"content"`
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
