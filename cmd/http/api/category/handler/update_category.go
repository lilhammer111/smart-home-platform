package handler

import (
	"context"

	category "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateCategoryService(Context context.Context, RequestContext *app.RequestContext) *UpdateCategoryService {
	return &UpdateCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateCategoryService) Do(req *category.CategoryInfo) (resp *category.CategoryInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
