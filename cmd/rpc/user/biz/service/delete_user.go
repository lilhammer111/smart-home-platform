package service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(userId int32) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
