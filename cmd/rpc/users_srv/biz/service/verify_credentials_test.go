package service

import (
	"context"
	users "git.zqbjj.top/pet/services/cmd/rpc/users_srv/dto/kitex_gen/users"
	"testing"
)

func TestVerifyCredentials_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyCredentialsService(ctx)
	// init req and assert value

	req := &users.CredentialsReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == false {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
