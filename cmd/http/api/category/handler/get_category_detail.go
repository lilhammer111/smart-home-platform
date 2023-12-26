package handler

import (
	"context"

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

func (h *GetCategoryDetailService) Do(req *common.Req) (resp *category.CategoryDetail, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
