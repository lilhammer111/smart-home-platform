package service

import (
	"context"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/dto/kitex_gen/users"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req int32) (resp *users.UserData, err error) {
	// Finish your business logic.

	return
}
