package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/model"
	rpcCommon "git.zqbjj.top/pet/services/cmd/http/kitex_gen/common"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetModelDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetModelDetailService(Context context.Context, RequestContext *app.RequestContext) *GetModelDetailService {
	return &GetModelDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetModelDetailService) Do(req *common.Req) (resp *model.ModelInfo, err error) {
	modelInfo, err := micro_product_cli.GetModelDetail(h.Context, &rpcCommon.Req{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &model.ModelInfo{}
	err = copier.Copy(resp, modelInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting model detail successes.")
	return resp, nil
}
