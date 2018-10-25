package confbz

// EmailConf 统一各种邮件服务的配置
type EmailConf struct {
	User       string   `toml:"user"`     // 账户名
	Password   string   `toml:"password"` // 密码
	Host       string   `toml:"host"`     // host
	Port       int      `toml:"port"`     // port
	AssetPaths []string `toml:assetPaths` // 邮件模板存放路径
}

// GetLomoMailConf 获取lomocoin邮件服务配置
func GetEmailConf() (conf EmailConf) {
	GetConf("email", &conf)
	return
}
