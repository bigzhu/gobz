package confbz

// TwitterConf twitter 配置文件
type TwitterConf struct {
	ConsumerKey       string `toml:"consumer_key"`
	ConsumerSecret    string `toml:"consumer_secret"`
	AccessToken       string `toml:"access_token"`
	AccessTokenSecret string `toml:"access_token_secret"`
}

// GetTwitterConf 获取数据库连接配置
func GetTwitterConf() (twitterConf TwitterConf) {
	GetConfNew(&twitterConf)
	return
}
