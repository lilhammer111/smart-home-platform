// Code generated by Kitex v0.6.1. DO NOT EDIT.

package microuser

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendSMSViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (r *common_rpc.RpcEmpty, err error)
	FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error)
	FreezePatrolAfterAuth(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error)
	VerifyCredentials(ctx context.Context, req *micro_user.RpcCredentialReq, callOptions ...callopt.Option) (r *common_rpc.RpcEmpty, err error)
	FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (r *micro_user.RpcUser, err error)
	QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (r []*micro_user.RpcUser, err error)
	UpsertUser(ctx context.Context, req *micro_user.RpcUser, callOptions ...callopt.Option) (r *micro_user.RpcUser, err error)
	DeleteUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (r *common_rpc.RpcEmpty, err error)
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

func (p *kMicroUserClient) SendSMSViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (r *common_rpc.RpcEmpty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendSMSViaAliyun(ctx, req)
}

func (p *kMicroUserClient) FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FreezePatrolBeforeAuth(ctx, req)
}

func (p *kMicroUserClient) FreezePatrolAfterAuth(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (r *micro_user.RpcFreezeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FreezePatrolAfterAuth(ctx, req)
}

func (p *kMicroUserClient) VerifyCredentials(ctx context.Context, req *micro_user.RpcCredentialReq, callOptions ...callopt.Option) (r *common_rpc.RpcEmpty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyCredentials(ctx, req)
}

func (p *kMicroUserClient) FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (r *micro_user.RpcUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindUser(ctx, req)
}

func (p *kMicroUserClient) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (r []*micro_user.RpcUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUsersWithFilter(ctx, req)
}

func (p *kMicroUserClient) UpsertUser(ctx context.Context, req *micro_user.RpcUser, callOptions ...callopt.Option) (r *micro_user.RpcUser, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpsertUser(ctx, req)
}

func (p *kMicroUserClient) DeleteUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (r *common_rpc.RpcEmpty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, req)
}