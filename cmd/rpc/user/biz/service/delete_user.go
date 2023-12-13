package service

import (
	"context"
	standard "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/standard"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *standard.Req) (resp *user_micro.EmptyResp, err error) {
	// Finish your business logic.

	return
}
