package handler

import (
	"context"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	product "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductDetailService(Context context.Context, RequestContext *app.RequestContext) *GetProductDetailService {
	return &GetProductDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductDetailService) Do(req *common.Req) (resp *product.ProductInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
