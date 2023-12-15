package service

import (
	"context"
	common_rpc "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/common_rpc"
	micro_user "git.zqbjj.top/pet/services/cmd/rpc/user/kitex_gen/micro_user"
	"testing"
)

func TestSendSMSViaAliyun_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendSMSViaAliyunService(ctx)
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