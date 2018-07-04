package libs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lomocoin/lomo_sdk/conf"
	"github.com/pkg/errors"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	// Dabot 大棒客机器人
	Dabot *TBot
)

// TBot 电报机器人
type TBot struct {
	botAPI *tgbotapi.BotAPI
}

// NewTBot 初始化一个电报机器人
func NewTBot(token string) (bot *TBot, err error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		err = errors.Wrap(err, "调用 NewBotAPI 时出错")
		return
	}
	bot = &TBot{
		botAPI: botAPI,
	}
	return
}

// TelegramSendMsgToChannel 机器人向群聊发送消息
func (b *TBot) TelegramSendMsgToChannel(chanID string, msg string) (err error) {
	if b == nil {
		err = errors.New("机器人还没有初始化")
		return
	}
	newMsg := tgbotapi.NewMessageToChannel(chanID, msg)
	newMsg.ParseMode = tgbotapi.ModeMarkdown
	_, err = b.botAPI.Send(newMsg)
	return
}

// InitDabot 初始化Dabot
func InitDabot() (err error) {
	Dabot, err = NewTBot(conf.DabankTBotToken)
	if err != nil {
		err = errors.Wrap(err, "Dabot启动失败")
		return
	}

	ips, ipErr := externalIP()
	if ipErr != nil {
		ips = []string{ipErr.Error()}
	}

	// 发送启动信息，确认信息可达
	msg := fmt.Sprintf(`** dabank服务启动啦 **

服务IP：%s

启动参数：%s

现在时间：%s

快 @Dabot 祝我顺顺利利！
（别试了，这个功能并没有开发 :P）
`, strings.Join(ips, ", "), fmt.Sprintf("`%s`", strings.Join(os.Args, " ")), time.Now().Format(time.RFC3339))

	err = Dabot.TelegramSendMsgToChannel(conf.DabankTelegramChannelID, msg)
	if err != nil {
		err = errors.Wrap(err, "Dabot启动信息发送失败")
		return
	}

	return
}
