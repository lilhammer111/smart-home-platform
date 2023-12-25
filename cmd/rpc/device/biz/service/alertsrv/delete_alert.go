package alertsrv

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	common "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	micro_device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type DeleteAlertService struct {
	ctx context.Context
} // NewDeleteAlertService new DeleteAlertService
func NewDeleteAlertService(ctx context.Context) *DeleteAlertService {
	return &DeleteAlertService{ctx: ctx}
}

// Run create note info
func (s *DeleteAlertService) Run(req *micro_device.RpcDeleteAlertReq) (resp *common.Empty, err error) {
	err = db.GetMysql().First(&model.Alert{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "alert info was not existed")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	err = db.GetMysql().Delete(&model.Alert{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	return &common.Empty{}, err
}
