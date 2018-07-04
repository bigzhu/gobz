package confbz

// Oauth 配置文件
type Oauth struct {
	RedirectURL  string `toml:"redirect_url"`
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}

// GetOauthConf tumblr 配置
func GetOauthConf() (oauth Oauth) {
	getConf("oauth", &oauth)
	return
}
