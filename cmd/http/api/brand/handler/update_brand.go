package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateBrandService(Context context.Context, RequestContext *app.RequestContext) *UpdateBrandService {
	return &UpdateBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateBrandService) Do(req *brand.BrandInfo) (resp *brand.BrandInfo, err error) {
	brandsReq := product.BrandInfo{}
	err = copier.Copy(&brandsReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	updatedBrandInfo, err := micro_product_cli.UpdateBrand(h.Context, &brandsReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &brand.BrandInfo{}
	err = copier.Copy(resp, updatedBrandInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Updating brand successes.")
	return resp, nil
}
