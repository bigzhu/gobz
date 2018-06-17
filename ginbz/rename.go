package ginbz

import (
	"log"

	"github.com/bigzhu/gobz/inspectbz"
)

// RName router name 统一加上 API
func RName(handlers interface{}) string {
	log.Println("/API" + inspectbz.FName(handlers))
	return "/API" + inspectbz.FName(handlers)
}
