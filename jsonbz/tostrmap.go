package jsonbz

import "encoding/json"

// ToStrMap 以 json 为媒介,转化为 string map
// 主要用于签名
func ToStrMap(s interface{}) (strMap map[string]string, err error) {
	sJSON, err := json.Marshal(s)
	if err != nil {
		return
	}
	json.Unmarshal(sJSON, &strMap)
	return
}
