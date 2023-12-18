package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type FindUserService struct {
	ctx context.Context
} // NewFindUserService new FindUserService
func NewFindUserService(ctx context.Context) *FindUserService {
	return &FindUserService{ctx: ctx}
}

// Run create note info
func (s *FindUserService) Run(req *micro_user.RpcFindUserReq) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	return
}
