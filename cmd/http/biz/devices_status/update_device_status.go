package devices_status

import (
	"context"

	devices_status "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/devices_status"
	resp "git.zqbjj.top/pet/services/cmd/http/dto/hertz_gen/resp"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateDeviceStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateDeviceStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateDeviceStatusService {
	return &UpdateDeviceStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateDeviceStatusService) Do(req *devices_status.DeviceStatusInfo) (resp *resp.StdResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return
}
