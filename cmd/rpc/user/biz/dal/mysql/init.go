package mysql

import (
	"database/sql"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
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
	sqlDB, err = DB.DB()
	if err != nil {
		panic(err)
	}

	// just check if the connection is maintaining
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
