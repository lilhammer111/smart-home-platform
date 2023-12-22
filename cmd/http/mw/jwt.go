package mw

import (
	"context"
	"errors"
	"git.zqbjj.top/pet/services/cmd/http/conf"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_user_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
	"net/http"
	"time"

	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
)

const (
	IdentityKey = "uid"

	PathMobileLogin      = "/api/auth/mobile_login"
	PathPwdLogin         = "/api/auth/pwd_login"
	PathMiniProgLogin    = "/api/auth/mini_prog_login"
	PathMobileRegister   = "/api/auth/mobile_register"
	PathUsernameRegister = "/api/auth/username_register"
	PathRefreshPath      = "/api/auth/refresh"

	AdminPathGetUserList = "/api/users/list"
)

var (
	ErrMobileExists = errors.New("mobile number already exists")
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
			c.JSON(http.StatusOK, responder.FormattedResp{
				Success: false,
				Code:    code,
				Message: message,
			})
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
			case PathUsernameRegister:
				return usernameRegisterAuthenticator(ctx, c)
			default:
				return nil, jwt.ErrFailedAuthentication
			}
		},
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				hlog.Info("payload func type assert success")
				hlog.Infof("v is %+v", v)
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			hlog.Infof("claims is %+v", claims)
			userInfo := model.User{}
			uid := claims[IdentityKey].(float64)
			userInfo.ID = int32(uid)
			hlog.Info("user info id is: %d", userInfo.ID)
			return &userInfo
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

// @Summary		user login by mobile number and sms code
// @id			LoginByMobile
// @Tags		auth
// @Accept    	json
// @Produce		json
// @Param		mobile_login	body	auth.MobileLoginReq	true	"mobile login form"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/auth/mobile_login [POST]
func mobileLoginAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	req := &auth.MobileLoginReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		return nil, jwt.ErrMissingLoginValues
	}

	// check if the account was frozen
	patrolResp, err := micro_user_cli.FreezePatrolBeforeVerify(ctx, &micro_user.RpcFreezeReq{Mobile: &req.Mobile})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	verifyResp, err := micro_user_cli.VerifySmsCode(ctx, &micro_user.RpcVerifyCodeReq{
		Mobile:  req.Mobile,
		SmsCode: req.SmsCode,
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	_, err = micro_user_cli.FreezePatrolAfterVerify(ctx, &micro_user.RpcAfterVerifyReq{
		UserId:       patrolResp.UserId,
		VerifyPassed: verifyResp.VerifyPassed,
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	findResp := &user.UserInfo{}
	findResp, err = micro_user_cli.FindUser(ctx, &micro_user.RpcFindUserReq{UserId: patrolResp.UserId})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	userInfo := model.User{}
	err = copier.Copy(&userInfo, findResp)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &userInfo, nil
}

// @Summary		user login by username and password
// @id			LoginByPwd
// @Tags		auth
// @Accept    	json
// @Produce		json
// @Param		pwd_login	body	auth.PwdLoginReq	true	"password"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/auth/pwd_login [POST]
func pwdLoginAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	// todo the pwdLoginReq struct may need to be modified
	req := &auth.PwdLoginReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		return nil, jwt.ErrMissingLoginValues
	}

	// check if the account was frozen
	patrolResp, err := micro_user_cli.FreezePatrolBeforeVerify(ctx, &micro_user.RpcFreezeReq{Username: &req.Username})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	verifyResp, err := micro_user_cli.VerifyUsernamePwd(ctx, &micro_user.RpcVerifyUsernamePwdReq{
		Username: req.Username,
		EntryPwd: req.Password,
	})
	if err != nil {
		hlog.Error(err)
		return nil, jwt.ErrFailedAuthentication
	}

	_, err = micro_user_cli.FreezePatrolAfterVerify(ctx, &micro_user.RpcAfterVerifyReq{
		UserId:       patrolResp.UserId,
		VerifyPassed: verifyResp.VerifyPassed,
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	findResp := &user.UserInfo{}
	findResp, err = micro_user_cli.FindUser(ctx, &micro_user.RpcFindUserReq{UserId: patrolResp.UserId})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	userInfo := model.User{}
	err = copier.Copy(&userInfo, findResp)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &userInfo, nil
}

// @Summary		user login by mini program
// @id			LoginByMiniProg
// @Tags		auth
// @Accept    	json
// @Produce		json
// @Param		mini_login	body		auth.MiniProgLoginReq	true	"mini program login"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/auth/mini_prog_login [POST]
func miniProgLoginAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	req := &auth.MiniProgLoginReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		return nil, jwt.ErrMissingLoginValues
	}

	miniProgResp, err := micro_user_cli.RequestOpenId(ctx, &micro_user.RpcRequestOpenIdReq{
		JsCode: req.WxCode,
		Secret: conf.GetConf().Wx.Secret,
		Appid:  conf.GetConf().Wx.Appid,
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	findResp := &user.UserInfo{}
	findResp, err = micro_user_cli.FindUserByOpenid(ctx, &micro_user.RpcFindUserByOpenidReq{Openid: miniProgResp.OpenId})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	userInfo := model.User{}
	err = copier.Copy(&userInfo, findResp)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &userInfo, nil
}

// @Summary		user register by mobile, sms code and password
// @id			RegisterByMobile
// @Accept    	json
// @Tags		auth
// @Produce		json
// @Param		mobile_register	body		auth.MobileRegisterReq	true	"mobile register form"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		409 			{object}		example.RespConflict				"account already exists"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/auth/mobile_register [POST]
func mobileRegisterAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error

	req := &auth.MobileRegisterReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		return nil, jwt.ErrMissingLoginValues
	}
	findResp := &user.UserInfo{}
	// If user has existed, generate an error of 'account exists'
	findResp, err = micro_user_cli.FindUserByMobile(ctx, &micro_user.RpcFindUserByMobileReq{Mobile: req.Mobile})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}
	if findResp != nil {
		hlog.Error(err)
		return nil, ErrMobileExists
	}

	_, err = micro_user_cli.VerifySmsCode(ctx, &micro_user.RpcVerifyCodeReq{
		Mobile:  req.Mobile,
		SmsCode: req.SmsCode,
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	createResp := &user.UserInfo{}
	createResp, err = micro_user_cli.CreateUser(ctx, &user.UserInfo{
		Mobile:   &req.Mobile,
		Username: &req.Username,
		Password: &req.Password,
	})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	userInfo := model.User{}
	err = copier.Copy(&userInfo, createResp)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return &userInfo, nil
}

// @Summary		user register by username and password
// @id			RegisterByUsername
// @Tags		auth
// @Accept 		json
// @Produce		json
// @Param		pwd_register_req	body		auth.UsernameRegisterReq	true	"register form"
// @Success		200				{object}		example.RespOk{data=example.UserData} "success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		409 			{object}		example.RespConflict				"account already exists"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
// @router /api/auth/username_register [POST]
func usernameRegisterAuthenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var err error
	req := &auth.UsernameRegisterReq{}
	err = c.BindAndValidate(req)
	if err != nil {
		hlog.Error(err)
		return nil, jwt.ErrMissingLoginValues
	}

	createResp := &user.UserInfo{}
	createResp, err = micro_user_cli.CreateUser(ctx, &user.UserInfo{Username: &req.Username, Password: &req.Password})
	hlog.Errorf("err is: %s", err)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	userInfo := model.User{}
	err = copier.Copy(&userInfo, createResp)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}
	return &userInfo, nil
}
