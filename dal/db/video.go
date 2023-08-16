package db

import (
	"gorm.io/gorm"
)

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	gorm.Model
	AuthorID      uint   `gorm:"column:author_id;not null" json:"author_id"`
	PlayURL       string `gorm:"column:play_url;not null" json:"play_url"`
	CoverURL      string `gorm:"column:cover_url;not null" json:"cover_url"`
	FavoriteCount uint   `gorm:"column:favorite_count;not null" json:"favorite_count"`
	CommentCount  uint   `gorm:"column:comment_count;not null" json:"comment_count"`
	Title         string `gorm:"column:title;not null" json:"title"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
