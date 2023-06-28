package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

// Bind 绑定基本路由
func Bind(e *gin.Engine) {
	r := e.Group("/")
	baseBind(r)
}

// BindGroup 绑定路由组
func BindGroup(r *gin.RouterGroup) {
	baseBind(r)
}

// baseBind 执行路由绑定
func baseBind(r *gin.RouterGroup) {
	for class, value := range Routes {
		for method, v := range value {
			path := "/" + class + "/" + method

			r.POST(path, HandlerFunc(v))
		}
	}
}

// HandlerFunc 将控制器方法转为 gin.HandlerFunc 方法
func HandlerFunc(v Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		arguments := make([]reflect.Value, 1)
		arguments[0] = reflect.ValueOf(c)
		v.Method.Call(arguments)
	}
}
