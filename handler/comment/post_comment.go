package comment

import (
	"log"
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
	klog.Info("comment action")
	//id, _ := c.Get("user_id")
	//userid, _ := id.(string)
	//userId, err := strconv.ParseInt(userid, 10, 64)
	//if err != nil {
	//	klog.Errorf("strconv.ParseInt error: %s", err)
	//	c.JSON(http.StatusOK, ActionResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "comment userId json invalid",
	//	})
	//	return
	//}
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
	service := new(comment_service.CommentService)
	//错误处理
	if err != nil || actionType < 1 || actionType > 2 {
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment actionType json invalid",
		})
		log.Println("CommentController-Comment_Action: return actionType json invalid") //评论类型数据无效
		return
	}

	if actionType == 1 {
		content := c.Query("comment_text")

		var sendComment repository.Comment
		//sendComment.UserId = int(userId)
		sendComment.VideoId = uint(videoId)
		sendComment.Content = content
		//发表评论
		commentInfo, err := service.CreateComment(sendComment)
		//发表评论失败
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				StatusCode: -1,
				StatusMsg:  "send comment failed",
			})
			log.Printf("CommentController-Comment_Action: return send comment failed, %v", err) //发表失败
			return
		}

		//发表评论成功:
		//返回结果
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: 0,
			StatusMsg:  "send comment success",
			Comment:    commentInfo,
		})
		klog.Info("CommentController-Comment_Action: return Send success") //发表评论成功，返回正确信息
		return
	} else {
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, ActionResponse{
				StatusCode: -1,
				StatusMsg:  "comment commentId json invalid",
			})
			log.Println("CommentController-Comment_Action: return commentId json invalid") //评论id数据无效
			return
		}
		err = service.DeleteComment(commentId)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, ActionResponse{
			StatusCode: 0,
			StatusMsg:  "delete comment success",
		})
		klog.Info("CommentController-Comment_Action: return delete success") //删除评论成功，返回正确信息
	}
}
