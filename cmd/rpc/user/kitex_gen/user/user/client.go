// Code generated by Kitex v0.6.1. DO NOT EDIT.

package user

import (
	"context"
	standard "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/standard"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserList(ctx context.Context, req *user.UsersFilter, callOptions ...callopt.Option) (r *user.UserListResp, err error)
	GetUserDetail(ctx context.Context, req *standard.Req, callOptions ...callopt.Option) (r *user.UserInfoResp, err error)
	UpdateUserInfo(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (r *user.UserInfoResp, err error)
	DeregisterUser(ctx context.Context, req *standard.Req, callOptions ...callopt.Option) (r *standard.Resp, err error)
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
	return &kUserClient{
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

type kUserClient struct {
	*kClient
}

func (p *kUserClient) GetUserList(ctx context.Context, req *user.UsersFilter, callOptions ...callopt.Option) (r *user.UserListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserList(ctx, req)
}

func (p *kUserClient) GetUserDetail(ctx context.Context, req *standard.Req, callOptions ...callopt.Option) (r *user.UserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserDetail(ctx, req)
}

func (p *kUserClient) UpdateUserInfo(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (r *user.UserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, req)
}

func (p *kUserClient) DeregisterUser(ctx context.Context, req *standard.Req, callOptions ...callopt.Option) (r *standard.Resp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeregisterUser(ctx, req)
}