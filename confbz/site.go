package confbz

// SiteConf 网站 配置文件
type SiteConf struct {
	Secret string `toml:"secret"`
}

// GetSiteConf 获取数据库连接配置
func GetSiteConf() (siteConf SiteConf) {
	getConf("site", &siteConf)
	return
}
