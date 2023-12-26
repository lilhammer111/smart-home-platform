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

type GetBrandListByCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBrandListByCategoryService(Context context.Context, RequestContext *app.RequestContext) *GetBrandListByCategoryService {
	return &GetBrandListByCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBrandListByCategoryService) Do(req *brand.BrandByCatReq) (resp *[]*brand.BrandInfo, err error) {
	brandInfos, err := micro_product_cli.GetBrandListByCategory(h.Context, &product.BrandByCatReq{CategoryId: req.CategoryId})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*brand.BrandInfo)
	err = copier.Copy(resp, brandInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting brands list successes.")
	return resp, nil
}
