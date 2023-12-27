package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetRelatedBrandsByCategoryIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetRelatedBrandsByCategoryIdService(Context context.Context, RequestContext *app.RequestContext) *GetRelatedBrandsByCategoryIdService {
	return &GetRelatedBrandsByCategoryIdService{RequestContext: RequestContext, Context: Context}
}

func (h *GetRelatedBrandsByCategoryIdService) Do(req *brand.BrandByCatReq) (resp *[]*brand.BrandInfo, err error) {
	brandInfos, err := micro_product_cli.GetRelatedBrandsByCategoryId(h.Context, &product.BrandByCatReq{CategoryId: req.CategoryId})
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
