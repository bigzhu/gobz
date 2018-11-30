package functionbz

import "reflect"

// GetStructName 获取结构体的名称
// 无论传入指针还是实体, 都能正确取到结构体名称
func GetStructName(s interface{}) string {
	val := reflect.Indirect(reflect.ValueOf(s))
	return val.Type().Name()
}
