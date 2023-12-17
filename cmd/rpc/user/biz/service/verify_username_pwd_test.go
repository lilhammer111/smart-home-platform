package service

import (
	"context"
	"testing"
)

func TestVerifyUsernamePwd_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyUsernamePwdService(ctx)
	// init req and assert value

	username := &string{}

	entryPwd := &string{}
	resp, err := s.Run(usernameentryPwd)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
