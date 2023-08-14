package comment_service

import (
	"github.com/houqingying/douyin-lite/repository"
	"k8s.io/klog"
)

// https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707528
type CommentServiceImpl struct {
	//UserService
}

var (
	CommentDao *repository.CommentDao = &repository.CommentDao{}
)

func (c *CommentServiceImpl) CreateComment(comment repository.Comment) (CommentInfo, error) {
	klog.Info("comment service create comment")
	commentInfo := &repository.Comment{
		VideoId: comment.VideoId,
		UserId:  comment.UserId,
		Comment: comment.Comment,
	}
	err := CommentDao.CreateComment(commentInfo)
	if err != nil {
		klog.Errorf("create comment error: %s", err)
		return CommentInfo{}, err
	}
	return CommentInfo{
		Id:         int64(commentInfo.ID),
		UserInfo:   User{},
		Content:    commentInfo.Comment,
		CreateDate: commentInfo.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (c *CommentServiceImpl) DeleteComment(commentId int64) error {
	klog.Info("comment service delete comment")
	err := CommentDao.DeleteComment(uint(commentId))
	if err != nil {
		klog.Errorf("delete comment error: %s", err)
		return err
	}
	return nil
}

// GetList get comment list
func (c *CommentServiceImpl) GetList(videoId int64) ([]CommentInfo, error) {
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
			Content:    comment.Comment,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return commentInfos, nil
}
