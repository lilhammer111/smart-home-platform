package devicesrv

import (
	"context"
	device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	"testing"
)

func TestCreateDevice_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateDeviceService(ctx)
	// init req and assert value

	req := &device.DeviceInfo{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
