package brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetBrandListService struct {
	ctx context.Context
}

// NewGetBrandListService new GetBrandListService
func NewGetBrandListService(ctx context.Context) *GetBrandListService {
	return &GetBrandListService{ctx: ctx}
}

// Run create note info
func (s *GetBrandListService) Run(req *product.PageFilter) (resp []*product.BrandInfo, err error) {
	// Finish your business logic.

	return
}
