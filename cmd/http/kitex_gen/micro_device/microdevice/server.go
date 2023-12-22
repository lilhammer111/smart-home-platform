// Code generated by Kitex v0.6.1. DO NOT EDIT.
package microdevice

import (
	micro_device "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler micro_device.MicroDevice, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
