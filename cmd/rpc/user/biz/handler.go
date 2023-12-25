package biz

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/service"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

// MicroUserImpl implements the last service interface defined in the IDL.
type MicroUserImpl struct{}

// SendSmsViaAliyun implements the MicroUserImpl interface.
func (s *MicroUserImpl) SendSmsViaAliyun(ctx context.Context, req *micro_user.RpcSmsReq) (resp *common.Empty, err error) {
	resp, err = service.NewSendSmsViaAliyunService(ctx).Run(req)

	return resp, err
}

// FreezePatrolBeforeVerify implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolBeforeVerify(ctx context.Context, req *micro_user.RpcFreezeReq) (resp *micro_user.RpcUserId, err error) {
	resp, err = service.NewFreezePatrolBeforeVerifyService(ctx).Run(req)

	return resp, err
}

// FreezePatrolAfterVerify implements the MicroUserImpl interface.
func (s *MicroUserImpl) FreezePatrolAfterVerify(ctx context.Context, req *micro_user.RpcAfterVerifyReq) (resp *common.Empty, err error) {
	resp, err = service.NewFreezePatrolAfterVerifyService(ctx).Run(req)

	return resp, err
}

// VerifySmsCode implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifySmsCode(ctx context.Context, req *micro_user.RpcVerifyCodeReq) (resp *micro_user.RpcVerifyResp, err error) {
	resp, err = service.NewVerifySmsCodeService(ctx).Run(req)

	return resp, err
}

// VerifyUsernamePwd implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyUsernamePwd(ctx context.Context, req *micro_user.RpcVerifyUsernamePwdReq) (resp *micro_user.RpcVerifyResp, err error) {
	resp, err = service.NewVerifyUsernamePwdService(ctx).Run(req)

	return resp, err
}

// VerifyEmailPwd implements the MicroUserImpl interface.
func (s *MicroUserImpl) VerifyEmailPwd(ctx context.Context, req *micro_user.RpcVerifyEmailPwdReq) (resp *micro_user.RpcVerifyResp, err error) {
	resp, err = service.NewVerifyEmailPwdService(ctx).Run(req)

	return resp, err
}

// FindUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUser(ctx context.Context, req *micro_user.RpcFindUserReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserService(ctx).Run(req)

	return resp, err
}

// FindUserByOpenid implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUserByOpenid(ctx context.Context, req *micro_user.RpcFindUserByOpenidReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserByOpenidService(ctx).Run(req)

	return resp, err
}

// FindUserByMobile implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUserByMobile(ctx context.Context, req *micro_user.RpcFindUserByMobileReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserByMobileService(ctx).Run(req)

	return resp, err
}

// FindUserByUsername implements the MicroUserImpl interface.
func (s *MicroUserImpl) FindUserByUsername(ctx context.Context, req *micro_user.RpcFindUserByUsernameReq) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserByUsernameService(ctx).Run(req)

	return resp, err
}

// QueryUsersWithFilter implements the MicroUserImpl interface.
func (s *MicroUserImpl) QueryUsersWithFilter(ctx context.Context, req *user.UsersFilter) (resp []*user.UserInfo, err error) {
	resp, err = service.NewQueryUsersWithFilterService(ctx).Run(req)

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

// DeleteUser implements the MicroUserImpl interface.
func (s *MicroUserImpl) DeleteUser(ctx context.Context, req *micro_user.RpcDeleteUserReq) (resp *common.Empty, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// RequestOpenId implements the MicroUserImpl interface.
func (s *MicroUserImpl) RequestOpenId(ctx context.Context, req *micro_user.RpcRequestOpenIdReq) (resp *micro_user.RpcRequestOpenIdResp, err error) {
	resp, err = service.NewRequestOpenIdService(ctx).Run(req)

	return resp, err
}
