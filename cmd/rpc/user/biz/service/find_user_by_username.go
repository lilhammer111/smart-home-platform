package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type FindUserByUsernameService struct {
	ctx context.Context
} // NewFindUserByUsernameService new FindUserByUsernameService
func NewFindUserByUsernameService(ctx context.Context) *FindUserByUsernameService {
	return &FindUserByUsernameService{ctx: ctx}
}

// Run create note info
func (s *FindUserByUsernameService) Run(username *micro_user.RpcFindUserByUsernameReq) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	return
}
