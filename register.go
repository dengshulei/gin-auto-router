package gin_auto_router

import (
	"reflect"
	"strings"
)

// Route 数据格式为：
// Method:控制器方法对象
// Args:方法的参数
type Route struct {
	Method reflect.Value
	Args   []reflect.Type
}

// Routes 存储的数据格式为：Routes[控制器名称][方法名称]Route对象
var Routes = make(map[string]map[string]Route)

// Register 注册路由方法
func Register(controller interface{}) bool {
	// 取得控制器的 reflect.Value，结果类似：&{}
	v := reflect.ValueOf(controller)
	// 取得控制器中导出的方法的数量，非控制器或无方法则直接返回，结果是个数字，比如：2
	if v.NumMethod() == 0 {
		return false
	}

	// 取得控制器的 reflect.Type，结果类似“*controller.Article”
	t := reflect.TypeOf(controller).String()
	// 设置控制器类名
	module := t
	if strings.Contains(t, ".") {
		// 此时结果类似“Article”
		module = t[strings.Index(t, ".")+1:]
	}

	// 遍历控制器的所有方法
	for i := 0; i < v.NumMethod(); i++ {
		// 获取第 i 个方法
		method := v.Method(i)

		// 获取第 i 个方法的名称
		action := v.Type().Method(i).Name

		// 遍历方法中传递过来的参数，保存到 params
		params := make([]reflect.Type, 0)
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}

		if Routes[module] == nil {
			// 新控制类，则初始化
			Routes[module] = make(map[string]Route)
		}
		Routes[module][action] = Route{method, params}
	}
	return true
}
