package confbz

// DBConf 数据库配置文件
type DBConf struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	DBName   string `toml:"db_name"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
}

// GetDBConf 获取数据库连接配置
func GetDBConf() (dbConf DBConf) {
	GetConf("db", &dbConf)
	if dbConf.Port == "" {
		dbConf.Port = "5432"
	}
	return
}
