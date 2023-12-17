package main

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/service"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
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

// UpdateUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) UpdateUser(ctx context.Context, req *user.UserInfo) (resp *user.UserInfo, err error) {
	resp, err = service.NewUpdateUserService(ctx).Run(req)

	return resp, err
}

// CreateUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) CreateUser(ctx context.Context, req *user.UserInfo) (resp *user.UserInfo, err error) {
	resp, err = service.NewCreateUserService(ctx).Run(req)

	return resp, err
}

// FindUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUser(ctx context.Context, req *common_rpc.RpcId) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserService(ctx).Run(req)

	return resp, err
}

// QueryUsersWithFilter implements the MicroUserImpl interface.
func (s *MicroUserImpl) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq) (resp []*user.UserInfo, err error) {
	resp, err = service.NewQueryUsersWithFilterService(ctx).Run(req)

	return resp, err
}

// SendSmsViaAliyun implements the MicroUserImpl interface.
func (s *MicroUserImpl) SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq) (resp *common_rpc.RpcEmpty, err error) {
	resp, err = service.NewSendSmsViaAliyunService(ctx).Run(req)

	return resp, err
}

// VerifySmsCode implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifySmsCode(ctx context.Context, req *micro_user.RpcMobileReq) (resp bool, err error) {
	resp, err = service.NewVerifySmsCodeService(ctx).Run(req)

	return resp, err
}

// VerifyUsernamePwd implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyUsernamePwd(ctx context.Context, req *micro_user.RpcUsernameReq) (resp bool, err error) {
	resp, err = service.NewVerifyUsernamePwdService(ctx).Run(req)

	return resp, err
}

// VerifyEmailPwd implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyEmailPwd(ctx context.Context, req *micro_user.RpcEmailReq) (resp bool, err error) {
	resp, err = service.NewVerifyEmailPwdService(ctx).Run(req)

	return resp, err
}
