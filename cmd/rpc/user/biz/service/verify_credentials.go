package service

import (
	"context"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
)

type VerifyCredentialsService struct {
	ctx context.Context
} // NewVerifyCredentialsService new VerifyCredentialsService
func NewVerifyCredentialsService(ctx context.Context) *VerifyCredentialsService {
	return &VerifyCredentialsService{ctx: ctx}
}

// Run create note info
func (s *VerifyCredentialsService) Run(req *user_micro.CredentialsReq) (resp *user_micro.EmptyResp, err error) {
	// Finish your business logic.

	return
}
