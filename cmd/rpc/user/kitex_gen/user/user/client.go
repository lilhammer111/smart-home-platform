// Code generated by Kitex v0.6.1. DO NOT EDIT.

package user

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FindUserByID(ctx context.Context, req int32, callOptions ...callopt.Option) (r *user.UserData, err error)
	GetUsersByFilter(ctx context.Context, req *user.UsersFilterReq, callOptions ...callopt.Option) (r []*user.UserData, err error)
	VerifyCredentials(ctx context.Context, req *user.CredentialsReq, callOptions ...callopt.Option) (r bool, err error)
	CreateOrUpdateUser(ctx context.Context, req *user.UserData, callOptions ...callopt.Option) (r *user.UserData, err error)
	DeleteUser(ctx context.Context, req int32, callOptions ...callopt.Option) (r *user.UserData, err error)
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

func (p *kUserClient) FindUserByID(ctx context.Context, req int32, callOptions ...callopt.Option) (r *user.UserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindUserByID(ctx, req)
}

func (p *kUserClient) GetUsersByFilter(ctx context.Context, req *user.UsersFilterReq, callOptions ...callopt.Option) (r []*user.UserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUsersByFilter(ctx, req)
}

func (p *kUserClient) VerifyCredentials(ctx context.Context, req *user.CredentialsReq, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyCredentials(ctx, req)
}

func (p *kUserClient) CreateOrUpdateUser(ctx context.Context, req *user.UserData, callOptions ...callopt.Option) (r *user.UserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateOrUpdateUser(ctx, req)
}

func (p *kUserClient) DeleteUser(ctx context.Context, req int32, callOptions ...callopt.Option) (r *user.UserData, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, req)
}
