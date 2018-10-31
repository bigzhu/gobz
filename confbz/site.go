package confbz

// SiteConf 网站 配置文件
type SiteConf struct {
	Secret     string `toml:"secret"`
	CookieSalt string `toml:"cookie_salt"`
}

// GetSiteConf 获取数据库连接配置
func GetSiteConf() (siteConf SiteConf) {
	GetConf("site", &siteConf)
	return
}
