package api

import (
	"douyin-lite/configs"
	"douyin-lite/internal/entity"
	favorite_service2 "douyin-lite/internal/service/favorite_service"
	"douyin-lite/internal/service/user_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

type FavoriteActionResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Favorite(c *gin.Context) {
	// klog.Info("post relation action")
	// get guest_id
	user_Id, got := c.Get("user_id")
	if !got {
		klog.Errorf("user_id didn't set properly, something may be wrong with the jwt")
		c.JSON(http.StatusOK, SendMessageResp{
			Code: 403,
			Msg:  "user_id is invalid",
		})
		return
	}
	user_IdInt64, _ := user_Id.(int64)
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		fmt.Println("actionType=?", actionType)
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: -1,
			StatusMsg:  "操作无效",
		})
		return
	}
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: -1,
			StatusMsg:  "video_id is invalid",
		})
		return
	}

	err = favorite_service2.Favorite_Action(user_IdInt64, video_id, actionType)
	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteActionResp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		})
	}
}

type FavoriteListResp struct {
	configs.Response
	FavoriteList []*Video `json:"video_list"`
}

func FavoriteVideoList(c *gin.Context) {
	// user_Id, got := c.Get("user_id")
	// if !got {
	// 	klog.Errorf("user_id didn't set properly, something may be wrong with the jwt")
	// 	c.JSON(http.StatusOK, SendMessageResp{
	// 		Code: 403,
	// 		Msg:  "user_id is invalid",
	// 	})
	// 	return
	// }
	// user_IdInt64, _ := user_Id.(int64)
	// followListResp, err := QueryFavoriteList(user_IdInt64)
	// if err != nil {
	// 	c.JSON(http.StatusOK, followListResp)
	// 	return
	// }
	// c.JSON(http.StatusOK, followListResp)

	//获取用户id
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	var video = entity.Favorite{}

	// 获取当前用户点赞的视频列表
	videos, err := video.GetFavoriteListResp(userId)
	fmt.Println("videos=?", videos)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResp{
			Response: configs.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  configs.ErrDatabaseQueryFailed,
			},
			FavoriteList: nil,
		})
		return
	}

	// 获取用户信息
	userInfo, err := user_service.QueryAUserInfo2(userId)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResp{
			Response: configs.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  fmt.Sprintf("获取用户信息失败，错误信息：%s", configs.ErrDatabaseQueryFailed),
			},
			FavoriteList: nil,
		})
		return
	}

	//处理视频获取列表
	var videoList = make([]*Video, len(videos))
	for i := 0; i < len(videos); i++ {
		videoList[i] = &Video{
			IsFavorite: true,
			Id:         videos[i].ID,
			//PlayUrl:       videos[i].PlayUrl,
			//CoverUrl:      videos[i].CoverUrl,
			//FavoriteCount: videos[i].FavoriteCount,
			//CommentCount:  videos[i].CommentCount,
			//Title:         videos[i].Title,
			Author: &Author{
				ID:              userInfo.ID,
				Name:            userInfo.Name,
				Avatar:          userInfo.Avatar,
				BackgroundImage: userInfo.BackgroundImage,
				Signature:       userInfo.Signature,
				FollowingCount:  userInfo.FollowingCount,
				FollowerCount:   userInfo.FollowerCount,
				IsFollow:        userInfo.IsFollow,
				TotalFavorited:  userInfo.TotalFavorited,
				WorkCount:       userInfo.WorkCount,
				FavoriteCount:   userInfo.FavoriteCount,
			},
		}
	}

	c.JSON(http.StatusOK, FavoriteListResp{
		Response: configs.Response{
			StatusCode: configs.StatusSuccess,
			StatusMsg:  configs.SuccessMessage,
		},
		FavoriteList: videoList,
	})
	return

}

// func QueryFavoriteList(hostInt64 int64) (*FavoriteListResp, error) {
// 	//hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
// 	//if err != nil {
// 	//	return &FavoriteListResp{
// 	//		Code: "403",
// 	//		Msg:  "找不到用户",
// 	//	}, err
// 	//}
// 	favoriteListData, err := favorite_service2.QueryFavoriteListInfo(hostInt64)
// 	if err != nil {
// 		return &FavoriteListResp{
// 			Code: "-1",
// 			Msg:  "查找失败",
// 		}, err
// 	}
// 	return &FavoriteListResp{
// 		Code:             "200",
// 		Msg:              "查找成功",
// 		FavoriteListInfo: favoriteListData,
// 	}, nil
// }
