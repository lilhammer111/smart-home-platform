package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/kitex_gen/micro_device"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_device_cli"
	"git.zqbjj.top/pet/services/cmd/http/utils/responder"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"

	common "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/common"
	device "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetDeviceDetailService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetDeviceDetailService(Context context.Context, RequestContext *app.RequestContext) *GetDeviceDetailService {
	return &GetDeviceDetailService{RequestContext: RequestContext, Context: Context}
}

func (h *GetDeviceDetailService) Do(req *common.Req) (resp *device.DeviceInfo, err error) {
	alertInfo, err := micro_device_cli.FindDevice(h.Context, &micro_device.RpcFindDeviceReq{Id: req.Id})
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = &device.DeviceInfo{}
	err = copier.Copy(resp, alertInfo)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	h.RequestContext.Set(responder.SuccessMessage, "Getting device info succeed!")
	return resp, nil
}
