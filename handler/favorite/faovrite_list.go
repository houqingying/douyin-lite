package favorite

import (
	"douyin-lite/service/favorite_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavoriteListResp struct {
	Code string `json:"status_code"`
	Msg  string `json:"status_msg"`
	*favorite_service.FavoriteListInfo
}

func QueryFavoriteListHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")
	followListResp, err := QueryFavoriteList(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, followListResp)
		return
	}
	c.JSON(http.StatusOK, followListResp)
}

func QueryFavoriteList(hostIdStr string) (*FavoriteListResp, error) {
	hostId, err := strconv.ParseInt(hostIdStr, 10, 64)
	if err != nil {
		return &FavoriteListResp{
			Code: "403",
			Msg:  "找不到用户",
		}, err
	}
	favoriteListData, err := favorite_service.QueryFavoriteListInfo(uint(hostId))
	if err != nil {
		return &FavoriteListResp{
			Code: "-1",
			Msg:  "查找失败",
		}, err
	}
	return &FavoriteListResp{
		Code:             "200",
		Msg:              "查找成功",
		FavoriteListInfo: favoriteListData,
	}, nil
}

// FavoriteList 获取列表方法
//func FavoriteList(c *gin.Context) {
//	// klog.Info("favorite list")
//	//user_id获取
//	userIdStr := c.Query("user_id")
//	// klog.Info(userIdStr)
//	userId, err := strconv.ParseInt(userIdStr, 10, 64)
//	if err != nil {
//		// klog.Errorf("strconv.ParseInt error: %s", err)
//		c.JSON(http.StatusOK, Response{
//			StatusCode: -1,
//			StatusMsg:  "comment userId json invalid",
//		})
//		return
//	}
//
//	var userIdHost uint
//	if v, ok := getUserId.(uint); ok {
//		userIdHost = v
//	}
//	userIdStr := c.Query("user_id") //自己id或别人id
//	userId, _ := strconv.ParseUint(userIdStr, 10, 10)
//	userIdNew := uint(userId)
//	if userIdNew == 0 {
//		userIdNew = userIdHost
//	}
//
//	//函数调用及响应
//	videoList, err := favorite_service.Favorite_List(userIdNew)
//	videoListNew := make([]FavoriteVideo, 0)
//	for _, m := range videoList {
//		var author = FavoriteAuthor{}
//		// var getUser = repository.User{}
//		user, err := UserDao.QueryUserById(m.AuthorId)
//		if err != nil {
//			c.JSON(http.StatusOK, Response{
//				StatusCode: 403,
//				StatusMsg:  "找不到作者！",
//			})
//			c.Abort()
//			return
//		}
//		//isfollowing
//		isfollowing, err := repository.NewFollowingDaoInstance().QueryisFollow(userIdHost, m.AuthorId)
//		//isfavorite
//		isfavorite := favorite_service.Check_Favorite(userIdHost, m.ID)
//		//作者信息
//		author.Id = user.ID
//		author.Name = user.Name
//		author.FollowCount = user.FollowingCount
//		author.FollowerCount = user.FollowerCount
//		author.IsFollow = isfollowing
//		//组装
//		var video = FavoriteVideo{}
//		video.Id = m.ID //类型转换
//		video.Author = author
//		video.PlayUrl = m.PlayUrl
//		video.CoverUrl = m.CoverUrl
//		video.FavoriteCount = m.FavoriteCount
//		video.CommentCount = m.CommentCount
//		video.IsFavorite = isfavorite
//		video.Title = m.Title
//
//		videoListNew = append(videoListNew, video)
//	}
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, FavoriteListResponse{
//			Response: Response{
//				StatusCode: 1,
//				StatusMsg:  "查找失败！",
//			},
//			VideoList: nil,
//		})
//	} else {
//		c.JSON(http.StatusOK, FavoriteListResponse{
//			Response: Response{
//				StatusCode: 0,
//				StatusMsg:  "找到列表！",
//			},
//			VideoList: videoListNew,
//		})
//	}
//}
