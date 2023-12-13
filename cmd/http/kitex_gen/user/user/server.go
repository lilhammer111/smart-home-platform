// Code generated by Kitex v0.6.1. DO NOT EDIT.
package user

import (
	user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user.User, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
