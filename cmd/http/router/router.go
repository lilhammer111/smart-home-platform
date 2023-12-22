package router

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/mw"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"net/http"
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
	r.GET(mw.PathRefreshPath, mw.JwtMiddleware.RefreshHandler)
	r.NoRoute(mw.JwtMiddleware.MiddlewareFunc(), func(ctx context.Context, c *app.RequestContext) {
		claims := jwt.ExtractClaims(ctx, c)
		hlog.Infof("NoRoute claims: %#v\n", claims)
		c.JSON(http.StatusOK, responder.FormattedResp{
			Success: false,
			Code:    404,
			Message: "page not found",
		})
	})
}
