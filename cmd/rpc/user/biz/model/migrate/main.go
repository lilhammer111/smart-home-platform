package main

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/db"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	err := db.GetMysql().AutoMigrate(&model.User{})
	if err != nil {
		klog.Errorf("failed to auto migrate user table: %s", err)
	}
}
