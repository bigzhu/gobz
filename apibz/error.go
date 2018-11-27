package apibz

import "fmt"

// Error  json 返回错误信息
type Error struct {
	Error string `json:"error"`
}

// NewE 新建错误类型
func NewE(errIn interface{}) (errOut Error) {
	errOut.Error = fmt.Sprintf("%+v", errIn)
	return
}
