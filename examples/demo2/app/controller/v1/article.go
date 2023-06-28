package v1

import (
	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	ginAutoRouter.Register(&Article{})
}

type Article struct{}

func (api *Article) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:List",
	})
}

func (api *Article) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:Test",
	})
}

func (api *Article) TestGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:Test",
	})
}
