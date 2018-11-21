package apibz

// ErrorBz  json 返回错误信息
type ErrorBz struct {
	ErrCode string      `json:"err_code"`
	Payload interface{} `json:"payload"`
}

// NewErr 新建错误
func NewErr(code string) *ErrorBz {
	return &ErrorBz{ErrCode: code}
}
