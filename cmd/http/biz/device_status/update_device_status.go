package device_status

import (
	"context"

	device_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/device_status"
	standard "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/standard"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateDeviceStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateDeviceStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateDeviceStatusService {
	return &UpdateDeviceStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateDeviceStatusService) Do(req *device_status.DeviceStatusInfo) (resp *standard.Resp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
