package main

import (
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	err := db.GetMysql().AutoMigrate(&model.Model{})
	if err != nil {
		klog.Errorf("failed to auto migrate user table: %s", err)
	}

	klog.Info("migration succeed")
}
