package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

// Route 控制器的一个方法将存储为一个 Route，具体数据格式为：
// ClassName:类名（控制器的对象名称，如：Admin）
// ActionName:方法名（控制器下的具体方法名称，如：ListGet）
// ActionObj:方法对象（方法的原始对象，不是方法名称）
// Args:方法的参数类型（方法包含的所有参数的类型）
// Ext:扩展参数
type Route struct {
	ClassName  string
	ActionName string
	ActionObj  reflect.Value
	Args       []reflect.Type
	Ext        ExtModel
}

// ExtModel Route 对象的扩展数据
// Nom:命名方法(nomenclature),支持: snake_case (下划线命名法，默认)、camelCase(驼峰命名法)和 PascalCase (帕斯卡命名法)
// HttpMethod:Http 请求方法，已支持：{"post", "get", "put", "patch", "head", "options", "delete", "any"}
// UrlClass: Url 路径 Path 部分，对应的是 controller 对象名称，如：/pathToClass/*****
// UrlAction:Url 路径 Action 部分，对应的是方法名称，如：/****/actionName
type ExtModel struct {
	Nom        string
	HttpMethod string
	UrlClass   string
	UrlAction  string
}

// Routes 存储所有方法的切片
var Routes = make([]Route, 0)

// Bind 执行路由绑定
func (r *Route) Bind(g *gin.RouterGroup) {
	r.Init()
	path := r.GetPath()
	handler := HandlerFunc(r.ActionObj)
	switch r.Ext.HttpMethod {
	case "get":
		g.GET(path, handler)
	case "put":
		g.PUT(path, handler)
	case "patch":
		g.PATCH(path, handler)
	case "head":
		g.HEAD(path, handler)
	case "options":
		g.OPTIONS(path, handler)
	case "delete":
		g.DELETE(path, handler)
	case "any":
		g.Any(path, handler)
	// Any registers a route that matches all the HTTP methods.
	// GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE.
	//case "post":    // The DEFAULT VALUE is "post".
	default:
		g.POST(path, handler)
	}
}

// Init 初始化参数
func (r *Route) Init() {
	// 字符串转为下划线分割，如：ListGet => list_get, infoPush => info_push
	urlAction := StringToSnakeCase(r.ActionName)

	// 分割 urlAction 字符串，取出最后一段，用来匹配请求类型，形如：list_get , info_push，则取出：get 和 push
	fields := strings.Split(urlAction, "_")
	httpMethod := fields[len(fields)-1]

	// 模式匹配上，则将 urlAction 的最后一段给去掉，即：将类似 "user_get" 变为 "user"
	if InArray(httpMethod, []string{"post", "get", "put", "patch", "head", "options", "delete", "any"}) {
		urlAction = urlAction[:len(urlAction)-len(httpMethod)-1]
		// TrimRight 耗时 是上面这种字符串截取耗时的大约 60 倍，能用上面这种就用这种
		//urlAction = strings.TrimRight(urlAction, "_"+httpMethod)
	} else {
		// 如果模式没有匹配上，则默认模式为：POST，并且 urlAction 保持原状
		httpMethod = "post"
	}
	r.Ext.HttpMethod = httpMethod
	r.Ext.UrlClass = r.GetStrByNom(StringToSnakeCase(r.ClassName))
	r.Ext.UrlAction = r.GetStrByNom(urlAction)
}

// GetPath 获取要绑定的路径
func (r *Route) GetPath() string {
	return "/" + r.Ext.UrlClass + "/" + r.Ext.UrlAction
}

// GetStrByNom 根据 Nom:命名方法(nomenclature) 获取字符串
// Nom 已支持：snake_case (下划线命名法，默认)、camelCase(驼峰命名法)和 PascalCase (帕斯卡命名法)
func (r *Route) GetStrByNom(str string) string {
	switch r.Ext.Nom {
	case "camelCase":
		return SnakeCaseToCamelCase(str)
	case "PascalCase":
		return SnakeCaseToPascalCase(str)
	//case "snake_case":
	default:
		return str
	}
}
