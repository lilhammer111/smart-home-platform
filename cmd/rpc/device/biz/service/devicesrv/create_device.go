package devicesrv

import (
	"context"
	device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
)

type CreateDeviceService struct {
	ctx context.Context
} // NewCreateDeviceService new CreateDeviceService
func NewCreateDeviceService(ctx context.Context) *CreateDeviceService {
	return &CreateDeviceService{ctx: ctx}
}

// Run create note info
func (s *CreateDeviceService) Run(req *device.DeviceInfo) (resp *device.DeviceInfo, err error) {
	// Finish your business logic.

	return
}
