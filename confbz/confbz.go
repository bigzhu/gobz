package confbz

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/bigzhu/gobz/functionbz"
)

// 各种东西的配置文件读取

func getConfFileNameByStruct(v interface{}) string {
	confFileName := functionbz.GetStructName(v)
	confFileName = strings.ToLower(confFileName)
	confFileName = strings.Replace(confFileName, "conf", "", 1)
	return confFileName
}

// GetConf 获取配置文件
func GetConf(confbzName string, v interface{}) {
	codePath, err := functionbz.GetRunPath()
	if err != nil {
		panic(err)
	}
	confPath := fmt.Sprintf("conf/%s.toml", confbzName)
	// 尝试向上找10级
	for i := 0; i < 10; i++ {
		_, err = toml.DecodeFile(codePath+"/"+confPath, v)
		if err == nil {
			return
		}
		confPath = "../" + confPath
	}
	if err != nil {
		panic(err)
	}
}

// GetConfNew 根据结构体, 自动确定文件名
func GetConfNew(v interface{}) {
	confFileName := getConfFileNameByStruct(v)
	GetConf(confFileName, v)
}
