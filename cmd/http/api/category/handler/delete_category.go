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

type DeleteCategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteCategoryService(Context context.Context, RequestContext *app.RequestContext) *DeleteCategoryService {
	return &DeleteCategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteCategoryService) Do(req *common.Req) (resp *common.Empty, err error) {
	_, err = micro_product_cli.DeleteCategory(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Deleting category successes.")
	return &common.Empty{}, nil
}
