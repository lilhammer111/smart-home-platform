package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddNewBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewBrandService(Context context.Context, RequestContext *app.RequestContext) *AddNewBrandService {
	return &AddNewBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewBrandService) Do(req *brand.NewBrand) (resp *brand.BrandInfo, err error) {
	rpcBrandReq := product.NewBrand_{}
	err = copier.Copy(&rpcBrandReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	brandInfo, err := micro_product_cli.AddNewBrand(h.Context, &rpcBrandReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &brand.BrandInfo{}
	err = copier.Copy(resp, brandInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Adding a new brand successes.")
	return resp, nil
}
