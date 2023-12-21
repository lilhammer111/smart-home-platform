package locationsrv

import (
	"context"
	micro_device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"testing"
)

func TestExpandLocationEnum_Run(t *testing.T) {
	ctx := context.Background()
	s := NewExpandLocationEnumService(ctx)
	// init req and assert value

	req := &micro_device.RpcExpandLocReq{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
