package users

import (
	"context"
	usersSrv "git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/kitex_gen/users"

	"git.zqbjj.top/pet/services/pet-feeder/cmd/http/dto/kitex_gen/users/users"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() users.Client
	Service() string
	FindUser(ctx context.Context, req int32, callOptions ...callopt.Option) (resp *usersSrv.UserData, err error)

	VerifyCredentials(ctx context.Context, req *usersSrv.CredentialsReq, callOptions ...callopt.Option) (resp bool, err error)

	CreateOrUpdateUser(ctx context.Context, req *usersSrv.UserData, callOptions ...callopt.Option) (resp *usersSrv.UserData, err error)

	DeleteUser(ctx context.Context, req int32, callOptions ...callopt.Option) (resp *usersSrv.UserData, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := users.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient users.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() users.Client {
	return c.kitexClient
}

func (c *clientImpl) FindUser(ctx context.Context, req int32, callOptions ...callopt.Option) (resp *usersSrv.UserData, err error) {
	return c.kitexClient.FindUser(ctx, req, callOptions...)
}

func (c *clientImpl) VerifyCredentials(ctx context.Context, req *usersSrv.CredentialsReq, callOptions ...callopt.Option) (resp bool, err error) {
	return c.kitexClient.VerifyCredentials(ctx, req, callOptions...)
}

func (c *clientImpl) CreateOrUpdateUser(ctx context.Context, req *usersSrv.UserData, callOptions ...callopt.Option) (resp *usersSrv.UserData, err error) {
	return c.kitexClient.CreateOrUpdateUser(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteUser(ctx context.Context, req int32, callOptions ...callopt.Option) (resp *usersSrv.UserData, err error) {
	return c.kitexClient.DeleteUser(ctx, req, callOptions...)
}
