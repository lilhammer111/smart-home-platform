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
	r.POST(mw.PathMobileLogin, mw.JwtMiddleware.LoginHandler)
	r.GET(mw.PathMiniProgLogin, mw.JwtMiddleware.LoginHandler)
	r.POST(mw.PathMobileRegister, mw.JwtMiddleware.LoginHandler)
	r.POST(mw.PathUsernameRegister, mw.JwtMiddleware.LoginHandler)
	r.POST(mw.PathPwdLogin, mw.JwtMiddleware.LoginHandler)
}
