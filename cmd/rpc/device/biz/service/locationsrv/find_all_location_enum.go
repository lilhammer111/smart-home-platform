package locationsrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
)

type FindAllLocationEnumService struct {
	ctx context.Context
}

// NewFindAllLocationEnumService new FindAllLocationEnumService
func NewFindAllLocationEnumService(ctx context.Context) *FindAllLocationEnumService {
	return &FindAllLocationEnumService{ctx: ctx}
}

// Run create note info
func (s *FindAllLocationEnumService) Run() (resp []*micro_device.LocationData, err error) {
	locInfos := make([]model.Location, 0)
	err = db.GetMysql().Find(&locInfos).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = make([]*micro_device.LocationData, 0)
	err = copier.Copy(resp, locInfos)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
