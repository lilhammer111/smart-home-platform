package service

import (
	"context"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"testing"
)

func TestSendSmsViaAliyun_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendSmsViaAliyunService(ctx)
	// init req and assert value

	req := &micro_user.RpcSmsReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
