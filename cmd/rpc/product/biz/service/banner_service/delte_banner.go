package banner_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
)

type DelteBannerService struct {
	ctx context.Context
}

// NewDelteBannerService new DelteBannerService
func NewDelteBannerService(ctx context.Context) *DelteBannerService {
	return &DelteBannerService{ctx: ctx}
}

// Run create note info
func (s *DelteBannerService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
