package confbz

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func getConf(confbzName string, v interface{}) {
	_, err := toml.DecodeFile(fmt.Sprintf("confbz/%s.toml", confbzName), v)
	if err != nil {
		_, err := toml.DecodeFile(fmt.Sprintf("../confbz/%s.toml", confbzName), v)
		if err != nil {
			panic(err)
		}
	}
	return
}

// DBConf 数据库配置文件
type DBConf struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	DBName   string `toml:"db_name"`
	Password string `toml:"password"`
}

// GetDBConf 获取数据库连接配置
func GetDBConf() (dbConf DBConf) {
	getConf("db", &dbConf)
	return
}

// TwitterConf twitter 配置文件
type TwitterConf struct {
	ConsumerKey       string `toml:"consumer_key"`
	ConsumerSecret    string `toml:"consumer_secret"`
	AccessToken       string `toml:"access_token"`
	AccessTokenSecret string `toml:"access_token_secret"`
}

// GetTwitterConf 获取数据库连接配置
func GetTwitterConf() (twitterConf TwitterConf) {
	getConf("twitter", &twitterConf)
	return
}

// SiteConf 网站 配置文件
type SiteConf struct {
	Secret string `toml:"secret"`
}

// GetSiteConf 获取数据库连接配置
func GetSiteConf() (siteConf SiteConf) {
	getConf("site", &siteConf)
	return
}

// TumblrConf 配置文件
type TumblrConf struct {
	ConsumerKey    string `toml:"consumer_key"`
	ConsumerSecret string `toml:"consumer_secret"`
}

// GetTumblrConf tumblr 配置
func GetTumblrConf() (tumblrConf TumblrConf) {
	getConf("tumblr", &tumblrConf)
	return
}
