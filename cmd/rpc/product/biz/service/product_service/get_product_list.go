package product_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetProductListService struct {
	ctx context.Context
}

// NewGetProductListService new GetProductListService
func NewGetProductListService(ctx context.Context) *GetProductListService {
	return &GetProductListService{ctx: ctx}
}

// Run create note info
func (s *GetProductListService) Run(req *product.ProductFilter) (resp []*product.ProductDetail, err error) {
	// Finish your business logic.

	return
}
