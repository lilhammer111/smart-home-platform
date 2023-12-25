package alertsrv

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type FindAlertService struct {
	ctx context.Context
}

// NewFindAlertService new FindAlertService
func NewFindAlertService(ctx context.Context) *FindAlertService {
	return &FindAlertService{ctx: ctx}
}

// Run create note info
func (s *FindAlertService) Run(req *micro_device.RpcFindAlertReq) (resp *alert.AlertInfo, err error) {
	resp = &alert.AlertInfo{}
	err = db.GetMysql().Model(&model.Alert{}).First(resp, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "alert info was not found")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	return resp, nil
}
