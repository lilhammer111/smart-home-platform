package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type FindUserByOpenidService struct {
	ctx context.Context
} // NewFindUserByOpenidService new FindUserByOpenidService
func NewFindUserByOpenidService(ctx context.Context) *FindUserByOpenidService {
	return &FindUserByOpenidService{ctx: ctx}
}

// Run create note info
func (s *FindUserByOpenidService) Run(req *micro_user.RpcFindUserByOpenidReq) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	return
}
