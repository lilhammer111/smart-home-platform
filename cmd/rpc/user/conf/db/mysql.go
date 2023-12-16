package db

import (
	"database/sql"
	"fmt"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	gormDB    *gorm.DB
	err       error
	mysqlOnce sync.Once
)

func GetMysql() *gorm.DB {
	// here once is from redis.go
	mysqlOnce.Do(initMysql)
	return gormDB
}

func initMysql() {
	DSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		binding.GetRemoteConf().Mysql.Username,
		binding.GetRemoteConf().Mysql.Password,
		binding.GetRemoteConf().Mysql.Addr,
		binding.GetRemoteConf().Mysql.DB,
	)
	gormDB, err = gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// the underlying sqlDB object
	var sqlDB *sql.DB
	sqlDB, err = gormDB.DB()
	if err != nil {
		panic(err)
	}

	// just check if the connection is maintaining
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
