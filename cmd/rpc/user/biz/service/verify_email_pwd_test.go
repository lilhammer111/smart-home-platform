package service

import (
	"context"
	"testing"
)

func TestVerifyEmailPwd_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyEmailPwdService(ctx)
	// init req and assert value

	email := &string{}

	entryPwd := &string{}
	resp, err := s.Run(emailentryPwd)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
