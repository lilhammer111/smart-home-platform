package brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetBrandByCategoryService struct {
	ctx context.Context
}

// NewGetBrandByCategoryService new GetBrandByCategoryService
func NewGetBrandByCategoryService(ctx context.Context) *GetBrandByCategoryService {
	return &GetBrandByCategoryService{ctx: ctx}
}

// Run create note info
func (s *GetBrandByCategoryService) Run(req *product.GetBrandByCatReq) (resp []*product.BrandInfo, err error) {
	// Finish your business logic.

	return
}
