package modelsbz

import (
	"testing"

	"github.com/bigzhu/gobz/confbz"
)

func TestGetDBConf(t *testing.T) {
	dbConf := confbz.GetDBConf()
	t.Log(dbConf)
}
