package modelsbz

import (
	"time"
)

// Base 基本模型的定义
type Base struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// GetInstance 获取一条记录
func (base *Base) GetInstance() (result interface{}) {
	return base
}

//GetFirst 条件就作为结果返回
func (base *Base) GetFirst(i interface{}) (err error) {
	err = GetFirst(i.(Model), i)
	return
}
