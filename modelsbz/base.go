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
