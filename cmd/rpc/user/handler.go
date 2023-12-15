package main

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/service"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

// MicroUserImpl implements the last service interface defined in the IDL.
type MicroUserImpl struct{}

// FreezePatrolBeforeAuth implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolBeforeAuth(ctx context.Context, req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = service.NewFreezePatrolBeforeAuthService(ctx).Run(req)

	return resp, err
}

// FreezePatrolAfterAuth implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolAfterAuth(ctx context.Context, req *common_rpc.RpcId) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = service.NewFreezePatrolAfterAuthService(ctx).Run(req)

	return resp, err
}

// VerifyCredentials implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyCredentials(ctx context.Context, req *micro_user.RpcCredentialReq) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = service.NewVerifyCredentialsService(ctx).Run(req)

	return resp, err
}

// FindUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUser(ctx context.Context, req *common_rpc.RpcId) (resp *micro_user.RpcUser, err error) {
	resp, err = service.NewFindUserService(ctx).Run(req)

	return resp, err
}

// QueryUsersWithFilter implements the MicroUserImpl interface.
func (s *MicroUserImpl) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq) (resp []*micro_user.RpcUser, err error) {
	resp, err = service.NewQueryUsersWithFilterService(ctx).Run(req)

	return resp, err
}

// UpsertUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) UpsertUser(ctx context.Context, req *micro_user.RpcUser) (resp *micro_user.RpcUser, err error) {
	resp, err = service.NewUpsertUserService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) DeleteUser(ctx context.Context, req *common_rpc.RpcId) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// SendSMSViaAliyun implements the MicroUserImpl interface.
func (s *MicroUserImpl) SendSMSViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = service.NewSendSMSViaAliyunService(ctx).Run(req)

	return resp, err
}
