package video_service

import (
	"errors"
	"fmt"
	"time"

	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
)

const (
	MaxVideoNum = 30
)

type VideoVO struct {
	Id            int64           `json:"id,omitempty"`
	Author        entity.UserInfo `json:"author,omitempty"`
	PlayUrl       string          `json:"play_url,omitempty"`
	CoverUrl      string          `json:"cover_url,omitempty"`
	FavoriteCount int64           `json:"favorite_count,omitempty"`
	CommentCount  int64           `json:"comment_count,omitempty"`
	Title         string          `json:"title,omitempty"`
	IsFavorite    bool            `json:"is_favorite,omitempty"`
}
type FeedVideoListVO struct {
	Videos   []VideoVO `json:"video_list,omitempty"`
	NextTime int64     `json:"next_time,omitempty"`
}

type QueryFeedVideoListFlow struct {
	latestTime time.Time
	userId     int64
	videos     []*entity.Video
	nextTime   int64
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoListVO, error) {
	return NewQueryFeedVideoListFlow(userId, latestTime).Do()
}
func NewQueryFeedVideoListFlow(userId int64, latestTime time.Time) *QueryFeedVideoListFlow {
	return &QueryFeedVideoListFlow{userId: userId, latestTime: latestTime}
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
	err := entity.NewVideoDAO().QueryVideoListByLimitAndTime(MaxVideoNum, q.latestTime, &q.videos)
	if err != nil {
		return err
	}
	latestTime, _ := FillVideoListFields(&q.videos, q.userId) //不是致命错误，不返回
	//准备好时间戳
	if latestTime != nil {
		fmt.Println(latestTime)
		q.nextTime = (latestTime).UnixNano() / 1e6
		return nil
	}
	q.nextTime = time.Now().UnixNano() / 1e6
	return nil
}
func (q *QueryFeedVideoListFlow) packData() (*FeedVideoListVO, error) {
	return &FeedVideoListVO{
		Videos:   q.Video2VideoVO(q.userId),
		NextTime: q.nextTime,
	}, nil
}
func (q *QueryFeedVideoListFlow) Video2VideoVO(userId int64) []VideoVO {
	var videoVOList = make([]VideoVO, len(q.videos))
	for i, video := range q.videos {
		videoVO := VideoVO{
			Id:            int64(video.ID),
			Author:        e2sUserInfo(video.Author),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			Title:         video.Title,
		}
		videoVO.IsFavorite, _ = entity.NewFavoriteDaoInstance().Query_Check_Favorite(userId, videoVO.Id)
		videoVOList[i] = videoVO
	}
	return videoVOList
}
func FillVideoListFields(videos *[]*entity.Video, userId int64) (*time.Time, error) {
	size := len(*videos)
	if nil == *videos || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	latestTime := (*videos)[size-1].CreatedAt //获取最近的投稿时间
	for i := 0; i < size; i++ {
		userInfo, err := user_service.QueryAUserInfo1(userId, int64((*videos)[i].AuthorId))
		if err != nil {
			continue
		}
		(*videos)[i].Author = s2eUserInfo(*userInfo)
	}
	return &latestTime, nil
}
func e2sUserInfo(info entity.UserInfo) entity.UserInfo {
	return entity.UserInfo{
		ID:              info.ID,
		Name:            info.Name,
		Avatar:          info.Avatar,
		BackgroundImage: info.BackgroundImage,
		Signature:       info.Signature,
		FollowingCount:  info.FollowingCount,
		FollowerCount:   info.FollowerCount,
		IsFollow:        info.IsFollow,
		TotalFavorited:  info.TotalFavorited,
		WorkCount:       info.WorkCount,
		FavoriteCount:   info.FavoriteCount,
	}
}
func s2eUserInfo(info entity.UserInfo) entity.UserInfo {
	return entity.UserInfo{
		ID:              info.ID,
		Name:            info.Name,
		Avatar:          info.Avatar,
		BackgroundImage: info.BackgroundImage,
		Signature:       info.Signature,
		FollowingCount:  info.FollowingCount,
		FollowerCount:   info.FollowerCount,
		IsFollow:        info.IsFollow,
		TotalFavorited:  info.TotalFavorited,
		WorkCount:       info.WorkCount,
		FavoriteCount:   info.FavoriteCount,
	}
}
