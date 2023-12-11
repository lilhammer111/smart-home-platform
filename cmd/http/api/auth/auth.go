package auth

import (
	"context"

	biz "git.zqbjj.top/pet/services/pet-feeder/cmd/http/biz/auth"
	"git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/pet-feeder/cmd/http/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendSms .
// @router /api/auth/send_sms [GET]
func SendSms(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.SendSmsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewSendSmsService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MobileRegister .
// @router /api/auth/mobile_register [POST]
func MobileRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewMobileRegisterService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MobileLogin .
// @router /api/auth/mobile_login [POST]
func MobileLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewMobileLoginService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// MiniProgLogin .
// @router /api/auth/mini_prog_login [POST]
func MiniProgLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := biz.NewMiniProgLoginService(ctx, c).Do(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
