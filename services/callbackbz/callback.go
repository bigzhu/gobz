// Package callback 针对一些需要 http 反复回调, 并且需要记录回调结果的情况
package callback

import (
	"log"
	"net/http"

	"github.com/bigzhu/gobz/httpbz"
	"github.com/bigzhu/gobz/modelsbz"
	"github.com/bigzhu/gobz/telegram"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	//Success 成功
	Success = "Success"
	//Failure callback 失败
	Failure = "Failure"
	// NoNeed 不再需要处理, 回调那里
	NoNeed = "NoNeed"
)

// Callback 执行回调, 并保存结果. 必须包含以下两个参数
// callback.URL 地址
// callback.Request 请求参数
func Callback(callback CallbackModel) {
	log
	response, statusCode, err := httpbz.Post(callback.URL, callback.Request, nil)
	if err != nil {
		callback.ErrorInfo = err.Error()
		err = nil
	}
	callback.StatusCode = statusCode
	callback.Response = response

	// 如果返回 200 就姑且认为回调成功
	if callback.StatusCode == http.StatusOK {
		callback.Status = Success
	} else {
		callback.Status = Failure
	}
	// 保存
	callbackData, err := GetCallbackByTypeAndKey(callback.Key, callback.Type)
	if err == nil {
		callback.Count = callbackData.Count + 1
		// 更新错误信息
		modelsbz.DB.Model(&callbackData).Updates(callback)
	} else if err == gorm.ErrRecordNotFound {
		modelsbz.DB.Create(&callback)
		err = nil
	} else {
		err = errors.Wrap(err, "查询回调列表时出错")
		telegram.Send(err.Error())
		log.Printf("%+v", err)
	}
	return
}

// RunFailureCallbacks 回调所有的之前失败的回调记录
func RunFailureCallbacks() (err error) {
	callbacks, err := GetFailureIntervalCallBacks()
	if err != nil {
		return
	}
	for _, callback := range callbacks {
		Callback(callback)
	}
	return
}
