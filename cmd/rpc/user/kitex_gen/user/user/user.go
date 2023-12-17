// Code generated by Kitex v0.6.1. DO NOT EDIT.

package user

import (
	"context"
	common_http "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_http"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceInfo
}

var userServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "user"
	handlerType := (*user.User)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUserList":    kitex.NewMethodInfo(getUserListHandler, newUserGetUserListArgs, newUserGetUserListResult, false),
		"GetUserDetail":  kitex.NewMethodInfo(getUserDetailHandler, newUserGetUserDetailArgs, newUserGetUserDetailResult, false),
		"UpdateUserInfo": kitex.NewMethodInfo(updateUserInfoHandler, newUserUpdateUserInfoArgs, newUserUpdateUserInfoResult, false),
		"DeregisterUser": kitex.NewMethodInfo(deregisterUserHandler, newUserDeregisterUserArgs, newUserDeregisterUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
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

func getUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserGetUserListArgs)
	realResult := result.(*user.UserGetUserListResult)
	success, err := handler.(user.User).GetUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserGetUserListArgs() interface{} {
	return user.NewUserGetUserListArgs()
}

func newUserGetUserListResult() interface{} {
	return user.NewUserGetUserListResult()
}

func getUserDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserGetUserDetailArgs)
	realResult := result.(*user.UserGetUserDetailResult)
	success, err := handler.(user.User).GetUserDetail(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserGetUserDetailArgs() interface{} {
	return user.NewUserGetUserDetailArgs()
}

func newUserGetUserDetailResult() interface{} {
	return user.NewUserGetUserDetailResult()
}

func updateUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserUpdateUserInfoArgs)
	realResult := result.(*user.UserUpdateUserInfoResult)
	success, err := handler.(user.User).UpdateUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserUpdateUserInfoArgs() interface{} {
	return user.NewUserUpdateUserInfoArgs()
}

func newUserUpdateUserInfoResult() interface{} {
	return user.NewUserUpdateUserInfoResult()
}

func deregisterUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserDeregisterUserArgs)
	realResult := result.(*user.UserDeregisterUserResult)
	success, err := handler.(user.User).DeregisterUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserDeregisterUserArgs() interface{} {
	return user.NewUserDeregisterUserArgs()
}

func newUserDeregisterUserResult() interface{} {
	return user.NewUserDeregisterUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetUserList(ctx context.Context, req *user.UsersFilter) (r *user.UserListResp, err error) {
	var _args user.UserGetUserListArgs
	_args.Req = req
	var _result user.UserGetUserListResult
	if err = p.c.Call(ctx, "GetUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserDetail(ctx context.Context, req *common_http.Req) (r *user.UserInfoResp, err error) {
	var _args user.UserGetUserDetailArgs
	_args.Req = req
	var _result user.UserGetUserDetailResult
	if err = p.c.Call(ctx, "GetUserDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserInfo(ctx context.Context, req *user.UserInfo) (r *user.UserInfoResp, err error) {
	var _args user.UserUpdateUserInfoArgs
	_args.Req = req
	var _result user.UserUpdateUserInfoResult
	if err = p.c.Call(ctx, "UpdateUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeregisterUser(ctx context.Context, req *common_http.Req) (r *common_http.Resp, err error) {
	var _args user.UserDeregisterUserArgs
	_args.Req = req
	var _result user.UserDeregisterUserResult
	if err = p.c.Call(ctx, "DeregisterUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
