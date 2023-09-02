package api

import (
	"douyin-lite/configs"
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service"
	"douyin-lite/internal/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	configs.Response
	VideoList []*Video `json:"video_list"`
}

type Video struct {
	Id            int64   `json:"id,omitempty"`
	Author        *Author `json:"author,omitempty"`
	PlayUrl       string  `json:"play_url,omitempty"`
	CoverUrl      string  `json:"cover_url,omitempty"`
	FavoriteCount int64   `json:"favorite_count,omitempty"`
	IsFavorite    bool    `json:"is_favorite,omitempty"`
	CommentCount  int64   `json:"comment_count,omitempty"`
	Title         string  `json:"title,omitempty"`
}

type Author struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

// Publish 是一个处理文件上传和发布的处理函数。
func Publish(c *gin.Context) {
	// 从请求中获取上传的文件
	fileHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, configs.Response{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  configs.ErrNoFileUploaded,
		})
		return
	}

	// 保存到数据库
	title := c.PostForm("title")
	userId := c.GetInt64("user_id")

	go service.Publish(c, userId, title, fileHeader)

	// 返回上传结果
	c.JSON(http.StatusOK, configs.Response{
		StatusCode: configs.StatusSuccess,
		StatusMsg:  configs.SuccessMessage,
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	//获取用户id
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	var video = entity.Video{}

	// 获取视频信息
	videos, err := video.GetVideoList(userId)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: configs.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  configs.ErrDatabaseQueryFailed,
			},
			VideoList: nil,
		})
		return
	}

	// 获取用户信息
	userInfo, err := user_service.QueryAUserInfo2(userId)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: configs.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  fmt.Sprintf("获取用户信息失败，错误信息：%s", configs.ErrDatabaseQueryFailed),
			},
			VideoList: nil,
		})
		return
	}

	//处理视频获取列表
	var videoList = make([]*Video, len(videos))
	for i := 0; i < len(videos); i++ {
		videoList[i] = &Video{
			IsFavorite:    true,
			Id:            videos[i].ID,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			Title:         videos[i].Title,
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

	c.JSON(http.StatusOK, VideoListResponse{
		Response: configs.Response{
			StatusCode: configs.StatusSuccess,
			StatusMsg:  configs.SuccessMessage,
		},
		VideoList: videoList,
	})
	return
}
