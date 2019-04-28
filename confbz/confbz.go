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
	err = loopFind(codePath, confPath, v)
	// 尝试不要 conf 子路径查找

	if err != nil {
		confPath = fmt.Sprintf("%s.toml", confbzName)
		err = loopFind(codePath, confPath, v)
	}
	if err != nil {
		panic(err)
	}
}

func loopFind(codePath string, confPath string, v interface{}) (err error) {
	// 尝试向上找10级
	for i := 0; i < 10; i++ {
		_, err = toml.DecodeFile(codePath+"/"+confPath, v)
		if err == nil {
			return
		}
		// 不是找不到文件的报错, 就直接退出
		if !strings.Contains(err.Error(), "no such file or directory") {
			return
		}

		confPath = "../" + confPath
	}
	return
}

// GetConfNew 根据结构体, 自动确定文件名
func GetConfNew(v interface{}) {
	confFileName := getConfFileNameByStruct(v)
	GetConf(confFileName, v)
}
