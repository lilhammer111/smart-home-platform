package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type QueryUsersWithFilterService struct {
	ctx context.Context
} // NewQueryUsersWithFilterService new QueryUsersWithFilterService
func NewQueryUsersWithFilterService(ctx context.Context) *QueryUsersWithFilterService {
	return &QueryUsersWithFilterService{ctx: ctx}
}

// Run create note info
func (s *QueryUsersWithFilterService) Run(req *micro_user.RpcUsersFilterReq) (resp []*user.UserInfo, err error) {
	// Finish your business logic.

	return
}
