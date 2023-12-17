package service

import (
	"context"
)

type VerifyEmailPwdService struct {
	ctx context.Context
} // NewVerifyEmailPwdService new VerifyEmailPwdService
func NewVerifyEmailPwdService(ctx context.Context) *VerifyEmailPwdService {
	return &VerifyEmailPwdService{ctx: ctx}
}

// Run create note info
func (s *VerifyEmailPwdService) Run(email string, entryPwd string) (resp bool, err error) {
	// Finish your business logic.

	return
}
