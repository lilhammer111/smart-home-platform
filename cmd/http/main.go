// Code generated by hertz generator.

package main

import (
	"fmt"
	"git.zqbjj.top/pet/services/cmd/http/conf"
	_ "git.zqbjj.top/pet/services/cmd/http/docs"
	"git.zqbjj.top/pet/services/cmd/http/mw"
	"git.zqbjj.top/pet/services/cmd/http/router"
	"git.zqbjj.top/pet/services/cmd/http/utils/regis"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	// init jwt
	mw.InitJwt()
	// init regis
	r := regis.Init()

	address := conf.GetConf().Hertz.Address
	hlog.Debugf("http address: %s", address)
	h := server.New(
		server.WithHostPorts(address),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "api",
			Addr:        utils.NewNetAddr("tcp", address),
			Weight:      10,
			Tags: map[string]string{
				"role": "gateway",
				"name": "pet_api",
			},
		}),
	)

	registerMiddleware(h)

	router.Register(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	//logger := hertzlogrus.NewLogger()
	//hlog.SetLogger(logger)
	//hlog.SetLevel(conf.LogLevel())
	//hlog.SetOutput(&lumberjack.Logger{
	//	Filename:   conf.GetConf().Hertz.LogFileName,
	//	MaxSize:    conf.GetConf().Hertz.LogMaxSize,
	//	MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
	//	MaxAge:     conf.GetConf().Hertz.LogMaxAge,
	//})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	// swagger
	url := swagger.URL(fmt.Sprintf("http://%s/swagger/doc.json", conf.GetConf().Hertz.Address))
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
}
