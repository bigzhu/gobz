package apibz

// ErrorBz  json 返回错误信息
type ErrorBz struct {
	ErrCode string      `json:"err_code"`
	Payload interface{} `json:"payload"`
}
