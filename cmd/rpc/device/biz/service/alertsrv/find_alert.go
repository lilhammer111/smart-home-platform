package alertsrv

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	alert "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
	micro_device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FindAlertService struct {
	ctx context.Context
} // NewFindAlertService new FindAlertService
func NewFindAlertService(ctx context.Context) *FindAlertService {
	return &FindAlertService{ctx: ctx}
}

// Run create note info
func (s *FindAlertService) Run(req *micro_device.RpcFindDeviceReq) (resp *alert.AlertInfo, err error) {
	alertInfo := model.Alert{}
	err = db.GetMysql().First(&alertInfo, req.DeviceId).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "alert is not existed")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &alert.AlertInfo{}
	err = copier.Copy(resp, alertInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
