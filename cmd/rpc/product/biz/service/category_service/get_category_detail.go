package category_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetCategoryDetailService struct {
	ctx context.Context
}

// NewGetCategoryDetailService new GetCategoryDetailService
func NewGetCategoryDetailService(ctx context.Context) *GetCategoryDetailService {
	return &GetCategoryDetailService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryDetailService) Run(req *common.Req) (resp *product.CategoryDetail, err error) {
	// Finish your business logic.

	return
}
