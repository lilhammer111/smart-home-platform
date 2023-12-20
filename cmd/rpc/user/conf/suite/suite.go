package suite

import (
	"fmt"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consulApi "github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/pkg/errors"
	"net"

	"git.zqbjj.top/lilhammer111/micro-kit/utils"
)

const (
	defaultPort = 1222
)

type BasicSuite struct {
	MicroserviceIp   string
	MicroservicePort int
}

func (bs *BasicSuite) Options() []server.Option {
	// Set kitex log
	//log.InitKitexLog()

	if binding.GetRemoteConf() == nil {
		klog.Fatalf("GetRemoteConf() returns a nil value: %s",
			errors.New("runtime error: invalid memory address or nil pointer dereference"),
		)
	}
	// Service info
	endpointBasicInfo := rpcinfo.EndpointBasicInfo{
		ServiceName: binding.GetRemoteConf().Microservice.Name,
	}

	// Determine which port
	var port int
	// if port is specified when initializing, just use the port
	if bs.MicroservicePort != 0 {
		port = bs.MicroservicePort
	} else {
		// if not, allocate an idle port randomly for the microservice
		var err error
		port, err = utils.AllocatePort()
		if err != nil {
			klog.Errorf("failed to allocate port: %s", err)
			// failed to allocate? Ok, let port be the default
			port = defaultPort
		}
	}

	// Determine which ip
	var TCPAddr string
	if bs.MicroserviceIp != "" {
		TCPAddr = bs.MicroserviceIp
	} else {
		TCPAddr = fmt.Sprintf("%s:%d", binding.GetRemoteConf().Microservice.IP, port)
	}
	addr, err := net.ResolveTCPAddr("tcp", TCPAddr)
	if err != nil {
		klog.Fatalf("failed to resolve tcp addr: %s", err)
	}

	// service register
	r, err := consul.NewConsulRegister(
		binding.GetRemoteConf().Consul.Addr,
		consul.WithCheck(&consulApi.AgentServiceCheck{
			Name:                           fmt.Sprintf("%s-check", endpointBasicInfo.ServiceName),
			Interval:                       "7s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "1m",
		}),
	)
	if err != nil {
		klog.Fatalf("failed to start register: %s", err)
	}

	return []server.Option{
		server.WithServerBasicInfo(&endpointBasicInfo),
		server.WithServiceAddr(addr),
		server.WithMuxTransport(),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: endpointBasicInfo.ServiceName,
			Addr:        addr,
		}),
	}
}
