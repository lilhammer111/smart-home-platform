package locationsrv

import (
	"context"
	micro_device "git.zqbjj.top/pet/services/cmd/rpc/device/kitex_gen/micro_device"
)

type FindLocationTitleService struct {
	ctx context.Context
} // NewFindLocationTitleService new FindLocationTitleService
func NewFindLocationTitleService(ctx context.Context) *FindLocationTitleService {
	return &FindLocationTitleService{ctx: ctx}
}

// Run create note info
func (s *FindLocationTitleService) Run(req *micro_device.RpcFindLocReq) (resp *micro_device.LocationData, err error) {
	// Finish your business logic.

	return
}
