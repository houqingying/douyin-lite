package user

import (
	"douyin-lite/service/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	registerInfo, err := user_service.RegisterUser(userName, password)

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
