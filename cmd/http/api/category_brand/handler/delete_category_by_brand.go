package handler

import (
	"context"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteCategoryByBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCategoryByBrandService(Context context.Context, RequestContext *app.RequestContext) *DeleteCategoryByBrandService {
	return &DeleteCategoryByBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCategoryByBrandService) Do(req *common.Req) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
