// Code generated by Kitex v0.8.0. DO NOT EDIT.
package brand

import (
	brand "git.zqbjj.top/pet/services/cmd/http/kitex_gen/brand"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler brand.Brand, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
