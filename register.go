package gin_auto_router

import (
	"reflect"
)

// Register 注册路由方法
func Register(controller interface{}) bool {
	// 通过反射，取得控制器的对象（reflect.Value），结果类似：&{}
	v := reflect.ValueOf(controller)
	// 取得控制器中导出的方法的数量，非控制器或无方法则直接返回，结果是个数字，比如：2
	if v.NumMethod() == 0 {
		return false
	}
	// 遍历控制器的所有方法
	for i := 0; i < v.NumMethod(); i++ {
		// 获取第 i 个方法（方法的原始对象，不是方法名称）
		method := v.Method(i)
		// 遍历方法中传递过来的参数，保存到 args
		args := make([]reflect.Type, 0)
		for j := 0; j < method.Type().NumIn(); j++ {
			args = append(args, method.Type().In(j))
		}
		// v.Type().Method(i).Name:获取第 i 个方法的名称
		Routes = append(Routes, Route{
			ClassName:  GetClass(controller),
			ActionName: v.Type().Method(i).Name,
			ActionObj:  method,
			Args:       args,
			Ext:        ExtModel{},
		})
	}
	return true
}
