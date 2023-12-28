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

type AddNewProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewProductService(Context context.Context, RequestContext *app.RequestContext) *AddNewProductService {
	return &AddNewProductService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewProductService) Do(req *product.AddProductReq) (resp *product.ProductInfo, err error) {
	rpcProductReq := rpcProduct.AddProductReq{}
	err = copier.Copy(&rpcProductReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	categoryInfo, err := micro_product_cli.AddNewProduct(h.Context, &rpcProductReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &product.ProductInfo{}
	err = copier.Copy(resp, categoryInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Adding a new product successes.")
	return resp, nil
}
