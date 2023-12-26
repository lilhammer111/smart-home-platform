package handler

import (
	"context"

	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddNewProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewProductService(Context context.Context, RequestContext *app.RequestContext) *AddNewProductService {
	return &AddNewProductService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewProductService) Do(req *product.NewProduct) (resp *product.ProductInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
