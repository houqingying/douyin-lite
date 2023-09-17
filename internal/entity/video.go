package entity

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"

	"douyin-lite/pkg/storage"
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

func (v *Video) GetVideoCount(userId int64) (count int64, err error) {
	err = storage.DB.Model(&v).Where("author_id = ?", userId).Count(&count).Error
	return
}

// UpdateVideoCommentCount add video count
func (*VideoDao) UpdateVideoCommentCount(id int64, num int64) error {
	return storage.DB.Model(&Video{}).Where("id = ?", id).Update("comment_count", gorm.Expr("comment_count + ?", num)).Error
}
