package service

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/utils/rpc_client/user_micro"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	"testing"
)

func TestFindUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFindUserService(ctx)
	// init req and assert value

	req := &common_rpc.RpcId{Id: 1}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test
	_, err = user_micro.DefaultClient().FindUser(ctx, req)
	if err != nil {
		t.Errorf("failed to get user detail: %s", err)
	}
}
