package handler

import (
	"context"

	category "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/category"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddNewCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewCategoryService(Context context.Context, RequestContext *app.RequestContext) *AddNewCategoryService {
	return &AddNewCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewCategoryService) Do(req *category.AddCategoryReq) (resp *category.CategoryInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
