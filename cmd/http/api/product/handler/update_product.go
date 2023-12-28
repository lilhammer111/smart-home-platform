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

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Do(req *product.ProductInfo) (resp *product.ProductInfo, err error) {
	rpcProductReq := rpcProduct.ProductInfo{}
	err = copier.Copy(&rpcProductReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	updatedProductInfo, err := micro_product_cli.UpdateProduct(h.Context, &rpcProductReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &product.ProductInfo{}
	err = copier.Copy(resp, updatedProductInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Updating product successes.")
	return resp, nil
}
