package service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common"
	"testing"
)

func TestDeleteUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteUserService(ctx)
	// init req and assert value

	userId := &int32{}
	resp, err := s.Run(userId)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
