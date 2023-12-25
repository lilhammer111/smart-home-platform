package main

import (
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/suite"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user/microuser"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	klog.Info("klog debugf works")
	svr := microuser.NewServer(
		&biz.MicroUserImpl{},
		server.WithSuite(&suite.BasicSuite{
			//MicroservicePort: 5555, // 5555 for test
		}),
	)

	err := svr.Run()
	if err != nil {
		klog.Error("############################   failed to run server    ############################\n")
		klog.Errorf("############################   %s   ############################", err)
	}
}
