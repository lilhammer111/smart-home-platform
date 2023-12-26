package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_device_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetAlertDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAlertDetailService(Context context.Context, RequestContext *app.RequestContext) *GetAlertDetailService {
	return &GetAlertDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAlertDetailService) Do(req *common.Req) (resp *alert.AlertInfo, err error) {
	alertInfo, err := micro_device_cli.FindAlert(h.Context, &micro_device.RpcFindAlertReq{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &alert.AlertInfo{}
	err = copier.Copy(resp, alertInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting alert info succeed!")
	return resp, nil
}
