package locationsrv

import (
	"context"
	"errors"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FindLocationTitleService struct {
	ctx context.Context
}

// NewFindLocationTitleService new FindLocationTitleService
func NewFindLocationTitleService(ctx context.Context) *FindLocationTitleService {
	return &FindLocationTitleService{ctx: ctx}
}

// Run create note info
func (s *FindLocationTitleService) Run(req *micro_device.RpcFindLocReq) (resp *micro_device.LocationData, err error) {
	locInfo := model.Location{}
	err = db.GetMysql().First(&locInfo, req.Id).Error
	if err != nil {
		klog.Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kerrors.NewBizStatusError(code.NotFound, "the location info is not existed")
		}
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	resp = &micro_device.LocationData{}
	err = copier.Copy(resp, locInfo)
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.InternalError, msg.InternalError)
	}

	return resp, nil
}
