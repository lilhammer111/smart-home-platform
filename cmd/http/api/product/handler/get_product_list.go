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

type GetProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductListService(Context context.Context, RequestContext *app.RequestContext) *GetProductListService {
	return &GetProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductListService) Do(req *product.ProductFilter) (resp *[]*product.ProductBasicInfo, err error) {
	productFilter := rpcProduct.ProductFilter{}
	err = copier.Copy(&productFilter, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	productInfos, err := micro_product_cli.GetProductList(h.Context, &productFilter)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*product.ProductBasicInfo)
	err = copier.Copy(resp, productInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting product list successes.")
	return resp, nil
}
