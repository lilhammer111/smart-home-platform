package main

import (
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	err := db.GetMysql().AutoMigrate(&model.Alert{})
	if err != nil {
		klog.Errorf("failed to migrate: %s", err)
	}

	klog.Info("migration succeed")
}
