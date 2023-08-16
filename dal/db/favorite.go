package db

import (
	"gorm.io/gorm"
)

const TableNameFavorite = "favorite"

// Favorite mapped from table <favorite>
type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"column:user_id;not null" json:"user_id"`
	VideoID uint `gorm:"column:video_id;not null" json:"video_id"`
	Status  uint `gorm:"column:status;not null" json:"status"`
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}
