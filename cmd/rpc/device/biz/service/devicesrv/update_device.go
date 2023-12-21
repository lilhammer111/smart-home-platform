package devicesrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type UpdateDeviceService struct {
	ctx context.Context
} // NewUpdateDeviceService new UpdateDeviceService
func NewUpdateDeviceService(ctx context.Context) *UpdateDeviceService {
	return &UpdateDeviceService{ctx: ctx}
}

// Run create note info
func (s *UpdateDeviceService) Run(req *device.DeviceInfo) (resp *device.DeviceInfo, err error) {
	deviceInfo := model.Device{}
	err = copier.Copy(&deviceInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	// Using Save for a full update, as it ensures all fields in device model, including zero values, are updated.
	// This is appropriate since all necessary fields are validated and included at the API gateway level.
	err = db.GetMysql().Save(&deviceInfo).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &device.DeviceInfo{}
	err = copier.Copy(resp, deviceInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
