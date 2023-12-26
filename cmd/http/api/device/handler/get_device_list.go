package handler

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	rpcDevice "git.zqbjj.top/pet/services/cmd/http/kitex_gen/device"
	"git.zqbjj.top/pet/services/cmd/http/utils/micro_device_cli"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jinzhu/copier"
)

type GetDeviceListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetDeviceListService(Context context.Context, RequestContext *app.RequestContext) *GetDeviceListService {
	return &GetDeviceListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetDeviceListService) Do(req *device.DeviceFilter) (resp *[]*device.DeviceInfo, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcReq := rpcDevice.DeviceFilter{}
	if err = copier.Copy(&rpcReq, req); err != nil {
		hlog.Error(err)
		return nil, err
	}

	deviceInfos, err := micro_device_cli.QueryDevicesWithFilter(h.Context, &rpcReq)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	resp = new([]*device.DeviceInfo)

	err = copier.Copy(resp, deviceInfos)
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return resp, nil
}
