package product_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
}

// NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.ProductInfo) (resp *product.ProductInfo, err error) {
	// Finish your business logic.

	return
}
