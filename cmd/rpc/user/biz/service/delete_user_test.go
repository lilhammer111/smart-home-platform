package service

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	"testing"
)

func TestDeleteUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	req := &common_rpc.RpcId{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
