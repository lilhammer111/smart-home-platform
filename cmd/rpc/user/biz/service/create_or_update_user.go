package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
)

type CreateOrUpdateUserService struct {
	ctx context.Context
} // NewCreateOrUpdateUserService new CreateOrUpdateUserService
func NewCreateOrUpdateUserService(ctx context.Context) *CreateOrUpdateUserService {
	return &CreateOrUpdateUserService{ctx: ctx}
}

// Run create note info
func (s *CreateOrUpdateUserService) Run(req *user.UserData) (resp *user.UserData, err error) {
	// Finish your business logic.

	return
}
