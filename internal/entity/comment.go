package entity

import (
	"douyin-lite/pkg/storage"
	"sync"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoId int64  `gorm:"column:video_id;not null" json:"video_id"`
	UserId  int64  `gorm:"column:user_id;not null" json:"user_id"`
	Content string `gorm:"column:content;not null" json:"content"`
}

type CommentDao struct {
}

var commentDao *CommentDao
var CommentOnce sync.Once

func NewCommentDao() *CommentDao {
	CommentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

// CreateComment creates a comment
func (c *CommentDao) CreateComment(comment *Comment) error {
	err := storage.DB.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteComment deletes a comment
func (c *CommentDao) DeleteComment(id uint) error {
	err := storage.DB.Where("id = ?", id).Delete(&Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCommentByVideoId deletes a comment by video id
func (c *CommentDao) DeleteCommentByVideoId(videoId int) error {
	err := storage.DB.Where("video_id = ?", videoId).Delete(&Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCommentByUserId deletes a comment by user id
func (c *CommentDao) DeleteCommentByUserId(userId int) error {
	err := storage.DB.Where("user_id = ?", userId).Delete(&Comment{}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateComment updates a comment
func (*CommentDao) UpdateComment(id uint, comment string) error {
	err := storage.DB.Model(&Comment{}).Where("id = ?", id).Update("comment", comment).Error
	if err != nil {
		return err
	}
	return nil
}

// QueryCommentById gets a comment by id
func (c *CommentDao) QueryCommentById(id uint) (*Comment, error) {
	var comment Comment
	err := storage.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// QueryCommentsByVideoId gets comments by video id
func (c *CommentDao) QueryCommentsByVideoId(videoId int) ([]Comment, error) {
	var comments []Comment
	// sort by created_at desc
	err := storage.DB.Where("video_id = ?", videoId).Order("created_at desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// QueryCommentsByUserId gets comments by user id
func (c *CommentDao) QueryCommentsByUserId(userId int) ([]Comment, error) {
	var comments []Comment
	err := storage.DB.Where("user_id = ?", userId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
