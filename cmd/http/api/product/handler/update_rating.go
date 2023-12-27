package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	rpcProduct "git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type UpdateRatingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateRatingService(Context context.Context, RequestContext *app.RequestContext) *UpdateRatingService {
	return &UpdateRatingService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateRatingService) Do(req *product.RatingReq) (resp *product.RatingInfo, err error) {
	rpcRatingReq := rpcProduct.RatingReq{}
	err = copier.Copy(&rpcRatingReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	updatedRatingInfo, err := micro_product_cli.UpdateRating(h.Context, &rpcRatingReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &product.RatingInfo{}
	err = copier.Copy(resp, updatedRatingInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Updating category successes.")
	return resp, nil
}
