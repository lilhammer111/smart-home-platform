package devicesrv

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
)

type QueryDevicesWithFilterService struct {
	ctx context.Context
}

// NewQueryDevicesWithFilterService new QueryDevicesWithFilterService
func NewQueryDevicesWithFilterService(ctx context.Context) *QueryDevicesWithFilterService {
	return &QueryDevicesWithFilterService{ctx: ctx}
}

// Run create note info
func (s *QueryDevicesWithFilterService) Run(req *device.DeviceFilter) (resp []*device.DeviceInfo, err error) {

	return
}
