package video_service

import (
	"douyin-lite/repository"
	"errors"
	"fmt"
	"time"
)

const (
	MaxVideoNum = 30
)

type VideoVO struct {
	Id            int64           `json:"id,omitempty"`
	Author        repository.User `json:"author,omitempty"`
	PlayUrl       string          `json:"play_url,omitempty"`
	CoverUrl      string          `json:"cover_url,omitempty"`
	FavoriteCount int64           `json:"favorite_count,omitempty"`
	CommentCount  int64           `json:"comment_count,omitempty"`
	Title         string          `json:"title,omitempty"`
}
type FeedVideoListVO struct {
	Videos   []VideoVO `json:"video_list,omitempty"`
	NextTime int64     `json:"next_time,omitempty"`
}

type QueryFeedVideoListFlow struct {
	latestTime time.Time

	videos   []*repository.Video
	nextTime int64
}

func QueryFeedVideoList(latestTime time.Time) (*FeedVideoListVO, error) {
	return NewQueryFeedVideoListFlow(latestTime).Do()
}
func NewQueryFeedVideoListFlow(latestTime time.Time) *QueryFeedVideoListFlow {
	return &QueryFeedVideoListFlow{latestTime: latestTime}
}

func (q *QueryFeedVideoListFlow) Do() (*FeedVideoListVO, error) {
	//所有传入的参数不填也应该给他正常处理
	q.checkNum()
	var feedVideoListVO *FeedVideoListVO
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	feedVideoListVO, err := q.packData()
	if err != nil {
		return nil, err
	}
	return feedVideoListVO, nil
}
func (q *QueryFeedVideoListFlow) checkNum() {
	if q.latestTime.IsZero() {
		q.latestTime = time.Now()
	}
}
func (q *QueryFeedVideoListFlow) prepareData() error {
	err := repository.NewVideoDAO().QueryVideoListByLimitAndTime(MaxVideoNum, q.latestTime, &q.videos)
	if err != nil {
		return err
	}
	latestTime, _ := FillVideoListFields(&q.videos) //不是致命错误，不返回
	//准备好时间戳
	if latestTime != nil {
		fmt.Println(*latestTime)
		q.nextTime = (*latestTime).UnixNano() / 1e9
		return nil
	}
	q.nextTime = time.Now().UnixNano() / 1e9
	return nil
}
func (q *QueryFeedVideoListFlow) packData() (*FeedVideoListVO, error) {
	return &FeedVideoListVO{
		Videos:   q.Video2VideoVO(),
		NextTime: q.nextTime,
	}, nil
}
func (q *QueryFeedVideoListFlow) Video2VideoVO() []VideoVO {
	var videoVOList = make([]VideoVO, len(q.videos))
	for i, video := range q.videos {
		videoVO := VideoVO{
			Id:            int64(video.ID),
			Author:        video.Author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			Title:         video.Title,
		}
		videoVOList[i] = videoVO
	}
	return videoVOList
}
func FillVideoListFields(videos *[]*repository.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	userDao := repository.NewUserDaoInstance()
	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	for i := 0; i < size; i++ {
		userInfo, err := userDao.QueryUserById((*videos)[i].AuthorId)
		if err != nil {
			continue
		}
		(*videos)[i].Author = *userInfo
	}
	return &latestTime, nil
}
