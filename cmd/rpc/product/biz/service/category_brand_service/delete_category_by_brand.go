package category_brand_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
)

type DeleteCategoryByBrandService struct {
	ctx context.Context
}

// NewDeleteCategoryByBrandService new DeleteCategoryByBrandService
func NewDeleteCategoryByBrandService(ctx context.Context) *DeleteCategoryByBrandService {
	return &DeleteCategoryByBrandService{ctx: ctx}
}

// Run create note info
func (s *DeleteCategoryByBrandService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
