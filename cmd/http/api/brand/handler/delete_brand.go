package handler

import (
	"context"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteBrandService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteBrandService(Context context.Context, RequestContext *app.RequestContext) *DeleteBrandService {
	return &DeleteBrandService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteBrandService) Do(req *common.Req) (resp *common.Empty, err error) {
	_, err = micro_product_cli.DeleteBrand(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Deleting brand successes.")
	return &common.Empty{}, nil
}
