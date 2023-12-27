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

type UpdateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateCategoryService(Context context.Context, RequestContext *app.RequestContext) *UpdateCategoryService {
	return &UpdateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateCategoryService) Do(req *category.CategoryInfo) (resp *category.CategoryInfo, err error) {
	rpcCategoryReq := product.CategoryInfo{}
	err = copier.Copy(&rpcCategoryReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	updatedCategoryInfo, err := micro_product_cli.UpdateCategory(h.Context, &rpcCategoryReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &category.CategoryInfo{}
	err = copier.Copy(resp, updatedCategoryInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Updating category successes.")
	return resp, nil
}
