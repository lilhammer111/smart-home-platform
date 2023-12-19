package router

import (
	"git.zqbjj.top/pet/services/cmd/http/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Register registers all routers.
func Register(r *server.Hertz) {

	GeneratedRegister(r)

	customizedRegister(r)
}

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/api/auth/mini_prog_login", mw.JwtMiddleware.LoginHandler)
	r.GET("/api/auth/mobile_login", mw.JwtMiddleware.LoginHandler)
	r.POST("/api/auth/mobile_register", mw.JwtMiddleware.LoginHandler)
	r.GET("/api/auth/pwd_login", mw.JwtMiddleware.LoginHandler)
}
