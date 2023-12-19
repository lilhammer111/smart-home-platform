package mw

import (
	"context"
	"errors"
	"fmt"
	"git.zqbjj.top/pet/services/cmd/http/conf"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"

	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
)

const (
	MsgLoginSuccess   = "login success"
	MsgLogoutSuccess  = "logout success"
	MsgRefreshSuccess = "refresh token success"
)

const (
	IdentityKey        = "identity"
	PathMobileLogin    = "/api/auth/mobile_login"
	PathPwdLogin       = "/api/auth/pwd_login"
	PathMiniProgLogin  = "/api/auth/mini_prog_login"
	PathMobileRegister = "/api/auth/mobile_register"

	AdminPathGetUserList = "/api/users/list"
)

var (
	ErrMobileExists        = errors.New("mobile number already exists")
	ErrWrongSmsCode        = errors.New("wrong sms code")
	ErrInternalError       = errors.New("internal error")
	ErrMiniProgLoginFailed = errors.New("mini program login failed")
)

type JwtRespData struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type AccountFrozenRespData struct {
	IsFrozen bool      `json:"isFrozen"`
	ThawedAt time.Time `json:"thawedAt,omitempty"`
}

func newJwtData(token string, expiredAt time.Time) *JwtRespData {
	return &JwtRespData{
		Token:     token,
		ExpiredAt: expiredAt,
	}
}

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			responder.SendSuccessResponse(ctx, c, code, newJwtData(token, expire))
		},
		LogoutResponse: func(ctx context.Context, c *app.RequestContext, code int) {
			responder.SendSuccessResponse(ctx, c, code, nil)
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			responder.SendSuccessResponse(ctx, c, code, newJwtData(token, expire))
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			// If the account is frozen, remind users how many minutes it will take to try to
			if c.GetBool("is_frozen") {
				// todo why thawed_at is of type string?
				c.Set(responder.ErrorMessage, fmt.Sprintf("account frozen, please try again in %s minutes", c.GetString("thawed_at")))
			}
			responder.SendErrResponse(ctx, c, code, err)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			switch string(c.Path()) {
			case PathMobileLogin:
				return mobileLoginAuthenticator(ctx, c)
			case PathPwdLogin:
				return pwdLoginAuthenticator(ctx, c)
			case PathMiniProgLogin:
				return miniProgLoginAuthenticator(ctx, c)
			case PathMobileRegister:
				return mobileRegisterAuthenticator(ctx, c)
			default:
				return nil, jwt.ErrFailedAuthentication
			}
		},
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.UserInfo); ok {
				return jwt.MapClaims{
					IdentityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &user.UserInfo{
				Id: claims[IdentityKey].(*int32),
			}
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			switch string(c.Path()) {
			case AdminPathGetUserList:
				if v, ok := data.(*user.UserInfo); ok {
					return *v.Id == 1
				}
			}
			return true
		},
	})
	if err != nil {
		panic(err)
	}
}

func mobileLoginAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	userInfo := &user.UserInfo{}
	req := &auth.MobileLoginReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	// check if the account was frozen
	patrolResp, err := micro_user_cli.FreezePatrolBeforeVerify(ctx, &micro_user.RpcFreezeReq{Mobile: &req.Mobile})
	if err != nil {
		return nil, err
	}

	verifyResp, err := micro_user_cli.VerifySmsCode(ctx, &micro_user.RpcVerifyCodeReq{
		Mobile:  req.Mobile,
		SmsCode: req.SmsCode,
	})
	if err != nil {
		return nil, err
	}

	_, err = micro_user_cli.FreezePatrolAfterVerify(ctx, &micro_user.RpcAfterVerifyReq{
		UserId:       patrolResp.UserId,
		VerifyPassed: verifyResp.VerifyPassed,
	})
	if err != nil {
		return nil, err
	}

	userInfo, err = micro_user_cli.FindUser(ctx, &micro_user.RpcFindUserReq{UserId: patrolResp.UserId})
	if err != nil {
		return nil, ErrInternalError
	}
	return userInfo, nil
}

func pwdLoginAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	userInfo := &user.UserInfo{}
	// todo the pwdLoginReq struct may need to be modified
	req := &auth.PwdLoginReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	// check if the account was frozen
	patrolResp, err := micro_user_cli.FreezePatrolBeforeVerify(ctx, &micro_user.RpcFreezeReq{Username: &req.Username})
	if err != nil {
		return nil, err
	}

	verifyResp, err := micro_user_cli.VerifyUsernamePwd(ctx, &micro_user.RpcVerifyUsernamePwdReq{
		Username: req.Username,
		EntryPwd: req.Password,
	})
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	_, err = micro_user_cli.FreezePatrolAfterVerify(ctx, &micro_user.RpcAfterVerifyReq{
		UserId:       patrolResp.UserId,
		VerifyPassed: verifyResp.VerifyPassed,
	})
	if err != nil {
		return nil, err
	}

	userInfo, err = micro_user_cli.FindUser(ctx, &micro_user.RpcFindUserReq{UserId: patrolResp.UserId})
	if err != nil {
		return nil, ErrInternalError
	}
	return userInfo, nil
}

func miniProgLoginAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	userInfo := &user.UserInfo{}
	req := &auth.MiniProgLoginReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	miniProgResp, err := micro_user_cli.RequestOpenId(ctx, &micro_user.RpcRequestOpenIdReq{
		JsCode: req.WxCode,
		Secret: conf.GetConf().Wx.Secret,
		Appid:  conf.GetConf().Wx.Appid,
	})

	userInfo, err = micro_user_cli.FindUserByOpenid(ctx, &micro_user.RpcFindUserByOpenidReq{Openid: miniProgResp.OpenId})
	if err != nil {
		hlog.Errorf("failed to get userinfo by openid: %s", err)
		return nil, err
	}
	return userInfo, nil
}

func mobileRegisterAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	userInfo := &user.UserInfo{}
	req := &auth.MobileRegisterReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	userInfo, err = micro_user_cli.FindUserByMobile(ctx, &micro_user.RpcFindUserByMobileReq{Mobile: req.Mobile})
	if err == nil && userInfo != nil {
		return nil, ErrMobileExists
	}

	_, err = micro_user_cli.VerifySmsCode(ctx, &micro_user.RpcVerifyCodeReq{
		Mobile:  req.Mobile,
		SmsCode: req.SmsCode,
	})
	if err != nil {
		return nil, ErrWrongSmsCode
	}

	userInfo, err = micro_user_cli.CreateUser(ctx, &user.UserInfo{
		Mobile:   &req.Mobile,
		Username: &req.Username,
		Password: &req.Password,
	})
	if err != nil {
		return nil, ErrInternalError
	}
	return userInfo, nil
}
