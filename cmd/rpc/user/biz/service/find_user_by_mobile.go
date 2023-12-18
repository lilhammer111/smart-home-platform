package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type FindUserByMobileService struct {
	ctx context.Context
} // NewFindUserByMobileService new FindUserByMobileService
func NewFindUserByMobileService(ctx context.Context) *FindUserByMobileService {
	return &FindUserByMobileService{ctx: ctx}
}

// Run create note info
func (s *FindUserByMobileService) Run(req *micro_user.RpcFindUserByMobileReq) (resp *user.UserInfo, err error) {
	// Finish your business logic.

	return
}
