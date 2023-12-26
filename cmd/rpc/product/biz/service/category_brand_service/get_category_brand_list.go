package category_brand_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetCategoryBrandListService struct {
	ctx context.Context
}

// NewGetCategoryBrandListService new GetCategoryBrandListService
func NewGetCategoryBrandListService(ctx context.Context) *GetCategoryBrandListService {
	return &GetCategoryBrandListService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryBrandListService) Run(req *common.Req) (resp []*product.CategoryBrandInfo, err error) {
	// Finish your business logic.

	return
}
