package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"strings"
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

	// 模式匹配上，则将 path 中 action 的最后一段给去掉
	// 即：将类似 "/admin/user_get" 变为 "/admin/user"
	if InArray(method, []string{"post", "get", "put", "patch", "head", "options", "delete", "any"}) {
		path = "/" + class + "/" + action[:len(action)-len(method)-1]
		// TrimRight 耗时 是上面这种字符串截取耗时的大约 60 倍，能用上面这种就用这种
		//path = "/" + class + "/" + strings.TrimRight(action, "_"+method)
	} else {
		// 如果模式没有匹配上，则默认模式为：POST，并且 url 保持原状
		method = "post"
	}

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
