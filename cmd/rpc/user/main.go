package main

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf"
	"git.zqbjj.top/pet/services/cmd/rpc/user/init/suite"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user/microuser"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	opts := kitexInit()

	svr := microuser.NewServer(new(MicroUserImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	// basic opts
	opts = append(opts, server.WithSuite(suite.BasicSuite{
		MicroserviceIp:   conf.GetConf().Kitex.Address,
		MicroservicePort: 3333, //3333 for test
	}))

	return
}
