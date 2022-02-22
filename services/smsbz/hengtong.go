package smsbz

import (
	"encoding/base64"
	"net/url"

	"github.com/bigzhu/gobz/confbz"
	"github.com/bigzhu/gobz/httpbz"
)

// Send 发送消息
func Send(phone string, msg string) (result string, err error) {
	hengTongConf := confbz.GetHengTongConf()
	apiURL := "http://api.hengtongexun.com:8090/protocol/sendMessage.do"
	form := make(url.Values)
	form["userid"] = []string{hengTongConf.UserID}
	form["password"] = []string{hengTongConf.Password}
	form["codetype"] = []string{hengTongConf.CodeType}
	form["encodeway"] = []string{hengTongConf.EncodeWay}

	form["phone"] = []string{phone}
	base64Msg := base64.StdEncoding.EncodeToString([]byte(hengTongConf.Tag + msg))
	form["message"] = []string{base64Msg}

	result, _, err = httpbz.PostForm(apiURL, form, nil)
	if err != nil {
		return
	}
	return
}
