package handler

import (
	"context"

	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProdShowcaseService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProdShowcaseService(Context context.Context, RequestContext *app.RequestContext) *UpdateProdShowcaseService {
	return &UpdateProdShowcaseService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProdShowcaseService) Do(req *product.ProductShowcase) (resp *product.ProductShowcase, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
