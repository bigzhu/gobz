package modelsbz

import (
	"log"

	"github.com/pkg/errors"
)

func unlogged(tableNames ...string) (err error) {
	for _, tableName := range tableNames {
		err = DB.Exec("alter table " + tableName + " set unlogged").Error
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}
	return
}

func getTableName(model interface{}) (tableName string) {
	return DB.NewScope(model).GetModelStruct().TableName(DB)
}

// SetUnlogged 将表设置为unlogged类型, 用于缓存
func SetUnlogged(models ...interface{}) (err error) {
	for k, theModel := range models {
		log.Println(k)
		tableName := getTableName(theModel)
		log.Println(theModel)
		log.Println("table name=", tableName)
		err = DB.Exec("alter table " + tableName + " set unlogged").Error
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}
	return
}
