package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type AddNewCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewCategoryService(Context context.Context, RequestContext *app.RequestContext) *AddNewCategoryService {
	return &AddNewCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewCategoryService) Do(req *category.NewCategory) (resp *category.CategoryInfo, err error) {
	rpcCategoryReq := product.NewCategory_{}
	err = copier.Copy(&rpcCategoryReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	categoryInfo, err := micro_product_cli.AddNewCategory(h.Context, &rpcCategoryReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &category.CategoryInfo{}
	err = copier.Copy(resp, categoryInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Adding a new category successes.")
	return resp, nil
}
