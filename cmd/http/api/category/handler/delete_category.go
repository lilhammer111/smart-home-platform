package handler

import (
	"context"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCategoryService(Context context.Context, RequestContext *app.RequestContext) *DeleteCategoryService {
	return &DeleteCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCategoryService) Do(req *common.Req) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
