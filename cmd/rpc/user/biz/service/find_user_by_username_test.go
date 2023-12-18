package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/user"
	"testing"
)

func TestFindUserByUsername_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFindUserByUsernameService(ctx)
	// init req and assert value

	username := &micro_user.RpcFindUserByUsernameReq{}
	resp, err := s.Run(username)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
