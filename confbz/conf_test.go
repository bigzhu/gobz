package confbz

import "testing"

func TestConf(t *testing.T) {
	dbConf := GetDBConf()
	t.Log(dbConf)
}
