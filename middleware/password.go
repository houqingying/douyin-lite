package middleware

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func SHA1(s string) string {

	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}

func SHAMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Query("password")
		if password == "" {
			password = c.PostForm("password")
		}
		c.Set("password", SHA1(password))
		c.Next()
	}
}
