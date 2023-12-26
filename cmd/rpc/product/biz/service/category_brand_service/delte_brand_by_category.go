package category_brand_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
)

type DelteBrandByCategoryService struct {
	ctx context.Context
}

// NewDelteBrandByCategoryService new DelteBrandByCategoryService
func NewDelteBrandByCategoryService(ctx context.Context) *DelteBrandByCategoryService {
	return &DelteBrandByCategoryService{ctx: ctx}
}

// Run create note info
func (s *DelteBrandByCategoryService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
