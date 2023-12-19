// Code generated by Kitex v0.6.1. DO NOT EDIT.

package microuser

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	FreezePatrolBeforeVerify(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error)
	FreezePatrolAfterVerify(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error)
	VerifySmsCode(ctx context.Context, req *micro_user.RpcVerifyCodeReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	VerifyUsernamePwd(ctx context.Context, req *micro_user.RpcVerifyUsernamePwdReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	VerifyEmailPwd(ctx context.Context, req *micro_user.RpcVerifyEmailPwdReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	FindUser(ctx context.Context, req *micro_user.RpcFindUserReq, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	FindUserByOpenid(ctx context.Context, req *micro_user.RpcFindUserByOpenidReq, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	FindUserByMobile(ctx context.Context, req *micro_user.RpcFindUserByMobileReq, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	FindUserByUsername(ctx context.Context, req *micro_user.RpcFindUserByUsernameReq, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (r []*user.UserInfo, err error)
	UpdateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	CreateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (r *user.UserInfo, err error)
	DeleteUser(ctx context.Context, userId int32, callOptions ...callopt.Option) (r *common.Empty, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMicroUserClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMicroUserClient struct {
	*kClient
}

func (p *kMicroUserClient) SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendSmsViaAliyun(ctx, req)
}

func (p *kMicroUserClient) FreezePatrolBeforeVerify(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FreezePatrolBeforeVerify(ctx, req)
}

func (p *kMicroUserClient) FreezePatrolAfterVerify(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FreezePatrolAfterVerify(ctx, req)
}

func (p *kMicroUserClient) VerifySmsCode(ctx context.Context, req *micro_user.RpcVerifyCodeReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifySmsCode(ctx, req)
}

func (p *kMicroUserClient) VerifyUsernamePwd(ctx context.Context, req *micro_user.RpcVerifyUsernamePwdReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyUsernamePwd(ctx, req)
}

func (p *kMicroUserClient) VerifyEmailPwd(ctx context.Context, req *micro_user.RpcVerifyEmailPwdReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyEmailPwd(ctx, req)
}

func (p *kMicroUserClient) FindUser(ctx context.Context, req *micro_user.RpcFindUserReq, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindUser(ctx, req)
}

func (p *kMicroUserClient) FindUserByOpenid(ctx context.Context, req *micro_user.RpcFindUserByOpenidReq, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindUserByOpenid(ctx, req)
}

func (p *kMicroUserClient) FindUserByMobile(ctx context.Context, req *micro_user.RpcFindUserByMobileReq, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindUserByMobile(ctx, req)
}

func (p *kMicroUserClient) FindUserByUsername(ctx context.Context, req *micro_user.RpcFindUserByUsernameReq, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindUserByUsername(ctx, req)
}

func (p *kMicroUserClient) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (r []*user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUsersWithFilter(ctx, req)
}

func (p *kMicroUserClient) UpdateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, req)
}

func (p *kMicroUserClient) CreateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (r *user.UserInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kMicroUserClient) DeleteUser(ctx context.Context, userId int32, callOptions ...callopt.Option) (r *common.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, userId)
}
