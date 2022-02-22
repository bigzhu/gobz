package modelsbz

import (
	"testing"

	"github.com/bigzhu/gobz/confbz"
	"github.com/pkg/errors"
)

// Result 条目数
type Result struct {
	Count int
}

func TestGetDBConf(t *testing.T) {
	var result Result
	dbConf := confbz.GetDBConf()
	t.Log(dbConf)
	Connect()
	sql := `
	select 4 as count
	`
	err := DB.Raw(sql).Scan(&result).Error
	if err != nil {
		err = errors.WithStack(err)
		t.Fatal(err)
	}
	t.Log(result)
}
