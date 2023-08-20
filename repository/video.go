package repository

import (
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId      uint
	PlayUrl       string
	CoverUrl      string
	FavoriteCount uint
	CommentCount  uint
	Title         string
	Author        User `json:"author,omitempty" gorm:"-"`
}
type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDAO() *VideoDao {
	videoOnce.Do(func() {
		videoDao = new(VideoDao)
	})
	return videoDao
}

// QueryVideoListByLimitAndTime  返回按投稿时间倒序的视频列表，并限制为最多limit个//feed
func (*VideoDao) QueryVideoListByLimitAndTime(limit int, latestTime time.Time, videoList *[]*Video) error {
	if videoList == nil {
		return errors.New("QueryVideoListByLimit videoList 空指针")
	}
	return db.Model(&Video{}).Where("created_at<?", latestTime).
		Order("created_at ASC").Limit(limit).
		Select([]string{"id", "author_id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "created_at"}).
		Find(videoList).Error
}
