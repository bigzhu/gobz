package confbz

// OauthBase 配置文件, 通用基础的
type OauthBase struct {
	RedirectURL  string `toml:"redirect_url"`
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}

// Oauth 读取配置文件
type Oauth struct {
	Google OauthBase
	GitHub OauthBase
}

// GetOauthConf oauth 配置
func GetOauthConf() (oauth Oauth) {
	GetConfNew(&oauth)
	return
}
