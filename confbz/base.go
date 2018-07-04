package confbz

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

func getConf(confbzName string, v interface{}) {
	_, err := toml.DecodeFile(fmt.Sprintf("./conf/%s.toml", confbzName), v)
	if err != nil {
		log.Println(err)
		_, err := toml.DecodeFile(fmt.Sprintf("../conf/%s.toml", confbzName), v)
		if err != nil {
			panic(err)
		}
	}
	return
}
