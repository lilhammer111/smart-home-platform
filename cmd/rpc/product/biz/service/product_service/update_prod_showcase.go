package product_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateProdShowcaseService struct {
	ctx context.Context
}

// NewUpdateProdShowcaseService new UpdateProdShowcaseService
func NewUpdateProdShowcaseService(ctx context.Context) *UpdateProdShowcaseService {
	return &UpdateProdShowcaseService{ctx: ctx}
}

// Run create note info
func (s *UpdateProdShowcaseService) Run(req *product.ProductShowcase) (resp *product.ProductShowcase, err error) {
	// Finish your business logic.

	return
}
