package biz

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/service/alertsrv"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/service/devicesrv"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/service/locationsrv"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
)

// MicroDeviceImpl implements the last service interface defined in the IDL.
type MicroDeviceImpl struct{}

// FindDevice implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) FindDevice(ctx context.Context, req *micro_device.RpcFindDeviceReq) (resp *device.DeviceInfo, err error) {
	resp, err = devicesrv.NewFindDeviceService(ctx).Run(req)

	return resp, err
}

// QueryDevicesWithFilter implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) QueryDevicesWithFilter(ctx context.Context, req *device.DeviceFilter) (resp []*device.DeviceInfo, err error) {
	resp, err = devicesrv.NewQueryDevicesWithFilterService(ctx).Run(req)

	return resp, err
}

// CreateDevice implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) CreateDevice(ctx context.Context, req *device.DeviceInfo) (resp *device.DeviceInfo, err error) {
	resp, err = devicesrv.NewCreateDeviceService(ctx).Run(req)

	return resp, err
}

// UpdateDevice implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) UpdateDevice(ctx context.Context, req *device.DeviceInfo) (resp *device.DeviceInfo, err error) {
	resp, err = devicesrv.NewUpdateDeviceService(ctx).Run(req)

	return resp, err
}

// DeleteDevice implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) DeleteDevice(ctx context.Context, req *micro_device.RpcDeleteDeviceReq) (resp *common.Empty, err error) {
	resp, err = devicesrv.NewDeleteDeviceService(ctx).Run(req)

	return resp, err
}

// FindLocationTitle implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) FindLocationTitle(ctx context.Context, req *micro_device.RpcFindLocReq) (resp *micro_device.LocationData, err error) {
	resp, err = locationsrv.NewFindLocationTitleService(ctx).Run(req)

	return resp, err
}

// FindAllLocationEnum implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) FindAllLocationEnum(ctx context.Context) (resp []*micro_device.LocationData, err error) {
	resp, err = locationsrv.NewFindAllLocationEnumService(ctx).Run()

	return resp, err
}

// ReduceLocationEnum implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) ReduceLocationEnum(ctx context.Context, req *micro_device.RpcReduceLocReq) (resp *common.Empty, err error) {
	resp, err = locationsrv.NewReduceLocationEnumService(ctx).Run(req)

	return resp, err
}

// ExpandLocationEnum implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) ExpandLocationEnum(ctx context.Context, req *micro_device.RpcExpandLocReq) (resp *common.Empty, err error) {
	resp, err = locationsrv.NewExpandLocationEnumService(ctx).Run(req)

	return resp, err
}

// FindAlert implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) FindAlert(ctx context.Context, req *micro_device.RpcFindDeviceReq) (resp *alert.AlertInfo, err error) {
	resp, err = alertsrv.NewFindAlertService(ctx).Run(req)

	return resp, err
}

// QueryAlertsWithFilter implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) QueryAlertsWithFilter(ctx context.Context, req *alert.AlertFilter) (resp []*alert.AlertInfo, err error) {
	resp, err = alertsrv.NewQueryAlertsWithFilterService(ctx).Run(req)

	return resp, err
}

// CreateAlert implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) CreateAlert(ctx context.Context, req *alert.AlertInfo) (resp *alert.AlertInfo, err error) {
	resp, err = alertsrv.NewCreateAlertService(ctx).Run(req)

	return resp, err
}

// UpdateAlert implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) UpdateAlert(ctx context.Context, req *alert.AlertInfo) (resp *alert.AlertInfo, err error) {
	resp, err = alertsrv.NewUpdateAlertService(ctx).Run(req)

	return resp, err
}

// DeleteAlert implements the MicroDeviceImpl interface.
func (s *MicroDeviceImpl) DeleteAlert(ctx context.Context, req *micro_device.RpcDeleteAlertReq) (resp *common.Empty, err error) {
	resp, err = alertsrv.NewDeleteAlertService(ctx).Run(req)

	return resp, err
}
