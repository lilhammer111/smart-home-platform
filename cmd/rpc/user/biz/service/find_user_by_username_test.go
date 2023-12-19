package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"testing"
)

func TestFindUserByUsername_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFindUserByUsernameService(ctx)
	// init req and assert value

	req := &micro_user.RpcFindUserByUsernameReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}