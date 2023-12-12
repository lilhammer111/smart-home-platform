package main

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user_srv/biz/service"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// FindUserByID implements the UserImpl interface.
func (s *UserImpl) FindUserByID(ctx context.Context, req int32) (resp *user.UserData, err error) {
	resp, err = service.NewFindUserByIDService(ctx).Run(req)

	return resp, err
}

// GetUsersByFilter implements the UserImpl interface.
func (s *UserImpl) GetUsersByFilter(ctx context.Context, req *user.UsersFilterReq) (resp []*user.UserData, err error) {
	resp, err = service.NewGetUsersByFilterService(ctx).Run(req)

	return resp, err
}

// VerifyCredentials implements the UserImpl interface.
func (s *UserImpl) VerifyCredentials(ctx context.Context, req *user.CredentialsReq) (resp bool, err error) {
	resp, err = service.NewVerifyCredentialsService(ctx).Run(req)

	return resp, err
}

// CreateOrUpdateUser implements the UserImpl interface.
func (s *UserImpl) CreateOrUpdateUser(ctx context.Context, req *user.UserData) (resp *user.UserData, err error) {
	resp, err = service.NewCreateOrUpdateUserService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserImpl interface.
func (s *UserImpl) DeleteUser(ctx context.Context, req int32) (resp *user.UserData, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}
