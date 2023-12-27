package auth

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/api/auth/handler"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/auth"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
)

// SendSms .
// @id			SendSms
// @Summary		ask aliyun to send sms
// @Tags		captcha
// @Produce		json
// @Param		mobile	query	string	true	"mobile number"
// @Success		200				{object}		example.RespOk 						"success"
// @Failure		400 			{object}		example.RespBadRequest				"bad request"
// @Failure     404  			{object}		example.RespNotFound				"not found"
// @Failure		500 			{object}		example.RespInternal				"internal error"
// @Failure		401 			{object}		example.RespUnauthorized			"authentication failed"
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

	resp, err := handler.NewSendSmsService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

func MobileRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewMobileRegisterService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

func MobileLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MobileLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewMobileLoginService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

func MiniProgLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.MiniProgLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewMiniProgLoginService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

func PwdLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.PwdLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewPwdLoginService(ctx, c).Do(&req)
	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

func UsernameRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.UsernameRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.Set(responder.ErrorCode, http.StatusBadRequest)
		c.Set(responder.ErrorMessage, "Invalid parameter.")
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := handler.NewUsernameRegisterService(ctx, c).Do(&req)

	if err != nil {
		responder.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	responder.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
