package banner_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateBannerService struct {
	ctx context.Context
}

// NewUpdateBannerService new UpdateBannerService
func NewUpdateBannerService(ctx context.Context) *UpdateBannerService {
	return &UpdateBannerService{ctx: ctx}
}

// Run create note info
func (s *UpdateBannerService) Run(req *product.BannerInfo) (resp *product.BannerInfo, err error) {
	// Finish your business logic.

	return
}
