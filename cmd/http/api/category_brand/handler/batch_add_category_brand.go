package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	category_brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	"github.com/cloudwego/hertz/pkg/app"
)

type BatchAddCategoryBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBatchAddCategoryBrandService(Context context.Context, RequestContext *app.RequestContext) *BatchAddCategoryBrandService {
	return &BatchAddCategoryBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *BatchAddCategoryBrandService) Do(req *category_brand.NewCategoryBrand) (resp *[]*category_brand.CategoryBrandInfo, err error) {
	newCategoryBrands := product.NewCategoryBrand_{}
	err = copier.Copy(&newCategoryBrands, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	addedCategoryBrand, err := micro_product_cli.BatchAddCategoryBrand(h.Context, &newCategoryBrands)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*category_brand.CategoryBrandInfo)
	err = copier.Copy(resp, addedCategoryBrand)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Adding all related categories successes.")
	return resp, nil
}
