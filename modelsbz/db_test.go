package modelsbz

import (
	"testing"

	"github.com/bigzhu/gobz/conf"
)

func TestGetDBConf(t *testing.T) {
	dbConf := conf.GetDBConf()
	t.Log(dbConf)
}
