package service

import (
	"context"
	user "git.zqbjj.top/pet/services/cmd/rpc/user_srv/kitex_gen/user"
)

type VerifyCredentialsService struct {
	ctx context.Context
} // NewVerifyCredentialsService new VerifyCredentialsService
func NewVerifyCredentialsService(ctx context.Context) *VerifyCredentialsService {
	return &VerifyCredentialsService{ctx: ctx}
}

// Run create note info
func (s *VerifyCredentialsService) Run(req *user.CredentialsReq) (resp bool, err error) {
	// Finish your business logic.

	return
}
