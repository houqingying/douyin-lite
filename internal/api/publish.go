package api

import (
	"douyin-lite/configs"
	"douyin-lite/pkg/fastdfs"
	"github.com/astaxie/beego/httplib"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type VideoListResponse struct {
	configs.Response
	VideoList []configs.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	file, _ := c.FormFile("data")
	if !fastdfs.IsDirExists("tmp") {
		err := os.Mkdir("tmp", 0777)
		if err != nil {
			c.JSON(http.StatusOK, configs.Response{
				StatusCode: 1,
				StatusMsg:  "创建临时目录失败",
			})
			return
		}
	}

	filePath := "tmp/" + file.Filename
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: 1,
			StatusMsg:  "保存临时文件失败",
		})
		return
	}
	// 截取图片

	// 上传到dfs
	peersUrl, _ := getPeersUrl(c)

	var obj map[string]interface{}
	req := httplib.Post(peersUrl + fastdfs.ApiUpload)
	req.PostFile("file", filePath)
	req.Param("output", "json")
	req.Param("scene", fastdfs.Scene)
	req.Param("path", fastdfs.Path)
	err = req.ToJSON(&obj)
	if err != nil {
		c.JSON(http.StatusOK, configs.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	obj["url"] = getShowUrlNotGroup(c) + obj["path"].(string)
	err = os.Remove(filePath)
	c.JSON(http.StatusOK, obj)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: configs.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}

var DemoVideos = []configs.Video{
	{
		Id:            1,
		Author:        configs.User{},
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

// PublishList all users have same publish video list
/*func PublishList(c *gin.Context) {
	//获取用户id
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: configs.Response{
				StatusCode: 1,
				StatusMsg:  "类型转换错误",
			},
			VideoList: nil,
		})
	}
	publishList, err := GetVideoService().PublishList(userId)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: configs.Response{
				StatusCode: 1,
				StatusMsg:  "数据库异常",
			},
			VideoList: nil,
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: configs.Response{
			StatusCode: 0,
			StatusMsg:  "查询成功",
		},
		VideoList: publishList,
	})
}
*/
