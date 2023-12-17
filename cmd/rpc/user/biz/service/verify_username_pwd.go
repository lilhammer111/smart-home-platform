package service

import (
	"context"
)

type VerifyUsernamePwdService struct {
	ctx context.Context
} // NewVerifyUsernamePwdService new VerifyUsernamePwdService
func NewVerifyUsernamePwdService(ctx context.Context) *VerifyUsernamePwdService {
	return &VerifyUsernamePwdService{ctx: ctx}
}

// Run create note info
func (s *VerifyUsernamePwdService) Run(username string, entryPwd string) (resp bool, err error) {
	// Finish your business logic.

	return
}
