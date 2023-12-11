package users_srv

import (
	"context"
	"git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/biz/service"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/dto/kitex_gen/users"
)

// UsersImpl implements the last service interface defined in the IDL.
type UsersImpl struct{}

// FindUser implements the UsersImpl interface.
func (s *UsersImpl) FindUser(ctx context.Context, req int32) (resp *users.UserData, err error) {
	resp, err = service.NewFindUserService(ctx).Run(req)

	return resp, err
}

// VerifyCredentials implements the UsersImpl interface.
func (s *UsersImpl) VerifyCredentials(ctx context.Context, req *users.CredentialsReq) (resp bool, err error) {
	resp, err = service.NewVerifyCredentialsService(ctx).Run(req)

	return resp, err
}

// CreateOrUpdateUser implements the UsersImpl interface.
func (s *UsersImpl) CreateOrUpdateUser(ctx context.Context, req *users.UserData) (resp *users.UserData, err error) {
	resp, err = service.NewCreateOrUpdateUserService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UsersImpl interface.
func (s *UsersImpl) DeleteUser(ctx context.Context, req int32) (resp *users.UserData, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}
