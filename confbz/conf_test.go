package confbz

// 这里放置配置文件, 仅仅是为了测试

import (
	"testing"
)

func TestConf(t *testing.T) {
	dbConf := GetDBConf()
	t.Log(dbConf)
	telegramConf := GetTelegramConf()
	t.Log(telegramConf)
	oauthConf := GetOauthConf()
	t.Log(oauthConf)
}
