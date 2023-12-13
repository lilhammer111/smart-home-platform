package device

import (
	"context"

	device "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateDeviceInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateDeviceInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateDeviceInfoService {
	return &UpdateDeviceInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateDeviceInfoService) Do(req *device.DeviceInfo) (resp *device.DeviceInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
