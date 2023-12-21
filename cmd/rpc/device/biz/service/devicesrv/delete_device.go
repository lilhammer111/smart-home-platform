package devicesrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteDeviceService struct {
	ctx context.Context
} // NewDeleteDeviceService new DeleteDeviceService
func NewDeleteDeviceService(ctx context.Context) *DeleteDeviceService {
	return &DeleteDeviceService{ctx: ctx}
}

// Run create note info
func (s *DeleteDeviceService) Run(req *micro_device.RpcDeleteDeviceReq) (resp *common.Empty, err error) {
	err = db.GetMysql().Delete(&model.Device{}, req.DeviceId).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, "failed to delete devicesrv")
	}

	return &common.Empty{}, nil
}
