package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/product"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetProductDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductDetailService(Context context.Context, RequestContext *app.RequestContext) *GetProductDetailService {
	return &GetProductDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductDetailService) Do(req *common.Req) (resp *product.ProductInfo, err error) {
	productInfo, err := micro_product_cli.GetProductDetail(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &product.ProductInfo{}
	err = copier.Copy(resp, productInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting product details successes.")
	return resp, nil
}
