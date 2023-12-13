package service

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type FindUserService struct {
	ctx context.Context
} // NewFindUserService new FindUserService
func NewFindUserService(ctx context.Context) *FindUserService {
	return &FindUserService{ctx: ctx}
}

// Run create note info
func (s *FindUserService) Run(req *common_rpc.RpcId) (resp *micro_user.RpcUser, err error) {
	// Finish your business logic.

	return
}
