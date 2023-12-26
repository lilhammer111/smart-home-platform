package product_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateProductBasicInfoService struct {
	ctx context.Context
}

// NewUpdateProductBasicInfoService new UpdateProductBasicInfoService
func NewUpdateProductBasicInfoService(ctx context.Context) *UpdateProductBasicInfoService {
	return &UpdateProductBasicInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductBasicInfoService) Run(req *product.BasicProdInfo) (resp *product.BasicProdInfo, err error) {
	// Finish your business logic.

	return
}
