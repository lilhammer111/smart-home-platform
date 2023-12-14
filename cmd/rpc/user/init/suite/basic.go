package suite

import (
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"

	"git.zqbjj.top/pet/public-repo/utils"
)

const defaultPort = 1222

type BasicSuite struct {
	MicroserviceIp   string
	MicroservicePort int
}

func (bs *BasicSuite) Options() []server.Option {
	// address
	var err error
	bs.MicroservicePort, err = utils.AllocatePort()
	if err != nil {
		klog.Errorf("failed to allocate port: %s", err)
	}

	if bs.MicroservicePort == 0 {
		bs.MicroservicePort = defaultPort
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}

	// service info
	endpointBasicInfo := rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}

	return []server.Option{
		server.WithServerBasicInfo(&endpointBasicInfo),
		server.WithServiceAddr(addr),
		server.WithMuxTransport(),
	}
}
