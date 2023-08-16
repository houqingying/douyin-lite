package db

import (
	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	gorm.Model
	Name            string `gorm:"column:name;not null" json:"name"`
	Password        string `gorm:"column:password;not null" json:"password"`
	Avatar          string `gorm:"column:avatar;not null" json:"avatar"`
	BackgroundImage string `gorm:"column:background_image;not null" json:"background_image"`
	Signature       string `gorm:"column:signature;not null" json:"signature"`
	FollowCount     uint   `gorm:"column:follow_count;not null" json:"follow_count"`
	FollowerCount   uint   `gorm:"column:follower_count;not null" json:"follower_count"`
	TotalFavorited  uint   `gorm:"column:total_favorited;not null" json:"total_favorited"`
	WorkCount       uint   `gorm:"column:work_count;not null" json:"work_count"`
	FavoriteCount   uint   `gorm:"column:favorite_count;not null" json:"favorite_count"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
