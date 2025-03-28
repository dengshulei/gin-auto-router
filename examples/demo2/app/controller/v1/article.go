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

func (api *Article) ListGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:List",
	})
}

func (api *Article) InfoGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoGet",
	})
}

func (api *Article) InfoPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoPost",
	})
}

func (api *Article) InfoDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoDelete",
	})
}
func (api *Article) InfoPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoPut",
	})
}
