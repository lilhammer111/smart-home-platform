package regis

import (
	"git.zqbjj.top/pet/services/cmd/http/conf"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

func Init() registry.Registry {
	// define consul client config
	config := consulapi.DefaultConfig()
	config.Address = conf.GetConf().Consul.Addr
	// build a consul client
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		hlog.Fatalf("failed to create consul client: %s", err)
	}
	// build a consul register with the consul client
	return consul.NewConsulRegister(consulClient)
}
