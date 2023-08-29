package api

import (
	"douyin-lite/configs"
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/user_service"
	fastDFS "douyin-lite/pkg/fastdfs"
	"douyin-lite/pkg/file"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strconv"
)

type VideoListResponse struct {
	configs.Response
	VideoList []*Video `json:"video_list"`
}

type Video struct {
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
	Author        *Author
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

	// 构建文件保存的路径
	filePath := configs.TmpFileDir + "/" + fileHeader.Filename
	imagePath := configs.TmpFileDir + "/" + uuid.New().String() + ".jpeg"

	// 检查是否存在临时目录，如果不存在则创建
	if !file.IsDirExists(configs.TmpFileDir) {
		err := os.Mkdir(configs.TmpFileDir, 0777)
		if err != nil {
			c.JSON(http.StatusOK, configs.Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  configs.ErrCreateTmpDirectory,
			})
			return
		}
	}

	// 将上传的文件保存到本地临时目录
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  configs.ErrSaveTmpFile,
		})
		return
	}

	//在服务器上执行ffmpeg 从视频流中获取第一帧截图，并上传图片服务器，保存图片链接
	//向队列中添加消息
	/*if err := ffmpeg.ReadFrameAsJpeg(filePath, imagePath, configs.FrameNum); err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  fmt.Sprintf("%s,错误信息:%s", configs.ErrReadFrameFailure, err.Error()),
		})
		return
	}*/

	// 上传到dfs
	obj, err := fastDFS.FDClient.UploadGoFastDFS(filePath)
	if err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  fmt.Sprintf("%s,错误信息:%s", configs.ErrUploadToDFS, err.Error()),
		})
		return
	}

	// 保存到数据库
	title := c.PostForm("title")
	userId := c.GetInt64("user_id")
	//保存视频在数据库中
	var video = entity.Video{
		AuthorId:      userId,
		PlayUrl:       obj["url"].(string),
		CoverUrl:      imagePath,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	if err = video.SaveVideo(); err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  configs.ErrDatabaseInsertFailed,
		})
		return
	}

	// 删除本地临时文件
	if err = os.Remove(filePath); err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  configs.ErrDeleteTmpFile,
		})
		return
	}

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
			AuthorId:      videos[i].AuthorId,
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
