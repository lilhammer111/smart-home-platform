// Code generated by Kitex v0.6.1. DO NOT EDIT.

package auth

import (
	"context"
	auth "git.zqbjj.top/pet/services/cmd/http/kitex_gen/auth"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return authServiceInfo
}

var authServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "auth"
	handlerType := (*auth.Auth)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendSms":        kitex.NewMethodInfo(sendSmsHandler, newAuthSendSmsArgs, newAuthSendSmsResult, false),
		"MobileRegister": kitex.NewMethodInfo(mobileRegisterHandler, newAuthMobileRegisterArgs, newAuthMobileRegisterResult, false),
		"MobileLogin":    kitex.NewMethodInfo(mobileLoginHandler, newAuthMobileLoginArgs, newAuthMobileLoginResult, false),
		"MiniProgLogin":  kitex.NewMethodInfo(miniProgLoginHandler, newAuthMiniProgLoginArgs, newAuthMiniProgLoginResult, false),
		"PwdLogin":       kitex.NewMethodInfo(pwdLoginHandler, newAuthPwdLoginArgs, newAuthPwdLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "auth",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func sendSmsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthSendSmsArgs)
	realResult := result.(*auth.AuthSendSmsResult)
	success, err := handler.(auth.Auth).SendSms(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthSendSmsArgs() interface{} {
	return auth.NewAuthSendSmsArgs()
}

func newAuthSendSmsResult() interface{} {
	return auth.NewAuthSendSmsResult()
}

func mobileRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthMobileRegisterArgs)
	realResult := result.(*auth.AuthMobileRegisterResult)
	success, err := handler.(auth.Auth).MobileRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthMobileRegisterArgs() interface{} {
	return auth.NewAuthMobileRegisterArgs()
}

func newAuthMobileRegisterResult() interface{} {
	return auth.NewAuthMobileRegisterResult()
}

func mobileLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthMobileLoginArgs)
	realResult := result.(*auth.AuthMobileLoginResult)
	success, err := handler.(auth.Auth).MobileLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthMobileLoginArgs() interface{} {
	return auth.NewAuthMobileLoginArgs()
}

func newAuthMobileLoginResult() interface{} {
	return auth.NewAuthMobileLoginResult()
}

func miniProgLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthMiniProgLoginArgs)
	realResult := result.(*auth.AuthMiniProgLoginResult)
	success, err := handler.(auth.Auth).MiniProgLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthMiniProgLoginArgs() interface{} {
	return auth.NewAuthMiniProgLoginArgs()
}

func newAuthMiniProgLoginResult() interface{} {
	return auth.NewAuthMiniProgLoginResult()
}

func pwdLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*auth.AuthPwdLoginArgs)
	realResult := result.(*auth.AuthPwdLoginResult)
	success, err := handler.(auth.Auth).PwdLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAuthPwdLoginArgs() interface{} {
	return auth.NewAuthPwdLoginArgs()
}

func newAuthPwdLoginResult() interface{} {
	return auth.NewAuthPwdLoginResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendSms(ctx context.Context, req *auth.SendSmsReq) (r *common.Empty, err error) {
	var _args auth.AuthSendSmsArgs
	_args.Req = req
	var _result auth.AuthSendSmsResult
	if err = p.c.Call(ctx, "SendSms", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MobileRegister(ctx context.Context, req *auth.MobileRegisterReq) (r *auth.AuthInfo, err error) {
	var _args auth.AuthMobileRegisterArgs
	_args.Req = req
	var _result auth.AuthMobileRegisterResult
	if err = p.c.Call(ctx, "MobileRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MobileLogin(ctx context.Context, req *auth.MobileLoginReq) (r *auth.AuthInfo, err error) {
	var _args auth.AuthMobileLoginArgs
	_args.Req = req
	var _result auth.AuthMobileLoginResult
	if err = p.c.Call(ctx, "MobileLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MiniProgLogin(ctx context.Context, req *auth.MiniProgLoginReq) (r *auth.AuthInfo, err error) {
	var _args auth.AuthMiniProgLoginArgs
	_args.Req = req
	var _result auth.AuthMiniProgLoginResult
	if err = p.c.Call(ctx, "MiniProgLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PwdLogin(ctx context.Context, req *auth.PwdLoginReq) (r *auth.AuthInfo, err error) {
	var _args auth.AuthPwdLoginArgs
	_args.Req = req
	var _result auth.AuthPwdLoginResult
	if err = p.c.Call(ctx, "PwdLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
