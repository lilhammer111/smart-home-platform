package alertsrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/alert"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type CreateAlertService struct {
	ctx context.Context
}

// NewCreateAlertService new CreateAlertService
func NewCreateAlertService(ctx context.Context) *CreateAlertService {
	return &CreateAlertService{ctx: ctx}
}

// Run create note info
func (s *CreateAlertService) Run(req *alert.AlertInfo) (resp *alert.AlertInfo, err error) {
	alertInfo := model.Alert{}
	err = copier.Copy(&alertInfo, req)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	err = db.GetMysql().Create(&alertInfo).Error
	if err != nil {
		klog.Error(err)
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
