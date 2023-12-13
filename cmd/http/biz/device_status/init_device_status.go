package device_status

import (
	"context"

	device_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device_status"
	"github.com/cloudwego/hertz/pkg/app"
)

type InitDeviceStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewInitDeviceStatusService(Context context.Context, RequestContext *app.RequestContext) *InitDeviceStatusService {
	return &InitDeviceStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *InitDeviceStatusService) Do(req *device_status.DeviceStatusInfo) (resp *device_status.DeviceStatusInfoResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
