package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type FreezePatrolAfterAuthService struct {
	ctx context.Context
} // NewFreezePatrolAfterAuthService new FreezePatrolAfterAuthService
func NewFreezePatrolAfterAuthService(ctx context.Context) *FreezePatrolAfterAuthService {
	return &FreezePatrolAfterAuthService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolAfterAuthService) Run(req *micro_user.RpcUserId) (resp *micro_user.RpcFreezeResp, err error) {
	// Finish your business logic.

	return
}
