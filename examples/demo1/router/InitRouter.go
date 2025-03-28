package router

import (
	_ "demo1/app/controller" //一定要导入这个Controller包，用来注册需要访问的方法
	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//初始化路由
	r := gin.Default()

	//自动绑定 驼峰命名法 路由
	ginAutoRouter.Bind(r, "camelCase")

	return r
}
