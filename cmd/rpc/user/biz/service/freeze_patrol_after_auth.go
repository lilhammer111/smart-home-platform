package service

import (
	"context"
	standard "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/standard"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
)

type FreezePatrolAfterAuthService struct {
	ctx context.Context
} // NewFreezePatrolAfterAuthService new FreezePatrolAfterAuthService
func NewFreezePatrolAfterAuthService(ctx context.Context) *FreezePatrolAfterAuthService {
	return &FreezePatrolAfterAuthService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolAfterAuthService) Run(req *standard.Req) (resp *user_micro.FreezeResp, err error) {
	// Finish your business logic.

	return
}
