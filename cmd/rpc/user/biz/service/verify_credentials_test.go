package service

import (
	"context"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
	"testing"
)

func TestVerifyCredentials_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyCredentialsService(ctx)
	// init req and assert value

	req := &user_micro.CredentialsReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
