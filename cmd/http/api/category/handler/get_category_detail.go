package handler

import (
	"context"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	category "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCategoryDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryDetailService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryDetailService {
	return &GetCategoryDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryDetailService) Do(req *common.Req) (resp *category.CategoryInfo, err error) {
	categoryInfo, err := micro_product_cli.GetCategoryDetail(h.Context, &rpcCommon.Req{Id: req.Id})
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

	h.RequestContext.Set(responder.SuccessMessage, "Getting category details successes.")
	return resp, nil
}
