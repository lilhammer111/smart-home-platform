package handler

import (
	"context"

	category "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCategoryListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCategoryListService(Context context.Context, RequestContext *app.RequestContext) *GetCategoryListService {
	return &GetCategoryListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCategoryListService) Do(req *common.PageFilter) (resp *[]*category.CategoryBasicInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
