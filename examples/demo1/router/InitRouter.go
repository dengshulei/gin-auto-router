package router

import (
	_ "demo1/app/controller" //一定要导入这个Controller包，用来注册需要访问的方法
	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//初始化路由
	r := gin.Default()

	//绑定基本路由，访问路径：/article/list
	ginAutoRouter.Bind(r)

	return r
}
