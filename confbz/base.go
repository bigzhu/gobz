package confbz

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// GetConf 获取配置文件
func GetConf(confbzName string, v interface{}) {
	var err error
	codePath, err := getRunPath()
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
	return
}
func getRunPath() (path string, err error) {
	path, err = os.Getwd()
	log.Println(path)
	return
}
