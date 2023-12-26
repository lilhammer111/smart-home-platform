package brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateBrandService struct {
	ctx context.Context
}

// NewUpdateBrandService new UpdateBrandService
func NewUpdateBrandService(ctx context.Context) *UpdateBrandService {
	return &UpdateBrandService{ctx: ctx}
}

// Run create note info
func (s *UpdateBrandService) Run(req *product.BrandInfo) (resp *product.BrandInfo, err error) {
	// Finish your business logic.

	return
}
