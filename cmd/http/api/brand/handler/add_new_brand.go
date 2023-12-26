package handler

import (
	"context"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddNewBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewBrandService(Context context.Context, RequestContext *app.RequestContext) *AddNewBrandService {
	return &AddNewBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewBrandService) Do(req *brand.NewBrand) (resp *brand.BrandInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
