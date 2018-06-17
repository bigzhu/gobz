package inspectbz

import (
	"reflect"
	"runtime"
	"strings"
)

// FName 获取函数名称
// func FName(handlers gin.HandlerFunc) string {
func FName(f interface{}) string {
	pathName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	splitName := strings.Split(pathName, ".")
	name := splitName[len(splitName)-1]
	return name
}
