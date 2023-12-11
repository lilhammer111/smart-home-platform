package service

import (
	"context"
	users "git.zqbjj.top/pet/services/pet-feeder/cmd/rpc/users_srv/dto/kitex_gen/users"
)

type CreateOrUpdateUserService struct {
	ctx context.Context
} // NewCreateOrUpdateUserService new CreateOrUpdateUserService
func NewCreateOrUpdateUserService(ctx context.Context) *CreateOrUpdateUserService {
	return &CreateOrUpdateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateOrUpdateUserService) Run(req *users.UserData) (resp *users.UserData, err error) {
	// Finish your business logic.

	return
}
