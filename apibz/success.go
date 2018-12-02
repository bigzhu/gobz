package apibz

// Result json 返回错误信息
type Result struct {
	Result string `json:"result"`
}

// Success 成功的结构体
var Success = Result{Result: "SUCCESS"}
