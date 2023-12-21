package devicesrv

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/device"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FindDeviceService struct {
	ctx context.Context
}

// NewFindDeviceService new FindDeviceService
func NewFindDeviceService(ctx context.Context) *FindDeviceService {
	return &FindDeviceService{ctx: ctx}
}

// Run create note info
func (s *FindDeviceService) Run(req *micro_device.RpcFindDeviceReq) (resp *device.DeviceInfo, err error) {
	deviceInfo := model.Device{}
	err = db.GetMysql().First(&deviceInfo, req.DeviceId).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "device is not existed")
		}
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
