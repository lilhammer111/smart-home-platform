package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type UpsertUserService struct {
	ctx context.Context
} // NewUpsertUserService new UpsertUserService
func NewUpsertUserService(ctx context.Context) *UpsertUserService {
	return &UpsertUserService{ctx: ctx}
}

// Run create note info
func (s *UpsertUserService) Run(req *micro_user.RpcUser) (resp *micro_user.RpcUser, err error) {
	// Finish your business logic.

	return
}
