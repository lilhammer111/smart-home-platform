package alert

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_device_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteAlertService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteAlertService(Context context.Context, RequestContext *app.RequestContext) *DeleteAlertService {
	return &DeleteAlertService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteAlertService) Do(req *common.Req) (resp *common.Empty, err error) {
	_, err = micro_device_cli.DeleteAlert(h.Context, &micro_device.RpcDeleteAlertReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "deleting alert info succeed")
	return nil, nil
}
