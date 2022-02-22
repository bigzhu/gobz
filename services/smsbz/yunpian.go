package smsbz

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"gobz/confbz"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

var (
	isAlphanumeric = regexp.MustCompile(`^[[:alnum:]]+$`).MatchString
)

// SendIntl 发送国际短信.手机号码必须包含国际地区前缀号码.
func SendIntl(phone string, msg string) (result *ypclnt.Result, err error) {
	// 排除国内手机号. 只当做ascii处理应该是安全的
	if len(phone) < 4 || phone[:1] != "+" || phone[:3] == "+86" {
		err = fmt.Errorf("%s不是有效的国际手机号", phone)
		return
	}
	// 模板目前只支持验证码,传别的内容调用接口会报错
	if !isAlphanumeric(msg) {
		err = errors.New("国际短信目前只支持发送验证码")
		return
	}
	conf := confbz.GetYunPianConf()
	client := ypclnt.New(conf.APIKey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phone
	param[ypclnt.TEXT] = msg
	param[ypclnt.TPL_ID] = conf.TplID
	tplValue := url.Values{
		//"#company#": []string{conf.Tag},
		"#code#": []string{msg},
	}
	param[ypclnt.TPL_VALUE] = tplValue.Encode()
	if result = client.Sms().TplSingleSend(param); !result.IsSucc() {
		err = errors.New(result.String())
		return
	}
	return
}
