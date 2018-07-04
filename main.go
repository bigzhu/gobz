package main

import (
	"fmt"
	"log"

	"github.com/bigzhu/gobz/confbz"
)

func main() {
	dbConf := confbz.GetDBConf()
	log.Println(dbConf)
	fmt.Println("bigzhu's lib")
}
