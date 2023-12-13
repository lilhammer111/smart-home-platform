package main

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/biz/service"
	standard "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/standard"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// FreezePatrolBeforeAuth implements the UserImpl interface.
func (s *UserImpl) FreezePatrolBeforeAuth(ctx context.Context, req *user_micro.AuthQuery) (resp *user_micro.FreezeResp, err error) {
	resp, err = service.NewFreezePatrolBeforeAuthService(ctx).Run(req)

	return resp, err
}

// FreezePatrolAfterAuth implements the UserImpl interface.
func (s *UserImpl) FreezePatrolAfterAuth(ctx context.Context, req *standard.Req) (resp *user_micro.FreezeResp, err error) {
	resp, err = service.NewFreezePatrolAfterAuthService(ctx).Run(req)

	return resp, err
}

// VerifyCredentials implements the UserImpl interface.
func (s *UserImpl) VerifyCredentials(ctx context.Context, req *user_micro.CredentialsReq) (resp *user_micro.EmptyResp, err error) {
	resp, err = service.NewVerifyCredentialsService(ctx).Run(req)

	return resp, err
}

// FindUser implements the UserImpl interface.
func (s *UserImpl) FindUser(ctx context.Context, req *standard.Req) (resp *user.UserInfo, err error) {
	resp, err = service.NewFindUserService(ctx).Run(req)

	return resp, err
}

// QueryUsersWithFilter implements the UserImpl interface.
func (s *UserImpl) QueryUsersWithFilter(ctx context.Context, req *user.UsersFilter) (resp []*user.UserInfo, err error) {
	resp, err = service.NewQueryUsersWithFilterService(ctx).Run(req)

	return resp, err
}

// UpsertUser implements the UserImpl interface.
func (s *UserImpl) UpsertUser(ctx context.Context, req *user.UserInfo) (resp *user.UserInfo, err error) {
	resp, err = service.NewUpsertUserService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserImpl interface.
func (s *UserImpl) DeleteUser(ctx context.Context, req *standard.Req) (resp *user_micro.EmptyResp, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}
