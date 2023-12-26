package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	rpcAlert "git.zqbjj.top/pet/services/cmd/http/kitex_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_device_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetAlertListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetAlertListService(Context context.Context, RequestContext *app.RequestContext) *GetAlertListService {
	return &GetAlertListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetAlertListService) Do(req *alert.AlertFilter) (resp *[]*alert.AlertInfo, err error) {
	rpcReq := rpcAlert.AlertFilter{}
	if err = copier.Copy(&rpcReq, req); err != nil {
		hlog.Error(err)
		return nil, err
	}

	deviceInfos, err := micro_device_cli.QueryAlertsWithFilter(h.Context, &rpcReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*alert.AlertInfo)

	err = copier.Copy(resp, deviceInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "getting alert list succeed")
	return resp, nil
}
