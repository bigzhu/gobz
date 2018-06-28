package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm/dialects/postgres"
)

//ToJsonb 转为 jsonb
func ToJsonb(v interface{}) (j postgres.Jsonb) {
	jsonByte, _ := json.Marshal(v)
	j = postgres.Jsonb{RawMessage: json.RawMessage(jsonByte)}
	return
}
