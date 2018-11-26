package modelsbz

import (
	"time"
)

// Base 基本模型的定义
type Base struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetInstance 获取一条记录
func (base *Base) GetInstance() (result interface{}) {
	return base
}

//Get 条件就作为结果返回, 只返回一条
func (base *Base) Get(i interface{}) (err error) {
	err = GetFirst(i.(Model), i)
	return
}
