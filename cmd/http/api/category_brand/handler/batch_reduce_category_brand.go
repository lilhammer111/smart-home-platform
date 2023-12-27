package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	category_brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type BatchReduceCategoryBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewBatchReduceCategoryBrandService(Context context.Context, RequestContext *app.RequestContext) *BatchReduceCategoryBrandService {
	return &BatchReduceCategoryBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *BatchReduceCategoryBrandService) Do(req *category_brand.NewCategoryBrand) (resp *common.Empty, err error) {
	newCategoryBrands := product.NewCategoryBrand_{}
	err = copier.Copy(&newCategoryBrands, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	_, err = micro_product_cli.BatchReduceCategoryBrand(h.Context, &newCategoryBrands)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Reducing related categories successes.")
	return &common.Empty{}, nil
}
