package category_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateCategoryService struct {
	ctx context.Context
}

// NewUpdateCategoryService new UpdateCategoryService
func NewUpdateCategoryService(ctx context.Context) *UpdateCategoryService {
	return &UpdateCategoryService{ctx: ctx}
}

// Run create note info
func (s *UpdateCategoryService) Run(req *product.CategoryInfo) (resp *product.CategoryInfo, err error) {
	// Finish your business logic.

	return
}
