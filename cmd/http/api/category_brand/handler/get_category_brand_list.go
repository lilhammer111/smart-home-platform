package handler

import (
	"context"

	category_brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCategoryBrandListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryBrandListService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryBrandListService {
	return &GetCategoryBrandListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryBrandListService) Do(req *common.Req) (resp *[]*category_brand.CategoryBrandInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
