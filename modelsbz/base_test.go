package modelsbz

import (
	"testing"
)

// Base 基本模型的定义
type Son struct {
	Base
	SonName string `json:"son_name"`
}

// GetInstance 获取一条记录
func (son *Son) GetInstance() (result interface{}) {
	return son
}

func TestBase_GetFirst(t *testing.T) {
	Connect()
	DB.AutoMigrate(
		&Base{},
		&Son{},
	)
	GetFirst(&Base{ID: 1})
	GetFirst(&Son{SonName: "kakazhu"})
}
