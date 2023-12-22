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

type ReduceLocationEnumService struct {
	ctx context.Context
}

// NewReduceLocationEnumService new ReduceLocationEnumService
func NewReduceLocationEnumService(ctx context.Context) *ReduceLocationEnumService {
	return &ReduceLocationEnumService{ctx: ctx}
}

// Run create note info
func (s *ReduceLocationEnumService) Run(req *micro_device.RpcReduceLocReq) (resp *common.Empty, err error) {
	err = db.GetMysql().Delete(&model.Location{}, req.Id).Error
	if err != nil {
		klog.Error(err)
		return nil, kerrors.NewBizStatusError(code.ExternalError, msg.InternalError)
	}

	return &common.Empty{}, nil
}
