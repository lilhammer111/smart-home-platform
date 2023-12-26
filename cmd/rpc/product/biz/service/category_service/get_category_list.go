package category_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetCategoryListService struct {
	ctx context.Context
}

// NewGetCategoryListService new GetCategoryListService
func NewGetCategoryListService(ctx context.Context) *GetCategoryListService {
	return &GetCategoryListService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryListService) Run(req *product.PageFilter) (resp []*product.CategoryBasicInfo, err error) {
	// Finish your business logic.

	return
}
