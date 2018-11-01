package telegram

import (
	"errors"
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// SendToChannel 发送到频道
func SendToChannel(msg string, channelID string) (err error) {
	log.Println("SendToChannel")
	newMsg := tgbotapi.NewMessageToChannel(channelID, msg)
	newMsg.ParseMode = tgbotapi.ModeMarkdown
	_, err = Bzbot.Send(newMsg)
	return
}

// SendToUser 发送给某个用户
// 要拿到 chatID
func SendToUser(msg string, chatID int64) (err error) {
	newMsg := tgbotapi.NewMessage(chatID, msg)
	newMsg.ParseMode = tgbotapi.ModeMarkdown
	_, err = Bzbot.Send(newMsg)
	return
}

// Send 默认发送
// 优先看是否配置了 chatID, 没有再找 channelID
func Send(msg string) (err error) {
	switch {
	case telegramConf.ChatID != 0:
		err = SendToUser(msg, telegramConf.ChatID)
	case telegramConf.ChannelID != "":
		err = SendToChannel(msg, telegramConf.ChannelID)
	default:
		err = errors.New("telegram 配置文件缺少ChatID 和 ChannelID, 无法默认发送")
	}
	return
}
