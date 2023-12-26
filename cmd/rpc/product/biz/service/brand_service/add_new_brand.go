package brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type AddNewBrandService struct {
	ctx context.Context
}

// NewAddNewBrandService new AddNewBrandService
func NewAddNewBrandService(ctx context.Context) *AddNewBrandService {
	return &AddNewBrandService{ctx: ctx}
}

// Run create note info
func (s *AddNewBrandService) Run(req *product.NewBrand_) (resp *product.BrandInfo, err error) {
	// Finish your business logic.

	return
}
