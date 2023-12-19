package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type FreezePatrolAfterVerifyService struct {
	ctx context.Context
} // NewFreezePatrolAfterVerifyService new FreezePatrolAfterVerifyService
func NewFreezePatrolAfterVerifyService(ctx context.Context) *FreezePatrolAfterVerifyService {
	return &FreezePatrolAfterVerifyService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolAfterVerifyService) Run(req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	// Finish your business logic.

	return
}
