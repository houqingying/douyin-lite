package comment

import (
	"net/http"
	"strconv"

	"douyin-lite/repository"
	"douyin-lite/service/comment_service"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

// ActionResponse comment action response
type ActionResponse struct {
	StatusCode int32                       `json:"status_code"`
	StatusMsg  string                      `json:"status_msg,omitempty"`
	Comment    comment_service.CommentInfo `json:"comment"`
}

// Action comment action
// @Router /douyin/comment/action/ [post]
func Action(c *gin.Context) {
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

		var sendComment repository.Comment
		sendComment.UserId = uint(userId)
		sendComment.VideoId = uint(videoId)
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
