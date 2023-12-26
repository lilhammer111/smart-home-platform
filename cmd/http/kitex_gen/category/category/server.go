// Code generated by Kitex v0.8.0. DO NOT EDIT.
package category

import (
	category "git.zqbjj.top/pet/services/cmd/http/kitex_gen/category"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler category.Category, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
