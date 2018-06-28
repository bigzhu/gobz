package modelsbz

import (
	"fmt"
	"log"

	"github.com/bigzhu/gobz/confbz"
	"github.com/jinzhu/gorm"
	// 必须导入
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB db connect
var DB *gorm.DB

func init() {
	var err error
	DB, err = GetDB()
	if err != nil {
		panic(err)
	}
}

// TX my tx
type TX struct {
	// DB 原本的DB对象
	*gorm.DB
	// fired 标记事务是否已经commit或rollback
	fired bool
}

// Rollback 更稳定而流畅的回滚
func (tx *TX) Rollback() *TX {
	// 如果事务已经commit或rollback过，则静默返回
	if tx.fired {
		return tx
	}

	var err error
	defer func() {
		// 无条件记录触发
		tx.fired = true
		// 如果tx.DB无效（为nil或其他奇怪的值），会panic
		// 这里recover会做捕捉
		if recoverErr := recover(); recoverErr != nil {
			log.Printf("Recovered from panic，事务无法回滚： %v", recoverErr)
			return
		}
	}()

	err = tx.DB.Rollback().Error
	if err != nil {
		log.Printf("事务无法回滚： %v", err)
	}
	return tx
}

// Commit 更稳定而流畅的提交
func (tx *TX) Commit() *TX {
	// 如果事务已经commit或rollback过，则静默返回
	if tx.fired {
		return tx
	}

	var err error
	defer func() {
		// 无条件记录触发
		tx.fired = true
		// 如果tx.DB无效（为nil或其他奇怪的值），会panic
		// 这里recover会做捕捉
		recoverErr := recover()
		if recoverErr != nil {
			log.Printf("Recovered from panic，事务无法提交： %v", recoverErr)
			return
		}
	}()

	err = tx.DB.Commit().Error
	if err != nil {
		log.Printf("事务无法提交： %v", err)
	}
	return tx
}

// GetTx get the tx
func GetTx() *TX {
	tx := &TX{
		DB:    DB.Begin(),
		fired: false,
	}
	return tx
}

//GetDB connect to postgresql
func GetDB() (*gorm.DB, error) {
	dbConf := confbz.GetDBConf()
	str := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbConf.Host, dbConf.User, dbConf.DBName, dbConf.Password)
	DB, err := gorm.Open("postgres", str)
	return DB, err
}
