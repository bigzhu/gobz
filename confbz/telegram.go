package confbz

// TelegramConf 配置文件
type TelegramConf struct {
	Token     string `toml:"token"`
	ChannelID string `toml:"channel_id"`
}

// GetTelegramConf telegram 配置
func GetTelegramConf() (telegramConf TelegramConf) {
	getConf("telegram", &telegramConf)
	return
}
