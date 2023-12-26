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

type DeleteModelService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteModelService(Context context.Context, RequestContext *app.RequestContext) *DeleteModelService {
	return &DeleteModelService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteModelService) Do(req *common.Req) (resp *common.Empty, err error) {

	_, err = micro_product_cli.DeleteModel(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "deleting product model succeed")
	return &common.Empty{}, nil
}
