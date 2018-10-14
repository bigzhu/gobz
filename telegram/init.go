package telegram

import (
	"log"

	"github.com/bigzhu/gobz/confbz"
	"github.com/pkg/errors"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

//Bzbot 实例化
var Bzbot *tgbotapi.BotAPI
var telegramConf confbz.TelegramConf

func init() {
	telegramConf = confbz.GetTelegramConf()
	var err error
	Bzbot, err = tgbotapi.NewBotAPI(telegramConf.Token)
	if err != nil {
		err = errors.WithStack(err)
		log.Printf("%+v", err)
		return
	}
}
