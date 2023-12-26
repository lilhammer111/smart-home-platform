package banner_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type AddNewBannerService struct {
	ctx context.Context
}

// NewAddNewBannerService new AddNewBannerService
func NewAddNewBannerService(ctx context.Context) *AddNewBannerService {
	return &AddNewBannerService{ctx: ctx}
}

// Run create note info
func (s *AddNewBannerService) Run(req *product.NewBanner_) (resp *product.BannerInfo, err error) {
	// Finish your business logic.

	return
}
