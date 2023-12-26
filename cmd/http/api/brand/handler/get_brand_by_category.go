package handler

import (
	"context"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetBrandByCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBrandByCategoryService(Context context.Context, RequestContext *app.RequestContext) *GetBrandByCategoryService {
	return &GetBrandByCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBrandByCategoryService) Do(req *brand.GetBrandByCatReq) (resp *[]*brand.BrandInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
