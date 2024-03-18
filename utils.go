package gin_auto_router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sort"
	"unicode"
)

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
