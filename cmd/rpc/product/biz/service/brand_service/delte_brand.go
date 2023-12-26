package brand_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
)

type DelteBrandService struct {
	ctx context.Context
}

// NewDelteBrandService new DelteBrandService
func NewDelteBrandService(ctx context.Context) *DelteBrandService {
	return &DelteBrandService{ctx: ctx}
}

// Run create note info
func (s *DelteBrandService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
