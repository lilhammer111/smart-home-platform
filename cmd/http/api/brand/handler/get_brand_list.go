package handler

import (
	"context"

	brand "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/brand"
	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetBrandListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetBrandListService(Context context.Context, RequestContext *app.RequestContext) *GetBrandListService {
	return &GetBrandListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetBrandListService) Do(req *common.PageFilter) (resp *[]*brand.BrandInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
