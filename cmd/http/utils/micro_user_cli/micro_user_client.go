package micro_user_cli

import (
	"context"
	common_rpcSrv "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common_rpc"
	micro_userSrv "git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user"
	userSrv "git.zqbjj.top/pet/services/cmd/http/kitex_gen/user"

	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_user/microuser"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() microuser.Client
	Service() string
	SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error)

	FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error)

	FreezePatrolAfterAuth(ctx context.Context, id *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error)

	VerifySmsCode(ctx context.Context, mobile string, smsCode string, callOptions ...callopt.Option) (resp bool, err error)

	VerifyUsernamePwd(ctx context.Context, username string, entryPwd string, callOptions ...callopt.Option) (resp bool, err error)

	VerifyEmailPwd(ctx context.Context, email string, entryPwd string, callOptions ...callopt.Option) (resp bool, err error)

	FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *user.UserInfo, err error)

	QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (resp []*user.UserInfo, err error)

	UpdateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error)

	CreateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error)

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

func (c *clientImpl) SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	return c.kitexClient.SendSmsViaAliyun(ctx, req, callOptions...)
}

func (c *clientImpl) FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	return c.kitexClient.FreezePatrolBeforeAuth(ctx, req, callOptions...)
}

func (c *clientImpl) FreezePatrolAfterAuth(ctx context.Context, id *common_rpc.RpcId, callOptions ...callopt.Option) (resp *micro_user.RpcFreezeResp, err error) {
	return c.kitexClient.FreezePatrolAfterAuth(ctx, id, callOptions...)
}

func (c *clientImpl) VerifySmsCode(ctx context.Context, mobile string, smsCode string, callOptions ...callopt.Option) (resp bool, err error) {
	return c.kitexClient.VerifySmsCode(ctx, mobile, smsCode, callOptions...)
}

func (c *clientImpl) VerifyUsernamePwd(ctx context.Context, username string, entryPwd string, callOptions ...callopt.Option) (resp bool, err error) {
	return c.kitexClient.VerifyUsernamePwd(ctx, username, entryPwd, callOptions...)
}

func (c *clientImpl) VerifyEmailPwd(ctx context.Context, email string, entryPwd string, callOptions ...callopt.Option) (resp bool, err error) {
	return c.kitexClient.VerifyEmailPwd(ctx, email, entryPwd, callOptions...)
}

func (c *clientImpl) FindUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	return c.kitexClient.FindUser(ctx, req, callOptions...)
}

func (c *clientImpl) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq, callOptions ...callopt.Option) (resp []*user.UserInfo, err error) {
	return c.kitexClient.QueryUsersWithFilter(ctx, req, callOptions...)
}

func (c *clientImpl) UpdateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	return c.kitexClient.UpdateUser(ctx, req, callOptions...)
}

func (c *clientImpl) CreateUser(ctx context.Context, req *user.UserInfo, callOptions ...callopt.Option) (resp *user.UserInfo, err error) {
	return c.kitexClient.CreateUser(ctx, req, callOptions...)
}

func (c *clientImpl) DeleteUser(ctx context.Context, req *common_rpc.RpcId, callOptions ...callopt.Option) (resp *common_rpc.RpcEmpty, err error) {
	return c.kitexClient.DeleteUser(ctx, req, callOptions...)
}
