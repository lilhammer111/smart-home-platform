package main

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/suite"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user/microuser"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := microuser.NewServer(
		new(MicroUserImpl),
		server.WithSuite(&suite.BasicSuite{
			MicroservicePort: 5555, // 5555 for test
		}),
	)

	err := svr.Run()
	if err != nil {
		klog.Error("############################   failed to run server    ############################\n")
		klog.Errorf("############################   %s   ############################", err)
	}
}
