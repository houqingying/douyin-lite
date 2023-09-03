package api

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/comment_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

// ActionResponse comment action response
type ActionResponse struct {
	StatusCode int32                       `json:"status_code"`
	StatusMsg  string                      `json:"status_msg,omitempty"`
	Comment    comment_service.CommentInfo `json:"comment"`
}

// Comment comment action
// @Router /douyin/comment/action/ [post]
func Comment(c *gin.Context) {
	id, _ := c.Get("user_id")
	userId, _ := id.(int64)

	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		klog.Errorf("strconv.ParseInt error: %s", err)
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment userId json invalid",
		})
		return
	}
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)

	if err != nil {
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment actionType json invalid",
		})
		klog.Infof("CommentController Action: comment actionType json invalid")
	}

	service := new(comment_service.CommentService)
	if actionType < 1 || actionType > 2 {
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment actionType json invalid",
		})
		klog.Info("CommentController Action: actionType json invalid")
		return
	}

	// 1 : send comment
	if actionType == 1 {
		content := c.Query("comment_text")

		var sendComment entity.Comment
		sendComment.UserId = userId
		sendComment.VideoId = videoId
		sendComment.Content = content

		// do send comment
		commentInfo, err := service.CreateComment(sendComment)
		// send comment failed
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				StatusCode: -1,
				StatusMsg:  "send comment failed",
			})
			klog.Infof("CommentController Action: return send comment failed, %v", err) //发表失败
			return
		}

		// send comment success
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: 0,
			StatusMsg:  "send comment success",
			Comment:    commentInfo,
		})
		klog.Info("CommentController Action: return Send success")
		return
	} else { // 2 : delete comment
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				StatusCode: -1,
				StatusMsg:  "comment commentId json invalid",
			})
			klog.Infof("CommentController Action: return commentId json invalid")
			return
		}
		// do delete comment
		err = service.DeleteComment(commentId)
		// delete comment failed
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				StatusCode: -1,
				StatusMsg:  "delete comment failed",
			})
			klog.Infof("CommentController Action: return delete failed, %v", err)
			return
		}
		// delete comment success
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: 0,
			StatusMsg:  "delete comment success",
		})
		klog.Info("CommentController Action: return delete success")
	}
}

// ListResponse comment list request
type ListResponse struct {
	StatusCode  int32                         `json:"status_code"`
	StatusMsg   string                        `json:"status_msg,omitempty"`
	CommentList []comment_service.CommentInfo `json:"comment_list,omitempty"`
}

// CommentList comment list
// @Router /douyin/comment/list/ [get]
func CommentList(c *gin.Context) {
	klog.Info("comment list")
	id := c.Query("video_id")
	klog.Info(id)
	videoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		klog.Errorf("strconv.ParseInt error: %s", err)
		c.JSON(http.StatusOK, ListResponse{
			StatusCode: -1,
			StatusMsg:  "comment videoId json invalid",
		})
		return
	}
	service := new(comment_service.CommentService)
	commentList, err := service.GetList(int64(int(videoId)))
	if err != nil {
		c.JSON(http.StatusOK, ListResponse{
			StatusCode: -1,
			StatusMsg:  "comment list failed",
		})
		klog.Info("CommentController List: return comment list failed")
		return
	}
	c.JSON(http.StatusOK, ListResponse{
		StatusCode:  0,
		StatusMsg:   "comment list success",
		CommentList: commentList,
	})
	klog.Info("CommentController List: return comment list success")
}
