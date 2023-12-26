package product_service

import (
	"context"
	product "git.zqbjj.top/pet/services/cmd/rpc/product/kitex_gen/product"
)

type UpdateRatingService struct {
	ctx context.Context
}

// NewUpdateRatingService new UpdateRatingService
func NewUpdateRatingService(ctx context.Context) *UpdateRatingService {
	return &UpdateRatingService{ctx: ctx}
}

// Run create note info
func (s *UpdateRatingService) Run(req *product.RatingReq) (resp *product.RatingResp, err error) {
	// Finish your business logic.

	return
}
