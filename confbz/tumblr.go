package confbz

// TumblrConf 配置文件
type TumblrConf struct {
	ConsumerKey    string `toml:"consumer_key"`
	ConsumerSecret string `toml:"consumer_secret"`
}

// GetTumblrConf tumblr 配置
func GetTumblrConf() (tumblrConf TumblrConf) {
	GetConfNew(&tumblrConf)
	return
}
