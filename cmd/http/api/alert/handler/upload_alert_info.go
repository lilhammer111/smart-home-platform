package handler

import (
	"context"
	rpcAlert "git.zqbjj.top/pet/services/cmd/http/kitex_gen/alert"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_device_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	alert "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/alert"
	"github.com/cloudwego/hertz/pkg/app"
)

type UploadAlertInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadAlertInfoService(Context context.Context, RequestContext *app.RequestContext) *UploadAlertInfoService {
	return &UploadAlertInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadAlertInfoService) Do(req *alert.AlertInfo) (resp *alert.AlertInfo, err error) {
	rpcReq := rpcAlert.AlertInfo{}
	err = copier.Copy(&rpcReq, req)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	alertInfo, err := micro_device_cli.CreateAlert(h.Context, &rpcReq)
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

	h.RequestContext.Set(responder.SuccessMessage, "uploading alert info succeed")
	return resp, nil
}
