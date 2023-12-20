package auth

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	biz "git.zqbjj.top/pet/services/cmd/http/biz/auth"
)

// SendSms .
// @router /api/auth/send_sms [GET]
func SendSms(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.SendSmsReq
	err = c.BindAndValidate(&req)
	hlog.Errorf("***************** bind error: %s", err)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewSendSmsService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MobileRegister .
// @router /api/auth/mobile_register [POST]
func MobileRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewMobileRegisterService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MobileLogin .
// @router /api/auth/mobile_login [GET]
func MobileLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewMobileLoginService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MiniProgLogin .
// @router /api/auth/mini_prog_login [GET]
func MiniProgLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MiniProgLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewMiniProgLoginService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// PwdLogin .
// @router /api/auth/pwd_login [GET]
func PwdLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.PwdLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewPwdLoginService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UsernameRegister .
// @router /api/auth/username_register [POST]
func UsernameRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.UsernameRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewUsernameRegisterService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
