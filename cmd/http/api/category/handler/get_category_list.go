package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetCategoryListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryListService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryListService {
	return &GetCategoryListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryListService) Do(req *common.PageFilter) (resp *[]*category.CategoryInfo, err error) {
	pageFilter := rpcCommon.PageFilter{}
	err = copier.Copy(&pageFilter, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	categoryInfos, err := micro_product_cli.GetCategoryList(h.Context, &pageFilter)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*category.CategoryInfo)
	err = copier.Copy(resp, categoryInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting category list successes.")
	return resp, nil
}
