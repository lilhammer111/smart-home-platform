package main

import (
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/suite"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device/microdevice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := microdevice.NewServer(
		&biz.MicroDeviceImpl{},
		server.WithSuite(&suite.BasicSuite{}),
	)

	err := svr.Run()
	if err != nil {
		klog.Info("############################   failed to run server    ############################\n")
		klog.Errorf("############################   %s   ############################", err)
	}
}
