package banner_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type GetAllBannersService struct {
	ctx context.Context
}

// NewGetAllBannersService new GetAllBannersService
func NewGetAllBannersService(ctx context.Context) *GetAllBannersService {
	return &GetAllBannersService{ctx: ctx}
}

// Run create note info
func (s *GetAllBannersService) Run() (resp []*product.BannerInfo, err error) {
	// Finish your business logic.

	return
}
