package callback

import (
	"time"

	"github.com/bigzhu/gobz/modelsbz"
)

// CallbackModel 记录 callback 状态
type CallbackModel struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Request    string
	Response   string
	URL        string
	Count      int    // 调用次数
	Status     string //Success or Failure or NoNeed
	StatusCode int
	ErrorInfo  string
	Type       string // 考虑兼容多种回调, 根据 Type 做不同操作
	Key        string `sql:"type:JSONB;"` // 不同Type需要存放一些内容以用于重发取数据, 比如 transferID
}

// GetCallbackByTypeAndKey 查询回调记录
func GetCallbackByTypeAndKey(key string, callbackType string) (callback Callback, err error) {
	err = modelsbz.DB.
		Where("key = ? and type=?", key, callbackType).
		First(&callback).
		Error
	return
}

// GetFailureIntervalCallBacks 查出失败的回调记录
func GetFailureIntervalCallBacks() (callbacks []Callback, err error) {
	err = odelsbz.DB.
		Where("status = ? ", libs.Failure).
		// 次数为n, 2的n次方为秒, 间隔这么长时间再调用
		Where("(updated_at + 2^count * interval '1 second') < now()").
		Find(&callBacks).
		Error
	return
}
