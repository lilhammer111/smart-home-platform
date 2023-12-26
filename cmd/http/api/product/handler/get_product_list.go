package handler

import (
	"context"

	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductListService(Context context.Context, RequestContext *app.RequestContext) *GetProductListService {
	return &GetProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductListService) Do(req *product.ProductFilter) (resp *[]*product.BasicProdInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
