package handler

import (
	"context"

	category_brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category_brand"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateCategoryBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateCategoryBrandService(Context context.Context, RequestContext *app.RequestContext) *UpdateCategoryBrandService {
	return &UpdateCategoryBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateCategoryBrandService) Do(req *category_brand.NewCategoryBrand) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
