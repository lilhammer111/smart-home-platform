package product_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetProductDetailService struct {
	ctx context.Context
}

// NewGetProductDetailService new GetProductDetailService
func NewGetProductDetailService(ctx context.Context) *GetProductDetailService {
	return &GetProductDetailService{ctx: ctx}
}

// Run create note info
func (s *GetProductDetailService) Run(req *common.Req) (resp *product.ProductDetail, err error) {
	// Finish your business logic.

	return
}
