package service

import (
	"context"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
)

type FreezePatrolBeforeAuthService struct {
	ctx context.Context
} // NewFreezePatrolBeforeAuthService new FreezePatrolBeforeAuthService
func NewFreezePatrolBeforeAuthService(ctx context.Context) *FreezePatrolBeforeAuthService {
	return &FreezePatrolBeforeAuthService{ctx: ctx}
}

// Run create note info
func (s *FreezePatrolBeforeAuthService) Run(req *user_micro.AuthQuery) (resp *user_micro.FreezeResp, err error) {
	// Finish your business logic.

	return
}
