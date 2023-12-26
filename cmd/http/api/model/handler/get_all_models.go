package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	model "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/model"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetAllModelsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAllModelsService(Context context.Context, RequestContext *app.RequestContext) *GetAllModelsService {
	return &GetAllModelsService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAllModelsService) Do(req *common.Empty) (resp *[]*model.ModelInfo, err error) {
	modelInfos, err := micro_product_cli.GetAllModels(h.Context)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*model.ModelInfo)
	err = copier.Copy(resp, modelInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting all models succeed!")
	return resp, nil
}
