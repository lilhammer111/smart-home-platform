package devicesrv

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type DeleteDeviceService struct {
	ctx context.Context
} // NewDeleteDeviceService new DeleteDeviceService
func NewDeleteDeviceService(ctx context.Context) *DeleteDeviceService {
	return &DeleteDeviceService{ctx: ctx}
}

// Run create note info
func (s *DeleteDeviceService) Run(req *micro_device.RpcDeleteDeviceReq) (resp *common.Empty, err error) {
	err = db.GetMysql().First(&model.Device{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "device info was not existed")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	err = db.GetMysql().Delete(&model.Device{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, "failed to delete device")
	}

	return &common.Empty{}, nil
}
