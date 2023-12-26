package category_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type AddNewCategoryService struct {
	ctx context.Context
}

// NewAddNewCategoryService new AddNewCategoryService
func NewAddNewCategoryService(ctx context.Context) *AddNewCategoryService {
	return &AddNewCategoryService{ctx: ctx}
}

// Run create note info
func (s *AddNewCategoryService) Run(req *product.AddCategoryReq) (resp *product.CategoryInfo, err error) {
	// Finish your business logic.

	return
}
