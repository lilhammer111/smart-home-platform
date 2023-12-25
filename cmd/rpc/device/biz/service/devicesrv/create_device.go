package devicesrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type CreateDeviceService struct {
	ctx context.Context
} // NewCreateDeviceService new CreateDeviceService
func NewCreateDeviceService(ctx context.Context) *CreateDeviceService {
	return &CreateDeviceService{ctx: ctx}
}

// Run create note info
func (s *CreateDeviceService) Run(req *device.DeviceInfo) (resp *device.DeviceInfo, err error) {
	req.Id = nil
	deviceInfo := model.Device{}
	err = copier.Copy(&deviceInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&deviceInfo).Error
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
