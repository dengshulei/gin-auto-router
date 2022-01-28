package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = make(map[string]interface{})
		var code = 200

		//需要自己实现一些验证逻辑
		adminId := c.PostForm("admin_id")
		if adminId == "" {
			code = 306
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg"  : "登录信息验证失败，请重新登录",
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}