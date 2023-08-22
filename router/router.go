package router

import (
	"douyin-lite/internal/api"
	"douyin-lite/middleware"

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

	douyinGroup := r.Group("/douyin")
	{
		// user路由组
		userGroup := douyinGroup.Group("/user")
		{
			userGroup.GET("/", middleware.JWTMiddleWare(), api.UserInfoHandler)
			userGroup.POST("/login/", middleware.SHAMiddleWare(), api.LoginUserHandler)
			userGroup.POST("/register/", middleware.SHAMiddleWare(), api.RegisterUserHandler)
		}
		// relation路由组
		relationGroup := douyinGroup.Group("relation")
		{
			relationGroup.POST("/action/", middleware.JWTMiddleWare(), api.RelationActionHandler)
			relationGroup.GET("/follow/list/", api.QueryFollowListHandler)
			relationGroup.GET("/follower/list/", api.QueryFollowerListHandler)
			relationGroup.GET("/friend/list/", middleware.JWTMiddleWare(), api.QueryFriendListHandler)
		}
		// message 路由组
		messageGroup := douyinGroup.Group("/message")
		{
			messageGroup.POST("/action/", middleware.JWTMiddleWare(), api.Message)
			messageGroup.GET("/chat/", middleware.JWTMiddleWare(), api.MessageList)
		}
		// comment路由组
		commentGroup := douyinGroup.Group("/comment")
		{
			commentGroup.POST("/action/", middleware.JWTMiddleWare(), api.Comment)
			commentGroup.GET("/list/", api.CommentList)
		}
		// favorite路由组
		favoriteGroup := douyinGroup.Group("/favorite")
		{
			favoriteGroup.POST("/action/", middleware.JWTMiddleWare(), api.Favorite)
			favoriteGroup.GET("/list/", api.FavoriteList)
		}

		// favorite路由组
		publishGroup := douyinGroup.Group("/publish")
		{
			publishGroup.POST("/action/", middleware.JWTMiddleWare(), api.Publish)
			publishGroup.GET("/list/", api.PublishList)
		}

		douyinGroup.GET("/feed/", api.Feed)
	}

	return r
}
