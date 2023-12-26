package handler

import (
	"context"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateBrandService(Context context.Context, RequestContext *app.RequestContext) *UpdateBrandService {
	return &UpdateBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateBrandService) Do(req *brand.BrandInfo) (resp *brand.BrandInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
