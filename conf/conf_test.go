package confbz

// 这里放置配置文件, 仅仅是为了测试

import (
	"testing"

	"github.com/bigzhu/gobz/confbz"
)

func TestConf(t *testing.T) {
	// dbConf := confbz.GetDBConf()
	// t.Log(dbConf)
	telegramConf := confbz.GetTelegramConf()
	t.Log(telegramConf)
}
