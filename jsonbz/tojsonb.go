package jsonbz

import (
	"encoding/json"

	//"gorm.io/driver/postgres"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//ToJsonb 结构体转为 jsonb
func ToJsonb(v interface{}) (j postgres.Jsonb) {
	jsonByte, _ := json.Marshal(v)
	j = postgres.Jsonb{RawMessage: json.RawMessage(jsonByte)}
	return
}
