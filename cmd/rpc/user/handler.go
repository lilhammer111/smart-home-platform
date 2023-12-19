package main

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/service"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
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

// QueryUsersWithFilter implements the MicroUserImpl interface.
func (s *MicroUserImpl) QueryUsersWithFilter(ctx context.Context, req *micro_user.RpcUsersFilterReq) (resp []*user.UserInfo, err error) {
	resp, err = service.NewQueryUsersWithFilterService(ctx).Run(req)

	return resp, err
}

// SendSmsViaAliyun implements the MicroUserImpl interface.
func (s *MicroUserImpl) SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq) (resp *common.Empty, err error) {
	resp, err = service.NewSendSmsViaAliyunService(ctx).Run(req)

	return resp, err
}

// FreezePatrolAfterAuth implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolAfterAuth(ctx context.Context, userId int32) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = service.NewFreezePatrolAfterAuthService(ctx).Run(userId)

	return resp, err
}

// VerifySmsCode implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifySmsCode(ctx context.Context, mobile string, smsCode string) (resp *common.Empty, err error) {
	resp, err = service.NewVerifySmsCodeService(ctx).Run(mobile, smsCode)

	return resp, err
}

// VerifyUsernamePwd implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyUsernamePwd(ctx context.Context, username string, entryPwd string) (resp *common.Empty, err error) {
	resp, err = service.NewVerifyUsernamePwdService(ctx).Run(username, entryPwd)

	return resp, err
}

// VerifyEmailPwd implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyEmailPwd(ctx context.Context, email string, entryPwd string) (resp *common.Empty, err error) {
	resp, err = service.NewVerifyEmailPwdService(ctx).Run(email, entryPwd)

	return resp, err
}

// FindUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUser(ctx context.Context, userId int32) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserService(ctx).Run(userId)

	return resp, err
}

// DeleteUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) DeleteUser(ctx context.Context, userId int32) (resp *common.Empty, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(userId)

	return resp, err
}

// FindUserByMobile implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUserByMobile(ctx context.Context, req *micro_user.RpcFindUserByMobileReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserByMobileService(ctx).Run(req)

	return resp, err
}

// FindUserByUsername implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUserByUsername(ctx context.Context, username *micro_user.RpcFindUserByUsernameReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserByUsernameService(ctx).Run(username)

	return resp, err
}

// FindUserByOpenid implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUserByOpenid(ctx context.Context, req *micro_user.RpcFindUserByOpenidReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserByOpenidService(ctx).Run(req)

	return resp, err
}

// FreezePatrolBeforeVerify implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolBeforeVerify(ctx context.Context, req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = service.NewFreezePatrolBeforeVerifyService(ctx).Run(req)

	return resp, err
}

// FreezePatrolAfterVerify implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolAfterVerify(ctx context.Context, req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	resp, err = service.NewFreezePatrolAfterVerifyService(ctx).Run(req)

	return resp, err
}
