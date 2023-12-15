package suite

import (
	"fmt"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/log"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net"

	"git.zqbjj.top/pet/public-repo/utils"
)

const defaultPort = 1222

type BasicSuite struct {
	MicroserviceIp   string
	MicroservicePort int
	MysqlCli         *gorm.DB
	RedisCli         *redis.Client
}

func (bs *BasicSuite) Options() []server.Option {
	// set kitex log
	log.InitKitexLog()

	// service info
	endpointBasicInfo := rpcinfo.EndpointBasicInfo{
		ServiceName: binding.GetRemoteConf().Microservice.Name,
	}

	// determine port
	var port int
	if bs.MicroservicePort != 0 {
		port = bs.MicroservicePort
	} else {
		var err error
		port, err = utils.AllocatePort()
		if err != nil {
			klog.Errorf("failed to allocate port: %s", err)
			port = defaultPort
		}
	}
	// determine ip
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

	return []server.Option{
		server.WithServerBasicInfo(&endpointBasicInfo),
		server.WithServiceAddr(addr),
		server.WithMuxTransport(),
	}
}
