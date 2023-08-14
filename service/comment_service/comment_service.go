package comment_service

import (
	"time"

	"github.com/houqingying/douyin-lite/repository"
)

type CommentService interface {
	// CountFromVideoId count comment from video id
	CountFromVideoId(id int64) (int64, error)
	// CreateComment create comment
	CreateComment(comment repository.Comment) (CommentInfo, error)
	// DeleteComment delete comment
	DeleteComment(commentId int64) error
	// GetList get comment list
	GetList(videoId int64, userId int64) ([]CommentInfo, error)
}

// User 最终封装后,controller返回的User结构体
type User struct {
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	FollowCount    int64  `json:"follow_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
}

// CommentInfo 查看评论-传出的结构体-service
type CommentInfo struct {
	Id         int64  `json:"id,omitempty"`
	UserInfo   User   `json:"user,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentData struct {
	Id            int64     `json:"id,omitempty"`
	UserId        int64     `json:"user_id,omitempty"`
	Name          string    `json:"name,omitempty"`
	FollowCount   int64     `json:"follow_count"`
	FollowerCount int64     `json:"follower_count"`
	IsFollow      bool      `json:"is_follow"`
	Content       string    `json:"content,omitempty"`
	CreateDate    time.Time `json:"create_date,omitempty"`
}
