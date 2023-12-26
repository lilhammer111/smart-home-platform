package handler

import (
	"context"

	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductBasicInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductBasicInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductBasicInfoService {
	return &UpdateProductBasicInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductBasicInfoService) Do(req *product.BasicProdInfo) (resp *product.BasicProdInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
