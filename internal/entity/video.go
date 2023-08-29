package entity

import (
	"douyin-lite/pkg/storage"
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	gorm.Model
	AuthorId      int64
	ID            int64 `json:"id" gorm:"id,omitempty"`
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	Author        UserInfo `json:"author,omitempty" gorm:"-"`
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
	return storage.DB.Model(&Video{}).Where("created_at<?", latestTime).
		Order("created_at DESC").Limit(limit).
		Select([]string{"id", "author_id", "play_url", "cover_url", "favorite_count", "comment_count", "title", "created_at"}).
		Find(videoList).Error
}

func (v *Video) SaveVideo() error {
	return storage.DB.Create(&v).Error
}

func (v *Video) GetVideoList(userId int64) (videos []*Video, err error) {
	err = storage.DB.Model(&v).Where("author_id = ?", userId).Find(&videos).Error
	return
}
