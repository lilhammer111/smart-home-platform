package locationsrv

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/common"
	micro_device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
)

type ReduceLocationEnumService struct {
	ctx context.Context
} // NewReduceLocationEnumService new ReduceLocationEnumService
func NewReduceLocationEnumService(ctx context.Context) *ReduceLocationEnumService {
	return &ReduceLocationEnumService{ctx: ctx}
}

// Run create note info
func (s *ReduceLocationEnumService) Run(req *micro_device.RpcReduceLocReq) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
