package dal

import (
	"git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/biz/dal/mysql"
	"git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
