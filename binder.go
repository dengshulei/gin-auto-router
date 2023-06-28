package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"unicode"
)

// Bind 绑定基本路由，外部可以直接使用
func Bind(e *gin.Engine) {
	r := e.Group("/")
	BindGroup(r)
}

// BindGroup 绑定路由组，外部可以直接使用
func BindGroup(r *gin.RouterGroup) {
	for class, value := range Routes {
		for action, v := range value {
			baseBind(r, class, action, HandlerFunc(v))
		}
	}
}

// baseBind 执行路由绑定
func baseBind(r *gin.RouterGroup, class string, action string, handler gin.HandlerFunc) {
	// 驼峰方式全部转为下划线分割，如：ListGet => list_get, InfoPush => info_push
	class = CamelCaseToUnderscore(class)
	action = CamelCaseToUnderscore(action)
	path := "/" + class + "/" + action

	// action 中不包含下划线"_"，直接匹配 POST模式
	if !strings.Contains(action, "_") {
		r.POST(path, handler)
		return
	}
	// 分割 action字符串，取出最后一段，用来匹配请求类型，形如：list_get , info_push，则取出：get 和 push
	fields := strings.Split(action, "_")
	method := fields[len(fields)-1]
	path = strings.Replace(path, "_"+method, "", 1)
	switch method {
	case "get":
		r.GET(path, handler)
	case "put":
		r.PUT(path, handler)
	case "patch":
		r.PATCH(path, handler)
	case "head":
		r.HEAD(path, handler)
	case "options":
		r.OPTIONS(path, handler)
	case "delete":
		r.DELETE(path, handler)
	case "any":
		r.Any(path, handler)
	// Any registers a route that matches all the HTTP methods.
	// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
	//case "post":    // The DEFAULT VALUE is "post".
	default:
		r.POST(path, handler)
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

// CamelCaseToUnderscore 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
