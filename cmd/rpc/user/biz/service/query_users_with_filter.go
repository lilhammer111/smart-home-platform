package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
)

type QueryUsersWithFilterService struct {
	ctx context.Context
} // NewQueryUsersWithFilterService new QueryUsersWithFilterService
func NewQueryUsersWithFilterService(ctx context.Context) *QueryUsersWithFilterService {
	return &QueryUsersWithFilterService{ctx: ctx}
}

// Run create note info
func (s *QueryUsersWithFilterService) Run(req *user.UsersFilter) (resp []*user.UserInfo, err error) {
	// Finish your business logic.

	return
}
