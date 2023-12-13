package service

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type FreezePatrolAfterAuthService struct {
	ctx context.Context
} // NewFreezePatrolAfterAuthService new FreezePatrolAfterAuthService
func NewFreezePatrolAfterAuthService(ctx context.Context) *FreezePatrolAfterAuthService {
	return &FreezePatrolAfterAuthService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolAfterAuthService) Run(req *common_rpc.RpcId) (resp *micro_user.RpcFreezeResp, err error) {
	// Finish your business logic.

	return
}
