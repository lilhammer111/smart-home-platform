// Code generated by Kitex v0.6.1. DO NOT EDIT.
package brandservice

import (
	product "git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler product.BrandService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
