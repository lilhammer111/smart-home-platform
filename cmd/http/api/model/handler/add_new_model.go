package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/product"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_product_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	model "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/model"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddNewModelService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddNewModelService(Context context.Context, RequestContext *app.RequestContext) *AddNewModelService {
	return &AddNewModelService{RequestContext: RequestContext, Context: Context}
}

func (h *AddNewModelService) Do(req *model.NewModel) (resp *model.ModelInfo, err error) {
	modelInfoPtr, err := micro_product_cli.AddNewModel(h.Context, &product.NewModel_{Name: req.Name})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &model.ModelInfo{}
	err = copier.Copy(resp, modelInfoPtr)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Adding a new product model successes.")
	return resp, nil
}
