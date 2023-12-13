package dal

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/dal/mysql"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
