package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// StringToSnakeCase  字符串转为 snake_case (下划线命名法)
// 如：ListGet => list_get;  infoPush => info_push;  code => code;  Setting => setting
func StringToSnakeCase(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			// 首字母小写
			output = append(output, unicode.ToLower(r))
			continue
		}
		if r == 95 {
			// 跳过下划线"_"
			continue
		}
		if unicode.IsUpper(r) {
			// 如果是大写字母，就加一条下划线
			output = append(output, '_')
		}
		// 字符转小写
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// SnakeCaseToCamelCase snake_case (下划线命名法) 转 camelCase(驼峰命名法)
// 如：list_get => listGet;  info_push => infoPush;  code => code;  setting => setting
func SnakeCaseToCamelCase(s string) string {
	return SnakeCaseToOther(s, false)
}

// SnakeCaseToPascalCase snake_case (下划线命名法) 转 PascalCase (帕斯卡命名法)
// 如：list_get => ListGet;  info_push => InfoPush;  code => Code;  setting => Setting
func SnakeCaseToPascalCase(s string) string {
	return SnakeCaseToOther(s, true)
}

// SnakeCaseToOther snake_case (下划线命名法) 转 其他命名法
// s: 需要转变的字符串
// lastFlag: 值为 true ，首字母大写(PascalCase帕斯卡命名法)；值为 false 首字母小写（camelCase驼峰命名法）
func SnakeCaseToOther(s string, lastFlag bool) string {
	var output []rune
	for _, r := range s {
		if r == 95 {
			// 下划线"_"，标记并跳过
			lastFlag = true
			continue
		}
		if lastFlag == true {
			// 上一个字符是下划线，则本字符转大写
			output = append(output, unicode.ToUpper(r))
			lastFlag = false
			continue
		}
		output = append(output, r)
	}
	return string(output)
}

// InArray 判断此字符串是否存在于给定的数组中
func InArray(str string, strArray []string) bool {
	// 给数组排序
	sort.Strings(strArray)
	// 二分法查找，找到以后返回索引
	index := sort.SearchStrings(strArray, str)
	if index < len(strArray) && strArray[index] == str {
		return true
	}
	return false
}

// GetNom 获取命名方法(nomenclature),支持: snake_case (下划线命名法，默认)、camelCase(驼峰命名法)和 PascalCase (帕斯卡命名法)
func GetNom(args ...string) (nom string) {
	nom = GetArg(0, args...)
	switch nom {
	case "camelCase":
	case "PascalCase":
	default:
		nom = "snake_case"
	}
	return nom
}

// GetArg 在一个批量 args 里面获取指定的 arg
// i : 要获取的 arg 所在的索引，从 0 开始计算
func GetArg(i int, args ...string) (arg string) {
	if len(args) > i {
		arg = args[i]
	}
	return
}

// GetClass 获取控制器对象的名称（类名）
func GetClass(controller interface{}) (class string) {
	// 取得控制器的类型（reflect.Type），结果类似“*controller.Article”
	class = reflect.TypeOf(controller).String()
	if strings.Contains(class, ".") {
		// 此时结果类似“Article”
		class = class[strings.Index(class, ".")+1:]
	}
	return
}

// HandlerFunc 将控制器方法转为 gin.HandlerFunc 方法
func HandlerFunc(v reflect.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		arguments := make([]reflect.Value, 1)
		arguments[0] = reflect.ValueOf(c)
		v.Call(arguments)
	}
}
