package video_service

import (
	"douyin-lite/internal/entity"
)

type PublishVideoListVO struct {
	Videos []VideoVO `json:"video_list,omitempty"`
}

type QueryPublishVideoList struct {
	userId int64
	videos []*entity.Video
}

func QueryPublishList(userId int64) (*FeedVideoListVO, error) {
	return NewQueryPublishList(userId).Do()
}
func NewQueryPublishList(userId int64) *QueryPublishVideoList {
	return &QueryPublishVideoList{userId: userId}
}

func (q *QueryPublishVideoList) Do() (*FeedVideoListVO, error) {
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
func (q *QueryPublishVideoList) checkNum() {

}
func (q *QueryPublishVideoList) prepareData() error {
	v := entity.Video{}
	q.videos, _ = v.GetVideoList(q.userId)
	//err := entity.NewVideoDAO().QueryPublishVideoList(q.userId, &q.videos)
	//if err != nil {
	//	return err
	//}
	return nil
}
func (q *QueryPublishVideoList) packData() (*FeedVideoListVO, error) {
	return &FeedVideoListVO{
		Videos: q.Video2VideoVO(q.userId),
	}, nil
}

func (q *QueryPublishVideoList) Video2VideoVO(userId int64) []VideoVO {
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
