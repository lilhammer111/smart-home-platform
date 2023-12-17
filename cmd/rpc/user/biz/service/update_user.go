package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type UpdateUserService struct {
	ctx context.Context
} // NewUpdateUserService new UpdateUserService
func NewUpdateUserService(ctx context.Context) *UpdateUserService {
	return &UpdateUserService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserService) Run(req *user.UserInfo) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	return
}
