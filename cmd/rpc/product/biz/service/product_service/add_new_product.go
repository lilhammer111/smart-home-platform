package product_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type AddNewProductService struct {
	ctx context.Context
}

// NewAddNewProductService new AddNewProductService
func NewAddNewProductService(ctx context.Context) *AddNewProductService {
	return &AddNewProductService{ctx: ctx}
}

// Run create note info
func (s *AddNewProductService) Run(req *product.NewProduct_) (resp *product.ProductInfo, err error) {
	// Finish your business logic.

	return
}
