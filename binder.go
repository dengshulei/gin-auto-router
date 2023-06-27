package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

var paths = make(map[int]string)

// Bind 绑定基本路由
func Bind(e *gin.Engine) {
	pathInit()
	for _, path := range paths {
		e.POST(path, match(path))
	}
}

// BindGroup 绑定路由组
func BindGroup(r *gin.RouterGroup) {
	pathInit()
	for _, path := range paths {
		r.POST(path, match(path))
	}
}

// pathInit 组装path
func pathInit() {
	i := 0
	for class, value := range Routes {
		for method := range value {
			path := "/" + class + "/" + method
			paths[i] = path
			i += 1
		}
	}
}

// match 根据path匹配对应的方法
func match(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		if len(fields) < 3 {
			return
		}
		v, ok := Routes[fields[1]][fields[2]]
		if ok {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			v.Method.Call(arguments)
		}
	}
}
