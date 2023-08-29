package api

import (
	"douyin-lite/internal/entity"
	user_service2 "douyin-lite/internal/service/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryUserInfoResp struct {
	Code string           `json:"status_code"`
	Msg  string           `json:"status_msg"`
	User *entity.UserInfo `json:"user"`
}

func UserInfoHandler(c *gin.Context) {
	userIdStr := c.Query("user_id")
	//userToken := c.Query("token")
	userInfoResp, err := QueryUserInfo(userIdStr)
	if err != nil {
		c.JSON(http.StatusOK, userInfoResp)
		return
	}
	c.JSON(http.StatusOK, userInfoResp)
}

func QueryUserInfo(userIdStr string) (*QueryUserInfoResp, error) {
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return &QueryUserInfoResp{
			Code: "-1",
			Msg:  err.Error(),
			User: nil,
		}, err
	}
	userInfo, err := user_service2.QueryUserInfo(int64(uint(userId)))
	if err != nil {
		return &QueryUserInfoResp{
			Code: "-1",
			Msg:  err.Error(),
			User: nil,
		}, err
	}
	return &QueryUserInfoResp{
		Code: "0",
		Msg:  "success",
		User: userInfo,
	}, nil
}

type LoginResp struct {
	Code   string `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func LoginUserHandler(c *gin.Context) {
	userName := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusOK, &LoginResp{
			Code: "-1",
			Msg:  "密码解析错误",
		})
	}
	loginInfo, err := user_service2.LoginUser(userName, password)
	if err != nil {
		c.JSON(http.StatusOK, &LoginResp{
			Code: "-1",
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &RegisterResp{
		Code:   "0",
		Msg:    "success",
		UserId: loginInfo.UserId,
		Token:  loginInfo.Token,
	})
}

type RegisterResp struct {
	Code   string `json:"status_code"`
	Msg    string `json:"status_msg"`
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func RegisterUserHandler(c *gin.Context) {
	userName := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusOK, &RegisterResp{
			Code: "-1",
			Msg:  "密码解析错误",
		})
	}
	registerInfo, err := user_service2.RegisterUser(userName, password)

	if err != nil {
		c.JSON(http.StatusOK, &RegisterResp{
			Code: "-1",
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &RegisterResp{
		Code:   "0",
		Msg:    "success",
		UserId: registerInfo.UserId,
		Token:  registerInfo.Token,
	})
}
