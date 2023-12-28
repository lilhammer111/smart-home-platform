package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Do(req *common.Req) (resp *common.Empty, err error) {
	_, err = micro_product_cli.DeleteProduct(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Deleting product successes.")
	return &common.Empty{}, nil
}
