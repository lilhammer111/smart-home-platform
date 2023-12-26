package category_brand_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateCategoryBrandService struct {
	ctx context.Context
}

// NewUpdateCategoryBrandService new UpdateCategoryBrandService
func NewUpdateCategoryBrandService(ctx context.Context) *UpdateCategoryBrandService {
	return &UpdateCategoryBrandService{ctx: ctx}
}

// Run create note info
func (s *UpdateCategoryBrandService) Run(req *product.CategoryBrandInfo) (resp *product.CategoryBrandInfo, err error) {
	// Finish your business logic.

	return
}
