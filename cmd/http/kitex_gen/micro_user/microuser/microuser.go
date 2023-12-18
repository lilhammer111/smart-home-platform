// Code generated by Kitex v0.6.1. DO NOT EDIT.

package microuser

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return microUserServiceInfo
}

var microUserServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "micro_user"
	handlerType := (*micro_user.MicroUser)(nil)
	methods := map[string]kitex.MethodInfo{
		"SendSmsViaAliyun":       kitex.NewMethodInfo(sendSmsViaAliyunHandler, newMicroUserSendSmsViaAliyunArgs, newMicroUserSendSmsViaAliyunResult, false),
		"FreezePatrolBeforeAuth": kitex.NewMethodInfo(freezePatrolBeforeAuthHandler, newMicroUserFreezePatrolBeforeAuthArgs, newMicroUserFreezePatrolBeforeAuthResult, false),
		"FreezePatrolAfterAuth":  kitex.NewMethodInfo(freezePatrolAfterAuthHandler, newMicroUserFreezePatrolAfterAuthArgs, newMicroUserFreezePatrolAfterAuthResult, false),
		"VerifySmsCode":          kitex.NewMethodInfo(verifySmsCodeHandler, newMicroUserVerifySmsCodeArgs, newMicroUserVerifySmsCodeResult, false),
		"VerifyUsernamePwd":      kitex.NewMethodInfo(verifyUsernamePwdHandler, newMicroUserVerifyUsernamePwdArgs, newMicroUserVerifyUsernamePwdResult, false),
		"VerifyEmailPwd":         kitex.NewMethodInfo(verifyEmailPwdHandler, newMicroUserVerifyEmailPwdArgs, newMicroUserVerifyEmailPwdResult, false),
		"FindUser":               kitex.NewMethodInfo(findUserHandler, newMicroUserFindUserArgs, newMicroUserFindUserResult, false),
		"FindUserByOpenid":       kitex.NewMethodInfo(findUserByOpenidHandler, newMicroUserFindUserByOpenidArgs, newMicroUserFindUserByOpenidResult, false),
		"FindUserByMobile":       kitex.NewMethodInfo(findUserByMobileHandler, newMicroUserFindUserByMobileArgs, newMicroUserFindUserByMobileResult, false),
		"FindUserByUsername":     kitex.NewMethodInfo(findUserByUsernameHandler, newMicroUserFindUserByUsernameArgs, newMicroUserFindUserByUsernameResult, false),
		"QueryUsersWithFilter":   kitex.NewMethodInfo(queryUsersWithFilterHandler, newMicroUserQueryUsersWithFilterArgs, newMicroUserQueryUsersWithFilterResult, false),
		"UpdateUser":             kitex.NewMethodInfo(updateUserHandler, newMicroUserUpdateUserArgs, newMicroUserUpdateUserResult, false),
		"CreateUser":             kitex.NewMethodInfo(createUserHandler, newMicroUserCreateUserArgs, newMicroUserCreateUserResult, false),
		"DeleteUser":             kitex.NewMethodInfo(deleteUserHandler, newMicroUserDeleteUserArgs, newMicroUserDeleteUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "micro_user",
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

func sendSmsViaAliyunHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserSendSmsViaAliyunArgs)
	realResult := result.(*micro_user.MicroUserSendSmsViaAliyunResult)
	success, err := handler.(micro_user.MicroUser).SendSmsViaAliyun(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserSendSmsViaAliyunArgs() interface{} {
	return micro_user.NewMicroUserSendSmsViaAliyunArgs()
}

func newMicroUserSendSmsViaAliyunResult() interface{} {
	return micro_user.NewMicroUserSendSmsViaAliyunResult()
}

func freezePatrolBeforeAuthHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserFreezePatrolBeforeAuthArgs)
	realResult := result.(*micro_user.MicroUserFreezePatrolBeforeAuthResult)
	success, err := handler.(micro_user.MicroUser).FreezePatrolBeforeAuth(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserFreezePatrolBeforeAuthArgs() interface{} {
	return micro_user.NewMicroUserFreezePatrolBeforeAuthArgs()
}

func newMicroUserFreezePatrolBeforeAuthResult() interface{} {
	return micro_user.NewMicroUserFreezePatrolBeforeAuthResult()
}

func freezePatrolAfterAuthHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserFreezePatrolAfterAuthArgs)
	realResult := result.(*micro_user.MicroUserFreezePatrolAfterAuthResult)
	success, err := handler.(micro_user.MicroUser).FreezePatrolAfterAuth(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserFreezePatrolAfterAuthArgs() interface{} {
	return micro_user.NewMicroUserFreezePatrolAfterAuthArgs()
}

func newMicroUserFreezePatrolAfterAuthResult() interface{} {
	return micro_user.NewMicroUserFreezePatrolAfterAuthResult()
}

func verifySmsCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserVerifySmsCodeArgs)
	realResult := result.(*micro_user.MicroUserVerifySmsCodeResult)
	success, err := handler.(micro_user.MicroUser).VerifySmsCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserVerifySmsCodeArgs() interface{} {
	return micro_user.NewMicroUserVerifySmsCodeArgs()
}

func newMicroUserVerifySmsCodeResult() interface{} {
	return micro_user.NewMicroUserVerifySmsCodeResult()
}

func verifyUsernamePwdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserVerifyUsernamePwdArgs)
	realResult := result.(*micro_user.MicroUserVerifyUsernamePwdResult)
	success, err := handler.(micro_user.MicroUser).VerifyUsernamePwd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserVerifyUsernamePwdArgs() interface{} {
	return micro_user.NewMicroUserVerifyUsernamePwdArgs()
}

func newMicroUserVerifyUsernamePwdResult() interface{} {
	return micro_user.NewMicroUserVerifyUsernamePwdResult()
}

func verifyEmailPwdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserVerifyEmailPwdArgs)
	realResult := result.(*micro_user.MicroUserVerifyEmailPwdResult)
	success, err := handler.(micro_user.MicroUser).VerifyEmailPwd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserVerifyEmailPwdArgs() interface{} {
	return micro_user.NewMicroUserVerifyEmailPwdArgs()
}

func newMicroUserVerifyEmailPwdResult() interface{} {
	return micro_user.NewMicroUserVerifyEmailPwdResult()
}

func findUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserFindUserArgs)
	realResult := result.(*micro_user.MicroUserFindUserResult)
	success, err := handler.(micro_user.MicroUser).FindUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserFindUserArgs() interface{} {
	return micro_user.NewMicroUserFindUserArgs()
}

func newMicroUserFindUserResult() interface{} {
	return micro_user.NewMicroUserFindUserResult()
}

func findUserByOpenidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserFindUserByOpenidArgs)
	realResult := result.(*micro_user.MicroUserFindUserByOpenidResult)
	success, err := handler.(micro_user.MicroUser).FindUserByOpenid(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserFindUserByOpenidArgs() interface{} {
	return micro_user.NewMicroUserFindUserByOpenidArgs()
}

func newMicroUserFindUserByOpenidResult() interface{} {
	return micro_user.NewMicroUserFindUserByOpenidResult()
}

func findUserByMobileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserFindUserByMobileArgs)
	realResult := result.(*micro_user.MicroUserFindUserByMobileResult)
	success, err := handler.(micro_user.MicroUser).FindUserByMobile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserFindUserByMobileArgs() interface{} {
	return micro_user.NewMicroUserFindUserByMobileArgs()
}

func newMicroUserFindUserByMobileResult() interface{} {
	return micro_user.NewMicroUserFindUserByMobileResult()
}

func findUserByUsernameHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserFindUserByUsernameArgs)
	realResult := result.(*micro_user.MicroUserFindUserByUsernameResult)
	success, err := handler.(micro_user.MicroUser).FindUserByUsername(ctx, realArg.Username)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserFindUserByUsernameArgs() interface{} {
	return micro_user.NewMicroUserFindUserByUsernameArgs()
}

func newMicroUserFindUserByUsernameResult() interface{} {
	return micro_user.NewMicroUserFindUserByUsernameResult()
}

func queryUsersWithFilterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserQueryUsersWithFilterArgs)
	realResult := result.(*micro_user.MicroUserQueryUsersWithFilterResult)
	success, err := handler.(micro_user.MicroUser).QueryUsersWithFilter(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserQueryUsersWithFilterArgs() interface{} {
	return micro_user.NewMicroUserQueryUsersWithFilterArgs()
}

func newMicroUserQueryUsersWithFilterResult() interface{} {
	return micro_user.NewMicroUserQueryUsersWithFilterResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserUpdateUserArgs)
	realResult := result.(*micro_user.MicroUserUpdateUserResult)
	success, err := handler.(micro_user.MicroUser).UpdateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserUpdateUserArgs() interface{} {
	return micro_user.NewMicroUserUpdateUserArgs()
}

func newMicroUserUpdateUserResult() interface{} {
	return micro_user.NewMicroUserUpdateUserResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserCreateUserArgs)
	realResult := result.(*micro_user.MicroUserCreateUserResult)
	success, err := handler.(micro_user.MicroUser).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserCreateUserArgs() interface{} {
	return micro_user.NewMicroUserCreateUserArgs()
}

func newMicroUserCreateUserResult() interface{} {
	return micro_user.NewMicroUserCreateUserResult()
}

func deleteUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*micro_user.MicroUserDeleteUserArgs)
	realResult := result.(*micro_user.MicroUserDeleteUserResult)
	success, err := handler.(micro_user.MicroUser).DeleteUser(ctx, realArg.UserId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMicroUserDeleteUserArgs() interface{} {
	return micro_user.NewMicroUserDeleteUserArgs()
}

func newMicroUserDeleteUserResult() interface{} {
	return micro_user.NewMicroUserDeleteUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq) (r *common.Empty, err error) {
	var _args micro_user.MicroUserSendSmsViaAliyunArgs
	_args.Req = req
	var _result micro_user.MicroUserSendSmsViaAliyunResult
	if err = p.c.Call(ctx, "SendSmsViaAliyun", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq) (r *micro_user.RpcFreezeResp, err error) {
	var _args micro_user.MicroUserFreezePatrolBeforeAuthArgs
	_args.Req = req
	var _result micro_user.MicroUserFreezePatrolBeforeAuthResult
	if err = p.c.Call(ctx, "FreezePatrolBeforeAuth", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FreezePatrolAfterAuth(ctx context.Context, req *micro_user.RpcUserId) (r *micro_user.RpcFreezeResp, err error) {
	var _args micro_user.MicroUserFreezePatrolAfterAuthArgs
	_args.Req = req
	var _result micro_user.MicroUserFreezePatrolAfterAuthResult
	if err = p.c.Call(ctx, "FreezePatrolAfterAuth", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifySmsCode(ctx context.Context, req *micro_user.RpcVerifyCodeReq) (r *common.Empty, err error) {
	var _args micro_user.MicroUserVerifySmsCodeArgs
	_args.Req = req
	var _result micro_user.MicroUserVerifySmsCodeResult
	if err = p.c.Call(ctx, "VerifySmsCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyUsernamePwd(ctx context.Context, req *micro_user.RpcVerifyUsernamePwdReq) (r *common.Empty, err error) {
	var _args micro_user.MicroUserVerifyUsernamePwdArgs
	_args.Req = req
	var _result micro_user.MicroUserVerifyUsernamePwdResult
	if err = p.c.Call(ctx, "VerifyUsernamePwd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyEmailPwd(ctx context.Context, req *micro_user.RpcVerifyEmailPwdReq) (r *common.Empty, err error) {
	var _args micro_user.MicroUserVerifyEmailPwdArgs
	_args.Req = req
	var _result micro_user.MicroUserVerifyEmailPwdResult
	if err = p.c.Call(ctx, "VerifyEmailPwd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindUser(ctx context.Context, req *micro_user.RpcFindUserReq) (r *user.UserInfo, err error) {
	var _args micro_user.MicroUserFindUserArgs
	_args.Req = req
	var _result micro_user.MicroUserFindUserResult
	if err = p.c.Call(ctx, "FindUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindUserByOpenid(ctx context.Context, req *micro_user.RpcFindUserByOpenidReq) (r *user.UserInfo, err error) {
	var _args micro_user.MicroUserFindUserByOpenidArgs
	_args.Req = req
	var _result micro_user.MicroUserFindUserByOpenidResult
	if err = p.c.Call(ctx, "FindUserByOpenid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindUserByMobile(ctx context.Context, req *micro_user.RpcFindUserByMobileReq) (r *user.UserInfo, err error) {
	var _args micro_user.MicroUserFindUserByMobileArgs
	_args.Req = req
	var _result micro_user.MicroUserFindUserByMobileResult
	if err = p.c.Call(ctx, "FindUserByMobile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FindUserByUsername(ctx context.Context, username *micro_user.RpcFindUserByUsernameReq) (r *user.UserInfo, err error) {
	var _args micro_user.MicroUserFindUserByUsernameArgs
	_args.Username = username
	var _result micro_user.MicroUserFindUserByUsernameResult
	if err = p.c.Call(ctx, "FindUserByUsername", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq) (r []*user.UserInfo, err error) {
	var _args micro_user.MicroUserQueryUsersWithFilterArgs
	_args.Req = req
	var _result micro_user.MicroUserQueryUsersWithFilterResult
	if err = p.c.Call(ctx, "QueryUsersWithFilter", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, req *user.UserInfo) (r *user.UserInfo, err error) {
	var _args micro_user.MicroUserUpdateUserArgs
	_args.Req = req
	var _result micro_user.MicroUserUpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, req *user.UserInfo) (r *user.UserInfo, err error) {
	var _args micro_user.MicroUserCreateUserArgs
	_args.Req = req
	var _result micro_user.MicroUserCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteUser(ctx context.Context, userId int32) (r *common.Empty, err error) {
	var _args micro_user.MicroUserDeleteUserArgs
	_args.UserId = userId
	var _result micro_user.MicroUserDeleteUserResult
	if err = p.c.Call(ctx, "DeleteUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
