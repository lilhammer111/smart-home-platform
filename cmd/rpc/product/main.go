package main

import (
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/suite"
	"git.zqbjj.top/pet/services/cmd/rpc/product/biz"
	"git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product/combineservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := combineservice.NewServer(
		new(biz.CombineServiceImpl),
		server.WithSuite(&suite.BasicSuite{}),
	)

	err := svr.Run()
	if err != nil {
		klog.Info("############################   failed to run server    ############################\n")
		klog.Errorf("############################   %s   ############################", err)
	}
}
