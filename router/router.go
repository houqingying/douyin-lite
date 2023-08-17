package router

import (
	"douyin-lite/handler/comment"
	"douyin-lite/handler/follow"
	"douyin-lite/handler/message"
	"douyin-lite/handler/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})

	//r.Static("static", config.Global.StaticSourcePath)
	baseGroup := r.Group("/douyin")

	baseGroup.GET("/relation/follow/list/query/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		//tokenStr := c.Param("token")
		followListResp, err := follow.QueryFollowListHandler(userIdStr)
		if err != nil {
			c.JSON(http.StatusOK, followListResp)
			return
		}
		c.JSON(http.StatusOK, followListResp)
	})

	baseGroup.GET("/relation/follower/list/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		//tokenStr := c.Param("token")
		followListResp, err := follow.QueryFollowerListHandler(userIdStr)
		if err != nil {
			c.JSON(http.StatusOK, followListResp)
			return
		}
		c.JSON(http.StatusOK, followListResp)
	})

	baseGroup.POST("/user/register/", func(c *gin.Context) {
		userName := c.Query("username")
		userPassword := c.Query("password")
		registerResp, err := user.RegisterUserHandler(userName, userPassword)
		if err != nil {
			c.JSON(http.StatusOK, registerResp)
			return
		}
		c.JSON(http.StatusOK, registerResp)
	})

	baseGroup.POST("/user/login/", func(c *gin.Context) {
		userName := c.Query("username")
		userPassword := c.Query("password")
		loginResp, err := user.LoginUserHandler(userName, userPassword)
		if err != nil {
			c.JSON(http.StatusOK, loginResp)
			return
		}
		c.JSON(http.StatusOK, loginResp)
	})

	baseGroup.GET("/user/", func(c *gin.Context) {
		userIdStr := c.Query("user_id")
		//userToken := c.Query("token")
		userInfoResp, err := user.QueryUserInfoHandler(userIdStr)
		if err != nil {
			c.JSON(http.StatusOK, userInfoResp)
			return
		}
		c.JSON(http.StatusOK, userInfoResp)
	})

	baseGroup.POST("/message/action/", message.SendMessageHandler)
	baseGroup.GET("/message/chat/", message.QueryMessageHandler)
	baseGroup.POST("/comment/action/", comment.Action)
	baseGroup.GET("/comment/list/", comment.List)
	return r
}
