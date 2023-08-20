package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"

	"douyin-lite/service/follow_service"
)

type RelationActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func RelationActionHandler(c *gin.Context) {
	klog.Info("post relation action")
	// get guest_id
	guestId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		klog.Errorf("to_user_id strconv.ParseInt error: %v", err)
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: -1,
			StatusMsg:  "to_user_id is invalid",
		})
		return
	}
	// get action_type
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		klog.Errorf("action_type strconv.ParseInt error: %v", err)
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: -1,
			StatusMsg:  "action_type is invalid",
		})
		return
	}
	hostId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		klog.Errorf("to_user_id strconv.ParseInt error: %v", err)
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: -1,
			StatusMsg:  "user_id is invalid",
		})
		return
	}

	err = follow_service.FollowAction(uint(hostId), uint(guestId), uint(actionType))
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, RelationActionResp{
			StatusCode: 0,
			StatusMsg:  "success",
		})
	}
}
