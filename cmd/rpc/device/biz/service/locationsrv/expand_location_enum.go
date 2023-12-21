package locationsrv

import (
	"context"
	"git.zqbjj.top/lilhammer111/micro-kit/error/code"
	"git.zqbjj.top/lilhammer111/micro-kit/error/msg"
	"git.zqbjj.top/lilhammer111/micro-kit/initializer/db"
	"git.zqbjj.top/pet/services/cmd/rpc/device/biz/model"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ExpandLocationEnumService struct {
	ctx context.Context
} // NewExpandLocationEnumService new ExpandLocationEnumService
func NewExpandLocationEnumService(ctx context.Context) *ExpandLocationEnumService {
	return &ExpandLocationEnumService{ctx: ctx}
}

// Run create note info
func (s *ExpandLocationEnumService) Run(req *micro_device.RpcExpandLocReq) (resp *common.Empty, err error) {
	locInfo := model.Location{Title: req.Title}

	err = db.GetMysql().Create(&locInfo).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	return &common.Empty{}, nil
}
