package discovery

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/discovery"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	once     sync.Once
	resolver discovery.Resolver
	err      error
)

func GetDiscoveryResolver() discovery.Resolver {
	once.Do(func() {
		resolver, err = consul.NewConsulResolver("127.0.0.1:8500")
		if err != nil {
			hlog.Fatalf("failed to create a resovler: %s", err)
		}
	})
	return resolver
}
