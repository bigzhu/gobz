package telegram

import tgbotapi "gopkg.in/telegram-bot-api.v4"

// SendToChannel 发送到频道
func SendToChannel(msg string) (err error) {
	newMsg := tgbotapi.NewMessageToChannel(telegramConf.ChannelID, msg)
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
