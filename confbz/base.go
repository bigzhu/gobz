package confbz

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func getConf(confbzName string, v interface{}) {
	confPath := fmt.Sprintf("conf/%s.toml", confbzName)
	var err error
	// 尝试向上找10级
	for i := 0; i < 10; i++ {
		_, err = toml.DecodeFile(confPath, v)
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
