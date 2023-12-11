package service

import (
	"context"
	users "git.zqbjj.top/pet/services/cmd/rpc/users_srv/dto/kitex_gen/users"
)

type VerifyCredentialsService struct {
	ctx context.Context
} // NewVerifyCredentialsService new VerifyCredentialsService
func NewVerifyCredentialsService(ctx context.Context) *VerifyCredentialsService {
	return &VerifyCredentialsService{ctx: ctx}
}

// Run create note info
func (s *VerifyCredentialsService) Run(req *users.CredentialsReq) (resp bool, err error) {
	// Finish your business logic.

	return
}
