package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type UpsertUserService struct {
	ctx context.Context
} // NewUpsertUserService new UpsertUserService
func NewUpsertUserService(ctx context.Context) *UpsertUserService {
	return &UpsertUserService{ctx: ctx}
}

// Run create note info
func (s *UpsertUserService) Run(req *user.UserInfo) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	return
}
