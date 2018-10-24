package confbz

// HengTongConf 配置文件
type HengTongConf struct {
	Tag       string `toml:"tag"`       // 模板标签
	UserID    string `toml:"userid"`    // 账户名
	Password  string `toml:"password"`  // 密码
	CodeType  string `toml:"codetype"`  //
	EncodeWay string `toml:"encodeway"` // 编码方式
}

// GetHengTongConf telegram 配置
func GetHengTongConf() (conf HengTongConf) {
	GetConf("hengtong", &conf)
	return
}

// SMSConf 短信配置文件
type SMSConf struct {
	ID       string `toml:"id"`
	password string `toml:"password"`
}

// GetSMSConf telegram 配置
func GetSMSConf() (conf SMSConf) {
	GetConf("sms", &conf)
	return
}
