package comment_service

import (
	"time"

	"douyin-lite/repository"

	"k8s.io/klog"
)

// User UserVO
type User struct {
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	FollowCount    int64  `json:"follow_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
}

type CommentInfo struct {
	Id         int64  `json:"id,omitempty"`
	UserInfo   User   `json:"user,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

// CommentData CommentVO
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

const (
	DateFormat = "2006-01-02 15:04:05"
)

var CommentDao = repository.NewCommentDaoInstance()

type CommentService struct {
}

func (c *CommentService) CreateComment(comment repository.Comment) (CommentInfo, error) {
	klog.Info("comment service create comment")
	commentInfo := &repository.Comment{
		VideoId: comment.VideoId,
		UserId:  comment.UserId,
		Content: comment.Content,
	}
	err := CommentDao.CreateComment(commentInfo)
	if err != nil {
		klog.Errorf("create comment error: %s", err)
		return CommentInfo{}, err
	}
	return CommentInfo{
		Id: int64(commentInfo.ID),
		// TODO: add user info
		UserInfo:   User{},
		Content:    commentInfo.Content,
		CreateDate: commentInfo.CreatedAt.Format(DateFormat),
	}, nil
}

func (c *CommentService) DeleteComment(commentId int64) error {
	klog.Info("comment service delete comment")
	err := CommentDao.DeleteComment(uint(commentId))
	if err != nil {
		klog.Errorf("delete comment error: %s", err)
		return err
	}
	return nil
}

// GetList get comment list
func (c *CommentService) GetList(videoId int64) ([]CommentInfo, error) {
	klog.Info("comment service get comment list")
	comments, err := CommentDao.QueryCommentsByVideoId(int(videoId))
	if err != nil {
		klog.Errorf("get comment list error: %s", err)
		return nil, err
	}
	var commentInfos []CommentInfo
	for _, comment := range comments {
		commentInfos = append(commentInfos, CommentInfo{
			Id: int64(comment.ID),
			// TODO: add user info
			UserInfo:   User{},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format(DateFormat),
		})
	}
	return commentInfos, nil
}
