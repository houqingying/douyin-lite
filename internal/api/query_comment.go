package api

import (
	"douyin-lite/internal/service/comment_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

// ListResponse comment list request
type ListResponse struct {
	StatusCode  int32                         `json:"status_code"`
	StatusMsg   string                        `json:"status_msg,omitempty"`
	CommentList []comment_service.CommentInfo `json:"comment_list,omitempty"`
}

// List comment list
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
