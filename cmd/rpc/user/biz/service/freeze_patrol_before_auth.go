package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
)

type FreezePatrolBeforeAuthService struct {
	ctx context.Context
} // NewFreezePatrolBeforeAuthService new FreezePatrolBeforeAuthService
func NewFreezePatrolBeforeAuthService(ctx context.Context) *FreezePatrolBeforeAuthService {
	return &FreezePatrolBeforeAuthService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolBeforeAuthService) Run(req *micro_user.RpcFreezeReq) (resp *micro_user.RpcFreezeResp, err error) {
	// Finish your business logic.

	return
}
