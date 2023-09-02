package service

import (
	"douyin-lite/configs"
	"douyin-lite/internal/entity"
	fastDFS "douyin-lite/pkg/fastdfs"
	"douyin-lite/pkg/ffmpeg"
	"douyin-lite/pkg/file"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"os"
)

func Publish(c *gin.Context, userId int64, title string, fileHeader *multipart.FileHeader) {
	// 构建文件保存的路径
	filePath := configs.TmpFileDir + "/" + fileHeader.Filename

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
	err := c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  configs.ErrSaveTmpFile,
		})
		return
	}

	// 上传到dfs
	obj, err := fastDFS.FDClient.UploadGoFastDFS(filePath)
	if err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  fmt.Sprintf("%s,错误信息:%s", configs.ErrUploadToDFS, err.Error()),
		})
		return
	}

	//在服务器上执行ffmpeg 从视频流中获取第一帧截图，并上传图片服务器，保存图片链接
	imagePath := uuid.New().String() + ".jpg"
	remoteFilePath := "/root/go-fastdfs/fastdfs/data/files/video/" + fileHeader.Filename
	remoteImagePath := "/root/go-fastdfs/fastdfs/data/files/video/" + imagePath
	if err := ffmpeg.ReadFrameAsJpeg(remoteFilePath, remoteImagePath); err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  fmt.Sprintf("%s,错误信息:%s", configs.ErrReadFrameFailure, err.Error()),
		})
		return
	}

	//保存视频在数据库中
	imagePath = "http://47.102.185.103:8085/tiktok/video/" + imagePath
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

	//更改用户数据在数据库中
	/*if err = video.SaveVideo(); err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  configs.ErrDatabaseInsertFailed,
		})
		return
	}*/

	// 删除本地临时文件
	if err = os.Remove(filePath); err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  configs.ErrDeleteTmpFile,
		})
		return
	}
}
