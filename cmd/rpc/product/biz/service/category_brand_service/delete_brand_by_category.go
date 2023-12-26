package category_brand_service

import (
	"context"
	common "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/common"
)

type DeleteBrandByCategoryService struct {
	ctx context.Context
}

// NewDeleteBrandByCategoryService new DeleteBrandByCategoryService
func NewDeleteBrandByCategoryService(ctx context.Context) *DeleteBrandByCategoryService {
	return &DeleteBrandByCategoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteBrandByCategoryService) Run(req *common.Req) (resp *common.Empty, err error) {
	// Finish your business logic.

	return
}
