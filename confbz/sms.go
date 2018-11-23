package confbz

// HengTongConf 配置文件
type HengTongConf struct {
	Tag       string `toml:"tag"`       // 模板标签
	UserID    string `toml:"userid"`    // 账户名
	Password  string `toml:"password"`  // 密码
	CodeType  string `toml:"codetype"`  //
	EncodeWay string `toml:"encodeway"` // 编码方式
}

// YunPianConf 云片的配置
type YunPianConf struct {
	Tag    string `toml:"tag"`     // 模板标签
	APIKey string `toml:"api_key"` // 用户唯一标识
	TplID  string `toml:"tplID"`   // 模板ID
}

// GetHengTongConf telegram 配置
func GetHengTongConf() (conf HengTongConf) {
	GetConf("hengtong", &conf)
	return
}

// GetYunPianConf 获取云片配置
func GetYunPianConf() (conf YunPianConf) {
	GetConf("yunpian", &conf)
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
