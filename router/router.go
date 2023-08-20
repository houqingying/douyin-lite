package router

import (
	"douyin-lite/handler/comment"
	"douyin-lite/handler/message"
	"douyin-lite/middleware"

	"github.com/gin-gonic/gin"

	"douyin-lite/handler/follow"
	"douyin-lite/handler/user"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})

	douyinGroup := r.Group("/douyin")
	{
		// user路由组
		userGroup := douyinGroup.Group("/user")
		{
			userGroup.GET("/", middleware.JWTMiddleWare(), user.UserInfoHandler)
			userGroup.POST("/login/", middleware.SHAMiddleWare(), user.LoginUserHandler)
			userGroup.POST("/register/", middleware.SHAMiddleWare(), user.RegisterUserHandler)
		}
		// relation路由组
		relationGroup := douyinGroup.Group("relation")
		{
			relationGroup.POST("/action/", middleware.JWTMiddleWare(), follow.RelationActionHandler)
			relationGroup.GET("/follow/list/", follow.QueryFollowListHandler)
			relationGroup.GET("/follower/list/", follow.QueryFollowerListHandler)
			relationGroup.GET("/friend/list/", middleware.JWTMiddleWare(), follow.QueryFriendListHandler)
		}
		// message 路由组
		messageGroup := douyinGroup.Group("/message")
		{
			messageGroup.POST("/action/", middleware.JWTMiddleWare(), message.SendMessageHandler)
			messageGroup.GET("/chat/", middleware.JWTMiddleWare(), message.QueryMessageHandler)
		}
		// comment路由组
		commentGroup := douyinGroup.Group("/comment")
		{
			commentGroup.POST("/action/", middleware.JWTMiddleWare(), comment.Action)
			commentGroup.GET("/list/", comment.List)
		}
	}

	return r
}
