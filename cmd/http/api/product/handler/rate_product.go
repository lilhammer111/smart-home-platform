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

type RateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRateProductService(Context context.Context, RequestContext *app.RequestContext) *RateProductService {
	return &RateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *RateProductService) Do(req *product.RatingReq) (resp *product.RatingResp, err error) {
	rpcRatingReq := rpcProduct.RatingReq{}
	err = copier.Copy(&rpcRatingReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	updatedRatingInfo, err := micro_product_cli.RateProduct(h.Context, &rpcRatingReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &product.RatingResp{}
	err = copier.Copy(resp, updatedRatingInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Updating category successes.")
	return resp, nil
}
