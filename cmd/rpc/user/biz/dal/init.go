package dal

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user_srv/biz/dal/mysql"
	"git.zqbjj.top/pet/services/cmd/rpc/user_srv/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
