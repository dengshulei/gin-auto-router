package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data = make(map[string]interface{})
		var code = 200

		// You need to implement some verification logic yourself
		adminId := c.PostForm("admin_id")
		if adminId == "" {
			code = 306
		}

		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "Login verification failed, please login again",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
