package service

import (
	"context"
	standard "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/standard"
	user_micro "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user_micro"
	"testing"
)

func TestFreezePatrolAfterAuth_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFreezePatrolAfterAuthService(ctx)
	// init req and assert value

	req := &standard.Req{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
