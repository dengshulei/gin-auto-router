package gin_auto_router

import (
	"reflect"
	"strings"
)

type Route struct {
	Method reflect.Value
	Args   []reflect.Type
}

var Routes = make(map[string]map[string]Route)

// Register 注册路由方法
func Register(controller interface{}) bool {
	v := reflect.ValueOf(controller)
	//非控制器或无方法则直接返回
	if v.NumMethod() == 0 {
		return false
	}

	//取得控制器名称
	tmp := reflect.TypeOf(controller).String()
	module := tmp
	if strings.Contains(tmp, ".") {
		module = tmp[strings.Index(tmp, ".")+1:]
	}

	//遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name

		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}

		if Routes[module] == nil {
			Routes[module] = make(map[string]Route)
		}
		Routes[module][action] = Route{method, params}
	}
	return true
}
