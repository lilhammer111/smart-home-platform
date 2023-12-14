package user_micro

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/kitex_gen/micro_user/microuser"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() microuser.Client
	Service() string
	FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error)

	FreezePatrolAfterAuth(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error)

	VerifyCredentials(ctx context.Context, req *micro_user.RpcCredentialReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error)

	FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcUser, err error)

	QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (resp []*micro_user.RpcUser, err error)

	UpsertUser(ctx context.Context, req *micro_user.RpcUser, callOptions ...callopt.Option) (resp *micro_user.RpcUser, err error)

	DeleteUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := microuser.NewClient(dstService, opts...)
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
	kitexClient microuser.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() microuser.Client {
	return c.kitexClient
}

func (c *clientImpl) FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	return c.kitexClient.FreezePatrolBeforeAuth(ctx, req, callOptions...)
}

func (c *clientImpl) FreezePatrolAfterAuth(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	return c.kitexClient.FreezePatrolAfterAuth(ctx, req, callOptions...)
}

func (c *clientImpl) VerifyCredentials(ctx context.Context, req *micro_user.RpcCredentialReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	return c.kitexClient.VerifyCredentials(ctx, req, callOptions...)
}

func (c *clientImpl) FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcUser, err error) {
	return c.kitexClient.FindUser(ctx, req, callOptions...)
}

func (c *clientImpl) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (resp []*micro_user.RpcUser, err error) {
	return c.kitexClient.QueryUsersWithFilter(ctx, req, callOptions...)
}

func (c *clientImpl) UpsertUser(ctx context.Context, req *micro_user.RpcUser, callOptions ...callopt.Option) (resp *micro_user.RpcUser, err error) {
	return c.kitexClient.UpsertUser(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	return c.kitexClient.DeleteUser(ctx, req, callOptions...)
}
