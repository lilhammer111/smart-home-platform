package service

import (
	"context"
)

type VerifySmsCodeService struct {
	ctx context.Context
} // NewVerifySmsCodeService new VerifySmsCodeService
func NewVerifySmsCodeService(ctx context.Context) *VerifySmsCodeService {
	return &VerifySmsCodeService{ctx: ctx}
}

// Run create note info
func (s *VerifySmsCodeService) Run(mobile string, smsCode string) (resp bool, err error) {
	// Finish your business logic.

	return
}
